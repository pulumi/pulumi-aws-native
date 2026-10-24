# ref-parser

Historical helper. The supported generation path is now `mise exec -- make ref-db`, using the `ref-db` operation in `pulumi-gen-aws-native`. See `meta/README.md`. Do not use this tool to overwrite the generated database.

This is a quick tool to assist with data cleaning tasks of building out the `ref-db.json` reference metadata on the
behavior of the Cloud Formation Ref intrinsic. This code itself is not used by the provider in production scenarios,
only the generated data is.
