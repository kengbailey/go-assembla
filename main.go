package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// vars
	key := os.Getenv("ASSEMBLA_KEY")
	secret := os.Getenv("ASSEMBLA_SECRET")

	// Make request
	var client AssemblaClient
	err := client.Connect(key, secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	// get spaces
	spaces, err := client.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	// get users by space di
	users, err := client.GetSpaceByID(spaces[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users.Name)

}
