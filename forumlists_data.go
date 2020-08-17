package bggquery

import "encoding/xml"

// ForumListsItems contains all response data from a ForumListsQuery
type ForumListsItems struct {
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

// Write writes response body to ForumListsItems and fulfills io.Writer interface
func (fli *ForumListsItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, fli)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
