package h3

type h3IndexFat struct {
	mode     int
	res      int
	baseCell int
	index    [MAX_H3_RES]int
}

func initH3IndexFat(res int) h3IndexFat {
	return h3IndexFat{
		mode:     H3_HEXAGON_MODE,
		res:      res,
		baseCell: -1,
		index:    [MAX_H3_RES]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
	}
}

func (h h3IndexFat) toH3() H3Index {
	result := H3_INIT
	result.setMode(h.mode)
	result.setRes(h.res)
	result.setBaseCell(h.baseCell)
	for r := 1; r < h.res+1; r++ {
		result.setIndexDigit(r, h.index[r-1])
	}
	return result
}

func (h h3IndexFat) toGeo() GeoCoord {
	f := h.toFaceIJK()
	return f.toGeo(h.res)
}

func (h h3IndexFat) toFaceIJK() faceIJK {
	return faceIJK{}
}

func (h h3IndexFat) leadingNonZeroDigit() int {
	for r := 0; r < h.res; r++ {
		if h.index[r] != 0 {
			return h.index[r]
		}
	}
	return 0
}

func (h *h3IndexFat) rotate60cw() {
	rotDigit := []int{
		0, // original digit 0
		3, // original digit 1
		6, // original digit 2
		2, // original digit 3
		5, // original digit 4
		1, // original digit 5
		4, // original digit 6
	}

	for r := 0; r < h.res; r++ {
		h.index[r] = rotDigit[h.index[r]]
	}
}

func (h *h3IndexFat) rotate60ccw() {
	rotDigit := []int{
		0, // original digit 0
		5, // original digit 1
		3, // original digit 2
		1, // original digit 3
		6, // original digit 4
		4, // original digit 5
		2, // original digit 6
	}

	for r := 0; r < h.res; r++ {
		h.index[r] = rotDigit[h.index[r]]
	}
}

func (h *h3IndexFat) rotatePent60ccw() {
	// rotate in place; skips any leading 1 digits (k-axis)
	rotDigit := []int{
		0, // original digit 0
		5, // original digit 1
		3, // original digit 2
		1, // original digit 3
		6, // original digit 4
		4, // original digit 5
		2, // original digit 6
	}

	foundFirstNonZeroDigit := false
	for r := 0; r < h.res; r++ {
		// rotate this digit
		h.index[r] = rotDigit[h.index[r]]

		// look for the first non-zero digit so we
		// can adjust for deleted k-axes sequence
		// if neccessary
		if !foundFirstNonZeroDigit && h.index[r] != 0 {
			foundFirstNonZeroDigit = true

			// adjust for deleted k-axes sequence
			if h.leadingNonZeroDigit() == K_AXES_DIGIT {
				h.rotate60ccw()
			}
		}
	}
}
