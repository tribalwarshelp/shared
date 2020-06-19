package models

import "encoding/xml"

type BuildingConfig struct {
	XMLName xml.Name `xml:"config" json:"xmlName" gqlgen:"xmlName"`
	Text    string   `xml:",chardata" json:"text" gqlgen:"text"`
	Main    struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"main" json:"main" gqlgen:"main"`
	Barracks struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"barracks" json:"barracks" gqlgen:"barracks"`
	Stable struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"stable" json:"stable" gqlgen:"stable"`
	Garage struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"garage" json:"garage" gqlgen:"garage"`
	Watchtower struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"watchtower" json:"watchtower" gqlgen:"watchtower"`
	Snob struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Smith struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"smith" json:"smith" gqlgen:"smith"`
	Place struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"place" json:"place" gqlgen:"place"`
	Statue struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"statue" json:"statue" gqlgen:"statue"`
	Market struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"market" json:"market" gqlgen:"market"`
	Wood struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"wood" json:"wood" gqlgen:"wood"`
	Stone struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"stone" json:"stone" gqlgen:"stone"`
	Iron struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"iron" json:"iron" gqlgen:"iron"`
	Farm struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"farm" json:"farm" gqlgen:"farm"`
	Storage struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"storage" json:"storage" gqlgen:"storage"`
	Hide struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"hide" json:"hide" gqlgen:"hide"`
	Wall struct {
		Text            string `xml:",chardata" json:"text" gqlgen:"text"`
		MaxLevel        string `xml:"max_level" json:"maxLevel" gqlgen:"maxLevel"`
		MinLevel        string `xml:"min_level" json:"minLevel" gqlgen:"minLevel"`
		Wood            string `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           string `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            string `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             string `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      string `xml:"wood_factor" json:"woodFactor" gqlgen:"woodFactor"`
		StoneFactor     string `xml:"stone_factor" json:"stoneFactor" gqlgen:"stoneFactor"`
		IronFactor      string `xml:"iron_factor" json:"ironFactor" gqlgen:"ironFactor"`
		PopFactor       string `xml:"pop_factor" json:"popFactor" gqlgen:"popFactor"`
		BuildTime       string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		BuildTimeFactor string `xml:"build_time_factor" json:"buildTimeFactor" gqlgen:"buildTimeFactor"`
	} `xml:"wall" json:"wall" gqlgen:"wall"`
}
