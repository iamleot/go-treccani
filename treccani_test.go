package treccani

import (
	"regexp"
	"testing"
)

// TestLookupTerm calls treccani.TestLookup checking that it gets a definition.
func TestLookupTerm(t *testing.T) {
	term := "esempio"
	want := regexp.MustCompile(`^\b` + term + `\b`)
	definition := LookupTerm("esempio")
	if !want.MatchString(definition) {
		t.Fatalf(`LookupTerm("esempio") = %q want match for %#q`, definition, want)
	}
}

// TestLookupInexistentTerm calls treccani.TestLookup checking that it does not
// get a definition for an inexistent term.
func TestLookupInexistentTerm(t *testing.T) {
	definition := LookupTerm("inexistent")
	if definition != "" {
		t.Fatalf(`LookupTerm("inexistent") = %q want ""`, definition)
	}
}

// TestTerms calls treccani.Terms checking that one or more definitions
// are returned.
func TestTerms(t *testing.T) {
	term := "birba"
	definitions := Terms(term)
	if len(definitions) == 0 {
		t.Fatalf(`Terms("birba") should return some definitions`)
	}
}

// TestTermsMultiple calls treccani.Terms checking that more than one
// definition is returned.
func TestTermsMultiple(t *testing.T) {
	term := "lira"
	definitions := Terms(term)
	if len(definitions) <= 1 {
		t.Fatalf(`Terms("lira") should return several definitions`)
	}
}

// TestTermsSingle calls treccani.Terms checking that exactly a single
// definition is returned.
func TestTermsSingle(t *testing.T) {
	term := "ciao"
	definitions := Terms(term)
	if len(definitions) != 1 {
		t.Fatalf(`Terms("ciao") should return a single definition`)
	}
}

// TestInexistentTerms calls treccani.Terms checking that for inexistent terms
// no definitions are returned.
func TestInexistentTerms(t *testing.T) {
	term := "inexistent"
	definitions := Terms(term)
	if len(definitions) != 0 {
		t.Fatalf(`Terms("inexistent") should return no definitions`)
	}
}
