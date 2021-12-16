package h3

import "math"

var (
	faceCenterGeo = [NUM_ICOSA_FACES]GeoCoord{
		{0.803582649718989942, 1.248397419617396099},   // face  0
		{1.307747883455638156, 2.536945009877921159},   // face  1
		{1.054751253523952054, -1.347517358900396623},  // face  2
		{0.600191595538186799, -0.450603909469755746},  // face  3
		{0.491715428198773866, 0.401988202911306943},   // face  4
		{0.172745327415618701, 1.678146885280433686},   // face  5
		{0.605929321571350690, 2.953923329812411617},   // face  6
		{0.427370518328979641, -1.888876200336285401},  // face  7
		{-0.079066118549212831, -0.733429513380867741}, // face  8
		{-0.230961644455383637, 0.506495587332349035},  // face  9
		{0.079066118549212831, 2.408163140208925497},   // face 10
		{0.230961644455383637, -2.635097066257444203},  // face 11
		{-0.172745327415618701, -1.463445768309359553}, // face 12
		{-0.605929321571350690, -0.187669323777381622}, // face 13
		{-0.427370518328979641, 1.252716453253507838},  // face 14
		{-0.600191595538186799, 2.690988744120037492},  // face 15
		{-0.491715428198773866, -2.739604450678486295}, // face 16
		{-0.803582649718989942, -1.893195233972397139}, // face 17
		{-1.307747883455638156, -0.604647643711872080}, // face 18
		{-1.054751253523952054, 1.794075294689396615},  // face 19
	}

	faceAxesAzRadsCII = [NUM_ICOSA_FACES][3]float64{
		{5.619958268523939882, 3.525563166130744542, 1.431168063737548730}, // face  0
		{5.760339081714187279, 3.665943979320991689, 1.571548876927796127}, // face  1
		{0.780213654393430055, 4.969003859179821079, 2.874608756786625655}, // face  2
		{0.430469363979999913, 4.619259568766391033, 2.524864466373195467}, // face  3
		{6.130269123335111400, 4.035874020941915804, 1.941478918548720291}, // face  4
		{2.692877706530642877, 0.598482604137447119, 4.787272808923838195}, // face  5
		{2.982963003477243874, 0.888567901084048369, 5.077358105870439581}, // face  6
		{3.532912002790141181, 1.438516900396945656, 5.627307105183336758}, // face  7
		{3.494305004259568154, 1.399909901866372864, 5.588700106652763840}, // face  8
		{3.003214169499538391, 0.908819067106342928, 5.097609271892733906}, // face  9
		{5.930472956509811562, 3.836077854116615875, 1.741682751723420374}, // face 10
		{0.138378484090254847, 4.327168688876645809, 2.232773586483450311}, // face 11
		{0.448714947059150361, 4.637505151845541521, 2.543110049452346120}, // face 12
		{0.158629650112549365, 4.347419854898940135, 2.253024752505744869}, // face 13
		{5.891865957979238535, 3.797470855586042958, 1.703075753192847583}, // face 14
		{2.711123289609793325, 0.616728187216597771, 4.805518392002988683}, // face 15
		{3.294508837434268316, 1.200113735041072948, 5.388903939827463911}, // face 16
		{3.804819692245439833, 1.710424589852244509, 5.899214794638635174}, // face 17
		{3.664438879055192436, 1.570043776661997111, 5.758833981448388027}, // face 18
		{2.361378999196363184, 0.266983896803167583, 4.455774101589558636}, // face 19
	}
)

type GeoCoord struct {
	Latitude, Longitude float64
}

func (c *GeoCoord) deg2rad() {
	c.Latitude = c.Latitude * deg2rad
	c.Longitude = c.Longitude * deg2rad
}

func (c *GeoCoord) rad2deg() {
	c.Latitude = c.Latitude * rad2deg
	c.Longitude = c.Longitude * rad2deg
}

func (c GeoCoord) isNotValid() bool {
	return isNotFinite(c.Latitude) || isNotFinite(c.Longitude)
}

func isNotFinite(n float64) bool {
	return math.IsInf(n, 0) || math.IsNaN(n)
}

func FromGeo(geo GeoCoord, res int) H3Index {
	if res < 0 || res > MAX_H3_RES {
		return 0
	}
	geo.deg2rad()
	if geo.isNotValid() {
		return 0
	}
	return geo.toH3(res)
}

func (g GeoCoord) toH3(res int) H3Index {
	h3Fat := g.toH3Fat(res)
	return h3Fat.toH3()
}

func (g GeoCoord) toH3Fat(res int) h3IndexFat {
	f := g.toFaceIJK(res)
	return f.toH3Fat(res)
}

func (g GeoCoord) toFaceIJK(res int) faceIJK {
	vec2d, face := g.toVec2d(res)
	return faceIJK{
		face:  face,
		coord: vec2d.toCoordIJK(),
	}
}

