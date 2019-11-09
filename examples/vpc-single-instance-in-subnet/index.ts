import * as pulumi from "@pulumi/pulumi";
import * as cloudformation from "@pulumi/cloudformation";

const region = cloudformation.region!;
const stack = cloudformation.stack!;

const config = new pulumi.Config();
const instanceType = config.get("instanceType") || "t3.small";
const keyName = config.get("keyName");

const region2Examples: any = {
	"us-east-1":      "https://s3.amazonaws.com/cloudformation-examples-us-east-1",
	"us-west-2":      "https://s3-us-west-2.amazonaws.com/cloudformation-examples-us-west-2",
	"us-west-1":      "https://s3-us-west-1.amazonaws.com/cloudformation-examples-us-west-1",
	"eu-west-1":      "https://s3-eu-west-1.amazonaws.com/cloudformation-examples-eu-west-1",
	"eu-west-2":      "https://s3-eu-west-2.amazonaws.com/cloudformation-examples-eu-west-2",
	"eu-west-3":      "https://s3-eu-west-3.amazonaws.com/cloudformation-examples-eu-west-3",
	"eu-north-1":     "https://s3-eu-north-1.amazonaws.com/cloudformation-examples-eu-north-1",
	"eu-central-1":   "https://s3-eu-central-1.amazonaws.com/cloudformation-examples-eu-central-1",
	"ap-southeast-1": "https://s3-ap-southeast-1.amazonaws.com/cloudformation-examples-ap-southeast-1",
	"ap-northeast-1": "https://s3-ap-northeast-1.amazonaws.com/cloudformation-examples-ap-northeast-1",
	"ap-northeast-2": "https://s3-ap-northeast-2.amazonaws.com/cloudformation-examples-ap-northeast-2",
	"ap-northeast-3": "https://s3-ap-northeast-3.amazonaws.com/cloudformation-examples-ap-northeast-3",
	"ap-southeast-2": "https://s3-ap-southeast-2.amazonaws.com/cloudformation-examples-ap-southeast-2",
	"ap-south-1":     "https://s3-ap-south-1.amazonaws.com/cloudformation-examples-ap-south-1",
	"us-east-2":      "https://s3-us-east-2.amazonaws.com/cloudformation-examples-us-east-2",
	"ca-central-1":   "https://s3-ca-central-1.amazonaws.com/cloudformation-examples-ca-central-1",
	"sa-east-1":      "https://s3-sa-east-1.amazonaws.com/cloudformation-examples-sa-east-1",
	"cn-north-1":     "https://s3.cn-north-1.amazonaws.com.cn/cloudformation-examples-cn-north-1",
	"cn-northwest-1": "https://s3.cn-northwest-1.amazonaws.com.cn/cloudformation-examples-cn-northwest-1",
}

const awsInstanceType2Arch: any = {
	"t1.micro":    "HVM64",
	"t2.nano":     "HVM64",
	"t2.micro":    "HVM64",
	"t2.small":    "HVM64",
	"t2.medium":   "HVM64",
    "t2.large":    "HVM64",
    "t3.small":    "HVM64",
    "t3.micro":    "HVM64",
	"m1.small":    "HVM64",
	"m1.medium":   "HVM64",
	"m1.large":    "HVM64",
	"m1.xlarge":   "HVM64",
	"m2.xlarge":   "HVM64",
	"m2.2xlarge":  "HVM64",
	"m2.4xlarge":  "HVM64",
	"m3.medium":   "HVM64",
	"m3.large":    "HVM64",
	"m3.xlarge":   "HVM64",
	"m3.2xlarge":  "HVM64",
	"m4.large":    "HVM64",
	"m4.xlarge":   "HVM64",
	"m4.2xlarge":  "HVM64",
	"m4.4xlarge":  "HVM64",
	"m4.10xlarge": "HVM64",
	"c1.medium":   "HVM64",
	"c1.xlarge":   "HVM64",
	"c3.large":    "HVM64",
	"c3.xlarge":   "HVM64",
	"c3.2xlarge":  "HVM64",
	"c3.4xlarge":  "HVM64",
	"c3.8xlarge":  "HVM64",
	"c4.large":    "HVM64",
	"c4.xlarge":   "HVM64",
	"c4.2xlarge":  "HVM64",
	"c4.4xlarge":  "HVM64",
	"c4.8xlarge":  "HVM64",
	"g2.2xlarge":  "HVMG2",
	"g2.8xlarge":  "HVMG2",
	"r3.large":    "HVM64",
	"r3.xlarge":   "HVM64",
	"r3.2xlarge":  "HVM64",
	"r3.4xlarge":  "HVM64",
	"r3.8xlarge":  "HVM64",
	"i2.xlarge":   "HVM64",
	"i2.2xlarge":  "HVM64",
	"i2.4xlarge":  "HVM64",
	"i2.8xlarge":  "HVM64",
	"d2.xlarge":   "HVM64",
	"d2.2xlarge":  "HVM64",
	"d2.4xlarge":  "HVM64",
	"d2.8xlarge":  "HVM64",
	"hi1.4xlarge": "HVM64",
	"hs1.8xlarge": "HVM64",
	"cr1.8xlarge": "HVM64",
	"cc2.8xlarge": "HVM64",
}

