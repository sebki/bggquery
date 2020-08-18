package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// ForumItems contains all response data from a ForumsQuery
type ForumItems struct {
	XMLName      xml.Name `xml:"forum"`
	Text         string   `xml:",chardata"`
	ID           string   `xml:"id,attr"`
	Title        string   `xml:"title,attr"`
	Numthreads   string   `xml:"numthreads,attr"`
	Numposts     string   `xml:"numposts,attr"`
	Lastpostdate string   `xml:"lastpostdate,attr"`
	Noposting    string   `xml:"noposting,attr"`
	Termsofuse   string   `xml:"termsofuse,attr"`
	Threads      struct {
		Text   string `xml:",chardata"`
		Thread []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr"`
			Subject      string `xml:"subject,attr"`
			Author       string `xml:"author,attr"`
			Numarticles  string `xml:"numarticles,attr"`
			Postdate     string `xml:"postdate,attr"`
			Lastpostdate string `xml:"lastpostdate,attr"`
		} `xml:"thread"`
	} `xml:"threads"`
}

// Write unmarshals the response body to ForumItems
func (fi *ForumItems) Write(b *http.Response) error {
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
