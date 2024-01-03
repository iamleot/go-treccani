package treccani

import (
	"regexp"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

// TestLookupTerm calls treccani.TestLookup checking that it gets a definition.
func TestLookupTerm(t *testing.T) {
	r, err := recorder.New("fixtures/esempio")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	term := "esempio"
	want := regexp.MustCompile(`^\b` + term + `\b`)
	definition := LookupTerm("esempio", client)
	if !want.MatchString(definition) {
		t.Fatalf(`LookupTerm("esempio") = %q want match for %#q`, definition, want)
	}
}

// TestLookupInexistentTerm calls treccani.TestLookup checking that it does not
// get a definition for an inexistent term.
func TestLookupInexistentTerm(t *testing.T) {
	r, err := recorder.New("fixtures/inexistent")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	definition := LookupTerm("inexistent", client)
	if definition != "" {
		t.Fatalf(`LookupTerm("inexistent") = %q want ""`, definition)
	}
}

// TestTerms calls treccani.Terms checking that one or more definitions
// are returned.
func TestTerms(t *testing.T) {
	r, err := recorder.New("fixtures/birba")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	term := "birba"
	definitions := Terms(term, client)
	if len(definitions) == 0 {
		t.Fatalf(`Terms("birba") should return some definitions`)
	}
}

// TestTermsMultiple calls treccani.Terms checking that more than one
// definition is returned.
func TestTermsMultiple(t *testing.T) {
	r, err := recorder.New("fixtures/lira")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	term := "lira"
	definitions := Terms(term, client)
	if len(definitions) <= 1 {
		t.Fatalf(`Terms("lira") should return several definitions`)
	}
}

// TestTermsSingle calls treccani.Terms checking that exactly a single
// definition is returned.
func TestTermsSingle(t *testing.T) {
	r, err := recorder.New("fixtures/ciao")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	term := "ciao"
	definitions := Terms(term, client)
	if len(definitions) != 1 {
		t.Fatalf(`Terms("ciao") should return a single definition`)
	}
}

// TestInexistentTerms calls treccani.Terms checking that for inexistent terms
// no definitions are returned.
func TestInexistentTerms(t *testing.T) {
	r, err := recorder.New("fixtures/inexistents")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := r.GetDefaultClient()

	term := "inexistents"
	definitions := Terms(term, client)
	if len(definitions) != 0 {
		t.Fatalf(`Terms("inexistent") should return no definitions`)
	}
}
