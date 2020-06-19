package models

import "encoding/xml"

type UnitConfig struct {
	XMLName xml.Name `xml:"config" json:"-" gqlgen:"xmlName"`
	Text    string   `xml:",chardata" json:"-" gqlgen:"text"`
	Spear   struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"spear" json:"spear" gqlgen:"spear"`
	Sword struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"sword" json:"sword" gqlgen:"sword"`
	Axe struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"axe" json:"axe" gqlgen:"axe"`
	Archer struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"archer" json:"archer" gqlgen:"archer"`
	Spy struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"spy" json:"spy" gqlgen:"spy"`
	Light struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"light" json:"light" gqlgen:"light"`
	Marcher struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"marcher" json:"marcher" gqlgen:"marcher"`
	Heavy struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"heavy" json:"heavy" gqlgen:"heavy"`
	Ram struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"ram" json:"ram" gqlgen:"ram"`
	Catapult struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"catapult" json:"catapult" gqlgen:"catapult"`
	Knight struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"knight" json:"knight" gqlgen:"knight"`
	Snob struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Militia struct {
		Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
		BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
		Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
		Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"militia" json:"militia" gqlgen:"militia"`
}
