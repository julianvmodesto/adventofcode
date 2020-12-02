package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(line string) (a, b int, letter, password string, err error) {

	// example:
	// 7-9 l: vslmtglbc

	parts1 := strings.Split(line, " ")
	if len(parts1) != 3 {
		err = errors.New("expected 3 parts")
		return
	}

	parts2 := strings.Split(parts1[0], "-")
	if len(parts2) != 2 {
		err = errors.New("expected 2 parts")
		return
	}
	a64, err := strconv.ParseInt(parts2[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing min %q: %v", parts2[0], err)
		return
	}
	a = int(a64)
	b64, err := strconv.ParseInt(parts2[1], 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing max %q: %v", parts2[1], err)
		return
	}
	b = int(b64)

	if len(parts1[1]) != 2 {
		err = fmt.Errorf("error parsing letter %q", parts1[1])
		return
	}
	letter = parts1[1][:1]

	password = parts1[2]

	return
}

func valid1(min, max int, letter, password string) (bool, int) {
	var c int
	for _, r := range password {
		if strings.Compare(letter, string(r)) == 0 {
			c++
		}
	}
	return min <= c && c <= max, c
}

func main() {
	if len(os.Args) < 2 {
		log.Println("missing input file, provide filename")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read file %s", os.Args[1])
	}
	var lines = strings.Split(string(data), "\n")

	var p1 int
	var p2 int
	var e int
	var e2 int
	for i, line := range lines {
		if line == "" {
			continue
		}
		a, b, letter, password, err := parse(line)
		if err != nil {
			log.Printf("error parsing line %d %q: %v", i, line, err)
			e++
			continue
		}
		if v, _ := valid1(a, b, letter, password); v {
			p1++
		} else {
			// log.Printf("INVALID: line %d %q: min %d, max %d, count %d, letter %q, password %q", i, line, min, max, c, letter, password)
		}
		if a < 1 || a > len(password) || b < 1 || b > len(password) {
			e2++
		} else if (strings.Compare(string(password[a-1]), letter) == 0) != (strings.Compare(string(password[b-1]), letter) == 0) {
			// XOR
			p2++
		}
		// log.Printf("INVALID: line %d %q: a %d, b %d, letter %q, password %q, password[a] %q, password[b] %q", i, line, a, b, letter, password, string(password[a-1]), string(password[b-1]))
	}
	log.Printf("valid passwords: %d of %d\n", p1, len(lines)-1)
	log.Printf("valid passwords: %d of %d\n", p2, len(lines)-1)
	log.Printf("parse errors: %d of %d\n", e, len(lines)-1)
	log.Printf("rule errors: %d of %d\n", e2, len(lines)-1)
}
