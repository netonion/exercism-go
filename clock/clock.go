package clock

import "fmt"

const testVersion = 4
const minsInDay = 24 * 60

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	total := (hour*60 + minute) % minsInDay
	if total < 0 {
		total += minsInDay
	}
	return Clock{total / 60, total % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
