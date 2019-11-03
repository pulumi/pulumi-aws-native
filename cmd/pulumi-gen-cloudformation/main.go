package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pulumi/pulumi-cloudformation/pkg/gen"
	"github.com/pulumi/pulumi-cloudformation/pkg/schema"
)

func main() {
	var sch schema.CloudFormationSchema
	if err := json.NewDecoder(os.Stdin).Decode(&sch); err != nil {
		log.Fatalf("failed to decode CloudFormation schema: %v", err)
	}
	if err := gen.GenerateSDK(os.Stdout, sch); err != nil {
		log.Fatalf("failed to generate SDK: %v", err)
	}
}
