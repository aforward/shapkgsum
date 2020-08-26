package main

import (
	"testing"
)

func TestNewFingerprintOutputInvalidInputs(t *testing.T) {
	expected := fingerprint("")
	actual := newFingerprintOutput("hello world")

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestNewFingerprintOutputOk(t *testing.T) {

	intput := `
	Package "op_darwin_amd64_v1.5.0.pkg":
	Status: signed by a developer certificate issued by Apple for distribution
	Signed with a trusted timestamp on: 2020-08-18 15:47:29 +0000
	Certificate Chain:
	 1. Developer ID Installer: AgileBits Inc. (2BUA8C4S2C)
		Expires: 2024-10-23 17:10:43 +0000
		SHA256 Fingerprint:
			15 1D D8 7B 2B 23 12 11 F1 44 08 49 79 80 07 DF 62 1D E6 EB 3D AB 
			98 5B C9 64 EE 97 04 C4 A1 C1
		------------------------------------------------------------------------
	 2. Developer ID Certification Authority
		Expires: 2027-02-01 22:12:15 +0000
		SHA256 Fingerprint:
			7A FC 9D 01 A6 2F 03 A2 DE 96 37 93 6D 4A FE 68 09 0D 2D E1 8D 03 
			F2 9C 88 CF B0 B1 BA 63 58 7F
		------------------------------------------------------------------------
	 3. Apple Root CA
		Expires: 2035-02-09 21:40:36 +0000
		SHA256 Fingerprint:
			B0 B1 73 0E CB C7 FF 45 05 14 2C 49 F1 29 5E 6E DA 6B CA ED 7E 2C 
			68 C5 BE 91 B5 A1 10 01 F0 24`

	expected := fingerprint("15 1D D8 7B 2B 23 12 11 F1 44 08 49 79 80 07 DF 62 1D E6 EB 3D AB 98 5B C9 64 EE 97 04 C4 A1 C1")
	actual := newFingerprintOutput(intput)

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestNewFingerprintFromPkgutil(t *testing.T) {

	expected := fingerprint("14 1D D8 7B 2B 23 12 11 F1 44 08 49 79 80 07 DF 62 1D E6 EB 3D AB 98 5B C9 64 EE 97 04 C4 A1 C1")
	actual := newFingerprintFromPkgutil("./test/fixtures/op_darwin_amd64_v1.5.0.pkg")

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
