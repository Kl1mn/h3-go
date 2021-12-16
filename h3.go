package h3

type H3Index uint64

func (h *H3Index) setMode(mode int) {
	*h = H3Index((uint64(*h) & H3_MODE_MASK_NEGATIVE) | (uint64(mode) << H3_MODE_OFFSET))
}

func (h H3Index) getMode() int {
	return int((uint64(h) & H3_MODE_MASK) >> H3_MODE_OFFSET)
}

func (h *H3Index) setRes(res int) {
	*h = H3Index((uint64(*h) & H3_RES_MASK_NEGATIVE) | (uint64(res) << H3_RES_OFFSET))
}

func (h H3Index) getRes() int {
	return int((uint64(h) & H3_RES_MASK) >> H3_RES_OFFSET)
}

func (h *H3Index) setBaseCell(bc int) {
	*h = H3Index((uint64(*h) & H3_BC_MASK_NEGATIVE) | (uint64(bc) << H3_BC_OFFSET))
}

func (h H3Index) getBaseCell() int {
	return int((uint64(h) & H3_BC_MASK) >> H3_BC_OFFSET)
}

func (h *H3Index) setIndexDigit(res, digit int) {
	*h = H3Index((uint64(*h) & ^(H3_DIGIT_MASK << (uint64(MAX_H3_RES-res) * H3_PER_DIGIT_OFFSET))) | (uint64(digit) << (uint64(MAX_H3_RES-res) * H3_PER_DIGIT_OFFSET)))
}

func (h H3Index) getIndexDigit(res int) int {
	return int((uint64(h) >> (uint64(MAX_H3_RES-res) * H3_PER_DIGIT_OFFSET)) & H3_DIGIT_MASK)
}

func ToGeo(h3 H3Index) GeoCoord {
	hf := h3.toH3Fat()
	geo := hf.toGeo()
	geo.rad2deg()
	return geo
}

func (h H3Index) toH3Fat() h3IndexFat {
	hf := initH3IndexFat(h.getRes())
	hf.mode = h.getMode()
	hf.baseCell = h.getBaseCell()
	for r := 1; r <= MAX_H3_RES; r++ {
		hf.index[r-1] = h.getIndexDigit(r)
	}
	return hf
}
