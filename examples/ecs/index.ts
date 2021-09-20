import * as aws_native from "@pulumipreview/aws-native";
import * as old from "./unsupported";

// A basic resource - good to play with.
const logGroup = new aws_native.logs.LogGroup("test", {
    logGroupName: "LGN1",
    retentionInDays: 90,
});

export const arn = logGroup.arn;

// A nice example of a resource that isn't supported by the classic provider.
const cassandraKeyspace = new aws_native.cassandra.Keyspace("cassandra", {
    keyspaceName: "cassandrademo",
});

const cluster = new aws_native.ecs.Cluster("cluster", {
    clusterName: "cloud-api-cluster",
});

const wl = new aws_native.elasticloadbalancingv2.Listener("web", {
    loadBalancerArn: old.albArn,
    port: 80,
    protocol: "HTTP",
    defaultActions: [{
        type: "forward",
        targetGroupArn: old.atgArn,
    }],
});

const taskDefinition = new aws_native.ecs.TaskDefinition("app-task", {
    family: "fargate-task-definition",
    cpu: "256",
    memory: "512",
    networkMode: "awsvpc",
    requiresCompatibilities: ["FARGATE"],
    executionRoleArn: old.roleArn,
    containerDefinitions:[{
        name: "my-app",
        image: "nginx",
        portMappings: [{
            containerPort: 80,
            hostPort: 80,
            protocol: "tcp"
        }],
    }],
});

const service = new aws_native.ecs.Service("app-svc", {
    serviceName: "app-svc-cloud-api",
    cluster: cluster.arn,
    desiredCount: 1,
    launchType: "FARGATE",
    taskDefinition: taskDefinition.taskDefinitionArn,
    networkConfiguration: {
        awsvpcConfiguration: {
            assignPublicIp: "ENABLED",
            subnets: old.subnetIds,
            securityGroups: [old.securityGroupId],
        },
    },
    loadBalancers: [{
        targetGroupArn: old.atgArn,
        containerName: "my-app",
        containerPort: 80,
    }],
}, { dependsOn: [wl] });

export const url = old.albDnsName;
