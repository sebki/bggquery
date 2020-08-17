package bggquery

import "encoding/xml"

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

// Write writes response body to ForumItems and fulfills io.Writer interface
func (fi *ForumItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, fi)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
