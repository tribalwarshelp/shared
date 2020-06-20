package models

import "time"

type Ennoblement struct {
	tableName struct{} `pg:"?SERVER.ennoblements,alias:ennoblement"`

	VillageID       int       `gqlgen:"-" json:"villageID" xml:"villageID"`
	Village         *Village  `gqlgen:"-" json:"village,omitempty" xml:"village"`
	NewOwnerID      int       `gqlgen:"-" json:"newOwnerID" xml:"newOwnerID"`
	NewOwner        *Player   `gqlgen:"-" json:"newOwner,omitempty" xml:"newOwner"`
	NewOwnerTribeID int       `json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	NewOwnerTribe   *Tribe    `json:"newOwnerTribe,omitempty" gqlgen:"-" xml:"newOwnerTribe"`
	OldOwnerID      int       `gqlgen:"-" json:"oldOwnerID" xml:"oldOwnerID"`
	OldOwner        *Player   `gqlgen:"-" json:"oldOwner,omitempty" xml:"oldOwner"`
	OldOwnerTribeID int       `json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
	OldOwnerTribe   *Tribe    `json:"oldOwnerTribe,omitempty" gqlgen:"-" xml:"oldOwnerTribe"`
	EnnobledAt      time.Time `pg:"default:now(),use_zero" json:"ennobledAt" gqlgen:"ennobledAt" xml:"ennobledAt"`
}

type EnnoblementFilter struct {
	tableName struct{} `urlstruct:"ennoblement"`

	VillageID    []int `json:"villageID" gqlgen:"villageID" xml:"villageID"`
	VillageIdNEQ []int `json:"villageIDNEQ" gqlgen:"villageIDNEQ" xml:"villageIDNEQ"`

	NewOwnerID         []int `json:"newOwnerID" gqlgen:"newOwnerID" xml:"newOwnerID"`
	NewOwnerIdNEQ      []int `json:"newOwnerIdNEQ" gqlgen:"newOwnerIdNEQ" xml:"newOwnerIdNEQ"`
	NewOwnerTribeID    []int `json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	NewOwnerTribeIdNEQ []int `json:"newOwnerTribeIDNEQ" gqlgen:"newOwnerTribeIDNEQ" xml:"newOwnerTribeIDNEQ"`

	OldOwnerID         []int `json:"oldOwnerID" gqlgen:"oldOwnerID" xml:"oldOwnerID"`
	OldOwnerIdNEQ      []int `json:"oldOwnerIdNEQ" gqlgen:"oldOwnerIdNEQ" xml:"oldOwnerIdNEQ"`
	OldOwnerTribeID    []int `json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
	OldOwnerTribeIdNEQ []int `json:"oldOwnerTribeIDNEQ" gqlgen:"oldOwnerTribeIDNEQ" xml:"oldOwnerTribeIDNEQ"`

	EnnobledAt    time.Time `json:"ennobledAt" gqlgen:"ennobledAt" xml:"ennobledAt"`
	EnnobledAtGT  time.Time `json:"ennobledAtGT" gqlgen:"ennobledAtGT" xml:"ennobledAtGT"`
	EnnobledAtGTE time.Time `json:"ennobledAtGTE" gqlgen:"ennobledAtGTE" xml:"ennobledAtGTE"`
	EnnobledAtLT  time.Time `json:"ennobledAtLT" gqlgen:"ennobledAtLT" xml:"ennobledAtLT"`
	EnnobledAtLTE time.Time `json:"ennobledAtLTE" gqlgen:"ennobledAtLTE" xml:"ennobledAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
