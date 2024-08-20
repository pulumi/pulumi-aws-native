# Copyright 2024, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import ivs

advanced_sd_channel = ivs.Channel("channel-advanced-sd", type=ivs.ChannelType.ADVANCED_SD)
standard_channel = ivs.Channel("channel-standard", type=ivs.ChannelType.STANDARD)

pulumi.export("advanced_sd_arn", advanced_sd_channel.arn)
pulumi.export("standard_arn", standard_channel.arn)
