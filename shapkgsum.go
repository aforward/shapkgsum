package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type fingerprint string

func pkgutil(filename string) string {
	out, err := exec.Command("pkgutil", "--check-signature", filename).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return string(out[:])
}

func newFingerprintFromPkgutil(filename string) fingerprint {
	output := pkgutil(filename)
	return newFingerprintOutput(output)
}

func newFingerprintOutput(output string) fingerprint {
	isFound := false
	fingerprints := []string{}
	for _, v := range strings.Split(output, "\n") {
		firstMatch, _ := regexp.MatchString(`SHA256`, v)
		lastMatch, _ := regexp.MatchString(`-------`, v)

		if lastMatch {
			return fingerprint(strings.Join(fingerprints, " "))
		} else if isFound {
			fingerprints = append(fingerprints, strings.TrimSpace(v))
		} else if firstMatch {
			isFound = true
		}
	}

	return ""
}
