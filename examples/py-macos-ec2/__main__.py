import pulumi
import pulumi_cloudformation as aws

config = pulumi.config.Config()
key_name = config.get('keyName')
availability_zone = config.get('az', 'us-west-2a')
mac_ami_id = config.get('ami', 'ami-01cdf97a3a0ab6eac') # macOS Catalina 10.15.7 us-west-2

vpc = aws.ec2.VPC('macVpc',
    properties=aws.ec2.VPCPropertiesArgs(
        cidr_block='10.0.0.0/16',
        enable_dns_hostnames=True
    ))

subnet = aws.ec2.Subnet('macSubnet',
    properties=aws.ec2.SubnetPropertiesArgs(
        vpc_id=vpc.id,
        cidr_block='10.0.0.0/24',
        availability_zone=availability_zone,
    ))

security_group = aws.ec2.SecurityGroup('macSecurityGroup',
    properties=aws.ec2.SecurityGroupPropertiesArgs(
        vpc_id=vpc.id,
        group_description='Enable SSH access cia port 22',
        security_group_ingress=[
            aws.ec2.SecurityGroupIngressArgs(ip_protocol='tcp', from_port=22, to_port=22, cidr_ip='0.0.0.0/0'),
        ]
    ))

host = aws.ec2.Host('macHost',
    properties=aws.ec2.HostPropertiesArgs(
        availability_zone=availability_zone,
        instance_type='mac1.metal',
    ))

instance = aws.ec2.Instance('macInstance',
    properties=aws.ec2.InstancePropertiesArgs(
        image_id=mac_ami_id,
        instance_type='mac1.metal',
        host_id=host.id,
        key_name=key_name,
        network_interfaces=[
            aws.ec2.InstanceNetworkInterfaceArgs(
                group_set=[security_group.id],
                associate_public_ip_address=True,
                device_index='0',
                delete_on_termination=True,
                subnet_id=subnet.id,
            )],
    ))
