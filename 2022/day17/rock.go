package day17

type rock []coord

var (
	rock1 = rock{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}
	rock2 = rock{
		{1, 2},
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 0},
	}
	rock3 = rock{
		{2, 2},
		{2, 1},
		{0, 0},
		{1, 0},
		{2, 0},
	}
	rock4 = rock{
		{0, 3},
		{0, 2},
		{0, 1},
		{0, 0},
	}
	rock5 = rock{
		{0, 1},
		{1, 1},
		{0, 0},
		{1, 0},
	}
)

func (r rock) top() int {
	top := 0
	for _, c := range r {
		if c.y > top {
			top = c.y
		}
	}
	return top
}

func (r rock) bottom() int {
	bottom := 1000000
	for _, c := range r {
		if c.y < bottom {
			bottom = c.y
		}
	}
	return bottom
}

func (r rock) left() int {
	left := 1000000
	for _, c := range r {
		if c.x < left {
			left = c.x
		}
	}
	return left
}

func (r rock) right() int {
	right := 0
	for _, c := range r {
		if c.x > right {
			right = c.x
		}
	}
	return right
}

func (r rock) height() int {
	return r.top() - r.bottom() + 1
}

func (r rock) width() int {
	return r.left() - r.right() + 1
}
