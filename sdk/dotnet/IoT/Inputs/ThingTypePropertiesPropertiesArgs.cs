// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The thing type properties for the thing type to create. It contains information about the new thing type including a description, and a list of searchable thing attribute names. `ThingTypeProperties` can't be updated after the initial creation of the `ThingType` .
    /// </summary>
    public sealed class ThingTypePropertiesPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("searchableAttributes")]
        private InputList<string>? _searchableAttributes;

        /// <summary>
        /// A list of searchable thing attribute names.
        /// </summary>
        public InputList<string> SearchableAttributes
        {
            get => _searchableAttributes ?? (_searchableAttributes = new InputList<string>());
            set => _searchableAttributes = value;
        }

        /// <summary>
        /// The description of the thing type.
        /// </summary>
        [Input("thingTypeDescription")]
        public Input<string>? ThingTypeDescription { get; set; }

        public ThingTypePropertiesPropertiesArgs()
        {
        }
        public static new ThingTypePropertiesPropertiesArgs Empty => new ThingTypePropertiesPropertiesArgs();
    }
}
