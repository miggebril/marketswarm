package models

import (
	"encoding/json"
	"encoding/base64"
)

type Broker struct {
    BrokerID uint64 `json:"ID"`
    Name string
}
func (b *Broker) MarshalJSON() ([]byte, error) {
    type Alias Broker
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(b),
    })
}

func (u Broker) GetIDEncoded() string {
	return base64.URLEncoding.EncodeToString([]byte(u.Name))
}