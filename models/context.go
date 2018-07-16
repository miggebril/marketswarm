package models

import (
  "log"
  "fmt"
  "net/http"
  "marketswarm/helpers"
  "marketswarm/lib/gorp"
)

type Context struct {
  DbMap   *gorp.DbMap
  User    *Trader
}

func (ctx Context) GetTrader(id int64) Trader {
  var trader Trader
  err := ctx.DbMap.SelectOne(&trader, "select * from users where ID=?", GetIdEncoded(id))
  helpers.CheckErr(err, "Failed to get user")
  return trader
}

func (ctx Context) SetTrader(id int64) (error) {
  stmt, err := ctx.DbMap.Db.Prepare("exec Select_Traders_By_Id")
  if err != nil {
    log.Fatal("Prepare failed:", err.Error())
    return err
  }
  defer stmt.Close()

  row := stmt.QueryRow() 
  
  err = row.Scan(&ctx.User.TraderID, &ctx.User.UserName, &ctx.User.Email, &ctx.User.Password, &ctx.User.IsVerified)
  if err != nil {
    log.Fatal("Scan failed:", err.Error())
    return err
  }

  fmt.Printf("id:%d\n", ctx.User.TraderID)
  fmt.Printf("email:%s\n", ctx.User.Email)
  return nil
}

func (ctx *Context) LookupTrader(id int64) Trader {
  if id == 0 {
    return Trader{UserName: ""}
  }
  var user Trader
  err := ctx.DbMap.SelectOne(&user, "select * from users where ID=?", GetIdEncoded(id))
  helpers.CheckErr(err, "Failed to lookup user")
  return user

}

func NewContext(req *http.Request, dbmap *gorp.DbMap) (*Context, error) {
  ctx := &Context{
    DbMap:   dbmap,
  }
  
  return ctx, nil
}
