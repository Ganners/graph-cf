package main

import (
	"fmt"
	"log"
	"net/http"

	cf "github.com/Ganners/graph-cf"
)

func main() {
	ratings, err := cf.RatingsFromCSV("data/fixture.csv")
	if err != nil {
		log.Fatalf("unable to load csv: %v", err)
	}
	graph := cf.NewSimpleGraphCF()
	graph.AddRatings(ratings)

	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		if user == "" {
			http.Error(w, "user not specified", http.StatusBadRequest)
			return
		}

		topK, err := graph.UserTopK(user, 5)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("computed topK for %s: %#v\n", user, topK)

		fmt.Fprintf(w, ratingsToRecString(topK))
	})

	log.Println("visit http://localhost:8080/predict?user=John")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ratingsToRecString(ratings cf.Ratings) string {
	recString := ""
	numRatings := len(ratings)
	for i, r := range ratings {
		recString += fmt.Sprintf("[%s,%.5f]", r.Item, r.Rating)
		if i < numRatings-1 {
			recString += ","
		}
	}
	recString = "[" + recString + "]"
	return recString
}
