package signal

import (
	"encoding/json"
	"net/http"
)

// Pick will return fist element from provided array
func Pick(numbers []int) int {
	return numbers[0]
}

// Person holds personal information about the user
type Person struct {
	Name    string
	Address string
	Email   string
}

// Handler will take a http response handler and will write the content
// into it
func Handler(w http.ResponseWriter, r *http.Request) {
	p := &Person{
		Name:    "Sam",
		Address: "128-138 High Road",
		Email:   "sam@gmail.com",
	}
	jsonB, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonB)
}
