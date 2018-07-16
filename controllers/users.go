package controllers


import (
	"encoding/json"
	"net/http"
	"marketswarm/models"
	"marketswarm/helpers"
)

func UserInfo(w http.ResponseWriter, r *http.Request, ctx *models.Context) (err error) {
	// js, err := json.Marshal(&interface{})
 //    if err != nil {
 //        http.Error(w, err.Error(), http.StatusInternalServerError)
 //        return err
 //    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    //w.Write()
    return nil
}

func UsersIndex(w http.ResponseWriter, r *http.Request, ctx *models.Context) (err error) {
	var users []models.Trader
	_, err = ctx.DbMap.Select(&users, "SELECT * FROM Traders")
	helpers.CheckErr(err, "Failed to query all traders")

	js, err := json.Marshal(users)
	helpers.CheckErr(err, "Failed to marshal users")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(js)
    return nil
}