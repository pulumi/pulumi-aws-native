// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Inputs
{

    public sealed class DetectorEventTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the event type.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The description of the event type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("entityTypes")]
        private InputList<Inputs.DetectorEntityTypeArgs>? _entityTypes;

        /// <summary>
        /// The event type entity types.
        /// </summary>
        public InputList<Inputs.DetectorEntityTypeArgs> EntityTypes
        {
            get => _entityTypes ?? (_entityTypes = new InputList<Inputs.DetectorEntityTypeArgs>());
            set => _entityTypes = value;
        }

        [Input("eventVariables")]
        private InputList<Inputs.DetectorEventVariableArgs>? _eventVariables;

        /// <summary>
        /// The event type event variables.
        /// </summary>
        public InputList<Inputs.DetectorEventVariableArgs> EventVariables
        {
            get => _eventVariables ?? (_eventVariables = new InputList<Inputs.DetectorEventVariableArgs>());
            set => _eventVariables = value;
        }

        /// <summary>
        /// Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack. If the value is `true` , CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false` , CloudFormation will validate that the object exists and then use it within the resource without making changes to the object.
        /// 
        /// For example, when creating `AWS::FraudDetector::Detector` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will create/update/delete the Variables as part of stack operations. However, if you set `Inline=false` , CloudFormation will associate the variables to your detector but not execute any changes to the variables.
        /// </summary>
        [Input("inline")]
        public Input<bool>? Inline { get; set; }

        [Input("labels")]
        private InputList<Inputs.DetectorLabelArgs>? _labels;

        /// <summary>
        /// The event type labels.
        /// </summary>
        public InputList<Inputs.DetectorLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.DetectorLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// The name for the event type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.DetectorTagArgs>? _tags;

        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public InputList<Inputs.DetectorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DetectorTagArgs>());
            set => _tags = value;
        }

        public DetectorEventTypeArgs()
        {
        }
        public static new DetectorEventTypeArgs Empty => new DetectorEventTypeArgs();
    }
}
