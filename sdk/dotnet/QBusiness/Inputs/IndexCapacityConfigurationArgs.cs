// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class IndexCapacityConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of storage units configured for an Amazon Q Business index.
        /// </summary>
        [Input("units")]
        public Input<double>? Units { get; set; }

        public IndexCapacityConfigurationArgs()
        {
        }
        public static new IndexCapacityConfigurationArgs Empty => new IndexCapacityConfigurationArgs();
    }
}
