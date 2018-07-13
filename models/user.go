package models

import (
	//"labix.org/v2/mgo/bson"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"encoding/base64"
	"gemplay/helpers"
	"log"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

type Location struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type User struct {
	UserID       int64 
	Username string
	Password []byte	
	Email string
	Role int64
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}

//func (u *User) Rating() float64 {
	//return math.Ceil((float64(u.RatingSum)/float64(u.RatingCount))-0.5)
//}

//func (u User) Name() string {
//	return u.FirstName + " " + u.LastName
//}

func (u User) GetIDEncoded() string {
	return base64.URLEncoding.EncodeToString([]byte(u.Username))
}

//SetPassword takes a plaintext password and hashes it with bcrypt and sets the
//password field to the hash.
func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
}

//Login validates and returns a user object if they exist in the database.
func Login(ctx *Context, username string, password string) (b []byte, err error) {
	var u User
	log.Println("Logging in", username)
	err = ctx.DbMap.SelectOne(&u, "select * from [dbo].[User] where [Username]=?", username)
	helpers.CheckErr(err, "Failed to login.")
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err == nil {
		log.Println("Correct password")
		uid, err := helpers.ObjectIdFromString(u.GetIDEncoded())
		if err != nil {
			return []byte(""), err
		}

		token, err := ctx.Auth.GenerateToken(uid)
		if err != nil {
			return []byte(""), err
		} else {
			response, _ := json.Marshal(TokenAuthentication{token})
			return response, nil
		}
	}

	return []byte(""), err
}

func LoginNoPass(ctx *Context, email string) (u *User, err error) {
	err = ctx.DbMap.SelectOne(&u, "select * from User where email=?", email)
	return
}

func LoginFB(ctx *Context, fbid string) (u *User, err error) {
	// var users []User
	// err = ctx.DbMap.Select(&users, "select * from User where Director = ?", helpers.GetIDEncoded(u.ID))
	// if err != nil {
	// 	return
	// }

	return
}

func LoginTwitter(ctx *Context, twtid string) (u *User, err error) {
	// var users []User
	// err = ctx.DbMap.Select(&users, "select * from User where Twitter = ?", helpers.GetIDEncoded(u.ID))
	// if err != nil {
	// 	return
	// }

	return
}