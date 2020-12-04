package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear string
	IssueYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportID string
	CountryID string
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func parsePassport(text string) (passport Passport) {
	fields := strings.Split(text, " ")
	for _, field := range fields {
		parts := strings.Split(field, ":")
		switch parts[0] {
		case "byr":
			passport.BirthYear = parts[1]
		case "iyr":
			passport.IssueYear = parts[1]
		case "eyr":
			passport.ExpirationYear = parts[1]
		case "hgt":
			passport.Height = parts[1]
		case "hcl":
			passport.HairColor = parts[1]
		case "ecl":
			passport.EyeColor = parts[1]
		case "pid":
			passport.PassportID = parts[1]
		case "cid":
			passport.CountryID = parts[1]
		}
	}
	return
}

func getPassports(lines []string) (passports []Passport) {
	var passportLines string
	for _, line := range lines {
		if line == "" {
			passports = append(passports, parsePassport(passportLines))
			passportLines = ""
		} else {
			passportLines += " " + line
		}
	}
	passports = append(passports, parsePassport(passportLines))
	return
}

func isCharacterDigit(char int32) bool {
	return char >= 48 && char <= 57
}

func isCharacterHexDigit(char int32) bool {
	return isCharacterDigit(char) ||
		(char >= 65 && char <= 70) ||
		(char >= 97 && char <= 102)
}

func validateYear(year string, reqLength int, min int, max int) bool {
	if len(year) != reqLength {
		return false
	}
	val, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return val >= min && val <= max
}

func validateBirthYear(birthYear string) bool {
	return validateYear(birthYear, 4, 1920, 2002)
}

func validateIssueYear(issueYear string) bool {
	return validateYear(issueYear, 4, 2010, 2020)
}

func validateExpirationYear(expirationYear string) bool {
	return validateYear(expirationYear, 4, 2020, 2030)
}

func validateHeight(height string) bool {
	if len(height) < 4 {
		return false
	}

	units := height[len(height)-2:]
	val, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		return false
	}
	if units == "cm" {
		return val >= 150 && val <= 193
	} else if units == "in" {
		return val >= 59 && val <= 76
	}
	return false
}

func validateHairColor(hairColor string) bool {
	if len(hairColor) != 7 || hairColor[0] != '#' {
		return false
	}

	for i := 1; i < len(hairColor); i++ {
		if !isCharacterHexDigit(int32(hairColor[i])) {
			return false
		}
	}
	return true
}

var validEyeColors = map[string]bool {
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}
func validateEyeColor(eyeColor string) bool {
	if len(eyeColor) != 3 {
		return false
	}

	_, found := validEyeColors[eyeColor]
	return found
}

func validatePassportID(passportID string) bool {
	if len(passportID) != 9 {
		return false
	}

	for _, digit := range passportID {
		if !isCharacterDigit(digit) {
			return false
		}
	}
	return true
}

func validatePassport(passport Passport) bool {
	return validateBirthYear(passport.BirthYear) &&
		validateIssueYear(passport.IssueYear) &&
		validateExpirationYear(passport.ExpirationYear) &&
		validateHeight(passport.Height) &&
		validateHairColor(passport.HairColor) &&
		validateEyeColor(passport.EyeColor) &&
		validatePassportID(passport.PassportID)
}

func part1(passports []Passport) (result int) {
	for _, passport := range passports {
		if passport.BirthYear != "" &&
			passport.IssueYear != "" &&
			passport.ExpirationYear != "" &&
			passport.Height != "" &&
			passport.HairColor != "" &&
			passport.EyeColor != "" &&
			passport.PassportID != "" {
			result ++
		}
	}
	return
}

func part2(passports []Passport) (result int) {
	for _, passport := range passports {
		if validatePassport(passport) {
			result++
		}
	}
	return
}

func main() {
	lines, err := getInput("day4_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	passports := getPassports(lines)

	part1Result := part1(passports)
	part2Result := part2(passports)

	fmt.Printf("Part 1 answer: %d\n", part1Result)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
