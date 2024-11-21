package main

import (
	"strings"

	"fmt"
	metadata "github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
)

func reportOnRefAsPrimaryIdentifier(dbFile, schemaAbsPath string) error {

	db := new(metadata.RefDB)
	if err := db.LoadJSONFile(dbFile); err != nil {
		return err
	}

	refAsPID := map[string]struct{}{}

	for resID, res := range db.Resources {
		sch, err := findSchema(schemaAbsPath, resID)
		if err != nil {
			return err
		}

		switch {
		case res.RefReturns.Property != "":
			if len(sch.PrimaryIdentifier) == 1 {
				if strings.TrimPrefix(sch.PrimaryIdentifier[0], "/properties/") == res.RefReturns.Property {
					refAsPID[resID] = struct{}{}
				}
			}
		case len(res.RefReturns.Properties) == len(sch.PrimaryIdentifier):
			for i, p := range res.RefReturns.Properties {
				if strings.TrimPrefix(sch.PrimaryIdentifier[i], "/properties/") == p {
					refAsPID[resID] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("Total resources with schematized Ref behavior: %d\n", len(db.Resources))
	fmt.Printf("Resources where Ref returns the primary identifier: %d\n", len(refAsPID))

	for refID, sr := range db.Resources {
		if _, ok := refAsPID[refID]; ok {
			continue
		}

		sch, err := findSchema(schemaAbsPath, refID)
		if err != nil {
			return err
		}

		if sr.RefReturns.CfRefBehavior.Property != "" {
			fmt.Printf("Strange resource: %s --> %q (pid=#%v)\n", refID,
				sr.RefReturns.CfRefBehavior.Property,
				strings.Join(sch.PrimaryIdentifier, "|"))
		} else {
			fmt.Printf("Strange resource: %s --> %#v\n", refID, sr.RefReturns.CfRefBehavior)

		}
	}

	return nil
}
