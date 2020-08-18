package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// SearchItems contains all response data from a search query
type SearchItems struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Total      string   `xml:"total,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
		Name struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Value string `xml:"value,attr"`
		} `xml:"name"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
	} `xml:"item"`
}

// Write unmarshals the response body to SearchItems
func (si *SearchItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, si)
	if err != nil {
		return err
	}
	return nil
}
