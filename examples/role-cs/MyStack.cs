// Copyright 2021, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Pulumi.AwsNative.Iam;
using IamInputs = Pulumi.AwsNative.Iam.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var testRole = new Role("role-test", new RoleArgs
        {
            AssumeRolePolicyDocument = """
            {
                "Version": "2012-10-17",
                "Statement": [{
                    "Sid": "",
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "ec2.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }]
            }
            """,
            Policies = new[]
            {
                new IamInputs.RolePolicyArgs
                {
                    PolicyName = "test-policy",
                    PolicyDocument = new Dictionary<string, object?>
                    {
                        ["Version"] = "2012-10-17",
                        ["Statement"] = new List<object?>
                        {
                            new Dictionary<string, object?>
                            {
                                ["Effect"] = "Allow",
                                ["Action"] = "*",
                                ["Resource"] = "*",
                            },
                        },
                    },
                },
            },
        });

        this.Arn = testRole.Arn;
        this.Policy = testRole.Policies.Apply(policies => policies[0].PolicyDocument);
    }

    [Output]
    public Output<string> Arn { get; set; }
    [Output]
    public Output<object> Policy { get; set; }
}
