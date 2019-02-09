package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kengbailey/go-assembla/assembla"
)

func main() {

	// vars
	key := os.Getenv("ASSEMBLA_KEY")
	secret := os.Getenv("ASSEMBLA_SECRET")

	// create client
	client := assembla.NewClient(key, secret)

	// get spaces
	spaces, err := client.Spaces.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, space := range spaces {
		fmt.Println(space.Name)
		// get Active/followed tickets
		tickets, err := client.Tickets.GetFollowedTicketsBySpace(space.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("--> Tickets %d\n", len(tickets))
		for _, ticket := range tickets {
			fmt.Println(ticket.Summary)

			// get ticket comments
			comments, err := client.Comments.GetTicketComments(space.ID, ticket.Number)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("----> Comments %d\n", len(comments))

		}
		fmt.Println()
	}
}
