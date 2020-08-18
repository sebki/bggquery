package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// ForumListItems contains all response data from a ForumListsQuery
type ForumListItems struct {
	XMLName    xml.Name `xml:"forums"`
	Text       string   `xml:",chardata"`
	Type       string   `xml:"type,attr"`
	ID         string   `xml:"id,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Forum      []struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		Groupid      string `xml:"groupid,attr"`
		Title        string `xml:"title,attr"`
		Noposting    string `xml:"noposting,attr"`
		Description  string `xml:"description,attr"`
		Numthreads   string `xml:"numthreads,attr"`
		Numposts     string `xml:"numposts,attr"`
		Lastpostdate string `xml:"lastpostdate,attr"`
	} `xml:"forum"`
}

// Write unmarshals the response body to ForumListItems
func (fli *ForumListItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, fli)
	if err != nil {
		return err
	}
	return nil
}
