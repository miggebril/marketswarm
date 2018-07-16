package models

import (
	"errors"
	"strconv"
	"encoding/json"
	"marketswarm/helpers"
	"golang.org/x/crypto/bcrypt"
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
func GetIdEncoded(id int64) string {
	return strconv.FormatInt(int64(id), 10)
}
func (u *Trader) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
}
//Login validates and returns a user object if they exist in the database.
func Login(ctx *Context, email, password string) (u *Trader, err error) {
	err = ctx.DbMap.SelectOne(&u, "select * from users where email=?", email)
	helpers.CheckErr(err, "Failed to login.")
	if err != nil {
		return
	}

	if len(u.Password) == 0 {
		err = errors.New("Password not set.")
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		u = nil
	}
	return
}