package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// UserItems contains response data from a User query from boardgamegeek
type UserItems struct {
	XMLName    xml.Name `xml:"user"`
	Text       string   `xml:",chardata"`
	ID         string   `xml:"id,attr"`
	Name       string   `xml:"name,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Firstname  struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"firstname"`
	Lastname struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"lastname"`
	Avatarlink struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"avatarlink"`
	Yearregistered struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"yearregistered"`
	Lastlogin struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"lastlogin"`
	Stateorprovince struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"stateorprovince"`
	Country struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"country"`
	Webaddress struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"webaddress"`
	Xboxaccount struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"xboxaccount"`
	Wiiaccount struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"wiiaccount"`
	Psnaccount struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"psnaccount"`
	Battlenetaccount struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"battlenetaccount"`
	Steamaccount struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"steamaccount"`
	Traderating struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"traderating"`
	Marketrating struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"marketrating"`
	Buddies struct {
		Text  string `xml:",chardata"`
		Total string `xml:"total,attr"`
		Page  string `xml:"page,attr"`
		Buddy []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"buddy"`
	} `xml:"buddies"`
	Guilds struct {
		Text  string `xml:",chardata"`
		Total string `xml:"total,attr"`
		Page  string `xml:"page,attr"`
		Guild []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"guild"`
	} `xml:"guilds"`
	Top struct {
		Text   string `xml:",chardata"`
		Domain string `xml:"domain,attr"`
		Item   []struct {
			Text string `xml:",chardata"`
			Rank string `xml:"rank,attr"`
			Type string `xml:"type,attr"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"item"`
	} `xml:"top"`
}

// Write unmarshals the response body to UserItems
func (ui *UserItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, ui)
	if err != nil {
		return err
	}
	return nil
}
