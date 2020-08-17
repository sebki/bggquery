package bggquery

import "encoding/xml"

// ThreadItems contains all response data from a ThreadQuery
type ThreadItems struct {
	XMLName     xml.Name `xml:"thread"`
	Text        string   `xml:",chardata"`
	ID          string   `xml:"id,attr"`
	Numarticles string   `xml:"numarticles,attr"`
	Link        string   `xml:"link,attr"`
	Termsofuse  string   `xml:"termsofuse,attr"`
	Subject     string   `xml:"subject"`
	Articles    struct {
		Text    string `xml:",chardata"`
		Article []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Username string `xml:"username,attr"`
			Link     string `xml:"link,attr"`
			Postdate string `xml:"postdate,attr"`
			Editdate string `xml:"editdate,attr"`
			Numedits string `xml:"numedits,attr"`
			Subject  string `xml:"subject"`
			Body     string `xml:"body"`
		} `xml:"article"`
	} `xml:"articles"`
}

// Write writes response body to ThreadItems
func (ti *ThreadItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, ti)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
