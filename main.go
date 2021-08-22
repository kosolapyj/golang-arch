package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{
	// 	First: "Jenny",
	// }

	// p2 := person{
	// 	First: "James",
	// }

	// xp := []person{
	// 	p1,
	// 	p2,
	// }

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("JSON MARSHAL", string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)

	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("BACK TO SLICE", xp2)
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{
		p1,
		p2,
	}
	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var xp []person
	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("decoded bad data", err)
	}
	log.Println("Decoded p", xp)
}
