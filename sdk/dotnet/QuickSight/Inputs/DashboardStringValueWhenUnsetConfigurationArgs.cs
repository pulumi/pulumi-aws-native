// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardStringValueWhenUnsetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A custom value that's used when the value of a parameter isn't set.
        /// </summary>
        [Input("customValue")]
        public Input<string>? CustomValue { get; set; }

        /// <summary>
        /// The built-in options for default values. The value can be one of the following:
        /// 
        /// - `RECOMMENDED` : The recommended value.
        /// - `NULL` : The `NULL` value.
        /// </summary>
        [Input("valueWhenUnsetOption")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardValueWhenUnsetOption>? ValueWhenUnsetOption { get; set; }

        public DashboardStringValueWhenUnsetConfigurationArgs()
        {
        }
        public static new DashboardStringValueWhenUnsetConfigurationArgs Empty => new DashboardStringValueWhenUnsetConfigurationArgs();
    }
}
