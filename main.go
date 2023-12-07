package main

import (
	"crypto/rand"
	"flag"
	"log"
	"strings"
)

var (
	seed        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	seedSymbols = "`~!@#$%^&*()_-+={}[]\\|:;\"'<>,.?/"
)

const (
	FLAG_LENGTH_MIN = 8
	FLAG_LENGTH_MAX = 50
)

const HELP = `
	command
		go run main.go [flag]

	flag
		-help
			show manual.

		-l
			specifiy the length of password.

		-no-symbols
			symbols are excluded from the seed of password creation.

		-add-chars
			includes chars to the seed of password creation.

		-remove-chars
			excludes chars from the seed of password creation.
`

func main() {
	flagHelp := flag.Bool("help", false, "show manual.")
	flagLength := flag.Int("l", FLAG_LENGTH_MIN, "specifiy the length of password.")
	flagNoSymbols := flag.Bool("no-symbols", false, "symbols are excluded from the seed of password creation.")
	flagAddChars := flag.String("add-chars", "", "includes chars to the seed of password creation.")
	flagRemoveChars := flag.String("remove-chars", "", "excludes chars from the seed of password creation.")
	flag.Parse()

	if *flagHelp {
		println(HELP)

		return
	}

	if *flagLength < FLAG_LENGTH_MIN {
		log.Fatalf("must be flag l more than or equal to %d.", FLAG_LENGTH_MIN)
	}

	if *flagLength > FLAG_LENGTH_MAX {
		log.Fatalf("must be flag n less than or equal to %d.", FLAG_LENGTH_MAX)
	}

	if !*flagNoSymbols {
		seed += seedSymbols
	}

	if *flagAddChars != "" {
		for _, rn := range *flagAddChars {
			if !strings.Contains(seed, string(rn)) {
				seed += string(rn)
			}
		}
	}

	if *flagRemoveChars != "" {
		for _, rn := range *flagRemoveChars {
			if strings.Contains(seed, string(rn)) {
				seed = strings.ReplaceAll(seed, string(rn), "")
			}
		}
	}

	if seed == "" {
		log.Fatalf("no seed")
	}

	b := make([]byte, *flagLength)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	var pw string
	for _, v := range b {
		pw += string(seed[int(v)%len(seed)])
	}

	println(pw)
}
