// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Inputs
{

    /// <summary>
    /// An object that specifies information about a property definition.
    /// </summary>
    public sealed class EntityDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        private InputMap<string>? _configuration;

        /// <summary>
        /// An object that specifies information about a property configuration.
        /// </summary>
        public InputMap<string> Configuration
        {
            get => _configuration ?? (_configuration = new InputMap<string>());
            set => _configuration = value;
        }

        /// <summary>
        /// An object that contains information about the data type.
        /// </summary>
        [Input("dataType")]
        public Input<Inputs.EntityDataTypeArgs>? DataType { get; set; }

        /// <summary>
        /// An object that contains the default value.
        /// </summary>
        [Input("defaultValue")]
        public Input<Inputs.EntityDataValueArgs>? DefaultValue { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property ID comes from an external data store.
        /// </summary>
        [Input("isExternalId")]
        public Input<bool>? IsExternalId { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition can be updated.
        /// </summary>
        [Input("isFinal")]
        public Input<bool>? IsFinal { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition is imported from an external data store.
        /// </summary>
        [Input("isImported")]
        public Input<bool>? IsImported { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition is inherited from a parent entity.
        /// </summary>
        [Input("isInherited")]
        public Input<bool>? IsInherited { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property is required.
        /// </summary>
        [Input("isRequiredInEntity")]
        public Input<bool>? IsRequiredInEntity { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property is stored externally.
        /// </summary>
        [Input("isStoredExternally")]
        public Input<bool>? IsStoredExternally { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property consists of time series data.
        /// </summary>
        [Input("isTimeSeries")]
        public Input<bool>? IsTimeSeries { get; set; }

        public EntityDefinitionArgs()
        {
        }
        public static new EntityDefinitionArgs Empty => new EntityDefinitionArgs();
    }
}
