# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
import pulumi_random as random
from pulumi_aws_native import logs

name = random.RandomString(
    "name",
    length=8,
    special=False,
    upper=False)

log_group = logs.LogGroup(
    "log-test",
    log_group_name=name.result,
    retention_in_days=90)

pulumi.export("arn", log_group.arn)
