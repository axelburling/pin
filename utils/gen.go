package utils

import (
	"fmt"
	"math/rand"
)

const (
	separator = "-"
)

func isEven(num int) bool {
	return num%2 == 0
}

func randomEvenNumber(num int) int {
	if isEven(num) {
		return num
	}

	return randomEvenNumber(rand.Intn(10))
}

func randomOddNumber(num int) int {
	if !isEven(num) {
		return num
	}

	return randomOddNumber(rand.Intn(10))
}

func randomPlace() string {
	place := rand.Intn(99)

	if place < 10 {
		return fmt.Sprintf("0%d", place)
	} else {
		return fmt.Sprintf("%d", place)
	}
}

func Generate(bday string, gender bool, sep bool) (string, error) {

	var genderNum int

	num := rand.Intn(10)

	if gender {
		genderNum = randomEvenNumber(num)
	} else {
		genderNum = randomOddNumber(num)
	}

	place := randomPlace()

	var res string

	var val = fmt.Sprintf("%s%s%d", bday, place, genderNum)

	l := Luhn{Pin: val}
	checksum, err := l.Calculate()

	if err != nil {
		return "", err
	}

	if sep {
		res = fmt.Sprintf("%s%s%s%d%s", bday, separator, place, genderNum, checksum)
	} else {
		res = fmt.Sprintf("%s%s", val, checksum)
	}

	return res, nil
}
