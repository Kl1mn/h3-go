package h3

type faceIJK struct {
	face  int
	coord coordIJK
}

type faceOrientIJK struct {
	face      int
	translate coordIJK
	ccwRot60  int
}

var (
	maxDimByCIIres = []int{
		2,        // res  0
		-1,       // res  1
		14,       // res  2
		-1,       // res  3
		98,       // res  4
		-1,       // res  5
		686,      // res  6
		-1,       // res  7
		4802,     // res  8
		-1,       // res  9
		33614,    // res 10
		-1,       // res 11
		235298,   // res 12
		-1,       // res 13
		1647086,  // res 14
		-1,       // res 15
		11529602, // res 16
	}

	unitScaleByCIIres = []int{
		1,       // res  0
		-1,      // res  1
		7,       // res  2
		-1,      // res  3
		49,      // res  4
		-1,      // res  5
		343,     // res  6
		-1,      // res  7
		2401,    // res  8
		-1,      // res  9
		16807,   // res 10
		-1,      // res 11
		117649,  // res 12
		-1,      // res 13
		823543,  // res 14
		-1,      // res 15
		5764801, // res 16
	}

	faceNeighbors = [NUM_ICOSA_FACES][4]faceOrientIJK{
		{
			// face 0
			{0, coordIJK{0, 0, 0}, 0}, // central face
			{4, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{1, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{5, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 1
			{1, coordIJK{0, 0, 0}, 0}, // central face
			{0, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{2, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{6, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 2
			{2, coordIJK{0, 0, 0}, 0}, // central face
			{1, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{3, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{7, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 3
			{3, coordIJK{0, 0, 0}, 0}, // central face
			{2, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{4, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{8, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 4
			{4, coordIJK{0, 0, 0}, 0}, // central face
			{3, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{0, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{9, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 5
			{5, coordIJK{0, 0, 0}, 0},  // central face
			{10, coordIJK{2, 2, 0}, 3}, // ij quadrant
			{14, coordIJK{2, 0, 2}, 3}, // ki quadrant
			{0, coordIJK{0, 2, 2}, 3},  // jk quadrant
		},
		{
			// face 6
			{6, coordIJK{0, 0, 0}, 0},  // central face
			{11, coordIJK{2, 2, 0}, 3}, // ij quadrant
			{10, coordIJK{2, 0, 2}, 3}, // ki quadrant
			{1, coordIJK{0, 2, 2}, 3},  // jk quadrant
		},
		{
			// face 7
			{7, coordIJK{0, 0, 0}, 0},  // central face
			{12, coordIJK{2, 2, 0}, 3}, // ij quadrant
			{11, coordIJK{2, 0, 2}, 3}, // ki quadrant
			{2, coordIJK{0, 2, 2}, 3},  // jk quadrant
		},
		{
			// face 8
			{8, coordIJK{0, 0, 0}, 0},  // central face
			{13, coordIJK{2, 2, 0}, 3}, // ij quadrant
			{12, coordIJK{2, 0, 2}, 3}, // ki quadrant
			{3, coordIJK{0, 2, 2}, 3},  // jk quadrant
		},
		{
			// face 9
			{9, coordIJK{0, 0, 0}, 0},  // central face
			{14, coordIJK{2, 2, 0}, 3}, // ij quadrant
			{13, coordIJK{2, 0, 2}, 3}, // ki quadrant
			{4, coordIJK{0, 2, 2}, 3},  // jk quadrant
		},
		{
			// face 10
			{10, coordIJK{0, 0, 0}, 0}, // central face
			{5, coordIJK{2, 2, 0}, 3},  // ij quadrant
			{6, coordIJK{2, 0, 2}, 3},  // ki quadrant
			{15, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 11
			{11, coordIJK{0, 0, 0}, 0}, // central face
			{6, coordIJK{2, 2, 0}, 3},  // ij quadrant
			{7, coordIJK{2, 0, 2}, 3},  // ki quadrant
			{16, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 12
			{12, coordIJK{0, 0, 0}, 0}, // central face
			{7, coordIJK{2, 2, 0}, 3},  // ij quadrant
			{8, coordIJK{2, 0, 2}, 3},  // ki quadrant
			{17, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 13
			{13, coordIJK{0, 0, 0}, 0}, // central face
			{8, coordIJK{2, 2, 0}, 3},  // ij quadrant
			{9, coordIJK{2, 0, 2}, 3},  // ki quadrant
			{18, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 14
			{14, coordIJK{0, 0, 0}, 0}, // central face
			{9, coordIJK{2, 2, 0}, 3},  // ij quadrant
			{5, coordIJK{2, 0, 2}, 3},  // ki quadrant
			{19, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 15
			{15, coordIJK{0, 0, 0}, 0}, // central face
			{16, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{19, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{10, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 16
			{16, coordIJK{0, 0, 0}, 0}, // central face
			{17, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{15, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{11, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 17
			{17, coordIJK{0, 0, 0}, 0}, // central face
			{18, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{16, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{12, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 18
			{18, coordIJK{0, 0, 0}, 0}, // central face
			{19, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{17, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{13, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
		{
			// face 19
			{19, coordIJK{0, 0, 0}, 0}, // central face
			{15, coordIJK{2, 0, 2}, 1}, // ij quadrant
			{18, coordIJK{2, 2, 0}, 5}, // ki quadrant
			{14, coordIJK{0, 2, 2}, 3}, // jk quadrant
		},
	}
)

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
		if h3Fat.leadingNonZeroDigit() == K_AXES_DIGIT {
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
	v := f.coord.toVec2d()
	return v.toGeo(f.face, res, false)
}

func (f *faceIJK) adjustOverageClassII(res int, pentLeading4, substrate bool) int {

	overage := 0

	maxDim := maxDimByCIIres[res]
	if substrate {
		maxDim *= 3
	}

	if substrate && f.coord.i+f.coord.j+f.coord.k == maxDim {
		overage = 1
	} else if f.coord.i+f.coord.j+f.coord.k > maxDim {
		overage = 2

		var fijkOrient faceOrientIJK

		if f.coord.k > 0 {
			if f.coord.j > 0 {
				fijkOrient = faceNeighbors[f.face][JK]
			} else {
				fijkOrient = faceNeighbors[f.face][KI]

				if pentLeading4 {
					var origin coordIJK
					origin.set(maxDim, 0, 0)
					var tmp coordIJK
					tmp.sub(&f.coord, &origin)
					tmp.rotate60cw()
					f.coord.add(&tmp, &origin)
				}
			}
		} else {
			fijkOrient = faceNeighbors[f.face][IJ]
		}

		f.face = fijkOrient.face

		for i := 0; i < fijkOrient.ccwRot60; i++ {
			f.coord.rotate60ccw()
		}

		transVec := fijkOrient.translate

		unitScale := unitScaleByCIIres[res]
		if substrate {
			unitScale *= 3
		}

		transVec.scale(unitScale)
		f.coord.add(&f.coord, &transVec)
		f.coord.normalize()

		if substrate && f.coord.i+f.coord.j+f.coord.k == maxDim {
			overage = 1
		}
	}

	return overage
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
