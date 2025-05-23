// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Inputs
{

    /// <summary>
    /// The custom grouped component.
    /// </summary>
    public sealed class ApplicationCustomComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the component.
        /// </summary>
        [Input("componentName", required: true)]
        public Input<string> ComponentName { get; set; } = null!;

        [Input("resourceList", required: true)]
        private InputList<string>? _resourceList;

        /// <summary>
        /// The list of resource ARNs that belong to the component.
        /// </summary>
        public InputList<string> ResourceList
        {
            get => _resourceList ?? (_resourceList = new InputList<string>());
            set => _resourceList = value;
        }

        public ApplicationCustomComponentArgs()
        {
        }
        public static new ApplicationCustomComponentArgs Empty => new ApplicationCustomComponentArgs();
    }
}
