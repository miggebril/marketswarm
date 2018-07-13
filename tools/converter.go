package tools
	
import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)
	
const (
	base_url = "https://openexchangerates.org/api/"
	api_key = "61258132fb5c4037b541870ae13cebfa"
)


type OXResponse struct {
	Disclaimer string 
	License string 
	Timestamp int
	Base string
	Rates map[string]float64  
} 

func Convert(amount uint64, from string, to string) (bal uint64, err error) {

	method := "latest.json"
	key := "?app_id=" + api_key
	base := "&base=" + from
	cur := "&symbols=" + to

	params := base + cur
	endpoint := base_url + method + key + params 

	log.Println("API request to:", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("Error connecting to exchange API")
		return 0, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error:", err)
		return 0, err
	}

	var body = new(OXResponse)
	err = json.Unmarshal(respBytes, &body)
	if err != nil {
		log.Println("Error:", err)
		return 0, err
	}	
	fmt.Println("Response:", body)

	res := uint64(body.Rates[to] * float64(amount))
	log.Println("Result:", res)

	return 0, nil
}