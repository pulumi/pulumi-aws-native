package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	guide := flag.String("guide", "", "path to a folder with Cloud Formation user guide")
	schema := flag.String("schema", "", "path to a folder with Cloud Formation schema JSON files")
	dbFile := flag.String("db", "", "path to a database to play a guessing game")
	flag.Parse()

	if guide == nil || *guide == "" {
		log.Fatal("-guide is required")
	}

	if schema == nil || *schema == "" {
		log.Fatal("-schema is required")
	}

	guideAbsPath, err := filepath.Abs(*guide)
	if err != nil {
		log.Fatal(err)
	}

	schemaAbsPath, err := filepath.Abs(*schema)
	if err != nil {
		log.Fatal(err)
	}

	rFiles, err := resourceFiles(guideAbsPath)
	if err != nil {
		log.Fatal(err)
	}

	refCategoryCounts := map[string]int{}

	parsedFiles := []resourceFile{}

	for _, rfLoc := range rFiles {
		rf, err := parseResourceFile(rfLoc)
		if err != nil {
			log.Fatal(err)
		}
		cn := rf.Category.Name()
		refCategoryCounts[cn] = refCategoryCounts[cn] + 1
		parsedFiles = append(parsedFiles, rf)
	}

	if dbFile != nil {
		allResources := map[string]resourceFile{}
		for _, rf := range parsedFiles {
			allResources[rf.ResourceID] = rf
		}
		if err := game(schemaAbsPath, *dbFile, allResources); err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("Ref sections by category:")
	for sec, n := range refCategoryCounts {
		name := sec
		if sec == "" {
			name = "Uncategorized"
		}
		fmt.Println(name, n)
	}

	fmt.Println("Sampling of un-categorized files")
	for _, x := range parsedFiles {
		if x.Category.Name() != Uncategorized.Name() {
			continue
		}

		props, err := findProperties(schemaAbsPath, x.ResourceID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ResourceID: %s\n", x.ResourceID)
		fmt.Printf("Properties: %s\n", strings.Join(props, ", "))
		fmt.Printf("RefSection: %s\n\n", x.RefSection)
	}

	fmt.Println(len(rFiles))
}

type resourceFile struct {
	ResourceID string // e.g. AWS::EC2::VPC
	RefSection string // optional
	Category   Category
}

var resourceIdPattern = regexp.MustCompile(`^[#]\s*([^<\n]+)`)
var refPattern = regexp.MustCompile(`(?m)^[#]+[ ]*Ref[^\n]*[\n]([^#]+)`)

func parseResourceFile(file string) (resourceFile, error) {
	fbytes, err := os.ReadFile(file)
	if err != nil {
		return resourceFile{}, err
	}
	rf := resourceFile{}

	if m := resourceIdPattern.FindSubmatch(fbytes); m != nil {
		rf.ResourceID = string(m[1])
	}
	if rf.ResourceID == "" {
		return resourceFile{}, fmt.Errorf("Could not parse ResourceID from %q", file)
	}

	if m := refPattern.FindSubmatch(fbytes); m != nil {
		rf.RefSection = string(m[1])
		rf.Category = Categorize(rf.RefSection)
	} else {
		rf.Category = Undocumented
	}

	return rf, nil
}

func resourceFiles(guide string) ([]string, error) {
	var result []string
	d := filepath.Join(guide, "doc_source")
	entries, err := os.ReadDir(d)
	if err != nil {
		return nil, err
	}
	for _, file := range entries {
		if file.IsDir() {
			continue
		}
		if !strings.HasPrefix(file.Name(), "aws-resource") {
			continue
		}
		result = append(result, filepath.Join(d, file.Name()))
	}
	return result, nil
}
