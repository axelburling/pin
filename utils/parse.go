package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	length                     = 10
	lengthWithOutCheckDigit    = length - 1
	maxLength                  = 12
	maxLengthWithOutCheckDigit = maxLength - 1
	spaceChar                  = "-"
)

func (p *Pin) Parse(pin string) (string, error) {
	var century, year, check, gender, place, month, day string

	var parsedPin string

	new := strings.Replace(pin, spaceChar, "", -1)

	if len(new) == maxLength {

		century = new[0:2]
		year = new[2:4]
		month = new[4:6]
		day = new[6:8]
		place = new[8:10]
		gender = new[10:11]
		check = new[11:12]

		parsedPin = new[2:]
	} else if len(new) == length {

		year = new[0:2]
		month = new[2:4]
		day = new[4:6]
		place = new[6:8]
		gender = new[8:9]
		check = new[9:10]

		parsedPin = new
	} else {
		return "", errors.New("invalid length")
	}

	p.Century = century
	p.Year = year
	p.Month = month
	p.Day = day
	p.Place = place
	p.Gender = gender
	p.ControlDigit = check

	if p.Century == "" {
		now := time.Now().Year()
		centuryStr := strconv.Itoa(now - ((now - StringToInt(year)) % 100))

		p.Century = centuryStr[0:2]
	}

	p.FullYear = p.Century + p.Year

	return parsedPin, nil
}

func (p *Pin) ParseCheck(pin string) (string, error) {
	var parsedPin string

	new := strings.Replace(pin, spaceChar, "", -1)

	if len(new) == maxLengthWithOutCheckDigit {

		parsedPin = new[2:]

	} else if len(new) == lengthWithOutCheckDigit {
		parsedPin = new
	} else {
		return "", errors.New("invalid length")
	}

	return parsedPin, nil
}
