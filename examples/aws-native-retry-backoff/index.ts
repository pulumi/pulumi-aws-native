import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";


for (let i = 0; i < 100; i++) {
    new aws.iam.Role(`role-${i}`, {
        assumeRolePolicyDocument: {
            version: "2012-10-17",
            statement: [{
                effect: "Allow",
                principal: {
                    service: ["grafana.amazonaws.com"],
                },
                action: ["sts:AssumeRole"],
            }],
        },
    });

}
