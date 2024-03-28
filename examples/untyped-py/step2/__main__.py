# Copyright 2021, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_aws_native import ExtensionResource, Provider, ProviderDefaultTagsArgs, AutoNamingArgs

aws_native_config = pulumi.Config("aws-native")

aws_native_provider = Provider("aws_native_provider",
                               region=aws_native_config.require("region"),
                               default_tags=ProviderDefaultTagsArgs(tags={"defaultTag": "defaultTagValue"}))

resource = ExtensionResource("param",
                             type="AWS::SSM::Parameter",
                             properties={
                                 "Type": "String",
                                 "Value": "updatedValue",
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
                            #  Should default to "Tags" so we can omit it
                            #  tags_property="Tags",
                             tags_style="stringMap",
                             auto_naming=AutoNamingArgs(property_name="Name"),
                             opts=pulumi.ResourceOptions(provider=aws_native_provider))

pulumi.export("outputs", resource.outputs)
