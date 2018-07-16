package models

import (
    "encoding/json"
    "strings"
)

type Counterparty struct {
    CounterpartyID uint64 `json:"ID"`
    CounterpartyName string
}

type Exchange struct {
    ExchangeID uint64 `json:"ID"`
    Name string
}

type ContractPeriod struct {
    ContractPeriodName string `json:"Name"`
    StartMonthIndex int `json:"Index"`
    MonthLength int `json:"Length"`
    ContractPeriodCode string `json:"Code"`
}

type ContractPeriodResponseContent struct {
    Period ContractPeriod `json:"ContractPeriod"`
    Year int `json:"Year"`
}

type ContractPeriodResponse struct {
    ContractPeriods []ContractPeriod `json:"ContractPeriods"`
}

func (c *Counterparty) MarshalJSON() ([]byte, error) {
    type Alias Counterparty
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(c),
    })
}

func (e *Exchange) MarshalJSON() ([]byte, error) {
    type Alias Exchange
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(e),
    })
}


func (cp *ContractPeriod) MarshalJSON() ([]byte, error) {
    type Alias ContractPeriod
    return json.Marshal(&struct {
        ContractPeriodCode  string  `json:"Code"`
        *Alias
    }{
        ContractPeriodCode: strings.TrimSpace(cp.ContractPeriodCode),
        Alias: (*Alias)(cp),
    })
}