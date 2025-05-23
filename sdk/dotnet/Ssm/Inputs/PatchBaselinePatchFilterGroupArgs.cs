// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm.Inputs
{

    /// <summary>
    /// The patch filter group that defines the criteria for the rule.
    /// </summary>
    public sealed class PatchBaselinePatchFilterGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("patchFilters")]
        private InputList<Inputs.PatchBaselinePatchFilterArgs>? _patchFilters;

        /// <summary>
        /// The set of patch filters that make up the group.
        /// </summary>
        public InputList<Inputs.PatchBaselinePatchFilterArgs> PatchFilters
        {
            get => _patchFilters ?? (_patchFilters = new InputList<Inputs.PatchBaselinePatchFilterArgs>());
            set => _patchFilters = value;
        }

        public PatchBaselinePatchFilterGroupArgs()
        {
        }
        public static new PatchBaselinePatchFilterGroupArgs Empty => new PatchBaselinePatchFilterGroupArgs();
    }
}
