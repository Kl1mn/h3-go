package h3

type H3Index uint64

func (h *H3Index) setMode(mode int) {
	*h = H3Index((uint64(*h) & H3_MODE_MASK_NEGATIVE) | (uint64(mode) << H3_MODE_OFFSET))
}

func (h *H3Index) setRes(res int) {
	*h = H3Index((uint64(*h) & H3_RES_MASK_NEGATIVE) | (uint64(res) << H3_RES_OFFSET))
}

func (h *H3Index) setBaseCell(bc int) {
	*h = H3Index((uint64(*h) & H3_BC_MASK_NEGATIVE) | (uint64(bc) << H3_BC_OFFSET))
}

func (h *H3Index) setIndexDigit(res, digit int) {
	*h = H3Index((uint64(*h) & ^(H3_DIGIT_MASK << (uint64(MAX_H3_RES-res) * H3_PER_DIGIT_OFFSET))) | (uint64(digit) << (uint64(MAX_H3_RES-res) * H3_PER_DIGIT_OFFSET)))
}
