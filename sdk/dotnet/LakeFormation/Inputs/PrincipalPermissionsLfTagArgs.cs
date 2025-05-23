// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    /// <summary>
    /// The LF-tag key and values attached to a resource.
    /// </summary>
    public sealed class PrincipalPermissionsLfTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key-name for the LF-tag.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        [Input("tagValues")]
        private InputList<string>? _tagValues;

        /// <summary>
        /// A list of possible values of the corresponding ``TagKey`` of an LF-tag key-value pair.
        /// </summary>
        public InputList<string> TagValues
        {
            get => _tagValues ?? (_tagValues = new InputList<string>());
            set => _tagValues = value;
        }

        public PrincipalPermissionsLfTagArgs()
        {
        }
        public static new PrincipalPermissionsLfTagArgs Empty => new PrincipalPermissionsLfTagArgs();
    }
}
