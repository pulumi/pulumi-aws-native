// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Contains the configuration information of alarm state changes.
    /// </summary>
    public sealed class AlarmModelAlarmCapabilitiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to get notified for alarm state changes.
        /// </summary>
        [Input("acknowledgeFlow")]
        public Input<Inputs.AlarmModelAcknowledgeFlowArgs>? AcknowledgeFlow { get; set; }

        /// <summary>
        /// Specifies the default alarm state. The configuration applies to all alarms that were created based on this alarm model.
        /// </summary>
        [Input("initializationConfiguration")]
        public Input<Inputs.AlarmModelInitializationConfigurationArgs>? InitializationConfiguration { get; set; }

        public AlarmModelAlarmCapabilitiesArgs()
        {
        }
        public static new AlarmModelAlarmCapabilitiesArgs Empty => new AlarmModelAlarmCapabilitiesArgs();
    }
}
