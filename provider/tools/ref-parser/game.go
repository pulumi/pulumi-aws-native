package main

import (
	"fmt"
	"sort"
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
	for step := 1; true; step++ {
		r, ok := gs.nextResource()
		if !ok {
			fmt.Println("No resources left to map")
			return nil
		}

		fmt.Printf("Step %d: %d/%d resources done\n", step, gs.labelledResourceCount(), gs.totalResourceCount())

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

		primaryIdentifierIndex := func(prop string, notFound int) int {
			for i, p := range sch.PrimaryIdentifier {
				if strings.TrimPrefix(p, "/properties/") == prop {
					return i
				}
			}
			return notFound
		}

		// Going to reorder properties so that properties that are part of the PrimaryIdentifier come first.
		// This makes it easier for the player to select these as the answer, as they are very commonly the
		// answer.
		sort.Slice(props, func(i, j int) bool {
			if primaryIdentifierIndex(props[i], 1000) < primaryIdentifierIndex(props[j], 1000) {
				return true
			}
			return props[i] < props[j]
		})

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
		rInfo.RefReturns.Property = prop
		db.Resources[r] = rInfo

		if err := db.store(dbFile); err != nil {
			return err
		}
		fmt.Println("Stored your choice in the database: ", prop)
	}

	return nil
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

func (gs *gameState) totalResourceCount() int {
	return len(gs.allResources)
}

func (gs *gameState) labelledResourceCount() int {
	return len(gs.db.Resources)
}
