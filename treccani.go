// Search term definitions on Vocabolario Treccani.
package treccani

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// LookupTerm search a single term on Vocabolario Treccani. If no definition is
// found an empty string is returned, otherwise the definition is returned as
// text.
func LookupTerm(term string, client *http.Client) string {
	resp, err := client.Get(fmt.Sprintf("https://www.treccani.it/vocabolario/%s/",
		url.PathEscape(term)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Delete inline style that would only add noise
	doc.Find(".term-content").Find("style").Remove()

	return strings.ReplaceAll(strings.ReplaceAll(
		strings.TrimSpace(
			doc.Find(".term-content").Text()),
		"\n", " "),
		"   ", " ")
}

// Terms search all terms in Vocabolario Treccani. Returns a string slice of
// term definitions.
func Terms(term string, client *http.Client) []string {
	var terms []string

	if t := LookupTerm(term, client); t != "" {
		terms = append(terms, t)
	} else {
		for i := 1; ; i++ {
			t = LookupTerm(fmt.Sprintf("%s%d", term, i), client)
			if t != "" {
				terms = append(terms, t)
			} else {
				break
			}
		}
	}

	return terms
}
