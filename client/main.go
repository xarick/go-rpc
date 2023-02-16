package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"Birinchi", "Toxirov Bobur"}
	b := Item{"Ikkinchi", "Umarov Aziz"}
	c := Item{"Uchinchi", "Karimov Azamat"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("DB: ", db)

	client.Call("API.EditItem", Item{"Ikkinchi", "Umarov Azizbek"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("DB: ", db)

	client.Call("API.GetByName", "Birinchi", &reply)
	fmt.Println("birinchi malumot: ", reply)

}
