package twmodel

import "encoding/xml"

type Unit struct {
	Text           string  `xml:",chardata" json:"-" gqlgen:"text"`
	BuildTime      float64 `xml:"build_time" json:"build_time" gqlgen:"buildTime"`
	Pop            int     `xml:"pop" json:"pop" gqlgen:"pop"`
	Speed          float64 `xml:"speed" json:"speed" gqlgen:"speed"`
	Attack         int     `xml:"attack" json:"attack" gqlgen:"attack"`
	Defense        int     `xml:"defense" json:"defense" gqlgen:"defense"`
	DefenseCavalry int     `xml:"defense_cavalry" json:"defense_cavalry" gqlgen:"defenseCavalry"`
	DefenseArcher  int     `xml:"defense_archer" json:"defense_archer" gqlgen:"defenseArcher"`
	Carry          int     `xml:"carry" json:"carry" gqlgen:"carry"`
}

type UnitConfig struct {
	XMLName  xml.Name `xml:"config" json:"-" gqlgen:"xmlName"`
	Text     string   `xml:",chardata" json:"-" gqlgen:"text"`
	Spear    Unit     `xml:"spear" json:"spear" gqlgen:"spear"`
	Sword    Unit     `xml:"sword" json:"sword" gqlgen:"sword"`
	Axe      Unit     `xml:"axe" json:"axe" gqlgen:"axe"`
	Archer   Unit     `xml:"archer" json:"archer" gqlgen:"archer"`
	Spy      Unit     `xml:"spy" json:"spy" gqlgen:"spy"`
	Light    Unit     `xml:"light" json:"light" gqlgen:"light"`
	Marcher  Unit     `xml:"marcher" json:"marcher" gqlgen:"marcher"`
	Heavy    Unit     `xml:"heavy" json:"heavy" gqlgen:"heavy"`
	Ram      Unit     `xml:"ram" json:"ram" gqlgen:"ram"`
	Catapult Unit     `xml:"catapult" json:"catapult" gqlgen:"catapult"`
	Knight   Unit     `xml:"knight" json:"knight" gqlgen:"knight"`
	Snob     Unit     `xml:"snob" json:"snob" gqlgen:"snob"`
	Militia  Unit     `xml:"militia" json:"militia" gqlgen:"militia"`
}
