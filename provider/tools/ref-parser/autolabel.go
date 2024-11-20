package main

import "strings"

// Instead of humans playing the game, let the computer assign labels based on heuristics that seem to be fairly robust.
// This should add a label in the data indicating that it was assigned by a heuristic, in case this data point ends up
// being erroneous and need correction later.
func autoLabel(schemaAbsPath, dbFile string, allResources map[string]resourceFile) error {
	db := new(db)
	if err := db.load(dbFile); err != nil {
		return err
	}
	gs := &gameState{
		db:           db,
		allResources: allResources,
		dbFile:       dbFile,
	}
	for r, res := range allResources {

		if _, seen := db.Resources[r]; seen {
			continue
		}

		sch, err := findSchema(schemaAbsPath, r)
		if err != nil {
			continue
		}

		switch {
		case Categorize(res.RefSection).Name() == RefReturnsArn.Name() &&
			len(sch.PrimaryIdentifier) == 1 &&
			strings.ToLower(sch.PrimaryIdentifier[0]) == "/properties/arn":

			if err := gs.edit(r, "heuristic", func(ri *resourceInfo) {
				ri.RefReturns.Property = strings.TrimPrefix(sch.PrimaryIdentifier[0], "/properties/arn")
				ri.RefReturns.Heuristic = RefReturnsArn.Name()
			}); err != nil {
				return err
			}
		}

	}
	return nil
}
