package twmodel

import "github.com/go-pg/pg/v10/orm"

func init() {
	registerModels()
}

func registerModels() {
	models := []interface{}{
		(*SpecialServer)(nil),
		(*Server)(nil),
		(*Version)(nil),
		(*PlayerToServer)(nil),
		(*PlayerNameChange)(nil),
		(*Tribe)(nil),
		(*Player)(nil),
		(*Village)(nil),
		(*Ennoblement)(nil),
		(*ServerStats)(nil),
		(*TribeHistory)(nil),
		(*PlayerHistory)(nil),
		(*TribeChange)(nil),
		(*DailyPlayerStats)(nil),
		(*DailyTribeStats)(nil),
	}

	for _, model := range models {
		orm.RegisterTable(model)
	}
}
