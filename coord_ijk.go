package h3

import "math"

var (
	UNIT_VECS = []*coordIJK{
		{0, 0, 0}, // direction 0
		{0, 0, 1}, // direction 1
		{0, 1, 0}, // direction 2
		{0, 1, 1}, // direction 3
		{1, 0, 0}, // direction 4
		{1, 0, 1}, // direction 5
		{1, 1, 0}, // direction 6
	}
)

type coordIJK struct {
	i, j, k int
}

func (c coordIJK) toDigit() int {
	c.normalize()

	digit := INVALID_DIGIT
	for i := 0; i < 7; i++ {
		if c.matches(UNIT_VECS[i]) {
			digit = i
			break
		}
	}

	return digit
}

func (c1 coordIJK) matches(c2 *coordIJK) bool {
	return c1.i == c2.i && c1.j == c2.j && c1.k == c2.k
}

func (c *coordIJK) upAp7() {
	i := c.i - c.k
	j := c.j - c.k

	c.i = int(math.Round(float64(3*i-j) / 7))
	c.j = int(math.Round(float64(i+2*j) / 7))
	c.k = 0

	c.normalize()
}

func (c *coordIJK) upAp7r() {
	i := c.i - c.k
	j := c.j - c.k

	c.i = int(math.Round(float64(2*i+j) / 7))
	c.j = int(math.Round(float64(3*j-i) / 7))
	c.k = 0

	c.normalize()
}

func (c *coordIJK) downAp7() {
	iVec := &coordIJK{3, 0, 1}
	jVec := &coordIJK{1, 3, 0}
	kVec := &coordIJK{0, 1, 3}

	iVec.scale(c.i)
	jVec.scale(c.j)
	kVec.scale(c.k)

	c.add(iVec, jVec)
	c.add(c, kVec)

	c.normalize()
}

func (c *coordIJK) downAp7r() {
	iVec := &coordIJK{3, 1, 0}
	jVec := &coordIJK{0, 3, 1}
	kVec := &coordIJK{1, 0, 3}

	iVec.scale(c.i)
	jVec.scale(c.j)
	kVec.scale(c.k)

	c.add(iVec, jVec)
	c.add(c, kVec)

	c.normalize()
}

func (c *coordIJK) scale(factor int) {
	c.i *= factor
	c.j *= factor
	c.k *= factor
}

func (c *coordIJK) add(c1, c2 *coordIJK) {
	c.i = c1.i + c2.i
	c.j = c1.j + c2.j
	c.k = c1.k + c2.k
}

func (c *coordIJK) sub(c1, c2 *coordIJK) {
	c.i = c1.i - c2.i
	c.j = c1.j - c2.j
	c.k = c1.k - c2.k
}

func (c *coordIJK) normalize() {
	if c.i < 0 {
		c.j -= c.i
		c.k -= c.i
		c.i = 0
	}

	if c.j < 0 {
		c.i -= c.j
		c.k -= c.j
		c.j = 0
	}

	if c.k < 0 {
		c.i -= c.k
		c.j -= c.k
		c.k = 0
	}

	min := c.i
	if c.j < min {
		min = c.j
	}
	if c.k < min {
		min = c.k
	}
	if min > 0 {
		c.i -= min
		c.j -= min
		c.k -= min
	}
}
