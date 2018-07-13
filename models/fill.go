package models

import (
	"encoding/json"
)

type OppositeParties []map[string]interface{}

type BlotterFill struct {
	GID int64 `json:"GID"`
	Data FillDetails `json:"FillData"`
	Legs []BlotterFill `json:"Legs"`
}

type NewFillReqDetails struct {
	PortfolioID int64 `json:"Portfolio"`
	ProductID int64 `json:"Symbol"`
	Quantity float64 `json:"Quantity"`
	Period string `json:"Period"`
	Side string `json:"Side"`
}

type FillDetails struct {
	GID int64 `json:"GID"`
	Portfolio string `json:"Portfolio"`
	ProductID int64 `json:"Symbol"`
	Quantity float64 `json:"Quantity"`
	Side string `json:"Side"`
	Period string `json:"Period,omitempty"`
	Price float64 `json:"Price,omitempty"`
}

type FillLeg struct {
	ProductID int64 `json:"ProductID"`
	Quantity float64 `json:"Quantity"`
	Side string `json:"Side"`
	Period string `json:"Period"`
	OtherParties OppositeParties `json:"OtherSides,omitempty"` // valid request properties but optional to include
	OptionalParams BlotterFillOptionalParams `json:"Params,omitempty"`
}

type BlotterFillOptionalParams struct {
    Price float64 `json:"Price,omitempty"`
    RTPLID int64 `json:"RTPLID,omitempty"`
    UserID int64 `json:"UserID,omitempty"`
    FillNotes string `json:"Notes,omitempty"`
}

type Fill struct {
	ID       uint64 `json:"-"`
	GroupID uint64 

	PortfolioID uint64 
	ExchangeID uint64 

	FillActionID uint64 `json:"Action"`
	FillStateID uint64 `json:"State"`
	FillTypeID uint64  `json:"Type"`
	ReconciliationStateID uint64

	ProductID int64
	Timestamp string 
	Volume float64
	Side bool
	Price float64

	UserID uint64 `json:"-"`
	OppositePortfolioID uint64 `json:"-"`
	BrokerID uint64 
	CounterpartyID uint64
	Description string 
}

type OTCFill struct {
	ID       uint64
	GID uint64 

	Portfolio string 

	Action uint64 `json:"Action"`
	State uint64 `json:"State"`
	Type uint64  `json:"Type"`
	RecState uint64

	Product uint64
	Time string
	Amount float64
	Side bool
	Price float64

	Broker string
	BrokerID uint64
	CounterID uint64
	Counterparty string 
}

type FillState struct {
    FillStateID uint64 `json:"State"`
    Description string
}

type FillAction struct {
    FillActionID uint64 `json:"Action"`
    Description string
}

type FillType struct {
    FillTypeID uint64 `json:"Type"`
    Description string
}

type ReconciliationState struct {
	ReconciliationStateID uint64 `json:"RecState"`
	Description string
}

func (t *Fill) MarshalJSON() ([]byte, error) {
	type Alias Fill
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	})
}