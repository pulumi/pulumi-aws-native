// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class TargetGroupMatcherArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// You can specify values between 0 and 99. You can specify multiple values, or a range of values. The default value is 12.
        /// </summary>
        [Input("grpcCode")]
        public Input<string>? GrpcCode { get; set; }

        /// <summary>
        /// For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200. You can specify multiple values or a range of values. 
        /// </summary>
        [Input("httpCode")]
        public Input<string>? HttpCode { get; set; }

        public TargetGroupMatcherArgs()
        {
        }
        public static new TargetGroupMatcherArgs Empty => new TargetGroupMatcherArgs();
    }
}
