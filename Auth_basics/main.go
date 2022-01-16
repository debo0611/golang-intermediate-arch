package main

import (
	"encoding/json"
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

	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bs))

	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("back into Go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	json.NewEncoder(w).Encode(p1)
}

func bar(w http.ResponseWriter, r *http.Request) {
	pass
}
