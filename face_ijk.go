package h3

type faceIJK struct {
	face  int
	coord coordIJK
}

func (f faceIJK) toH3Fat(res int) h3IndexFat {

	h3Fat := initH3IndexFat(res)

	if res == 0 {
		h3Fat.baseCell = f.toBaseCell()
		h3Fat.res = 0
		return h3Fat
	}

	ijk := &f.coord

	for r := res - 1; r >= 0; r-- {
		lastIJK := *ijk
		var lastCenter coordIJK
		if isResClassIII(r + 1) {
			ijk.upAp7()
			lastCenter = *ijk
			lastCenter.downAp7()
		} else {
			ijk.upAp7r()
			lastCenter = *ijk
			lastCenter.downAp7r()
		}

		var diff coordIJK
		diff.sub(&lastIJK, &lastCenter)
		diff.normalize()

		h3Fat.index[r] = diff.toDigit()
	}

	h3Fat.baseCell = f.toBaseCell()

	numRots := f.toBaseCellCCWrot60()
	if isBaseCellPentagon(h3Fat.baseCell) {
		// force rotation out of missing k-axes sub-sequence
		if h3Fat.leadingNonZeroDigit() == K_AXES_DIGIT {
			// check for a cw/ccw offset face; default is ccw
			if baseCellsData[h3Fat.baseCell].cwOffsetPent[0] == f.face ||
				baseCellsData[h3Fat.baseCell].cwOffsetPent[1] == f.face {
				h3Fat.rotate60cw()
			} else {
				h3Fat.rotate60ccw()
			}
		}

		for i := 0; i < numRots; i++ {
			h3Fat.rotatePent60ccw()
		}
	} else {
		for i := 0; i < numRots; i++ {
			h3Fat.rotate60ccw()
		}
	}

	return h3Fat
}

func (f faceIJK) toGeo(res int) GeoCoord {
	return GeoCoord{}
}

func (f faceIJK) toBaseCell() int {
	return faceIJKBaseCells[f.face][f.coord.i][f.coord.j][f.coord.k].baseCell
}

func (f faceIJK) toBaseCellCCWrot60() int {
	return faceIJKBaseCells[f.face][f.coord.i][f.coord.j][f.coord.k].ccwRot60
}

func isResClassIII(res int) bool {
	return res%2 != 0
}
