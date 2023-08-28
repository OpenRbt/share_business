package entity

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
