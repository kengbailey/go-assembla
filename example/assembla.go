package assembla

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
	// TODO: add this to new client behavior
	spaces, err := client.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, space := range spaces {
		fmt.Println(space.Name)
		// get Active/followed tickets
		tickets, err := client.GetFollowedTicketsBySpace(space.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("--> Tickets %d\n", len(tickets))
		for _, ticket := range tickets {
			fmt.Println(ticket.Summary)

			// get ticket comments
			comments, err := client.GetTicketComments(space.ID, ticket.Number)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("----> Comments %d\n", len(comments))

		}
		fmt.Println()
	}
}
