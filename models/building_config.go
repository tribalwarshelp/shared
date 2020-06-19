package models

import "encoding/xml"

type BuildingConfig struct {
	XMLName xml.Name `xml:"config" json:"-" gqlgen:"xmlName"`
	Text    string   `xml:",chardata" json:"-" gqlgen:"text"`
	Main    struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"main" json:"main" gqlgen:"main"`
	Barracks struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"barracks" json:"barracks" gqlgen:"barracks"`
	Stable struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"stable" json:"stable" gqlgen:"stable"`
	Garage struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"garage" json:"garage" gqlgen:"garage"`
	Watchtower struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"watchtower" json:"watchtower" gqlgen:"watchtower"`
	Snob struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Smith struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"smith" json:"smith" gqlgen:"smith"`
	Place struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"place" json:"place" gqlgen:"place"`
	Statue struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"statue" json:"statue" gqlgen:"statue"`
	Market struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"market" json:"market" gqlgen:"market"`
	Wood struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"wood" json:"wood" gqlgen:"wood"`
	Stone struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"stone" json:"stone" gqlgen:"stone"`
	Iron struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"iron" json:"iron" gqlgen:"iron"`
	Farm struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"farm" json:"farm" gqlgen:"farm"`
	Storage struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"storage" json:"storage" gqlgen:"storage"`
	Hide struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"hide" json:"hide" gqlgen:"hide"`
	Wall struct {
		Text            string  `xml:",chardata" json:"-" gqlgen:"text"`
		MaxLevel        int     `xml:"max_level" json:"max_level" gqlgen:"maxLevel"`
		MinLevel        int     `xml:"min_level" json:"min_level" gqlgen:"minLevel"`
		Wood            int     `xml:"wood" json:"wood" gqlgen:"wood"`
		Stone           int     `xml:"stone" json:"stone" gqlgen:"stone"`
		Iron            int     `xml:"iron" json:"iron" gqlgen:"iron"`
		Pop             int     `xml:"pop" json:"pop" gqlgen:"pop"`
		WoodFactor      float32 `xml:"wood_factor" json:"wood_factor" gqlgen:"woodFactor"`
		StoneFactor     float32 `xml:"stone_factor" json:"stone_factor" gqlgen:"stoneFactor"`
		IronFactor      float32 `xml:"iron_factor" json:"iron_factor" gqlgen:"ironFactor"`
		PopFactor       float32 `xml:"pop_factor" json:"pop_factor" gqlgen:"popFactor"`
		BuildTime       float32 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		BuildTimeFactor float32 `xml:"build_time_factor" json:"build_time_factor" gqlgen:"buildTimeFactor"`
	} `xml:"wall" json:"wall" gqlgen:"wall"`
}
