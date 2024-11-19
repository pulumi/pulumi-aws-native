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
	flag.Parse()

	if guide == nil {
		log.Fatal("-guide is required")
	}

	guideAbsPath, err := filepath.Abs(*guide)
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

	fmt.Println("Ref sections by category:")
	for sec, n := range refCategoryCounts {
		name := sec
		if sec == "" {
			name = "Uncategorized"
		}
		fmt.Println(name, n)
	}

	fmt.Println("Sampling of un-categorized files")
	for _, x := range parsedFiles[0:100] {
		if x.Category.Name() != Uncategorized.Name() {
			continue
		}
		fmt.Printf("%q\n\n", x.RefSection)
	}

	fmt.Println(len(rFiles))
}

type resourceFile struct {
	RefSection string // optional
	Category   Category
}

var refPattern = regexp.MustCompile(`(?m)^[#]+[ ]*Ref[^\n]*[\n]([^#]+)`)

func parseResourceFile(file string) (resourceFile, error) {
	fbytes, err := os.ReadFile(file)
	if err != nil {
		return resourceFile{}, err
	}
	rf := resourceFile{}

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
