package puzzles

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]*PassVal
}

type PassVal struct {
	value string
	valid func(string) bool
}

func NewPassport() *Passport {
	fields := map[string]*PassVal{
		"byr": &PassVal{value: "", valid: isValidBYR},
		"iyr": &PassVal{value: "", valid: isValidIYR},
		"eyr": &PassVal{value: "", valid: isValidEYR},
		"hgt": &PassVal{value: "", valid: isValidHGT},
		"hcl": &PassVal{value: "", valid: isValidHCL},
		"ecl": &PassVal{value: "", valid: isValidECL},
		"pid": &PassVal{value: "", valid: isValidPID},
		"cid": &PassVal{value: "", valid: func(s string) bool { return true }},
	}
	p := Passport{fields: fields}
	return &p
}

func (p *Passport) AddField(kv string) {
	values := strings.Split(kv, ":")
	p.fields[values[0]].value = values[1]
}

func (p *Passport) IsValid() bool {
	for k, pv := range p.fields {
		if k == "cid" {
			continue
		}
		if !pv.valid(pv.value) {
			return false
		}
	}
	return true
}

func isValidBYR(byr string) bool {
	// byr four digits; at least 1920 and at most 2002.
	return isStringInRange(1920, 2002, byr)
}

func isValidIYR(iyr string) bool {
	return isStringInRange(2010, 2020, iyr)
}

func isValidEYR(eyr string) bool {
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	return isStringInRange(2020, 2030, eyr)
}

func isValidHGT(hgt string) bool {
	/*
		hgt (Height) - a number followed by either cm or in:
		If cm, the number must be at least 150 and at most 193.
		If in, the number must be at least 59 and at most 76.
	*/
	if len(hgt) < 3 {
		return false
	}
	height := hgt[0 : len(hgt)-2]
	units := hgt[len(hgt)-2 : len(hgt)]

	switch units {
	case "cm":
		if !isStringInRange(150, 193, height) {
			return false
		}
		break
	case "in":
		if !isStringInRange(59, 76, height) {
			return false
		}
		break
	default:
		return false
	}

	return true
}

func isValidHCL(hcl string) bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	return isMatch(`^#([0-9a-f]){6}$`, hcl)
}

func isValidECL(ecl string) bool {
	return isMatch("^(amb|blu|brn|gry|grn|hzl|oth)$", ecl)
}

func isValidPID(pid string) bool {
	return isMatch(`^([0-9]){9}$`, pid)
}

func (p *Passport) prettyPrint() {
	fmt.Println("******** PASSPORT ********")
	for k, pv := range p.fields {
		fmt.Printf("%s: %s\n", k, pv.value)
	}
	fmt.Println("***** END PASSPORT ********")
}

func isStringInRange(min int, max int, numb string) bool {
	if numb == "" {
		return false
	}
	n, _ := strconv.Atoi(numb)
	return n >= min && n <= max
}

func isMatch(re string, s string) bool {
	matched, err := regexp.MatchString(re, s)
	if err != nil || matched == false {
		return false
	}
	return true
}
