package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

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
		fmt.Fprintln(w, "Invalid request method")
	} else {
		query := r.URL.Query()
		if validateQuery(query, "customer", "type") {
			item := Item{
				StringToAdTypes(query["type"][0]),
			}
			checkout.Add(item)
			fmt.Fprintf(w, checkout.Show())
		} else {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Malformed query")
		}
	}
}

// deleteHandler handles request for removing ad
func (checkout *Checkout) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		fmt.Fprintln(w, "Invalid request method")
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
			fmt.Fprintln(w, "Malformed query")
		}
	}
}

// totalHandler handles request for total price
func (checkout *Checkout) totalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		fmt.Fprintln(w, "Invalid request method")
	} else {
		query := r.URL.Query()
		if validateQuery(query, "customer") {
			fmt.Fprintf(w, checkout.ShowTotal())
		} else {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Malformed query")
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
