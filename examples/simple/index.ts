import * as aws_native from "@pulumi/aws-native";
import * as old from "./unsupported";

const logGroup = new aws_native.logs.LogGroup("test", {
    LogGroupName: "LGN1",
    RetentionInDays: 90,
});

export const arn = logGroup.Arn;

const cluster = new aws_native.ecs.Cluster("cluster", {
    ClusterName: "cloud-api-cluster",
});

const wl = new aws_native.elasticloadbalancingv2.Listener("web", {
    LoadBalancerArn: old.albArn,
    Port: 80,
    Protocol: "HTTP",
    DefaultActions: [{
        Type: "forward",
        TargetGroupArn: old.atgArn,
    }],
});

const taskDefinition = new aws_native.ecs.TaskDefinition("app-task", {
    Family: "fargate-task-definition",
    Cpu: "256",
    Memory: "512",
    NetworkMode: "awsvpc",
    RequiresCompatibilities: ["FARGATE"],
    ExecutionRoleArn: old.roleArn,
    ContainerDefinitions:[{
        Name: "my-app",
        Image: "nginx",
        PortMappings: [{
            ContainerPort: 80,
            HostPort: 80,
            Protocol: "tcp"
        }],
    }],
});

const service = new aws_native.ecs.Service("app-svc", {
    ServiceName: "app-svc-cloud-api",
    Cluster: cluster.Arn,
    DesiredCount: 3,
    LaunchType: "FARGATE",
    TaskDefinition: taskDefinition.TaskDefinitionArn,
    NetworkConfiguration: {
        AwsvpcConfiguration: {
            AssignPublicIp: "ENABLED",
            Subnets: old.subnetIds,
            SecurityGroups: [old.securityGroupId],
        },
    },
    LoadBalancers: [{
        TargetGroupArn: old.atgArn,
        ContainerName: "my-app",
        ContainerPort: 80,
    }],
}, { dependsOn: [wl] });

export const url = old.albDnsName;
