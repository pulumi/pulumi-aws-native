// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions.Inputs
{

    public sealed class PolicyEntityIdentifierArgs : global::Pulumi.ResourceArgs
    {
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        public PolicyEntityIdentifierArgs()
        {
        }
        public static new PolicyEntityIdentifierArgs Empty => new PolicyEntityIdentifierArgs();
    }
}
