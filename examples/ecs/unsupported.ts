import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as aws_native from "@pulumipreview/aws-native";

const defaultVpc = pulumi.output(aws.ec2.getVpc({ default: true }));
const defaultVpcSubnets = defaultVpc.id.apply(id => aws.ec2.getSubnetIds({vpcId: id}));

const group = new aws_native.ec2.SecurityGroup("websecgrp", {
	vpcId: defaultVpc.id,
	groupDescription: "Enable HTTP access",
	securityGroupIngress: [{
		ipProtocol: "tcp",
		fromPort: 80,
		toPort: 80,
		cidrIp: "0.0.0.0/0",
    }],
	securityGroupEgress: [{
		ipProtocol:"-1",
		fromPort: 0,
		toPort: 0,
		cidrIp: "0.0.0.0/0",
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

const role = new aws_native.iam.Role("taskexecrole", {
	roleName: "task-exec-role-37273",
	assumeRolePolicyDocument: JSON.stringify({
		Version: "2008-10-17",
		Statement: [{
			Sid: "",
			Effect: "Allow",
			Principal: {
				Service: "ecs-tasks.amazonaws.com"
			},
			Action: "sts:AssumeRole",
		}],
	}),
});

const rpa = new aws.iam.RolePolicyAttachment("task-exec-policy", {
	role: role.roleName.apply(v => v!),
	policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});

export const albArn = alb.arn;
export const albDnsName = alb.dnsName;
export const atgArn = atg.arn;
export const roleArn = role.arn;
export const subnetIds = defaultVpcSubnets.ids;
export const securityGroupId = group.id;
