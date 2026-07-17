import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";

const config = new pulumi.Config();
const reverseEnvironment = config.getBoolean("reverseEnvironment") ?? false;
const environment = reverseEnvironment
  ? [{ name: "A_VAR", value: "a" }, { name: "Z_VAR", value: "z" }]
  : [{ name: "Z_VAR", value: "z" }, { name: "A_VAR", value: "a" }];

const role = new aws.iam.Role("unordered-policies", {
  assumeRolePolicyDocument: {
    Version: "2012-10-17",
    Statement: [{
      Effect: "Allow",
      Principal: { Service: "lambda.amazonaws.com" },
      Action: "sts:AssumeRole",
    }],
  },
  // IAM returns inline policy names alphabetically. Deliberately author them
  // in reverse order so refresh observes an order-only CloudControl change.
  policies: [
    {
      policyName: "z-policy",
      policyDocument: {
        Version: "2012-10-17",
        Statement: [{
          Effect: "Allow",
          Action: "logs:DescribeLogGroups",
          Resource: "*",
        }],
      },
    },
    {
      policyName: "a-policy",
      policyDocument: {
        Version: "2012-10-17",
        Statement: [{
          Effect: "Allow",
          Action: "logs:DescribeLogStreams",
          Resource: "*",
        }],
      },
    },
  ],
});

const taskDefinition = new aws.ecs.TaskDefinition("nested-unordered-environment", {
  containerDefinitions: [{
    name: "app",
    image: "public.ecr.aws/docker/library/busybox:latest",
    memory: 128,
    essential: true,
    // Exercise the nested unordered path
    // containerDefinitions/*/environment. The live test changes only this
    // ordering after refresh and expects a no-change preview.
    environment,
  }],
});

export const roleName = role.roleName;
export const taskDefinitionArn = taskDefinition.taskDefinitionArn;
