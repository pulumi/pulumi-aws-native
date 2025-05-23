// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class MembershipPaymentConfiguration
    {
        /// <summary>
        /// The payment responsibilities accepted by the collaboration member for job compute costs.
        /// </summary>
        public readonly Outputs.MembershipJobComputePaymentConfig? JobCompute;
        /// <summary>
        /// The payment responsibilities accepted by the collaboration member for machine learning costs.
        /// </summary>
        public readonly Outputs.MembershipMlPaymentConfig? MachineLearning;
        /// <summary>
        /// The payment responsibilities accepted by the collaboration member for query compute costs.
        /// </summary>
        public readonly Outputs.MembershipQueryComputePaymentConfig QueryCompute;

        [OutputConstructor]
        private MembershipPaymentConfiguration(
            Outputs.MembershipJobComputePaymentConfig? jobCompute,

            Outputs.MembershipMlPaymentConfig? machineLearning,

            Outputs.MembershipQueryComputePaymentConfig queryCompute)
        {
            JobCompute = jobCompute;
            MachineLearning = machineLearning;
            QueryCompute = queryCompute;
        }
    }
}
