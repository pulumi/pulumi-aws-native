// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Inputs
{

    /// <summary>
    /// An object that specifies a value for a property.
    /// </summary>
    public sealed class EntityDataValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Boolean value.
        /// </summary>
        [Input("booleanValue")]
        public Input<bool>? BooleanValue { get; set; }

        /// <summary>
        /// A double value.
        /// </summary>
        [Input("doubleValue")]
        public Input<double>? DoubleValue { get; set; }

        /// <summary>
        /// An expression that produces the value.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// An integer value.
        /// </summary>
        [Input("integerValue")]
        public Input<int>? IntegerValue { get; set; }

        [Input("listValue")]
        private InputList<Inputs.EntityDataValueArgs>? _listValue;

        /// <summary>
        /// A list of multiple values.
        /// </summary>
        public InputList<Inputs.EntityDataValueArgs> ListValue
        {
            get => _listValue ?? (_listValue = new InputList<Inputs.EntityDataValueArgs>());
            set => _listValue = value;
        }

        /// <summary>
        /// A long value.
        /// </summary>
        [Input("longValue")]
        public Input<double>? LongValue { get; set; }

        [Input("mapValue")]
        private InputMap<Inputs.EntityDataValueArgs>? _mapValue;

        /// <summary>
        /// An object that maps strings to multiple DataValue objects.
        /// </summary>
        public InputMap<Inputs.EntityDataValueArgs> MapValue
        {
            get => _mapValue ?? (_mapValue = new InputMap<Inputs.EntityDataValueArgs>());
            set => _mapValue = value;
        }

        /// <summary>
        /// A value that relates a component to another component.
        /// </summary>
        [Input("relationshipValue")]
        public Input<Inputs.EntityDataValueRelationshipValuePropertiesArgs>? RelationshipValue { get; set; }

        /// <summary>
        /// A string value.
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        public EntityDataValueArgs()
        {
        }
        public static new EntityDataValueArgs Empty => new EntityDataValueArgs();
    }
}