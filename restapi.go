package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	_ "github.com/lib/pq"
	"github.com/mtmoses/httprouter"
)


/*
Server is the core structure for http router
*/
type Server struct {
	router *httprouter.Router
}

/*
Request is the core structure for web service input
*/
type Request struct {
	InputOne string `json:"inputone"`
	InputTwo string `json:"inputtwo"`
}

/*
Response is the core structure for web service output
*/
type Response struct {
	Status     bool    `json:"status"`
	Data       string  `json:"data,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Message    string  `json:"message"`
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Access-Control-Allow-Origin, Token, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Headers, *")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	s.router.ServeHTTP(w, r)
}

func showSplashscreen() {
	screenImage := `
	API HEALTHY
`
	fmt.Println(screenImage)
	fmt.Println("===============")
	fmt.Println("API")
	fmt.Println("===============")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "HI the api is working")
}

func initializeRoutes() {
	port := "8050"
	url := "localhost"

	portString := ":" + port
	fmt.Println("Starting server on\n", url, portString)

	router := httprouter.New()
	router.GET("/", healthCheckHandler)
	router.POST("/user/v1/check", checkDegreeHandler)

	http.ListenAndServe(":8050", &Server{router})
}

/* comprehentHandler()- Reads the input
 */
func checkDegreeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var names *Request

	//Reading request from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
	}

	//Unmarshaling the request to stay struct
	err = json.Unmarshal(body, &names)
	if err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	
	res2D := Response{
		Status:  true,
		Message: "Successful",
		Data:    names.InputOne,
	}

	fmt.Fprintln(w, jSONResponse(res2D))
}

func main() {
	showSplashscreen()
	initializeRoutes()

}

func jSONResponse(resp Response) string {
	j, err := json.Marshal(&resp)
	if err != nil {
	}
	return string(j)
}
