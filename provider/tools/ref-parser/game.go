package main

import (
	"fmt"
	"strings"
)

// Run the game loop where the user is trying to guess what CloudFormation Ref intrinsic function will return and record
// this guess in a database.
func game(schemaAbsPath, dbFile string, allResources map[string]resourceFile) error {
	db := new(db)
	if err := db.load(dbFile); err != nil {
		return err
	}
	gs := &gameState{
		db:           db,
		allResources: allResources,
	}
	for {
		r, ok := gs.nextResource()
		if !ok {
			fmt.Println("No resources left to map")
			return nil
		}

		res, ok := allResources[r]
		if !ok {
			return fmt.Errorf("Resource not found: %s", res)
		}

		props, err := findProperties(schemaAbsPath, r)
		if err != nil {
			return err
		}

		sch, err := findSchema(schemaAbsPath, r)
		if err != nil {
			return err
		}

		fmt.Println("ResourceID: ", r)
		fmt.Println("PrimaryID : ", strings.Join(sch.PrimaryIdentifier, " "))
		fmt.Println("RefSection: ", res.RefSection)
		fmt.Println("Enter one of the following: ")

		fmt.Printf("%3d: It is unclear what Ref would return\n", 0)

		for i, p := range props {
			fmt.Printf("%3d: Ref returns the %q property\n", i+1, p)
		}

		fmt.Println("Your choice:")

		var choice int
		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			return err
		}

		if choice == 0 {
			continue
		}

		prop, ok := props[choice-1], ok
		if !ok {
			continue
		}

		rInfo := db.Resources[r]
		rInfo.RefReturns.Property = prop1
		db.Resources[r] = rInfo

		if err := db.store(dbFile); err != nil {
			return err
		}
		fmt.Println("Stored your choice in the database: ", prop)
	}
}

type gameState struct {
	db           *db
	allResources map[string]resourceFile
}

func (gs *gameState) nextResource() (string, bool) {
	// Non-deterministic here to make it more fun.
	for r := range gs.allResources {
		if _, seen := gs.db.Resources[r]; seen {
			continue
		}
		if strings.TrimSpace(gs.allResources[r].RefSection) == "" {
			continue
		}
		return r, true
	}
	return "", false
}
