// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisVisualCustomAction
    {
        /// <summary>
        /// A list of `VisualCustomActionOperations` .
        /// 
        /// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisVisualCustomActionOperation> ActionOperations;
        /// <summary>
        /// The ID of the `VisualCustomAction` .
        /// </summary>
        public readonly string CustomActionId;
        /// <summary>
        /// The name of the `VisualCustomAction` .
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the `VisualCustomAction` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisWidgetStatus? Status;
        /// <summary>
        /// The trigger of the `VisualCustomAction` .
        /// 
        /// Valid values are defined as follows:
        /// 
        /// - `DATA_POINT_CLICK` : Initiates a custom action by a left pointer click on a data point.
        /// - `DATA_POINT_MENU` : Initiates a custom action by right pointer click from the menu.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisualCustomActionTrigger Trigger;

        [OutputConstructor]
        private AnalysisVisualCustomAction(
            ImmutableArray<Outputs.AnalysisVisualCustomActionOperation> actionOperations,

            string customActionId,

            string name,

            Pulumi.AwsNative.QuickSight.AnalysisWidgetStatus? status,

            Pulumi.AwsNative.QuickSight.AnalysisVisualCustomActionTrigger trigger)
        {
            ActionOperations = actionOperations;
            CustomActionId = customActionId;
            Name = name;
            Status = status;
            Trigger = trigger;
        }
    }
}
