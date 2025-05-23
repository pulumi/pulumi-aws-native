// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ServerlessClusterIamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SASL/IAM authentication is enabled or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ServerlessClusterIamArgs()
        {
        }
        public static new ServerlessClusterIamArgs Empty => new ServerlessClusterIamArgs();
    }
}
