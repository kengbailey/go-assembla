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
	client, err := NewClient(key, secret)
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
	// users, err := client.GetSpaceByID(spaces[0].ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(users.Name)

	// test get custom reports by spaceid
	// reports, err := client.GetCustomReportsBySpaceID(spaces[0].ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(reports.TeamReports[0].ID)

	// get tickets by report space id -- FINISH
	tickets, err := client.GetTicketsBySpaceAndReport(0, spaces[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range tickets {
		fmt.Println(t.Summary)
	}
}
