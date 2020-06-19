package models

import "encoding/xml"

type UnitConfig struct {
	XMLName xml.Name `xml:"config" json:"xmlName" gqlgen:"xmlName"`
	Text    string   `xml:",chardata" json:"text" gqlgen:"text"`
	Spear   struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"spear" json:"spear" gqlgen:"spear"`
	Sword struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"sword" json:"sword" gqlgen:"sword"`
	Axe struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"axe" json:"axe" gqlgen:"axe"`
	Archer struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"archer" json:"archer" gqlgen:"archer"`
	Spy struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"spy" json:"spy" gqlgen:"spy"`
	Light struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"light" json:"light" gqlgen:"light"`
	Marcher struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"marcher" json:"marcher" gqlgen:"marcher"`
	Heavy struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"heavy" json:"heavy" gqlgen:"heavy"`
	Ram struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"ram" json:"ram" gqlgen:"ram"`
	Catapult struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"catapult" json:"catapult" gqlgen:"catapult"`
	Knight struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"knight" json:"knight" gqlgen:"knight"`
	Snob struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"snob" json:"snob" gqlgen:"snob"`
	Militia struct {
		Text           string `xml:",chardata" json:"text" gqlgen:"text"`
		BuildTime      string `xml:"build_time" json:"buildTime" gqlgen:"buildTime"`
		Pop            string `xml:"pop" json:"pop" gqlgen:"pop"`
		Speed          string `xml:"speed" json:"speed" gqlgen:"speed"`
		Attack         string `xml:"attack" json:"attack" gqlgen:"attack"`
		Defense        string `xml:"defense" json:"defense" gqlgen:"defense"`
		DefenseCavalry string `xml:"defense_cavalry" json:"defenseCavalry" gqlgen:"defenseCavalry"`
		DefenseArcher  string `xml:"defense_archer" json:"defenseArcher" gqlgen:"defenseArcher"`
		Carry          string `xml:"carry" json:"carry" gqlgen:"carry"`
	} `xml:"militia" json:"militia" gqlgen:"militia"`
}
