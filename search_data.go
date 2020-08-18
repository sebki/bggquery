package bggquery

import "encoding/xml"

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

// Write writes response body to SearchItems and fulfills io.Writer interface
func (si *SearchItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, si)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
