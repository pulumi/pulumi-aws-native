// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class DataSourceOpenSearchServiceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Region.
        /// </summary>
        [Input("awsRegion", required: true)]
        public Input<string> AwsRegion { get; set; } = null!;

        /// <summary>
        /// The endpoint.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        public DataSourceOpenSearchServiceConfigArgs()
        {
        }
        public static new DataSourceOpenSearchServiceConfigArgs Empty => new DataSourceOpenSearchServiceConfigArgs();
    }
}
