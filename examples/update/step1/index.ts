// Copyright 2016-2021, Pulumi Corporation.
import * as pulumi from "@pulumi/pulumi";
import * as awsClassic from "@pulumi/aws";
import * as aws from "@pulumi/aws-native";

const logGroup = new aws.s3.Bucket("bucket", {
    tags: [{
        key: "foo",
        value: "bar",
    }]
});

const lambdaRole = new awsClassic.iam.Role("lambdaRole", {
    assumeRolePolicy: awsClassic.iam.assumeRolePolicyForPrincipal({ Service: "lambda.amazonaws.com" }),
    inlinePolicies: [{
        name: "lambdaAccess", policy: JSON.stringify({
            Version: "2012-10-17",
            Statement: [{
                Effect: "Allow",
                Action: [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                ],
                Resource: "arn:aws:logs:*:*:*",
            }],
        })
    }]
});

class DelayResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: { ms: number }, opts?: pulumi.CustomResourceOptions) {
        const delay = props.ms;
        super({
            create: async (inputs) => {
                await new Promise(resolve => setTimeout(resolve, delay))
                return { id: "delay", outs: {} };
            }
        }, name, props ?? {}, opts,);
    }
}

const func = new aws.lambda.Function("function", {
    runtime: "nodejs12.x",
    role: lambdaRole.arn,
    handler: "index.handler",
    code: {
        zipFile: `
exports.handler =  async function(event, context) {
  console.log("Running")
  return "foo"
}
`,
    },
}, { dependsOn: new DelayResource("delay", { ms: 60 * 1000 }, { dependsOn: lambdaRole }) });
