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
    public sealed class ClusterVpcConnectivityClientAuthentication
    {
        /// <summary>
        /// Details for VpcConnectivity ClientAuthentication using SASL.
        /// </summary>
        public readonly Outputs.ClusterVpcConnectivitySasl? Sasl;
        /// <summary>
        /// Details for VpcConnectivity ClientAuthentication using TLS.
        /// </summary>
        public readonly Outputs.ClusterVpcConnectivityTls? Tls;

        [OutputConstructor]
        private ClusterVpcConnectivityClientAuthentication(
            Outputs.ClusterVpcConnectivitySasl? sasl,

            Outputs.ClusterVpcConnectivityTls? tls)
        {
            Sasl = sasl;
            Tls = tls;
        }
    }
}
