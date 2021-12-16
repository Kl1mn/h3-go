package h3

import "math"

type vec2d struct {
	x, y float64
}

func (v vec2d) toCoordIJK() coordIJK {

	var (
		a1, a2 float64
		x1, x2 float64
		m1, m2 int
		r1, r2 float64

		coord coordIJK
	)

	a1 = math.Abs(v.x)
	a2 = math.Abs(v.y)

	x2 = a2 / M_SIN60
	x1 = a1 + x2/2

	m1 = int(x1)
	m2 = int(x2)

	r1 = x1 - float64(m1)
	r2 = x2 - float64(m2)

	if r1 < 0.5 {
		if r1 < float64(1)/3 {
			if r2 < (1+r1)/2 {
				coord.i = m1
				coord.j = m2
			} else {
				coord.i = m1
				coord.j = m2 + 1
			}
		} else {
			if r2 < (1 - r1) {
				coord.j = m2
			} else {
				coord.j = m2 + 1
			}

			if (1-r1) <= r2 && r2 < (2*r1) {
				coord.i = m1 + 1
			} else {
				coord.i = m1
			}
		}
	} else {
		if r1 < float64(2)/3 {
			if r2 < (1 - r1) {
				coord.j = m2
			} else {
				coord.j = m2 + 1
			}

			if (2*r1-1) < r2 && r2 < (1-r1) {
				coord.i = m1
			} else {
				coord.i = m1 + 1
			}
		} else {
			if r2 < (r1 / 2) {
				coord.i = m1 + 1
				coord.j = m2
			} else {
				coord.i = m1 + 1
				coord.j = m2 + 1
			}
		}
	}

	if v.x < 0 {
		if (coord.j % 2) == 0 {
			axisi := int(coord.j / 2)
			diff := int(coord.i - axisi)
			coord.i = coord.i - 2*diff
		} else {
			axisi := int((coord.j + 1) / 2)
			diff := int(coord.i - axisi)
			coord.i = coord.i - (2*diff + 1)
		}
	}

	if v.y < 0 {
		coord.i = coord.i - (2*coord.j+1)/2
		coord.j = -1 * coord.j
	}

	coord.normalize()

	return coord
}

func (v vec2d) toGeo(face, res int, substrate bool) GeoCoord {

	r := v.mag()

	if r < EPSILON {
		return faceCenterGeo[face]
	}

	theta := math.Atan2(v.y, v.x)

	for i := 0; i < res; i++ {
		r /= M_SQRT7
	}

	if substrate {
		r /= 3.0
		if isResClassIII(res) {

			r /= M_SQRT7
		}
	}

	r *= RES0_U_GNOMONIC

	r = math.Atan(r)

	if !substrate && isResClassIII(res) {
		theta = posAngleRads(theta + M_AP7_ROT_RADS)
	}

	theta = posAngleRads(faceAxesAzRadsCII[face][0] - theta)

	return faceCenterGeo[face].azDistanceRads(theta, r)
}

func (v vec2d) mag() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
