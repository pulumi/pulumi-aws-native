// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Inputs
{

    /// <summary>
    /// Auto-naming specification for the resource.
    /// </summary>
    public sealed class AutoNamingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum length of the name.
        /// </summary>
        [Input("maxLength")]
        public Input<int>? MaxLength { get; set; }

        /// <summary>
        /// The minimum length of the name.
        /// </summary>
        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        /// <summary>
        /// The name of the property in the Cloud Control payload that is used to set the name of the resource.
        /// </summary>
        [Input("propertyName")]
        public Input<string>? PropertyName { get; set; }

        public AutoNamingArgs()
        {
        }
        public static new AutoNamingArgs Empty => new AutoNamingArgs();
    }
}
