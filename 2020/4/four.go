package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

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

	var valid1, valid2 int
	var total int
	var byr, iyr, eyr, hgt, hcl, ecl, pid bool
	var byr2, iyr2, eyr2, hgt2, hcl2, ecl2, pid2 bool
	for _, line := range lines {
		if len(line) == 0 {
			total++
			if byr && iyr && eyr && hgt && hcl && ecl && pid {
				valid1++
			}
			if byr2 && iyr2 && eyr2 && hgt2 && hcl2 && ecl2 && pid2 {
				valid2++
			}
			byr, iyr, eyr, hgt, hcl, ecl, pid = false, false, false, false, false, false, false
			byr2, iyr2, eyr2, hgt2, hcl2, ecl2, pid2 = false, false, false, false, false, false, false
			continue
		}
		kvs := strings.Split(line, " ")
		for _, s := range kvs {
			kv := strings.Split(s, ":")
			if len(kv) != 2 {
				log.Printf("invalid parts in key-value %q in line %q", s, line)
				continue
			}
			key := kv[0]
			value := kv[1]

			var err error
			switch key {
			case "byr":
				byr = true
				byr2, err = regexp.Match(`^19[2-9]\d|200[0-2]$`, []byte(value))
				if err != nil {
					log.Printf("error matching byr: %v\n", err)
				}
			case "iyr":
				iyr = true
				iyr2, err = regexp.Match(`^20(1\d|20)$`, []byte(value))
				if err != nil {
					log.Printf("error matching iyr: %v\n", err)
				}
			case "eyr":
				eyr = true
				eyr2, err = regexp.Match(`^20(2\d|30)$`, []byte(value))
				if err != nil {
					log.Printf("error matching eyr: %v\n", err)
				}
			case "hgt":
				hgt = true
				hgt2, err = regexp.Match(`^1([5-8]\d|9[0-3])cm|(59|6\d|7[0-6])in$`, []byte(value))
				if err != nil {
					log.Printf("error matching hgt: %v\n", err)
				}
			case "hcl":
				hcl = true
				hcl2, err = regexp.Match(`^#[0-9a-f]{6}$`, []byte(value))
				if err != nil {
					log.Printf("error matching hcl: %v\n", err)
				}
			case "ecl":
				ecl = true
				ecl2, err = regexp.Match(`^amb|blu|brn|gry|grn|hzl|oth$`, []byte(value))
				if err != nil {
					log.Printf("error matching ecl: %v\n", err)
				}
			case "pid":
				pid = true
				pid2, err = regexp.Match(`^\d{9}$`, []byte(value))
				if err != nil {
					log.Printf("error matching ecl: %v\n", err)
				}
			case "cid":
			default:
				log.Printf("unknown key %q in line %q\n", key, line)
			}
		}
	}

	log.Printf("part 1: valid passports %d of %d\n", valid1, total)
	log.Printf("part 2: valid passports %d of %d\n", valid2, total)
}