func (g GeoCoord) toVec2d(res int) (vec2d, int) {

	var (
		v    vec2d
		face int
	)

	r := faceCenterGeo[0].distRads(g)

	for f := 1; f < NUM_ICOSA_FACES; f++ {
		dist := faceCenterGeo[f].distRads(g)
		if dist < r {
			face = f
			r = dist
		}
	}

	if r < EPSILON {
		v.x = 0
		v.y = 0
		return v, face
	}

	theta := posAngleRads(faceAxesAzRadsCII[face][0] - posAngleRads(faceCenterGeo[face].geoAzimuthRads(g)))

	if isResClassIII(res) {
		theta = posAngleRads(theta - M_AP7_ROT_RADS)
	}

	r = math.Tan(r)

	r /= RES0_U_GNOMONIC
	for i := 0; i < res; i++ {
		r *= M_SQRT7
	}

	v.x = r * math.Cos(theta)
	v.y = r * math.Sin(theta)

	return v, face
}

func (p1 GeoCoord) distRads(p2 GeoCoord) float64 {
	bigC := math.Abs(p2.Longitude - p1.Longitude)
	if bigC > M_PI {
		lon1 := p1.Longitude
		if lon1 < 0 {
			lon1 += 2 * M_PI
		}
		lon2 := p2.Longitude
		if lon2 < 0 {
			lon2 += 2 * M_PI
		}

		bigC = math.Abs(lon2 - lon1)
	}

	b := M_PI_2 - p1.Latitude
	a := M_PI_2 - p2.Latitude

	cosc := math.Cos(a)*math.Cos(b) + math.Sin(a)*math.Sin(b)*math.Cos(bigC)
	if cosc > 1 {
		cosc = 1
	}
	if cosc < -1 {
		cosc = -1
	}

	return math.Acos(cosc)
}

func posAngleRads(rads float64) float64 {
	tmp := rads
	if rads < 0 {
		tmp = rads + M_2PI
	}
	if rads >= M_2PI {
		tmp -= M_2PI
	}
	return tmp
}

func constrainLng(lng float64) float64 {
	for lng > M_PI {
		lng = lng - (2 * M_PI)
	}
	for lng < -M_PI {
		lng = lng + (2 * M_PI)
	}
	return lng
}

func (p1 GeoCoord) geoAzimuthRads(p2 GeoCoord) float64 {
	return math.Atan2(math.Cos(p2.Latitude)*math.Sin(p2.Longitude-p1.Longitude),
		math.Cos(p1.Latitude)*math.Sin(p2.Latitude)-
			math.Sin(p1.Latitude)*math.Cos(p2.Latitude)*math.Cos(p2.Longitude-p1.Longitude))
}

func (p1 GeoCoord) azDistanceRads(az, distance float64) GeoCoord {

	if distance < EPSILON {
		return p1
	}

	var (
		sinlat float64
		sinlon float64
		coslon float64
		p2     GeoCoord
	)

	az = posAngleRads(az)

	if az < EPSILON || math.Abs(az-M_PI) < EPSILON {
		if az < EPSILON {
			p2.Latitude = p1.Latitude + distance
		} else {
			p2.Latitude = p1.Latitude - distance
		}

		if math.Abs(p2.Latitude-M_PI_2) < EPSILON {
			p2.Latitude = M_PI_2
			p2.Longitude = 0
		} else if math.Abs(p2.Latitude+M_PI_2) < EPSILON {
			p2.Latitude = -M_PI_2
			p2.Longitude = 0
		} else {
			p2.Longitude = constrainLng(p1.Longitude)
		}

	} else {
		sinlat = math.Sin(p1.Latitude)*math.Cos(distance) + math.Cos(p1.Latitude)*math.Sin(distance)*math.Cos(az)

		if sinlat > 1 {
			sinlat = 1
		}
		if sinlat < -1 {
			sinlat = -1
		}

		p2.Latitude = math.Asin(sinlat)
		if math.Abs(p2.Latitude-M_PI_2) < EPSILON {
			p2.Latitude = M_PI_2
			p2.Longitude = 0
		} else if math.Abs(p2.Latitude+M_PI_2) < EPSILON {
			p2.Latitude = -M_PI_2
			p2.Longitude = 0
		} else {
			sinlon = math.Sin(az) * math.Sin(distance) / math.Cos(p2.Latitude)
			coslon = (math.Cos(distance) - math.Sin(p1.Latitude)*math.Sin(p2.Latitude)) / math.Cos(p1.Latitude) / math.Cos(p2.Latitude)
			if sinlon > 1 {
				sinlon = 1
			}
			if sinlon < -1 {
				sinlon = -1
			}
			if coslon > 1 {
				sinlon = 1
			}
			if coslon < -1 {
				sinlon = -1
			}
			p2.Longitude = constrainLng(p1.Longitude + math.Atan2(sinlon, coslon))
		}
	}

	return p2
}
