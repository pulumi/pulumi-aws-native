// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";

const logGroup = new aws.logs.LogGroup("log-test", {
  retentionInDays: 90,
});

// Wait on the id to be populated so we don't fetch the log group before it's created.
// The name is immediately available because Pulumi generates it during the preview.
const fetchedLogGroup = aws.logs.getLogGroupOutput({
  logGroupName: logGroup.id.apply(() =>
    logGroup.logGroupName.apply((n) => n ?? "")
  ),
});

export const arns = pulumi
  .all([logGroup.arn, fetchedLogGroup.arn])
  .apply((arns) => {
    if (arns[0] !== arns[1]) {
      throw new Error("Fetched ARN doesn't match created ARN");
    }
    return arns;
  });
export const arn = logGroup.arn;
