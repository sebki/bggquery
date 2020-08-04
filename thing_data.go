package bggquery

import "encoding/xml"

// ThingItems contains all possible data response of a "thing"-query on Boardgamegeek
type ThingItems struct {
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
		Description   string `xml:"description"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
		Minplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplayers"`
		Maxplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplayers"`
		Poll []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Title      string `xml:"title,attr"`
			Totalvotes string `xml:"totalvotes,attr"`
			Results    []struct {
				Text       string `xml:",chardata"`
				Numplayers string `xml:"numplayers,attr"`
				Result     []struct {
					Text     string `xml:",chardata"`
					Value    string `xml:"value,attr"`
					Numvotes string `xml:"numvotes,attr"`
					Level    string `xml:"level,attr"`
				} `xml:"result"`
			} `xml:"results"`
		} `xml:"poll"`
		Playingtime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"playingtime"`
		Minplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplaytime"`
		Maxplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplaytime"`
		Minage struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minage"`
		Link []struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			ID    string `xml:"id,attr"`
			Value string `xml:"value,attr"`
		} `xml:"link"`
		Videos struct {
			Text  string `xml:",chardata"`
			Total string `xml:"total,attr"`
			Video []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Title    string `xml:"title,attr"`
				Category string `xml:"category,attr"`
				Language string `xml:"language,attr"`
				Link     string `xml:"link,attr"`
				Username string `xml:"username,attr"`
				Userid   string `xml:"userid,attr"`
				Postdate string `xml:"postdate,attr"`
			} `xml:"video"`
		} `xml:"videos"`
		Versions struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				ID        string `xml:"id,attr"`
				Thumbnail string `xml:"thumbnail"`
				Image     string `xml:"image"`
				Link      []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					ID      string `xml:"id,attr"`
					Value   string `xml:"value,attr"`
					Inbound string `xml:"inbound,attr"`
				} `xml:"link"`
				Name struct {
					Text      string `xml:",chardata"`
					Type      string `xml:"type,attr"`
					Sortindex string `xml:"sortindex,attr"`
					Value     string `xml:"value,attr"`
				} `xml:"name"`
				Yearpublished struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"yearpublished"`
				Productcode struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"productcode"`
				Width struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"width"`
				Length struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"length"`
				Depth struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"depth"`
				Weight struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"weight"`
			} `xml:"item"`
		} `xml:"versions"`
		Comments struct {
			Text       string `xml:",chardata"`
			Page       string `xml:"page,attr"`
			Totalitems string `xml:"totalitems,attr"`
			Comment    []struct {
				Text     string `xml:",chardata"`
				Username string `xml:"username,attr"`
				Rating   string `xml:"rating,attr"`
				Value    string `xml:"value,attr"`
			} `xml:"comment"`
		} `xml:"comments"`
		Marketplacelistings struct {
			Text    string `xml:",chardata"`
			Listing []struct {
				Text     string `xml:",chardata"`
				Listdate struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"listdate"`
				Price struct {
					Text     string `xml:",chardata"`
					Currency string `xml:"currency,attr"`
					Value    string `xml:"value,attr"`
				} `xml:"price"`
				Condition struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"condition"`
				Notes struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"notes"`
				Link struct {
					Text  string `xml:",chardata"`
					Href  string `xml:"href,attr"`
					Title string `xml:"title,attr"`
				} `xml:"link"`
			} `xml:"listing"`
		} `xml:"marketplacelistings"`
	} `xml:"item"`
}

func (ti *ThingItems) Write(b []byte) (n int, err error) {
	err = xml.Unmarshal(b, ti)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
