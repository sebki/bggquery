package bggquery

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// CollectionItems represents all response data from a CollectionQuery to Boardgamegeek
type CollectionItems struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Totalitems string   `xml:"totalitems,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Pubdate    string   `xml:"pubdate,attr"`
	Item       []struct {
		Text       string `xml:",chardata"`
		Objecttype string `xml:"objecttype,attr"`
		Objectid   string `xml:"objectid,attr"`
		Subtype    string `xml:"subtype,attr"`
		Collid     string `xml:"collid,attr"`
		Name       struct {
			Text      string `xml:",chardata"`
			Sortindex string `xml:"sortindex,attr"`
		} `xml:"name"`
		Yearpublished string `xml:"yearpublished"`
		Image         string `xml:"image"`
		Thumbnail     string `xml:"thumbnail"`
		Stats         struct {
			Text        string `xml:",chardata"`
			Minplayers  string `xml:"minplayers,attr"`
			Maxplayers  string `xml:"maxplayers,attr"`
			Minplaytime string `xml:"minplaytime,attr"`
			Maxplaytime string `xml:"maxplaytime,attr"`
			Playingtime string `xml:"playingtime,attr"`
			Numowned    string `xml:"numowned,attr"`
			Rating      struct {
				Text       string `xml:",chardata"`
				Value      string `xml:"value,attr"`
				Usersrated struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"usersrated"`
				Average struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"average"`
				Bayesaverage struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"bayesaverage"`
				Stddev struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"stddev"`
				Median struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"median"`
				Ranks struct {
					Text string `xml:",chardata"`
					Rank []struct {
						Text         string `xml:",chardata"`
						Type         string `xml:"type,attr"`
						ID           string `xml:"id,attr"`
						Name         string `xml:"name,attr"`
						Friendlyname string `xml:"friendlyname,attr"`
						Value        string `xml:"value,attr"`
						Bayesaverage string `xml:"bayesaverage,attr"`
					} `xml:"rank"`
				} `xml:"ranks"`
			} `xml:"rating"`
		} `xml:"stats"`
		Status struct {
			Text             string `xml:",chardata"`
			Own              string `xml:"own,attr"`
			Prevowned        string `xml:"prevowned,attr"`
			Fortrade         string `xml:"fortrade,attr"`
			Want             string `xml:"want,attr"`
			Wanttoplay       string `xml:"wanttoplay,attr"`
			Wanttobuy        string `xml:"wanttobuy,attr"`
			Wishlist         string `xml:"wishlist,attr"`
			Preordered       string `xml:"preordered,attr"`
			Lastmodified     string `xml:"lastmodified,attr"`
			Wishlistpriority string `xml:"wishlistpriority,attr"`
		} `xml:"status"`
		Numplays string `xml:"numplays"`
	} `xml:"item"`
}

// Write unmarshals the response body to CollectionItems
func (ci *CollectionItems) Write(b *http.Response) error {
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, ci)
	if err != nil {
		return err
	}
	return nil
}
