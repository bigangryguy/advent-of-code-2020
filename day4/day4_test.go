package main

import (
	"fmt"
	"testing"
)

func Test_isCharacterDigit(t *testing.T) {
	var actual bool

	good := "0123456789"
	for _, char := range good {
		actual = isCharacterDigit(char)
		if !actual {
			t.Errorf("isCharacterDigit = %v for %v, expected true", actual, char)
		}
	}

	bad := "+,-./:;<=abcdef "
	for _, char := range bad {
		actual = isCharacterDigit(char)
		if actual {
			t.Errorf("isCharacterDigit = %v for %v, expected false", actual, char)
		}
	}
}

func Test_isCharacterHexDigit(t *testing.T) {
	var actual bool

	good := "0123456789abcdefABCDEF"
	for _, char := range good {
		actual = isCharacterHexDigit(char)
		if !actual {
			t.Errorf("isCharacterHexDigit = %v for %v, expected true", actual, char)
		}
	}

	bad := "+,-./:;<=?@GH_`gh "
	for _, char := range bad {
		actual = isCharacterHexDigit(char)
		if actual {
			t.Errorf("isCharacterHexDigit = %v for %v, expected false", actual, char)
		}
	}
}

func Test_validateYear(t *testing.T) {
	var actual bool

	good := []string { "1998", "2002", "2012", "2020" }
	for _, year := range good {
		actual = validateYear(year, 4, 1998, 2020)
		if !actual {
			t.Errorf("validateYear = %v for %v, expected true", actual, year)
		}
	}

	bad := []string { "477", "1997", "2021", "32002", "" }
	for _, year := range bad {
		actual = validateYear(year, 4, 1998, 2020)
		if actual {
			t.Errorf("validateYear = %v for %v, expected false", actual, year)
		}
	}
}

func Test_validateBirthYear(t *testing.T) {
	var actual bool

	good := []string { "1920", "1944", "1982", "2002" }
	for _, year := range good {
		actual = validateYear(year, 4, 1920, 2002)
		if !actual {
			t.Errorf("validateBirthYear = %v for %v, expected true", actual, year)
		}
	}

	bad := []string { "477", "1919", "2003", "32002", "" }
	for _, year := range bad {
		actual = validateYear(year, 4, 1920, 2002)
		if actual {
			t.Errorf("validateBirthYear = %v for %v, expected false", actual, year)
		}
	}
}

func Test_validateIssueYear(t *testing.T) {
	var actual bool

	good := []string { "2010", "2012", "2017", "2020" }
	for _, year := range good {
		actual = validateYear(year, 4, 2010, 2020)
		if !actual {
			t.Errorf("validateIssueYear = %v for %v, expected true", actual, year)
		}
	}

	bad := []string { "477", "2009", "2021", "32002", "" }
	for _, year := range bad {
		actual = validateYear(year, 4, 2010, 2020)
		if actual {
			t.Errorf("validateIssueYear = %v for %v, expected false", actual, year)
		}
	}
}

func Test_validateExpirationYear(t *testing.T) {
	var actual bool

	good := []string { "2020", "2023", "2027", "2030" }
	for _, year := range good {
		actual = validateYear(year, 4, 2020, 2030)
		if !actual {
			t.Errorf("validateExpirationYear = %v for %v, expected true", actual, year)
		}
	}

	bad := []string { "477", "2019", "2031", "32002", "" }
	for _, year := range bad {
		actual = validateYear(year, 4, 2020, 2030)
		if actual {
			t.Errorf("validateExpirationYear = %v for %v, expected false", actual, year)
		}
	}
}

func Test_validateHeight(t *testing.T) {
	var actual bool

	good := []string { "150cm", "193cm", "162cm", "185cm", "59in", "76in", "62in", "72in" }
	for _, height := range good {
		actual = validateHeight(height)
		if !actual {
			t.Errorf("validateHeight = %v for %v, expected true", actual, height)
		}
	}

	bad := []string { "150", "193", "59", "76", "149cm", "194cm", "58in", "77in", "cm", "in", "1500", "15", "1", "" }
	for _, height := range bad {
		actual = validateHeight(height)
		if actual {
			t.Errorf("validateHeight = %v for %v, expected false", actual, height)
		}
	}
}

