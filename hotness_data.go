package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

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

// Write unmarshals the response body to HotItems
func (hi *HotItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, hi)
	if err != nil {
		return err
	}
	return nil
}
