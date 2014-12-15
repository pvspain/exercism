package clock

import "fmt"

func saneHour(hour int) int {
	if hour >= 0 {
		return hour % 24
	}
	return 0
}

func saneMinute(minute int) int {
	if minute >= 0 {
		return minute % 60
	}
	return 0
}

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{saneHour(hour), saneMinute(minute)}
}

func (clock *Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clock.hour, clock.minute)
}

func (clock *Clock) Add(minutes int) {
	sumMinutes := clock.minute + minutes
	clock.minute = saneMinute(sumMinutes)
	clock.hour = saneHour(clock.hour + sumMinutes/60)
}
