// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Inputs
{

    public sealed class MatchingWorkflowOutputAttributeArgs : global::Pulumi.ResourceArgs
    {
        [Input("hashed")]
        public Input<bool>? Hashed { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public MatchingWorkflowOutputAttributeArgs()
        {
        }
        public static new MatchingWorkflowOutputAttributeArgs Empty => new MatchingWorkflowOutputAttributeArgs();
    }
}