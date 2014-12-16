package clock

import "fmt"

type Clock struct {
	minutes int
}

func normalise(m int) int {
	minutes := m % (24 * 60)
	if minutes < 0 {
		minutes += (24 * 60)
	}
	return minutes
}

func New(h, m int) Clock {
	return Clock{normalise(h*60 + m)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c Clock) Add(m int) Clock {
	c.minutes = normalise(c.minutes + m)
	return c
}
