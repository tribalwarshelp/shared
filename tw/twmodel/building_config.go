package twmodel

import "encoding/xml"

type Building struct {
	Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
	MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
	MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
	Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
	Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
	Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
	Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
	WoodFactor      float64 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
	StoneFactor     float64 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
	IronFactor      float64 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
	PopFactor       float64 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
	BuildTime       float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
	BuildTimeFactor float64 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
}

type BuildingConfig struct {
	XMLName    xml.Name `xml:"config" json:"-" gqlgen:"xmlName"`
	Text       string   `xml:",chardata" json:"-" gqlgen:"text"`
	Main       Building `xml:"main" json:"main" gqlgen:"main"`
	Barracks   Building `xml:"barracks" json:"barracks" gqlgen:"barracks"`
	Stable     Building `xml:"stable" json:"stable" gqlgen:"stable"`
	Garage     Building `xml:"garage" json:"garage" gqlgen:"garage"`
	Watchtower Building `xml:"watchtower" json:"watchtower" gqlgen:"watchtower"`
	Snob       Building `xml:"snob" json:"snob" gqlgen:"snob"`
	Smith      Building `xml:"smith" json:"smith" gqlgen:"smith"`
	Place      Building `xml:"place" json:"place" gqlgen:"place"`
	Statue     Building `xml:"statue" json:"statue" gqlgen:"statue"`
	Market     Building `xml:"market" json:"market" gqlgen:"market"`
	Wood       Building `xml:"wood" json:"wood" gqlgen:"wood"`
	Stone      Building `xml:"stone" json:"stone" gqlgen:"stone"`
	Iron       Building `xml:"iron" json:"iron" gqlgen:"iron"`
	Farm       Building `xml:"farm" json:"farm" gqlgen:"farm"`
	Storage    Building `xml:"storage" json:"storage" gqlgen:"storage"`
	Hide       Building `xml:"hide" json:"hide" gqlgen:"hide"`
	Wall       Building `xml:"wall" json:"wall" gqlgen:"wall"`
}
