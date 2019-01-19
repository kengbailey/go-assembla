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

	// create client
	client := NewAssemblaClient(key, secret)

	// get spaces
	spaces, err := client.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, space := range spaces {
		fmt.Println(space.Name)
		// get ACtive/followed tickets
		tickets, err := client.GetFollowedTicketsBySpace(space.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(len(tickets))
		for _, ticket := range tickets {
			fmt.Println(ticket.Summary)
		}
		fmt.Println()
	}
}
