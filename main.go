package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	log.Println("Server is starting...")

	checkout := Checkout{
		pricingRules: GetRules("default"),
	}

	http.HandleFunc("/add", checkout.addHandler)
	http.HandleFunc("/delete", checkout.deleteHandler)
	http.HandleFunc("/total", checkout.totalHandler)

	http.ListenAndServe(":8080", nil)
}

// addHandler handles request for adding ad
func (checkout *Checkout) addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		log.Println(r.Method, "Invalid request method")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var jsonBodyItem struct {
			ItemType string
		}
		error := json.Unmarshal(body, &jsonBodyItem)

		if error != nil || !validItemText(jsonBodyItem.ItemType) {
			w.WriteHeader(405)
			log.Println("There was an error with the json input:", err)
		} else {
			item := Item{
				StringToAdTypes(jsonBodyItem.ItemType),
			}
			checkout.Add(item)
			fmt.Fprintf(w, checkout.Show())
		}
	}
}

// deleteHandler handles request for removing ad
func (checkout *Checkout) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		log.Println(r.Method, "Invalid request method")
	} else {
		query := r.URL.Query()
		if validateQuery(query, "customer", "type") {
			item := Item{
				StringToAdTypes(query["type"][0]),
			}
			checkout.Delete(item)
			fmt.Fprintf(w, checkout.Show())
		} else {
			w.WriteHeader(400)
			log.Println(query, "Malformed query")
		}
	}
}

// totalHandler handles request for total price
func (checkout *Checkout) totalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		log.Println(r.Method, "Invalid request method")
	} else {
		query := r.URL.Query()
		if validateQuery(query, "customer") {
			fmt.Fprintf(w, checkout.ShowTotal())
		} else {
			w.WriteHeader(400)
			log.Println(w, "Malformed query")
		}
	}
}

// validateQuery checks query values are correct
func validateQuery(v url.Values, args ...string) bool {
	for _, i := range args {
		_, queryPresent := v[i]
		if !queryPresent {
			return false
		}
	}
	return true
}

// validItemText checks that the item added is one of the known job types
func validItemText(item string) bool {
	for _, value := range AdTypesStrings {
		if value == item {
			return true
		}
	}
	return false
}
