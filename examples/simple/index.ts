import * as pulumi from "@pulumi/pulumi";
import * as aws_native from "@pulumi/aws-native";

const logGroup = new aws_native.logs.LogGroup("test", {
    properties: {
        LogGroupName: "LGN1",
        RetentionInDays: 90,
    },
});

export const value = logGroup.properties.RetentionInDays;
