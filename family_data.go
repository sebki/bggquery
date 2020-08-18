package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// FamilyItems contains bgg response from a family query
type FamilyItems struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description string `xml:"description"`
		Link        []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			ID      string `xml:"id,attr"`
			Value   string `xml:"value,attr"`
			Inbound string `xml:"inbound,attr"`
		} `xml:"link"`
	} `xml:"item"`
}

// Write unmarshals the response body to FamilyItems
func (fi *FamilyItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, fi)
	if err != nil {
		return err
	}
	return nil
}
