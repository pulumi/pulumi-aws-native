// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TargetGroupMatcher
    {
        public readonly string? GrpcCode;
        public readonly string? HttpCode;

        [OutputConstructor]
        private TargetGroupMatcher(
            string? grpcCode,

            string? httpCode)
        {
            GrpcCode = grpcCode;
            HttpCode = httpCode;
        }
    }
}