# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import logs

log_group = logs.LogGroup(
    "log-test",
    retention_in_days=90)

pulumi.export("arn", log_group.arn)
