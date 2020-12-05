package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//s := bufio.NewScanner(os.Stdin)
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	ps := make([]map[string]string, 0)
	//r := 0
	ps = append(ps, make(map[string]string))
	psIdx := 0
	for s.Scan() {
		l := s.Text()
		if l == "" {
			psIdx++
			ps = append(ps, make(map[string]string))
			continue
		}
		kvs := strings.Split(l, " ")
		for _, kv := range kvs {
			in := strings.Split(kv, ":")
			ps[psIdx][in[0]] = in[1]
		}
	}

	resultOne := 0
	for _, p := range ps {
		if p["byr"] != "" &&
			p["iyr"] != "" &&
			p["eyr"] != "" &&
			p["hgt"] != "" &&
			p["hcl"] != "" &&
			p["ecl"] != "" &&
			p["pid"] != "" {
		}
		resultOne++
	}

	fmt.Printf("result puzzle one: %d\n", resultOne)

	resultTwo := 0
	for _, p := range ps {
		if validateByr(p["byr"]) &&
			validateIyr(p["iyr"]) &&
			validateEyr(p["eyr"]) &&
			validateHgt(p["hgt"]) &&
			validateHcl(p["hcl"]) &&
			validateEcl(p["ecl"]) &&
			validatePid(p["pid"]) {
			resultTwo++
		}
	}
	fmt.Printf("result puzzle two: %d\n", resultTwo)
}

// cid (Country ID) - ignored, missing or not.

func validateByr(s string) (r bool) {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(s)
	if err == nil && byr >= 1920 && byr <= 2002 {
		r = true
	}
	return r
}

func validateIyr(s string) (r bool) {
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(s)
	if err == nil && iyr >= 2010 && iyr <= 2020 {
		r = true
	}
	return r
}

func validateEyr(s string) (r bool) {
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(s)
	if err == nil && eyr >= 2020 && eyr <= 2030 {
		r = true
	}
	return r
}

func validateHgt(s string) bool {
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	re, err := regexp.Compile("^([0-9]+)(cm|in)$")
	if err != nil {
		return false
	}
	r := re.FindStringSubmatch(s)
	if len(r) != 3 {
		return false
	}
	if r[2] == "cm" {
		num, err := strconv.Atoi(r[1])
		if err == nil && num >= 150 && num <= 193 {
			return true
		}
	}
	if r[2] == "in" {
		num, err := strconv.Atoi(r[1])
		if err == nil && num >= 59 && num <= 76 {
			return true
		}
	}

	return false
}

func validateHcl(s string) bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	re, err := regexp.Compile("^#[0-9a-z]{6}$")
	if err != nil {
		return false
	}
	r := re.MatchString(s)
	if r {
		return true
	}

	return false
}

func validateEcl(s string) bool {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if s == "amb" ||
		s == "blu" ||
		s == "brn" ||
		s == "gry" ||
		s == "grn" ||
		s == "hzl" ||
		s == "oth" {
		return true
	}
	return false
}

func validatePid(s string) bool {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	re, err := regexp.Compile("^[0-9]{9}$")
	if err != nil {
		return false
	}
	r := re.MatchString(s)
	if r {
		return true
	}

	return false
}
