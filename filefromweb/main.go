package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := string(bytes)
	tours := toursFromJson(data)
	for _, tour := range tours {
		fmt.Printf("%v : $%v\n", tour.Name, tour.Price)
	}
}

func toursFromJson(data string) []Tour {
	tours := make([]Tour, 0)
	decoder := json.NewDecoder(strings.NewReader(data))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name, Price string
}
