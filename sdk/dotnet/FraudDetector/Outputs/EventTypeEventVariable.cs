// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Outputs
{

    [OutputType]
    public sealed class EventTypeEventVariable
    {
        /// <summary>
        /// The event variable ARN.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        public readonly string? CreatedTime;
        /// <summary>
        /// The source of the event variable.
        /// 
        /// Valid values: `EVENT | EXTERNAL_MODEL_SCORE`
        /// 
        /// When defining a variable within a event type, you can only use the `EVENT` value for DataSource when the *Inline* property is set to true. If the *Inline* property is set false, you can use either `EVENT` or `MODEL_SCORE` for DataSource.
        /// </summary>
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataSource? DataSource;
        /// <summary>
        /// The data type of the event variable. For more information, see [Data types](https://docs.aws.amazon.com/frauddetector/latest/ug/variables.html#data-types) .
        /// </summary>
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataType? DataType;
        /// <summary>
        /// The default value of the event variable
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack. If the value is `true` , CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false` , CloudFormation will validate that the object exists and then use it within the resource without making changes to the object.
        /// 
        /// For example, when creating `AWS::FraudDetector::EventType` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will create/update/delete the Variables as part of stack operations. However, if you set `Inline=false` , CloudFormation will associate the variables to your event type but not execute any changes to the variables.
        /// </summary>
        public readonly bool? Inline;
        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        public readonly string? LastUpdatedTime;
        /// <summary>
        /// The name of the event variable.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventTypeTag> Tags;
        /// <summary>
        /// The type of event variable. For more information, see [Variable types](https://docs.aws.amazon.com/frauddetector/latest/ug/variables.html#variable-types) .
        /// </summary>
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableVariableType? VariableType;

        [OutputConstructor]
        private EventTypeEventVariable(
            string? arn,

            string? createdTime,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataSource? dataSource,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataType? dataType,

            string? defaultValue,

            string? description,

            bool? inline,

            string? lastUpdatedTime,

            string? name,

            ImmutableArray<Outputs.EventTypeTag> tags,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableVariableType? variableType)
        {
            Arn = arn;
            CreatedTime = createdTime;
            DataSource = dataSource;
            DataType = dataType;
            DefaultValue = defaultValue;
            Description = description;
            Inline = inline;
            LastUpdatedTime = lastUpdatedTime;
            Name = name;
            Tags = tags;
            VariableType = variableType;
        }
    }
}
