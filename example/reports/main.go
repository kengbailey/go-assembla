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

	spaces, err := client.Spaces.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	reports, err := client.Reports.GetCustomReportsBySpaceID(spaces[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(reports.TeamReports))
}
