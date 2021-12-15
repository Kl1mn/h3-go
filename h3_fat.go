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
