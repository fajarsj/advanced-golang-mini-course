// <--- Creating API --->
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type article struct {
// 	ID int
// 	Title string
// 	Content string
// }

// var data = []article{
// 	article{1, "Lorem", "lorem"},
// 	article{2, "Ipsum", "ipsum"},
// }

// func articles (w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(data)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(result)
// 		return
// 	}

// 	http.Error(w, " ", http.StatusBadRequest)
// }

// func main() {
// 	http.HandleFunc("/articles", articles)
// 	fmt.Println("starting web server at http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

// <--- Consuming RESTful API --->
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A response struct to map the entire esponse
type StarWarsPeople struct {
	Name string `json:"name"`
	Height string `json:"height"`
	Mass string `json:"mass"`
	HairColor string `json:"hair_color"`
}

func main() {
	response, _ := http.Get("https://swapi.dev/api/people/1/")
	
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var LukeSkywalker StarWarsPeople
	json.Unmarshal(responseData, &LukeSkywalker)

	fmt.Println("--- Print Field ---")
	fmt.Println(LukeSkywalker.Name)
	fmt.Println(LukeSkywalker.Height)
	fmt.Println(LukeSkywalker.Mass)
	fmt.Println(LukeSkywalker.HairColor)
}