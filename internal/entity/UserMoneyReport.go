package entity

import uuid "github.com/satori/go.uuid"

type UserMoneyReport struct {
	ID           int64
	StationID    int
	Banknotes    int
	CarsTotal    int
	Coins        int
	Electronical int
	Service      int
	Bonuses      int
	SessionID    uuid.UUID
	User         string
	Processed    bool
	UUID         uuid.UUID
}
