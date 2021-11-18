// Copyright 2016-2021, Pulumi Corporation.

import * as aws from "@pulumi/aws-native";

const logGroup = new aws.s3.Bucket("bucket", {
    tags: [{
        key: "foo",
        value: "buzz", // <-- this value has changed
    }]
});
