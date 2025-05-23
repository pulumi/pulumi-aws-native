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
    public sealed class MembershipMlPaymentConfig
    {
        /// <summary>
        /// The payment responsibilities accepted by the member for model inference.
        /// </summary>
        public readonly Outputs.MembershipModelInferencePaymentConfig? ModelInference;
        /// <summary>
        /// The payment responsibilities accepted by the member for model training.
        /// </summary>
        public readonly Outputs.MembershipModelTrainingPaymentConfig? ModelTraining;

        [OutputConstructor]
        private MembershipMlPaymentConfig(
            Outputs.MembershipModelInferencePaymentConfig? modelInference,

            Outputs.MembershipModelTrainingPaymentConfig? modelTraining)
        {
            ModelInference = modelInference;
            ModelTraining = modelTraining;
        }
    }
}
