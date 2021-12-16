package h3

var (
	rotate60cwDigits      = []int{0, 3, 6, 2, 5, 1, 4}
	rotate60ccwDigits     = []int{0, 5, 3, 1, 6, 4, 2}
	rotatePent60ccwDigits = []int{0, 5, 3, 1, 6, 4, 2}
)

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

	if isBaseCellPentagon(h.baseCell) && h.leadingNonZeroDigit() == 5 {
		h.rotate60cw()
	}

	var (
		fijk = baseCellsData[h.baseCell].homeFijk
		ok   bool
	)

	if fijk, ok = h.toFaceIjkWithInitializedFijk(fijk); !ok {
		return fijk
	}

	origCoord := fijk.coord

	res := h.res
	if isResClassIII(h.res) {
		fijk.coord.downAp7r()
		res++
	}

	pentLeading4 := isBaseCellPentagon(h.baseCell) && h.leadingNonZeroDigit() == 4
	if fijk.adjustOverageClassII(res, pentLeading4, false) > 0 {
		if isBaseCellPentagon(h.baseCell) {
			for {
				if fijk.adjustOverageClassII(res, false, false) == 0 {
					break
				}
			}
		}

		if res != h.res {
			fijk.coord.upAp7r()
		}
	} else if res != h.res {
		fijk.coord = origCoord
	}

	return fijk
}

func (h h3IndexFat) toFaceIjkWithInitializedFijk(base faceIJK) (faceIJK, bool) {

	coord := base.coord
	ok := true

	if !isBaseCellPentagon(h.baseCell) && (h.res == 0 || (base.coord.i == 0 && base.coord.j == 0 && base.coord.k == 0)) {
		ok = false
	}

	for r := 0; r < h.res; r++ {
		if isResClassIII(r + 1) {
			coord.downAp7()
		} else {
			coord.downAp7r()
		}
		coord.neighbor(h.index[r])
	}

	return faceIJK{
		face:  base.face,
		coord: coord,
	}, ok
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
	for r := 0; r < h.res; r++ {
		h.index[r] = rotate60cwDigits[h.index[r]]
	}
}

func (h *h3IndexFat) rotate60ccw() {
	for r := 0; r < h.res; r++ {
		h.index[r] = rotate60ccwDigits[h.index[r]]
	}
}

func (h *h3IndexFat) rotatePent60ccw() {
	foundFirstNonZeroDigit := false
	for r := 0; r < h.res; r++ {
		h.index[r] = rotatePent60ccwDigits[h.index[r]]

		if !foundFirstNonZeroDigit && h.index[r] != 0 {
			foundFirstNonZeroDigit = true

			if h.leadingNonZeroDigit() == K_AXES_DIGIT {
				h.rotate60ccw()
			}
		}
	}
}
