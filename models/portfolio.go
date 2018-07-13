package models

import (
    "encoding/json"
)
 
type Portfolio struct {
    PortfolioID int64
    Name string
}
 
func (r *Portfolio) MarshalJSON() ([]byte, error) {
    type Alias Portfolio
    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(r),
    })
}