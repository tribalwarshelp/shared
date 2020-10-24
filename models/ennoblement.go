package models

import "time"

type Ennoblement struct {
	tableName struct{} `pg:"?SERVER.ennoblements,alias:ennoblement"`

	VillageID       int       `pg:",use_zero" gqlgen:"-" json:"villageID" xml:"villageID"`
	Village         *Village  `gqlgen:"-" json:"village,omitempty" xml:"village" pg:"rel:has-one"`
	NewOwnerID      int       `pg:",use_zero" gqlgen:"-" json:"newOwnerID" xml:"newOwnerID"`
	NewOwner        *Player   `gqlgen:"-" json:"newOwner,omitempty" xml:"newOwner" pg:"rel:has-one"`
	NewOwnerTribeID int       `pg:",use_zero" json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	NewOwnerTribe   *Tribe    `json:"newOwnerTribe,omitempty" gqlgen:"-" xml:"newOwnerTribe" pg:"rel:has-one"`
	OldOwnerID      int       `pg:",use_zero" gqlgen:"-" json:"oldOwnerID" xml:"oldOwnerID"`
	OldOwner        *Player   `gqlgen:"-" json:"oldOwner,omitempty" xml:"oldOwner" pg:"rel:has-one"`
	OldOwnerTribeID int       `pg:",use_zero" json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
	OldOwnerTribe   *Tribe    `json:"oldOwnerTribe,omitempty" gqlgen:"-" xml:"oldOwnerTribe" pg:"rel:has-one"`
	EnnobledAt      time.Time `pg:"default:now(),use_zero" json:"ennobledAt" gqlgen:"ennobledAt" xml:"ennobledAt"`
}

type EnnoblementFilterOr struct {
	NewOwnerID      []int `json:"newOwnerID" gqlgen:"newOwnerID" xml:"newOwnerID"`
	NewOwnerTribeID []int `json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	OldOwnerID      []int `json:"oldOwnerID" gqlgen:"oldOwnerID" xml:"oldOwnerID"`
	OldOwnerTribeID []int `json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
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

	Or *EnnoblementFilterOr `urlstruct:",nowhere" json:"or" gqlgen:"or" xml:"or"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
