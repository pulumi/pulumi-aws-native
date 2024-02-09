# Automated Generation Reports

The purpose of these reports are to highlight information from within the generation process and allow tracking of how these change over time.

## Unexpected Tags Shapes

This report lists each resource which has a field called "Tags" but it doesn't match one of our expected heuristics for its shape (either a list of Tags objects or a map of string to string). The key is the Pulumi resource token and the value is the original JSON schema for the Tags property.
