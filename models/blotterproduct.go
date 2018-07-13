package models

import (
    "encoding/json"
)

type BlotterProduct struct {
    BlotterProductID int64 `json:"ID,omitempty"`
    BlotterProductSymbol string `json:"Symbol"`
    ContractSize int64 `json:"ContractSize"`
    ContractUnitID int64 `json:"ContractUnitID"`
}

type BlotterProductLeg struct {
    BlotterProductID int64 `json:"ID"`
    Rank int64 `json:"Rank"`
    ProductID int64 `json:"ProductID"`
    LegID int64 `json:"LegID,omitempty"`
}

type BlotterProductRequest struct {
    PartialSymbol string `json:"product"`
}

type BlotterProductResponseContent struct {
    BlotterProductID int64 `json:"ID"`
    BlotterProductSymbol string `json:"Symbol"`
}

type BlotterProductResponse struct {
    Products []BlotterProductResponseContent `json:"BlotterProducts"`
}

func (s *BlotterProduct) MarshalJSON() ([]byte, error) {
    type Alias BlotterProduct
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(s),
    })
}

func (t *BlotterProductLeg) MarshalJSON() ([]byte, error) {
    type Alias BlotterProductLeg
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(t),
    })
}