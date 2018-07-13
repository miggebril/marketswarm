package models

import (
	"encoding/json"
	"encoding/base64"
)
type Trader struct {
	TraderID       int64 
	UserName string
	Password []byte	
	Email string
	IsVerified bool
}

func (u *Trader) MarshalJSON() ([]byte, error) {
	type Alias Trader
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}
func (u Trader) GetIDEncoded() string {
	return base64.URLEncoding.EncodeToString([]byte(u.UserName))
}