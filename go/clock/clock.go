package clock

import "fmt"

func normaliseHour(h int) int {
	r := h % 24
	if h >= 0 {
		return r
	}
	if r != 0 {
		return 24 + r
	}
	return 0
}

func normaliseMinute(m int) (hours int, minutes int) {
	d, r := m/60, m%60
	if m >= 0 {
		return d, r
	}
	if r != 0 {
		return d - 1, 60 + r
	}
	return d, 0
}

type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	carry, minute := normaliseMinute(m)
	hour := normaliseHour(h + carry)
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	hours := 0
	hours, c.minute = normaliseMinute(c.minute + minutes)
	c.hour = normaliseHour(c.hour + hours)
	return c
}
