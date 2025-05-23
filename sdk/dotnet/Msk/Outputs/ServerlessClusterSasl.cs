// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ServerlessClusterSasl
    {
        /// <summary>
        /// Details for ClientAuthentication using IAM.
        /// </summary>
        public readonly Outputs.ServerlessClusterIam Iam;

        [OutputConstructor]
        private ServerlessClusterSasl(Outputs.ServerlessClusterIam iam)
        {
            Iam = iam;
        }
    }
}
