// Copyright 2016-2021, Pulumi Corporation.

package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSdkCfnSdkRoundtrip(t *testing.T) {
	cases := []string{
		"ipAddress",
		"anIpAddress",
		"aNewIpAddress",
		"ip",
		"ascii",
		"asciiAndMore",
		"anAsciiAndMore",
		"anIp",
		"a",
		"lowercase",
		"lowercaseA",
		"lowercaseACat",
		"aCatAndADog",
		"someArns",
		"ec2ManagedKey",
		"useEc2Please",
		"http2",
	}

	for _, c := range cases {
		assert.Equal(t, c, ToSdkName(ToCfnName(c, map[string]string{})))
	}
}

func TestCfnSdkCfnRoundtrip(t *testing.T) {
	casesSimple := []string{
		"A",
		"UppercaseACat",
		"ACatAndADog",
		"Ip",
		"AnIp",
		"IpAddress",
		"AnIpAddress",
		"ANewIpAddress",
		"Ascii",
		"AsciiAndMore",
		"SomeArns",
	}

	for _, c := range casesSimple {
		assert.Equal(t, c, ToCfnName(ToSdkName(c), map[string]string{}))
	}

	// Due to the presence of uppercase acronyms, these can only be round-tripped with a table
	casesLookup := []string{
		"IP",
		"AnIP",
		"IPAddress",
		"AnIPAddress",
		"ANewIPAddress",
		"ASCII",
		"ASCIIAndMore",
		"AnASCIIAndMore",
		"SomeARNs",
		"UseEC2Please",
		"EC2ManagedKey",
	}

	lookup := map[string]string{
		"ip":             "IP",
		"anIp":           "AnIP",
		"ipAddress":      "IPAddress",
		"anIpAddress":    "AnIPAddress",
		"aNewIpAddress":  "ANewIPAddress",
		"ascii":          "ASCII",
		"asciiAndMore":   "ASCIIAndMore",
		"anAsciiAndMore": "AnASCIIAndMore",
		"someArns":       "SomeARNs",
		"useEc2Please":   "UseEC2Please",
		"ec2ManagedKey":  "EC2ManagedKey",
	}

	for _, c := range casesLookup {
		assert.Equal(t, c, ToCfnName(ToSdkName(c), lookup))
	}
}

func TestHasUppercaseAcronym(t *testing.T) {
	casesTrue := []string{
		"IPAddress",
		"anIPAddress",
		"AnIPAddress",
		"ANewIPAddress",
		"IP",
		"IPI",
		"ASCIIA",
		"ASCIIAndMore",
		"anASCIIAndMore",
		"anIP",
		"UseS3Please",
		"HostIPs",
		"ThreeIPsPlease",
	}

	casesFalse := []string{
		"",
		"A",
		"lowercase",
		"lowercaseA",
		"lowercaseACat",
		"UppercaseACat",
		"ACatAndADog",
		"LessIsMore",
		"IAspire",
	}

	for _, c := range casesTrue {
		//testing.Logf(t, "text: %s, start: %d, end %d", c, s, e)
		assert.Truef(t, HasUppercaseAcronym(c), "Failed to detect uppercase acronym in: %s", c)
	}

	for _, c := range casesFalse {
		assert.Falsef(t, HasUppercaseAcronym(c), "Detected spurrious uppercase acronym in: %s", c)
	}
}

func TestFindFirstRunOfUppercase(t *testing.T) {
	cases := []struct {
		text          string
		minLength     int
		expectedStart int
		expectedEnd   int
	}{
		{
			text:          "IPAddress",
			minLength:     2,
			expectedStart: 0,
			expectedEnd:   3,
		},
		{
			text:          "IPAddress",
			minLength:     3,
			expectedStart: 0,
			expectedEnd:   3,
		},
		{
			text:          "IPAddress",
			minLength:     4,
			expectedStart: -1,
			expectedEnd:   -1,
		},
		{
			text:          "anIPAddress",
			minLength:     3,
			expectedStart: 2,
			expectedEnd:   5,
		},
		{
			text:          "anIPAddress",
			minLength:     4,
			expectedStart: -1,
			expectedEnd:   -1,
		},
		{
			text:          "ANewIPAddress",
			minLength:     2,
			expectedStart: 0,
			expectedEnd:   2,
		},
		{
			text:          "ANewIPAddress",
			minLength:     3,
			expectedStart: 4,
			expectedEnd:   7,
		},
		{
			text:          "AnEC2Instance",
			minLength:     3,
			expectedStart: 2,
			expectedEnd:   6,
		},
	}

	for _, c := range cases {
		s, e := findFirstRunOfUppercase(c.text, c.minLength)
		assert.Equalf(t, c.expectedStart, s, "'%s' Start", c.text)
		assert.Equalf(t, c.expectedEnd, e, "'%s' End", c.text)
	}
}
