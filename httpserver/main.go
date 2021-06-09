package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Request model for math operations
type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

//Response model for math operations
type MathResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

//create handler for math operations
func mathHandler(w http.ResponseWriter, r *http.Request) {
	//Decode request
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	request := &MathRequest{}
	if err := decoder.Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Do work
	response := &MathResponse{}
	switch request.Op {
	case "+":
		response.Result = request.Left + request.Right
	case "-":
		response.Result = request.Left - request.Right
	case "*":
		response.Result = request.Left * request.Right
	case "/":
		if request.Right == 0.0 {
			response.Error = "Invalid math operation, divide by 0"
		} else {
			response.Result = request.Left / request.Right
		}
	default:
		response.Error = "Unknown operation."
	}

	w.Header().Set("Content-Type", "application/json")
	if response.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	//Encode response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Print("Cannot encode respose object.")
	}

}

//Hello handler returns greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/math", mathHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server started.")
	}

}
