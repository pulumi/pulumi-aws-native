// Copyright 2016-2024, Pulumi Corporation.

import * as pulumi from '@pulumi/pulumi';
import * as aws from "@pulumi/aws-native";

const cidrBlock = '192.168.0.0/24';
const vpc = new aws.ec2.Vpc('vpc', {
    cidrBlock,
});

const ipv6Cidr = new aws.ec2.VpcCidrBlock('ipv6Cidr', {
    vpcId: vpc.vpcId,
    amazonProvidedIpv6CidrBlock: true,
});


const cidrs = aws.cidrOutput({
    count: 4,
    ipBlock: cidrBlock,
    cidrBits: 5
});
const ipvcCidrs = aws.cidrOutput({
    ipBlock: ipv6Cidr.ipv6CidrBlock.apply(cidr => cidr!),
    cidrBits: 64,
    count: 4,
});

ipvcCidrs.apply(cidr => {
    if (cidr.subnets.length !== 4) {
        throw new Error('Expected 4 ipv6 CIDRs');
    }
})

cidrs.apply(cidr => {
    if (cidr.subnets.length !== 4) {
        throw new Error('Expected 4 ipv4 CIDRs');
    }
})

pulumi.all([cidrs.subnets, ipvcCidrs.subnets]).apply(([ipv4, ipv6]) => {
    for (let i = 0; i < 4; i++) {
        new aws.ec2.Subnet(`subnet-${i}`, {
            vpcId: vpc.id,
            cidrBlock: ipv4[i],
            ipv6CidrBlock: ipv6[i]
        }, { dependsOn: [ipv6Cidr]});
    }
});
