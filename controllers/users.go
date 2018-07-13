package controllers


import (
	//"json"
	"net/http"
)

func UserInfo(w http.ResponseWriter, r *http.Request) (err error) {
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

func UsersIndex(w http.ResponseWriter, r *http.Request) (err error) {
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