const awsRegionArch2AMI: any = {
	"us-east-1":      {"HVM64": "ami-0080e4c5bc078760e", "HVMG2": "ami-0aeb704d503081ea6"},
	"us-west-2":      {"HVM64": "ami-01e24be29428c15b2", "HVMG2": "ami-0fe84a5b4563d8f27"},
	"us-west-1":      {"HVM64": "ami-0ec6517f6edbf8044", "HVMG2": "ami-0a7fc72dc0e51aa77"},
	"eu-west-1":      {"HVM64": "ami-08935252a36e25f85", "HVMG2": "ami-0d5299b1c6112c3c7"},
	"eu-west-2":      {"HVM64": "ami-01419b804382064e4", "HVMG2": "NOT_SUPPORTED"},
	"eu-west-3":      {"HVM64": "ami-0dd7e7ed60da8fb83", "HVMG2": "NOT_SUPPORTED"},
	"eu-central-1":   {"HVM64": "ami-0cfbf4f6db41068ac", "HVMG2": "ami-0aa1822e3eb913a11"},
	"eu-north-1":     {"HVM64": "ami-86fe70f8", "HVMG2": "ami-32d55b4c"},
	"ap-northeast-1": {"HVM64": "ami-00a5245b4816c38e6", "HVMG2": "ami-09d0e0e099ecabba2"},
	"ap-northeast-2": {"HVM64": "ami-00dc207f8ba6dc919", "HVMG2": "NOT_SUPPORTED"},
	"ap-northeast-3": {"HVM64": "ami-0b65f69a5c11f3522", "HVMG2": "NOT_SUPPORTED"},
	"ap-southeast-1": {"HVM64": "ami-05b3bcf7f311194b3", "HVMG2": "ami-0e46ce0d6a87dc979"},
	"ap-southeast-2": {"HVM64": "ami-02fd0b06f06d93dfc", "HVMG2": "ami-0c0ab057a101d8ff2"},
	"ap-south-1":     {"HVM64": "ami-0ad42f4f66f6c1cc9", "HVMG2": "ami-0244c1d42815af84a"},
	"us-east-2":      {"HVM64": "ami-0cd3dfa4e37921605", "HVMG2": "NOT_SUPPORTED"},
	"ca-central-1":   {"HVM64": "ami-07423fb63ea0a0930", "HVMG2": "NOT_SUPPORTED"},
	"sa-east-1":      {"HVM64": "ami-05145e0b28ad8e0b2", "HVMG2": "NOT_SUPPORTED"},
	"cn-north-1":     {"HVM64": "ami-053617c9d818c1189", "HVMG2": "NOT_SUPPORTED"},
	"cn-northwest-1": {"HVM64": "ami-0f7937761741dc640", "HVMG2": "NOT_SUPPORTED"},
}

const tags = [{Key: "Application", Value: cloudformation.getStackId()}];

const vpc = new cloudformation.ec2.VPC("vpc", {
    CidrBlock: "10.0.0.0/16",
    EnableDnsHostnames: true,
    Tags: tags,
});

const subnet = new cloudformation.ec2.Subnet("subnet", {
    VpcId: vpc.id,
    CidrBlock: "10.0.0.0/24",
    Tags: tags,
});

const internetGateway = new cloudformation.ec2.InternetGateway("internetGateway", {
    Tags: tags,
});

const attachGateway = new cloudformation.ec2.VPCGatewayAttachment("attachGateway", {
    VpcId: vpc.id,
    InternetGatewayId: internetGateway.id,
});

const routeTable = new cloudformation.ec2.RouteTable("routeTable", {
    VpcId: vpc.id,
    Tags: tags,
});

const route = new cloudformation.ec2.Route("route", {
    RouteTableId: routeTable.id,
    DestinationCidrBlock: "0.0.0.0/0",
    GatewayId: internetGateway.id,
});

const subnetRouteTableAssociation = new cloudformation.ec2.SubnetRouteTableAssociation("subnetRouteTableAssociation", {
    SubnetId: subnet.id,
    RouteTableId: routeTable.id,
});

const networkAcl = new cloudformation.ec2.NetworkAcl("networkAcl", {
    VpcId: vpc.id,
    Tags: tags,
});

const commonAclEntryProperties = {
    NetworkAclId: networkAcl.id,
    Protocol: 6,
    RuleAction: "allow",
    CidrBlock: "0.0.0.0/0",
}

const inboundHTTPNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("inboundHTTPNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 100,
    Egress: false,
    PortRange: {From: 80, To: 80},
})

const inboundSSHNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("inboundSSHNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 101,
    Egress: false,
    PortRange: {From: 22, To: 22},
})

const inboundResponsePortsNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("inboundResponsePortsNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 102,
    Egress: false,
    PortRange: {From: 1024, To: 65535},
})

const outBoundHTTPNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("outBoundHTTPNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 100,
    Egress: true,
    PortRange: {From: 80, To: 80},
})

const outBoundHTTPSNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("outBoundHTTPSNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 101,
    Egress: true,
    PortRange: {From: 443, To: 443},
})

const outBoundResponsePortsNetworkAclEntry = new cloudformation.ec2.NetworkAclEntry("outBoundResponsePortsNetworkAclEntry", {
    ...commonAclEntryProperties,
    RuleNumber: 102,
    Egress: true,
    PortRange: {From: 1024, To: 65535},
})

const subnetNetworkAclAssociation = new cloudformation.ec2.SubnetNetworkAclAssociation("subnetNetworkAclAssociation", {
    SubnetId: subnet.id,
    NetworkAclId: networkAcl.id,
});

const instanceSecurityGroup = new cloudformation.ec2.SecurityGroup("instanceSecurityGroup", {
    VpcId: vpc.id,
    GroupDescription: "Enable SSH access via port 22",
    SecurityGroupIngress: [
        {IpProtocol: "tcp", FromPort: 22, ToPort: 22, CidrIp: "0.0.0.0/0"},
        {IpProtocol: "tcp", FromPort: 80, ToPort: 80, CidrIp: "0.0.0.0/0"},
    ],
});

const indexHtml = `<img src="${region2Examples[region]}/cloudformation_graphic.png" alt="AWS CloudFormation Logo"/>
<h1>Congratulations, you have successfully launched the AWS CloudFormation sample.</h1>`;

const cfnHup = `[main]
stack=${stack}
region=${region}
`

const cfnAutoReloader =`[cfn-auto-reloader-hook]
triggers=post.update
path=Resources.webServerInstance.Metadata.AWS::CloudFormation::Init
action=/opt/aws/bin/cfn-init -v --stack ${stack} --resource webServerInstance --region ${region}
runas=root
`

const userData = `#!/bin/bash -xe
yum update -y aws-cfn-bootstrap 
/opt/aws/bin/cfn-init -v --stack ${stack} --resource webServerInstance --region ${region}
/opt/aws/bin/cfn-signal -e $? --stack ${stack} --resource webServerInstance --region ${region}
`

const webServerOptions = {
    logicalId: "webServerInstance",
    metadata: {
        Comment: "Install a simple application",
        "AWS::CloudFormation::Init": {
            config: {
                packages: {
                    yum: {
                        httpd: [],
                    },
                },
                files: {
                    "/var/www/html/index.html": {
                        content: indexHtml,
                        mode: "000644",
                        owner: "root",
                        group: "root",
                    },
                    "/etc/cfn/cfn-hup.conf": {
                        content: cfnHup,
                        mode: "000400",
                        owner: "root",
                        group: "root",
                    },
                    "/etc/cfn/hooks.d/cfn-auto-reloader.conf": {
                        content: cfnAutoReloader,
                        mode: "000400",
                        owner: "root",
                        group: "root",
                    },
                },
                services: {
                    sysvinit: {
                        httpd: {
                            enabled: "true",
                            ensureRunning: "true",
                        },
                        "cfn-hup": {
                            enabled: "true",
                            ensureRunning: "true",
                            files: ["/etc/cfn/cfn-hup.conf", "/etc/cfn/hooks.d/cfn-auto-reloader.conf"],
                        },
                    },
                },
            },
        },
    },
    creationPolicy: {
        ResourceSignal: {
            Timeout: "PT15M",
        },
    },
};

const webServerInstance = new cloudformation.ec2.Instance("webServerInstance", {
    ImageId: awsRegionArch2AMI[region][awsInstanceType2Arch[instanceType]],
    InstanceType: instanceType,
    KeyName: keyName,
    Tags: tags,
    NetworkInterfaces: [{
        GroupSet: [instanceSecurityGroup.id],
        AssociatePublicIpAddress: true,
        DeviceIndex: "0",
        DeleteOnTermination: true,
        SubnetId: subnet.id,
    }],
    UserData: new Buffer(userData).toString("base64"),
}, webServerOptions);

const ipAddress = new cloudformation.ec2.EIP("ipAddress", {
    Domain: "vpc",
    InstanceId: webServerInstance.id,
});

export const url = pulumi.interpolate`http://${webServerInstance.attributes.PublicDnsName}`;
