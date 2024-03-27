# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import ec2
from pulumi_command import local

aws_native_config = pulumi.Config("aws-native")

# Only 1 IPAM can exist per region, so this must be created manually as a pre-requisite and is therefore looked up here.
ipamId = local.run(
    command="aws ec2 describe-ipams --query 'Ipams[0].IpamId' --output text").stdout

ipam = ec2.get_ipam(ipam_id=ipamId)

ipamPool = ec2.IpamPool("ipamPool", address_family="ipv4",
                        ipam_scope_id=ipam.private_default_scope_id,
                        locale=aws_native_config.require("region"))

ipamPoolCidr = ec2.IpamPoolCidr(
    "ipamPoolCidr", ipam_pool_id=ipamPool.id, cidr="10.0.0.0/16")

vpc = ec2.Vpc("vpc",
              #   enable_dns_hostnames=True,
              ipv4_ipam_pool_id=ipamPool.id,
              ipv4_netmask_length=20,
              opts=pulumi.ResourceOptions(depends_on=[ipamPoolCidr]))
