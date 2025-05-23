// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class MembershipMlPaymentConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The payment responsibilities accepted by the member for model inference.
        /// </summary>
        [Input("modelInference")]
        public Input<Inputs.MembershipModelInferencePaymentConfigArgs>? ModelInference { get; set; }

        /// <summary>
        /// The payment responsibilities accepted by the member for model training.
        /// </summary>
        [Input("modelTraining")]
        public Input<Inputs.MembershipModelTrainingPaymentConfigArgs>? ModelTraining { get; set; }

        public MembershipMlPaymentConfigArgs()
        {
        }
        public static new MembershipMlPaymentConfigArgs Empty => new MembershipMlPaymentConfigArgs();
    }
}
