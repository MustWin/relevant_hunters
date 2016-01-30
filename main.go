package main

import (
	"fmt"
	"os"

	"github.com/MustWin/relevant_hunters/producthunter/api"
)

func main() {
	coll, err := api.GetCollection(35215)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Collection:\n%+v\n", coll)
}
