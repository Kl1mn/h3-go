package h3

import "math"

const (
	MAX_H3_RES                  = 15
	H3_INIT             H3Index = 35184372088831
	H3_HEXAGON_MODE             = 1
	H3_MODE_OFFSET              = 59
	H3_RES_OFFSET               = 52
	H3_BC_OFFSET                = 45
	H3_PER_DIGIT_OFFSET         = 3
	NUM_ICOSA_FACES             = 20
	INVALID_DIGIT               = -1
	NUM_BASE_CELLS              = 122
	K_AXES_DIGIT                = 1
	M_SQRT3_2                   = 0.8660254037844386467637231707529361834714
	M_SIN60                     = M_SQRT3_2
	M_PI                        = 3.14159265358979323846
	M_PI_2                      = 1.5707963267948966
	EPSILON                     = 0.0000000000000001
	M_2PI                       = 6.28318530717958647692528676655900576839433
	M_AP7_ROT_RADS              = 0.333473172251832115336090755351601070065900389
	RES0_U_GNOMONIC             = 0.38196601125010500003
	M_SQRT7                     = 2.6457513110645905905016157536392604257102
)

var (
	deg2rad = math.Pi / 180.0
	rad2deg = 180.0 / math.Pi

	H3_MODE_MASK          = uint64(15) << H3_MODE_OFFSET
	H3_MODE_MASK_NEGATIVE = ^H3_MODE_MASK

	H3_RES_MASK          = uint64(15) << H3_RES_OFFSET
	H3_RES_MASK_NEGATIVE = ^H3_RES_MASK

	H3_BC_MASK          = uint64(127) << H3_BC_OFFSET
	H3_BC_MASK_NEGATIVE = ^H3_BC_MASK

	H3_DIGIT_MASK = uint64(7)
)
