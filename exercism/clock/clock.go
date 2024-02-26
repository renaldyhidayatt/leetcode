package clock

import "fmt"

type Clock struct {
	m int
}

func New(h, m int) Clock {
	minutes := h*60 + m

	minutes %= 24 * 60

	if minutes < 0 {
		minutes += 24 * 60
	}

	return Clock{
		m: minutes,
	}
}

func (c Clock) Add(m int) Clock {
	c = New(c.hours(), c.minutes()-m)

	return c
}

func (c Clock) Subtract(m int) Clock {
	c = New(c.hours(), c.minutes()-m)

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours(), c.minutes())
}

func (c Clock) hours() int {
	return (c.m / 60) - (c.m / (60 * 24))
}

func (c Clock) minutes() int {
	return c.m - (c.hours() * 60)
}
