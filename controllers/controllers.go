package controllers


import (
	//"fmt"
	"sync"
	"net/http"
	"html/template"
	"path/filepath"
	"github.com/gorilla/pat"
)

var router *pat.Router
var cachedTemplates = map[string]*template.Template{}
var cachedMutex sync.Mutex

func Init(r *pat.Router) {
	router = r
}

func T(name string, pjax string) *template.Template {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	t := template.Must(template.New("_base.html").ParseFiles(
		"templates/_base.html",
		filepath.Join("templates", name),
	))

	return t
}

func Index(w http.ResponseWriter, r *http.Request) (err error) {
	return T("index.html", "na").Execute(w, map[string]interface{}{})
}

func Login(w http.ResponseWriter, r *http.Request) error {
	// email, password := r.FormValue("email"), r.FormValue("password")
	// fmt.Println("Login called on", email, password)
	// token, err := models.Login(ctx, email, password)
	// if err != nil {
	// 	http.Error(w, "Invalid password.", http.StatusInternalServerError)
	// 	return err
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(token)
	return nil
}