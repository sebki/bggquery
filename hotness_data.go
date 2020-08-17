package bggquery

import "encoding/xml"

// HotItems contains all response data from a HotQuery
type HotItems struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Rank      string `xml:"rank,attr"`
		Thumbnail struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"thumbnail"`
		Name struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"name"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
	} `xml:"item"`
}

// Write writes the response to HotItems and fulfills io.Writer interface
func (hi *HotItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, hi)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
