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
	http.HandleFunc("/remove", checkout.removeHandler)
	http.HandleFunc("/total", checkout.totalHandler)

	http.ListenAndServe(":8080", nil)
}

// addHandler handles request for adding ad
func (checkout *Checkout) addHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if validateQuery(query, "customer", "type") {
		item := Item{
			StringToAdTypes(query["type"][0]),
		}
		checkout.Add(item)
		fmt.Fprintf(w, checkout.Show())
	} else {
		w.WriteHeader(400)
		fmt.Fprintln(w, "malformed query")
	}
}

// removeHandler handles request for removing ad
func (checkout *Checkout) removeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if validateQuery(query, "customer", "type") {
		item := Item{
			StringToAdTypes(query["type"][0]),
		}
		checkout.Remove(item)
		fmt.Fprintf(w, checkout.Show())
	} else {
		w.WriteHeader(400)
		fmt.Fprintln(w, "malformed query")
	}
}

// totalHandler handles request for total price
func (checkout *Checkout) totalHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if validateQuery(query, "customer") {
		fmt.Fprintf(w, checkout.ShowTotal())
	} else {
		w.WriteHeader(400)
		fmt.Fprintln(w, "malformed query")
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
