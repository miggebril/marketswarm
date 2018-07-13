package main

import (
	"database/sql"
	 "encoding/gob"
	 "fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-goods/httpbuf"
	"github.com/gorilla/pat"
	// "github.com/gorilla/sessions"
	"log"
	//"context"
	// "net/url"
	"net/http"
	"marketswarm/models"
	"marketswarm/lib/gorp"
	"marketswarm/controllers"
	//"net/smtp" email service
	 "os"
	// "strconv"
	// "strings"
	// "time"
)

func init() {
	gob.Register(int64(0))
}

var converter models.MarketswarmTypeConverter
var dbmap *gorp.DbMap
var db *sql.DB
var database string
var router *pat.Router
//var auth *models.JWTAuthenticationBackend

type handler func(http.ResponseWriter, *http.Request) error


func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	log.Println(req)
	log.Println("Auth:", req.Header.Get("Authorization"))

	buf := new(httpbuf.Buffer)
	log.Println(buf)

	//ctx, _ := models.NewContext(req, database)

	err := h(buf, req)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	buf.Apply(w)
}


func main() {
	fmt.Printf("pid: %d\n", os.Getpid())
	var dialect gorp.Dialect
	var err error
	var db *sql.DB

	u := "server=localhost;user id=sa;password=maximum2;database=Markets"
	db, err = sql.Open("mssql", u)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	dbmap = &gorp.DbMap{Db: db, Dialect: dialect}
	dbmap.TypeConverter = converter

	dbmap.AddTableWithName(models.Trader{}, "Traders").SetKeys(true, "TraderID")
	converter = models.MarketswarmTypeConverter{}
	dbmap.TypeConverter = converter
	
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err.Error())
	} else
	{
		fmt.Printf("DB connection successful\n")
	}

	stmt, err := db.Prepare("exec Select_Traders")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query() 
	if err != nil {
		log.Fatal("Initial query failed", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var password interface{}
		var verified bool
		err = rows.Scan(&id, &name, &email, &password, &verified)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}

		fmt.Printf("id:%d\n", id)
		fmt.Printf("email:%s\n", email)
	}
	router = pat.New()
	controllers.Init(router)

	router.Add("POST", "/login", handler(controllers.Login)).Name("Login")

	router.Add("GET", "/users/{id}", handler(controllers.UserInfo))
	router.Add("GET", "/users", handler(controllers.UsersIndex))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.Add("GET", "/", handler(controllers.Index)).Name("index")

	if err := http.ListenAndServe(":8077", router); err != nil {
		panic(err)
	}
}