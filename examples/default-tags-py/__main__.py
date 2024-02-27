# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import logs, securityhub, TagArgs

log_group = logs.LogGroup(
    "log-test",
    retention_in_days=90,
    tags=[
        TagArgs(key="localTag", value="localTagValue"),
    ])

hub = securityhub.Hub(
  "hub",
  tags={"localTag": "localTagValue"},
)

pulumi.export("logGroupTags", log_group.tags)
pulumi.export("hubTags", hub.tags)
