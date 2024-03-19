# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import ExtensionResource, Provider, ProviderDefaultTagsArgs
from pulumi_random import RandomId

aws_native_config = pulumi.Config("aws-native")

aws_native_provider = Provider("aws_native_provider",
                               region=aws_native_config.require("region"),
                               default_tags=ProviderDefaultTagsArgs(tags={"defaultTag": "defaultTagValue"}))

paramName = RandomId("paramName", byte_length=8)

resource = ExtensionResource("param",
                             type="AWS::SSM::Parameter",
                             properties={
                                 "Name": paramName.id,
                                 "Type": "String",
                                 "Value": "originalValue",
                                 "Tags": {
                                     "localTag": "localTagValue",
                                 }
                             },
                             write_only=[
                                 "AllowedPattern",
                                 "Description",
                                 "Policies",
                                 "Tags",
                                 "Tier"
                             ],
                             tags_property="Tags",
                             tags_style="stringMap",
                             opts=pulumi.ResourceOptions(provider=aws_native_provider))

# Test applying default tags with array style
logGroupName = RandomId("logGroupName", byte_length=8)
resource2 = ExtensionResource("log-group",
                              type="AWS::Logs::LogGroup",
                                properties={
                                    "LogGroupName": logGroupName.id,
                                    "RetentionInDays": 90,
                                    "Tags": [{
                                      "Key": "localTag",
                                      "Value": "localTagValue",
                                    }],
                                },
                                tags_property="Tags",
                                opts=pulumi.ResourceOptions(provider=aws_native_provider))

pulumi.export("outputs", resource.outputs)
