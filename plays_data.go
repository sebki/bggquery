package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// PlaysItem contains response data from a Plays-request to boardgamegeek
type PlaysItems struct {
	XMLName    xml.Name `xml:"plays"`
	Text       string   `xml:",chardata"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Play       []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"id,attr"`
		Userid     string `xml:"userid,attr"`
		Date       string `xml:"date,attr"`
		Quantity   string `xml:"quantity,attr"`
		Length     string `xml:"length,attr"`
		Incomplete string `xml:"incomplete,attr"`
		Nowinstats string `xml:"nowinstats,attr"`
		Location   string `xml:"location,attr"`
		Item       struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Objecttype string `xml:"objecttype,attr"`
			Objectid   string `xml:"objectid,attr"`
			Subtypes   struct {
				Text    string `xml:",chardata"`
				Subtype []struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"subtype"`
			} `xml:"subtypes"`
		} `xml:"item"`
		Comments string `xml:"comments"`
		Players  struct {
			Text   string `xml:",chardata"`
			Player []struct {
				Text          string `xml:",chardata"`
				Username      string `xml:"username,attr"`
				Userid        string `xml:"userid,attr"`
				Name          string `xml:"name,attr"`
				Startposition string `xml:"startposition,attr"`
				Color         string `xml:"color,attr"`
				Score         string `xml:"score,attr"`
				New           string `xml:"new,attr"`
				Rating        string `xml:"rating,attr"`
				Win           string `xml:"win,attr"`
			} `xml:"player"`
		} `xml:"players"`
	} `xml:"play"`
}

// Write unmarshals the response body to PlaysItems
func (pi *PlaysItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, pi)
	if err != nil {
		return err
	}
	return nil
}
