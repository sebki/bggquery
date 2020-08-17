package bggquery

import "encoding/xml"

// GuildItems contains all response data from a GuildsQuery
type GuildItems struct {
	XMLName     xml.Name `xml:"guild"`
	Text        string   `xml:",chardata"`
	ID          string   `xml:"id,attr"`
	Name        string   `xml:"name,attr"`
	Created     string   `xml:"created,attr"`
	Termsofuse  string   `xml:"termsofuse,attr"`
	Category    string   `xml:"category"`
	Website     string   `xml:"website"`
	Manager     string   `xml:"manager"`
	Description string   `xml:"description"`
	Location    struct {
		Text            string `xml:",chardata"`
		Addr1           string `xml:"addr1"`
		Addr2           string `xml:"addr2"`
		City            string `xml:"city"`
		Stateorprovince string `xml:"stateorprovince"`
		Postalcode      string `xml:"postalcode"`
		Country         string `xml:"country"`
	} `xml:"location"`
	Members struct {
		Text   string `xml:",chardata"`
		Count  string `xml:"count,attr"`
		Page   string `xml:"page,attr"`
		Member []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Date string `xml:"date,attr"`
		} `xml:"member"`
	} `xml:"members"`
}

// Write writes response to GuildItems and fulfills io.Writer interface
func (gi *GuildItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, gi)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
