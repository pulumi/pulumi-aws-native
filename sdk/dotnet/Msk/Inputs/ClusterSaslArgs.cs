// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterSaslArgs : global::Pulumi.ResourceArgs
    {
        [Input("iam")]
        public Input<Inputs.ClusterIamArgs>? Iam { get; set; }

        [Input("scram")]
        public Input<Inputs.ClusterScramArgs>? Scram { get; set; }

        public ClusterSaslArgs()
        {
        }
        public static new ClusterSaslArgs Empty => new ClusterSaslArgs();
    }
}