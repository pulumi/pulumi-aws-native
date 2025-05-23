// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GlobalAccelerator.Outputs
{

    /// <summary>
    /// ARN of resource to share.
    /// </summary>
    [OutputType]
    public sealed class CrossAccountAttachmentResource
    {
        /// <summary>
        /// An IP address range, in CIDR format, that is specified as resource. The address must be provisioned and advertised in AWS Global Accelerator by following the bring your own IP address (BYOIP) process for Global Accelerator
        /// 
        /// For more information, see [Bring your own IP addresses (BYOIP)](https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in the AWS Global Accelerator Developer Guide.
        /// </summary>
        public readonly string? Cidr;
        /// <summary>
        /// The endpoint ID for the endpoint that is specified as a AWS resource.
        /// 
        /// An endpoint ID for the cross-account feature is the ARN of an AWS resource, such as a Network Load Balancer, that Global Accelerator supports as an endpoint for an accelerator.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// The AWS Region where a shared endpoint resource is located.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private CrossAccountAttachmentResource(
            string? cidr,

            string? endpointId,

            string? region)
        {
            Cidr = cidr;
            EndpointId = endpointId;
            Region = region;
        }
    }
}
