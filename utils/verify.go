package utils

import (
	"strconv"
)

var (
	monthDays = map[int]int{
		1:  31,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func (p *Pin) Verify(pin string) bool {
	// var wg sync.WaitGroup

	parsedPin, err := p.Parse(pin)

	if err != nil {
		return false
	}

	var t bool

	t = p.VerifyTime(parsedPin[0:6])

	if !t {
		return false
	}

	l := Luhn{Pin: parsedPin}

	t = l.Verify()

	return t
}

func (p *Pin) VerifyTime(t string) bool {
	year := StringToInt(t[0:2])
	month := StringToInt(t[2:4])
	day := StringToInt(t[4:6])
	if month != 2 {
		days, ok := monthDays[month]

		if !ok {
			return false
		}

		return day <= days
	}

	leapYear := year%4 == 0

	if leapYear {
		return day <= 29
	}

	return day <= 28
}
