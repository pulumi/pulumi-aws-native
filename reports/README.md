# Automated Generation Reports

The purpose of these reports are to highlight information from within the generation process and allow tracking of how these change over time.

## Unexpected Tags Shapes

This report lists each resource which has a field called "Tags" but it doesn't match one of our expected heuristics for its shape (either a list of Tags objects or a map of string to string). The key is the Pulumi resource token and the value is the original JSON schema for the Tags property.

## Missed Autonaming

This report lists each resource for which our autonaming has not managed to identify a field which is the logical "name" and therefore we won't be able to automatically create a name based on the Pulumi resource name. There are many resources where this is a valid limitation e.g. it's identifier is composed of two other resource's identifiers or it doesn't contain any kind of user-facing name.
