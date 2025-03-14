package main

import (
	"strings"

	metadata "github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
)

// Instead of humans playing the game, let the computer assign labels based on heuristics that seem to be fairly robust.
// This should add a label in the data indicating that it was assigned by a heuristic, in case this data point ends up
// being erroneous and need correction later.
func autoLabel(schemaAbsPath, dbFile string, allResources map[string]resourceFile) error {
	db := new(metadata.RefDB)
	if err := db.LoadJSONFile(dbFile); err != nil {
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
		// primaryIdentifier is "arn" and docs indicate Ref returns the ARN.
		case Categorize(res.RefSection).Name() == RefReturnsArn.Name() &&
			len(sch.PrimaryIdentifier) == 1 &&
			strings.Contains(strings.ToLower(sch.PrimaryIdentifier[0]), "arn"):

			if err := gs.edit(r, "heuristic", func(ri *metadata.RefDBResource) {
				ri.RefReturns.Property = strings.TrimPrefix(sch.PrimaryIdentifier[0], "/properties/")
				ri.RefReturns.Heuristic = RefReturnsArn.Name()
			}); err != nil {
				return err
			}
		// there is only one property called "ARN" or "FooARN" and docs indicate Ref returns the ARN.
		case Categorize(res.RefSection).Name() == RefReturnsArn.Name():
			if nameProp, ok := findUniqueProperty(sch, r, "ARN"); ok {
				if err := gs.edit(r, "heuristic", func(ri *metadata.RefDBResource) {
					ri.RefReturns.Property = nameProp
					ri.RefReturns.Heuristic = RefReturnsArn.Name()
				}); err != nil {
					return err
				}
			}
		// there is only one property called "Name" or "FooName" and docs indicate Ref returns the Name.
		case Categorize(res.RefSection).Name() == RefReturnsName.Name():
			if nameProp, ok := findUniqueProperty(sch, r, "Name"); ok {
				if err := gs.edit(r, "heuristic", func(ri *metadata.RefDBResource) {
					ri.RefReturns.Property = nameProp
					ri.RefReturns.Heuristic = RefReturnsName.Name()
				}); err != nil {
					return err
				}
			}
		// there is a property called "Id" or "FooId" and docs indicate Ref returns the ID.
		case Categorize(res.RefSection).Name() == RefReturnsID.Name():
			idProp := "Id"
			_, hasID := sch.Properties[idProp]
			if !hasID {
				idProp = simpleResName(r) + "Id"
				_, hasID = sch.Properties[idProp]
			}

			if hasID {
				if err := gs.edit(r, "heuristic", func(ri *metadata.RefDBResource) {
					ri.RefReturns.Property = idProp
					ri.RefReturns.Heuristic = RefReturnsID.Name()
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func simpleResName(r string) string {
	parts := strings.Split(r, "::")
	return parts[len(parts)-1]
}

func findUniqueProperty(s *schema, resourceName, propertyName string) (string, bool) {
	count := 0
	found := ""
	for p := range s.Properties {
		if p == propertyName {
			found = p
			count++
		}
		if p == simpleResName(resourceName)+propertyName {
			found = p
			count++
		}
	}
	if count == 1 {
		return found, true
	}
	return "", false
}
