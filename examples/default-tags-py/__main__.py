# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import logs, resiliencehub, TagArgs

log_group = logs.LogGroup(
    "log-test",
    retention_in_days=90,
    tags=[
        TagArgs(key="localTag", value="localTagValue"),
    ])

policy = resiliencehub.ResiliencyPolicy(
  "policy",
  tier=resiliencehub.ResiliencyPolicyTier.NON_CRITICAL,
  policy=resiliencehub.ResiliencyPolicyPolicyMapArgs(
    software=resiliencehub.ResiliencyPolicyFailurePolicyArgs(
      rto_in_secs=10,
      rpo_in_secs=10,
    ),
    hardware=resiliencehub.ResiliencyPolicyFailurePolicyArgs(
      rto_in_secs=10,
      rpo_in_secs=10,
    ),
    az=resiliencehub.ResiliencyPolicyFailurePolicyArgs(
      rto_in_secs=10,
      rpo_in_secs=10,
    ),
    region=resiliencehub.ResiliencyPolicyFailurePolicyArgs(
      rto_in_secs=10,
      rpo_in_secs=10,
    ),
  ),
  tags={"localTag": "localTagValue"},
)

pulumi.export("logGroupTags", log_group.tags)
pulumi.export("policyTags", policy.tags)
