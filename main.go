package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("PRINT JSON", string(bs))

	xp2 := []person{}
	if err = json.Unmarshal(bs, &xp2); err != nil {
		log.Panic(err)
	}
	fmt.Println("back into Go data structure", xp2)

}
