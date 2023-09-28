package entities

import uuid "github.com/satori/go.uuid"

type MoneyReport struct {
	StationID      int
	Banknotes      int
	CarsTotal      int
	Coins          int
	Electronical   int
	Service        int
	Bonuses        int
	SessionID      *uuid.UUID
	OrganizationID uuid.UUID
	Processed      bool
	UUID           uuid.UUID
}

type UserMoneyReport struct {
	ID             int64
	StationID      int
	Banknotes      int
	CarsTotal      int
	Coins          int
	Electronical   int
	Service        int
	Bonuses        int
	SessionID      uuid.UUID
	OrganizationID uuid.UUID
	User           string
	Processed      bool
	UUID           uuid.UUID
}