func Test_validateHairColor(t *testing.T) {
	var actual bool

	good := []string { "#123456", "#012345", "#0abcde", "#abcdef", "#a0b1c2" }
	for _, hairColor := range good {
		actual = validateHairColor(hairColor)
		if !actual {
			t.Errorf("validateHairColor = %v for %v, expected true", actual, hairColor)
		}
	}

	bad := []string { "123456", "012345", "abcdef", "#123", "#abc", "color", "#color", "#123456a", "" }
	for _, hairColor := range bad {
		actual = validateHairColor(hairColor)
		if actual {
			t.Errorf("validateHairColor = %v for %v, expected false", actual, hairColor)
		}
	}
}

func Test_validateEyeColor(t *testing.T) {
	var actual bool

	good := []string { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" }
	for _, eyeColor := range good {
		actual = validateEyeColor(eyeColor)
		if !actual {
			t.Errorf("validateEyeColor = %v for %v, expected true", actual, eyeColor)
		}
	}

	bad := []string { "amber", "blue", "br", "" }
	for _, eyeColor := range bad {
		actual = validateEyeColor(eyeColor)
		if actual {
			t.Errorf("validateEyeColor = %v for %v, expected false", actual, eyeColor)
		}
	}
}

func Test_validatePassportID(t *testing.T) {
	var actual bool

	good := []string { "123456789", "012345678", "000000000" }
	for _, passportID := range good {
		actual = validatePassportID(passportID)
		if !actual {
			t.Errorf("validatePassportID = %v for %v, expected true", actual, passportID)
		}
	}

	bad := []string { "12345", "1234567890", "123456abc", "" }
	for _, passportID := range bad {
		actual = validatePassportID(passportID)
		if actual {
			t.Errorf("validatePassportID = %v for %v, expected false", actual, passportID)
		}
	}
}

func Test_validatePassport(t *testing.T) {
	var actual bool

	good := []Passport {
		{
			BirthYear: "1980",
			IssueYear: "2012",
			ExpirationYear: "2030",
			Height: "74in",
			HairColor: "#623a2f",
			EyeColor: "grn",
			PassportID: "087499704",
		},
		{
			BirthYear: "1989",
			IssueYear: "2014",
			ExpirationYear: "2029",
			Height: "165cm",
			HairColor: "#a97842",
			EyeColor: "blu",
			PassportID: "896056539",
			CountryID: "129",
		},
	}
	for _, passport := range good {
		actual = validatePassport(passport)
		if !actual {
			t.Errorf("validatePassport = %v for %v, expected true", actual, passport)
		}
	}

	bad := []Passport {
		{
			BirthYear: "1926",
			IssueYear: "2018",
			ExpirationYear: "1972",
			Height: "170",
			HairColor: "#18171d",
			EyeColor: "amb",
			PassportID: "186cm",
			CountryID: "100",
		},
		{
			BirthYear: "1946",
			IssueYear: "2019",
			ExpirationYear: "1967",
			Height: "170cm",
			HairColor: "#602927",
			EyeColor: "grn",
			PassportID: "012533040",
		},
	}
	for _, passport := range bad {
		actual = validatePassport(passport)
		if actual {
			t.Errorf("validatePassport = %v for %v, expected false", actual, passport)
		}
	}
}

func arePassportsSame(p1 Passport, p2 Passport) bool {
	return p1.BirthYear == p2.BirthYear &&
		p1.IssueYear == p2.IssueYear &&
		p1.ExpirationYear == p2.ExpirationYear &&
		p1.Height == p2.Height &&
		p1.HairColor == p2.HairColor &&
		p1.EyeColor == p2.EyeColor &&
		p1.PassportID == p2.PassportID &&
		p1.CountryID == p2.CountryID
}

func Test_parsePassport(t *testing.T) {
	var actual Passport

	good := []string {
		"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
	expected := []Passport {
		{
			BirthYear: "2001",
			IssueYear: "2015",
			ExpirationYear: "2022",
			Height: "164cm",
			HairColor: "#888785",
			EyeColor: "hzl",
			PassportID: "545766238",
			CountryID: "88",
		},
		{
			BirthYear: "1944",
			IssueYear: "2010",
			ExpirationYear: "2021",
			Height: "158cm",
			HairColor: "#b6652a",
			EyeColor: "blu",
			PassportID: "093154719",
		},
	}
	for i, text := range good {
		actual = parsePassport(text)
		if !arePassportsSame(actual, expected[i]) {
			t.Errorf("parsePassport = %v for %v, expected %v", actual, text, expected[i])
		}
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day4_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	passports := getPassports(lines)

	actual := part1(passports)
	expected := 2
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	lines, err := getInput("day4_test_input2.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	passports := getPassports(lines)

	actual := part2(passports)
	expected := 4
	if actual != expected {
		t.Errorf("part2 = %d, expected %d", actual, expected)
	}
}
