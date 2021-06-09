import * as pulumi from "@pulumi/pulumi";
import { CdkComponent } from "./cdk-interop";

import * as cassandra from '@aws-cdk/aws-cassandra';

// A basic example of deploying a CDK resource from Pulumi.
// Note that we don't flow any inputs/outputs.
new CdkComponent("helloworld", stack => {
    new cassandra.CfnKeyspace(stack, "keyspace");
});

// The full program isn't that nice because we need to figure out a way to create
// dependencies across CDK resources and Pulumi components.
import * as ecs from '@aws-cdk/aws-ecs';
import * as elasticloadbalancingv2 from '@aws-cdk/aws-elasticloadbalancingv2';

import * as aws from "@pulumi/aws";

const defaultVpc = pulumi.output(aws.ec2.getVpc({ default: true }));
const defaultVpcSubnets = defaultVpc.id.apply(id => aws.ec2.getSubnetIds({vpcId: id}));

const group = new aws.ec2.SecurityGroup("web-secgrp", {
	vpcId: defaultVpc.id,
	description: "Enable HTTP access",
	ingress: [{
		protocol: "tcp",
		fromPort: 80,
		toPort: 80,
		cidrBlocks: ["0.0.0.0/0"],
    }],
  	egress: [{
		protocol:"-1",
		fromPort: 0,
		toPort: 0,
		cidrBlocks: ["0.0.0.0/0"],
    }],
});

const alb = new aws.lb.LoadBalancer("app-lb", {
	securityGroups: [group.id],
	subnets: defaultVpcSubnets.ids,
});

const atg = new aws.lb.TargetGroup("app-tg", {
    port: 80,
	protocol: "HTTP",
	targetType: "ip",
	vpcId: defaultVpc.id,
});

const role = new aws.iam.Role("task-exec-role", {
	assumeRolePolicy: {
		Version: "2008-10-17",
		Statement: [{
			Sid: "",
			Effect: "Allow",
			Principal: {
				Service: "ecs-tasks.amazonaws.com"
			},
			Action: "sts:AssumeRole",
		}],
	},
});

const rpa = new aws.iam.RolePolicyAttachment("task-exec-policy", {
	role: role.name,
	policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});

export const albArn = alb.arn;
export const albDnsName = alb.dnsName;
export const atgArn = atg.arn;
export const roleArn = role.arn;
export const subnetIds = defaultVpcSubnets.ids;
export const securityGroupId = group.id;

pulumi.all([albArn, atgArn, roleArn, subnetIds, securityGroupId]).apply(([albArn, atgArn, roleArn, subnetIds, securityGroupId]) => {
    new CdkComponent("test", stack => {
        new cassandra.CfnKeyspace(stack, "keyspace");

        const cluster = new ecs.CfnCluster(stack, "cluster");

        const wl = new elasticloadbalancingv2.CfnListener(stack, "web", {
            loadBalancerArn: albArn,
            port: 80,
            protocol: "HTTP",
            defaultActions: [{
                type: "forward",
                targetGroupArn: atgArn,
            }],
        });

        const taskDefinition = new ecs.CfnTaskDefinition(stack, "app-task", {
            family: "fargate-task-definition",
            cpu: "256",
            memory: "512",
            networkMode: "awsvpc",
            requiresCompatibilities: ["FARGATE"],
            executionRoleArn: roleArn,
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

        console.debug(cluster.attrArn);

        const service = new ecs.CfnService(stack, "app-svc", {
            serviceName: "app-svc-cloud-api",
            cluster: cluster.attrArn,
            desiredCount: 1,
            launchType: "FARGATE",
            taskDefinition: taskDefinition.attrTaskDefinitionArn,
            networkConfiguration: {
                awsvpcConfiguration: {
                    assignPublicIp: "ENABLED",
                    subnets: subnetIds,
                    securityGroups: [securityGroupId],
                },
            },
            loadBalancers: [{
                targetGroupArn: atgArn,
                containerName: "my-app",
                containerPort: 80,
            }],
        }/*, { dependsOn: [wl] }*/);
    });
});