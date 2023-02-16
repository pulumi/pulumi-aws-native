// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class OriginAccessControlConfig
    {
        public readonly string? Description;
        public readonly string Name;
        public readonly string OriginAccessControlOriginType;
        public readonly string SigningBehavior;
        public readonly string SigningProtocol;

        [OutputConstructor]
        private OriginAccessControlConfig(
            string? description,

            string name,

            string originAccessControlOriginType,

            string signingBehavior,

            string signingProtocol)
        {
            Description = description;
            Name = name;
            OriginAccessControlOriginType = originAccessControlOriginType;
            SigningBehavior = signingBehavior;
            SigningProtocol = signingProtocol;
        }
    }
}