package controllers


import (
	"fmt"
	"log"
	"sync"
	"net/http"
	"html/template"
	"path/filepath"
	"marketswarm/models"
	"github.com/gorilla/pat"
)

var router *pat.Router
var cachedTemplates = map[string]*template.Template{}
var cachedMutex sync.Mutex

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
  PageTitle        string
  PageRadioButtons []RadioButton
  Answer           string
}

func Init(r *pat.Router) {
	router = r
}

func T(name string, pjax string) *template.Template {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	t := template.Must(template.New("welcome.html").ParseFiles(
		"welcome.html",
		filepath.Join("templates", name),
	))

	return t
}

func reverse(name string, things ...interface{}) string {
	//convert the things to strings
	strs := make([]string, len(things))
	for i, th := range things {
		strs[i] = fmt.Sprint(th)
	}
	//grab the route
	u, err := router.GetRoute(name).URL(strs...)
	if err != nil {
		panic(err)
	}
	return u.Path
}

func Index(w http.ResponseWriter, r *http.Request, ctx *models.Context) (err error) {
	// Display some radio buttons to the user

   Title := "Which do you prefer?"
   MyRadioButtons := []RadioButton{
     RadioButton{"animalselect", "cats", false, false, "Cats"},
     RadioButton{"animalselect", "dogs", false, false, "Dogs"},
   }

  MyPageVariables := PageVariables{
    PageTitle: Title,
    PageRadioButtons : MyRadioButtons,
    }

   t, err := template.ParseFiles("views/welcome.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
   }

   return err
}

func LoginForm(w http.ResponseWriter, r *http.Request, ctx *models.Context) (err error) {
	form, err := template.ParseFiles("views/login.html")
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	MyPageVariables := PageVariables{}
	err = form.Execute(w, MyPageVariables)
	if err != nil {
		log.Print("Template execution error: ", err)
	}

	return err
}

func Login(w http.ResponseWriter, req *http.Request, ctx *models.Context) error {
	email, password := req.FormValue("email"), req.FormValue("password")

	user, err := models.Login(ctx, email, password)
	if err != nil {
		log.Println("Error on login: ", err)
	}

	err = ctx.SetTrader(user.TraderID)
	if err != nil {
		log.Println("Error setting new trader to context: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(token)
	return nil
}