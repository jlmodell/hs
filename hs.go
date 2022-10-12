package main

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strings"
)

func main() {
	documentFlag := flag.String("document", "", "Document type to clean")
	flag.Parse()

	if *documentFlag == "855" {
		CLEAN_eightfivefive()
	} else if *documentFlag == "856" {
		CLEAN_eightfivesix()
	} else {
		panic("Invalid document type")
	}

}

const (
	INPUT_855_FILE  = "/mnt/evision_out/m2k_855O.app"
	HS_855_FILE     = "/mnt/evision_out/m2k_855O_hs.app"
	OUTPUT_855_FILE = "/mnt/evision_out/m2k_855Oa.app"

	INPUT_856_FILE  = "/mnt/evision_out/m2k_856O.app"
	HS_856_FILE     = "/mnt/evision_out/m2k_856O_hs.app"
	OUTPUT_856_FILE = "/mnt/evision_out/m2k_856Oa.app"
)

func CLEAN_eightfivefive() {
	_, err := os.Stat(OUTPUT_855_FILE)
	if err != nil {
		_, err := os.Stat(INPUT_855_FILE)
		if err != nil {
			panic("Input file not found")
		} else {
			file, err := os.OpenFile(INPUT_855_FILE, os.O_RDONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			hs := []string{}
			out := []string{}

			move := false

			scanner := bufio.NewScanner(file)

			regexHS := regexp.MustCompile("^\"H\"~\"HENRY")
			regexNonHSorOWENS := regexp.MustCompile("^\"H\"~\"[A-G,I-Z]")

			for scanner.Scan() {
				line := scanner.Text()

				if regexHS.MatchString(line) {
					move = true
				} else if regexNonHSorOWENS.MatchString(line) {
					move = false
				}

				if move {
					hs = append(hs, line)
				} else {
					out = append(out, line)
				}
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}

			file2, err := os.OpenFile(OUTPUT_855_FILE, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file2.Close()
			file2.WriteString(strings.Join(out, "\n"))

			file3, err := os.OpenFile(HS_855_FILE, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file3.Close()
			file3.WriteString(strings.Join(hs, "\n"))

		}
	}

	if err := os.Remove(INPUT_855_FILE); err != nil {
		panic(err)
	}
}

func CLEAN_eightfivesix() {
	_, err := os.Stat(OUTPUT_856_FILE)
	if err != nil {
		_, err := os.Stat(INPUT_856_FILE)
		if err != nil {
			panic("Input file not found")
		} else {
			file, err := os.OpenFile(INPUT_856_FILE, os.O_RDONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			hs := []string{}
			out := []string{}

			move := false

			scanner := bufio.NewScanner(file)

			regexHS := regexp.MustCompile("^\"H\"~\"HENRY")
			regexNonHSorOWENS := regexp.MustCompile("^\"H\"~\"[A-G,I-Z]")

			for scanner.Scan() {
				line := scanner.Text()

				if regexHS.MatchString(line) {
					move = true
				} else if regexNonHSorOWENS.MatchString(line) {
					move = false
				}

				if move {
					hs = append(hs, line)
				} else {
					out = append(out, line)
				}
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}

			file2, err := os.OpenFile(OUTPUT_856_FILE, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file2.Close()
			file2.WriteString(strings.Join(out, "\n"))

			file3, err := os.OpenFile(HS_856_FILE, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file3.Close()
			file3.WriteString(strings.Join(hs, "\n"))

		}
	}

	if err := os.Remove(INPUT_856_FILE); err != nil {
		panic(err)
	}
}
