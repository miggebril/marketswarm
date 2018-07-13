package models

import (
    "encoding/json"
)
 
type SpreadType struct {
    SpreadID int64 `json:"ID,omitempty"`
    SpreadName string
    LegCount int64
}

type TradeType struct {
    TradeTypeID int64 `json:"ID,omitempty"`
    TradeTypeName string
}

type SpreadLeg struct {
    LegRank int64 `json:"Rank"`
    SpreadID int64 `json:"SpreadID"`
    LegPriceMultiplier float64 `json:"PriceMultiplier"`
    LegVolumeMultiplier float64 `json:"VolumeMultiplier"`
}

func (s *SpreadType) MarshalJSON() ([]byte, error) {
    type Alias SpreadType
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(s),
    })
}

func (t *TradeType) MarshalJSON() ([]byte, error) {
    type Alias TradeType
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(t),
    })
}

func (l *SpreadLeg) MarshalJSON() ([]byte, error) {
    type Alias SpreadLeg
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(l),
    })
}