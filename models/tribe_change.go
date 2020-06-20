package models

import "time"

type TribeChange struct {
	ID         int       `json:"id" gqlgen:"id" xml:"id"`
	PlayerID   int       `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player     *Player   `json:"player,omitempty" gqlgen:"-" xml:"player"`
	OldTribeID int       `json:"oldTribeID" gqlgen:"oldTribeID" xml:"oldTribeID"`
	OldTribe   *Tribe    `json:"oldTribe,omitempty" gqlgen:"-" xml:"oldTribe"`
	NewTribeID int       `json:"newTribeID" gqlgen:"newTribeID" xml:"newTribeID"`
	NewTribe   *Tribe    `json:"newTribe,omitempty" gqlgen:"-" xml:"newTribe"`
	CreatedAt  time.Time `pg:"default:now(),use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

type TribeChangeFilter struct {
	ID    []int `json:"id" gqlgen:"id" xml:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ" xml:"idNEQ"`

	OldTribeID    []int `json:"oldTribeID" gqlgen:"oldTribeID" xml:"oldTribeID"`
	OldTribeIdNEQ []int `json:"oldTribeIDNEQ" gqlgen:"oldTribeIDNEQ" xml:"oldTribeIDNEQ"`

	NewTribeID    []int `json:"newTribeID" gqlgen:"newTribeID" xml:"newTribeID"`
	NewTribeIdNEQ []int `json:"newTribeIDNEQ" gqlgen:"newTribeIDNEQ" xml:"newTribeIDNEQ"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
