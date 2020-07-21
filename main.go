package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Plays struct {
	XMLName    xml.Name `xml:"plays"`
	Text       string   `xml:",chardata"`
	Username   string   `xml:"username,attr"`
	Userid     string   `xml:"userid,attr"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Play       []Play   `xml:"play"`
}

type Play struct {
	Text       string  `xml:",chardata"`
	ID         string  `xml:"id,attr"`
	Date       string  `xml:"date,attr"`
	Quantity   string  `xml:"quantity,attr"`
	Length     string  `xml:"length,attr"`
	Incomplete string  `xml:"incomplete,attr"`
	Nowinstats string  `xml:"nowinstats,attr"`
	Location   string  `xml:"location,attr"`
	Item       Item    `xml:"item"`
	Players    Players `xml:"players"`
}

type Item struct {
	Text       string   `xml:",chardata"`
	Name       string   `xml:"name,attr"`
	Objecttype string   `xml:"objecttype,attr"`
	Objectid   string   `xml:"objectid,attr"`
	Subtypes   Subtypes `xml:"subtypes"`
}

type Subtypes struct {
	Text    string    `xml:",chardata"`
	Subtype []Subtype `xml:"subtype"`
}

type Subtype struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}

type Players struct {
	Text   string   `xml:",chardata"`
	Player []Player `xml:"player"`
}

type Player struct {
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
}

func main() {
	resp, err := http.Get("http://www.boardgamegeek.com/xmlapi2/plays?username=Sebki")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	p := Plays{}
	err = xml.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	for _, v := range p.Play {
		fmt.Println(v.Item.Name, v.Item.Objecttype)
	}
}
