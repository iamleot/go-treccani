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

const userAgent = "Mozilla/5.0 (Windows NT 6.1; rv:60.0) Gecko/20100101 Firefox/140.0"

// LookupTerm search a single term on Vocabolario Treccani. If no definition is
// found an empty string is returned, otherwise the definition is returned as
// text.
func LookupTerm(term string, client *http.Client) string {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.treccani.it/vocabolario/%s/",
			url.PathEscape(term)), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
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

	// Delete initial header
	doc.Find(`span:contains('DAL VOCABOLARIO')`).Remove()

	// Delete copyright note that would only add noise
	doc.Find(`p:contains(` +
		`'©  Istituto della Enciclopedia Italiana ` +
		`fondata da Giovanni Treccani - ` +
		`Riproduzione riservata')`).Remove()

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
