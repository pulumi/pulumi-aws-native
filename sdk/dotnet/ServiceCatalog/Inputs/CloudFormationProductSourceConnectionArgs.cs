// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog.Inputs
{

    public sealed class CloudFormationProductSourceConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectionParameters", required: true)]
        public Input<Inputs.CloudFormationProductConnectionParametersArgs> ConnectionParameters { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CloudFormationProductSourceConnectionArgs()
        {
        }
        public static new CloudFormationProductSourceConnectionArgs Empty => new CloudFormationProductSourceConnectionArgs();
    }
}