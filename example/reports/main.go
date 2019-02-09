package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kengbailey/go-assembla/assembla"
)

func main() {

	key := os.Getenv("ASSEMBLA_KEY")
	secret := os.Getenv("ASSEMBLA_SECRET")

	client := assembla.NewClient(key, secret)

	spaces, err := client.Spaces.GetUserSpaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, space := range spaces {
		reports, err := client.Reports.GetCustomReportsBySpaceID(space.ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(space.Name)
		fmt.Printf("Found Team Reports: %d\n", len(reports.TeamReports))
		fmt.Printf("Found User Reports: %d\n\n", len(reports.UserReports))
	}
}
