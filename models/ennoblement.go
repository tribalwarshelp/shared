package models

import "time"

type Ennoblement struct {
	VillageID  int       `json:"-" gqlgen:"-"`
	Village    *Village  `json:"-" gqlgen:"-"`
	NewOwnerID int       `json:"-" gqlgen:"-"`
	NewOwner   *Player   `json:"-" gqlgen:"-"`
	OldOwnerID int       `json:"-" gqlgen:"-"`
	OldOwner   *Player   `json:"-" gqlgen:"-"`
	EnnobledAt time.Time `json:"ennobledAt" gqlgen:"ennobledAt"`
}
