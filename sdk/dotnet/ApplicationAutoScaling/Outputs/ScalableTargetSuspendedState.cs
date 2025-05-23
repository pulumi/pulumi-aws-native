// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Outputs
{

    /// <summary>
    /// ``SuspendedState`` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies whether the scaling activities for a scalable target are in a suspended state.
    ///  For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the *Application Auto Scaling User Guide*.
    /// </summary>
    [OutputType]
    public sealed class ScalableTargetSuspendedState
    {
        /// <summary>
        /// Whether scale in by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is ``false``.
        /// </summary>
        public readonly bool? DynamicScalingInSuspended;
        /// <summary>
        /// Whether scale out by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is ``false``.
        /// </summary>
        public readonly bool? DynamicScalingOutSuspended;
        /// <summary>
        /// Whether scheduled scaling is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is ``false``.
        /// </summary>
        public readonly bool? ScheduledScalingSuspended;

        [OutputConstructor]
        private ScalableTargetSuspendedState(
            bool? dynamicScalingInSuspended,

            bool? dynamicScalingOutSuspended,

            bool? scheduledScalingSuspended)
        {
            DynamicScalingInSuspended = dynamicScalingInSuspended;
            DynamicScalingOutSuspended = dynamicScalingOutSuspended;
            ScheduledScalingSuspended = scheduledScalingSuspended;
        }
    }
}
