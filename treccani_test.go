package treccani

import (
	"regexp"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func TestLookupTerm(t *testing.T) {
	type testCase struct {
		term           string
		expectedRegexp *regexp.Regexp
		cassette       string
	}

	var tests = map[string]testCase{
		"term with a definition": {
			term:           "esempio",
			expectedRegexp: regexp.MustCompile(`^\besempio\b`),
			cassette:       "fixtures/esempio",
		},
		"term without a definition": {
			term:           "inexistent",
			expectedRegexp: regexp.MustCompile(`^$`),
			cassette:       "fixtures/inexistent",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, err := recorder.New(tc.cassette)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Stop()
			client := r.GetDefaultClient()

			definition := LookupTerm(tc.term, client)
			if !tc.expectedRegexp.MatchString(definition) {
				t.Fatalf(`LookupTerm("%s", ...) = %q want match for %#q`, tc.term, definition, tc.expectedRegexp)
			}
		})
	}
}

func TestTerms(t *testing.T) {
	type testCase struct {
		term                        string
		expectedNumberOfDefinitions int
		cassette                    string
	}

	var tests = map[string]testCase{
		"term with 2 definitions": {
			term:                        "birba",
			expectedNumberOfDefinitions: 2,
			cassette:                    "fixtures/birba",
		},
		"term with 1 definition": {
			term:                        "ciao",
			expectedNumberOfDefinitions: 1,
			cassette:                    "fixtures/ciao",
		},
		"term without a definition": {
			term:                        "inexistents",
			expectedNumberOfDefinitions: 0,
			cassette:                    "fixtures/inexistents",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, err := recorder.New(tc.cassette)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Stop()
			client := r.GetDefaultClient()

			definitions := Terms(tc.term, client)
			if len(definitions) != tc.expectedNumberOfDefinitions {
				t.Fatalf(`Terms("%s", ...) should return %d definitions but returned %d definitions`, tc.term, tc.expectedNumberOfDefinitions, len(definitions))
			}
		})
	}
}
