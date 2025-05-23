// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The rolling update policy for the inference component
    /// </summary>
    public sealed class InferenceComponentRollingUpdatePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The batch size for each rolling step in the deployment process. For each step, SageMaker AI provisions capacity on the new endpoint fleet, routes traffic to that fleet, and terminates capacity on the old endpoint fleet. The value must be between 5% to 50% of the copy count of the inference component.
        /// </summary>
        [Input("maximumBatchSize")]
        public Input<Inputs.InferenceComponentCapacitySizeArgs>? MaximumBatchSize { get; set; }

        /// <summary>
        /// The time limit for the total deployment. Exceeding this limit causes a timeout.
        /// </summary>
        [Input("maximumExecutionTimeoutInSeconds")]
        public Input<int>? MaximumExecutionTimeoutInSeconds { get; set; }

        /// <summary>
        /// The batch size for a rollback to the old endpoint fleet. If this field is absent, the value is set to the default, which is 100% of the total capacity. When the default is used, SageMaker AI provisions the entire capacity of the old fleet at once during rollback.
        /// </summary>
        [Input("rollbackMaximumBatchSize")]
        public Input<Inputs.InferenceComponentCapacitySizeArgs>? RollbackMaximumBatchSize { get; set; }

        /// <summary>
        /// The length of the baking period, during which SageMaker AI monitors alarms for each batch on the new fleet.
        /// </summary>
        [Input("waitIntervalInSeconds")]
        public Input<int>? WaitIntervalInSeconds { get; set; }

        public InferenceComponentRollingUpdatePolicyArgs()
        {
        }
        public static new InferenceComponentRollingUpdatePolicyArgs Empty => new InferenceComponentRollingUpdatePolicyArgs();
    }
}
