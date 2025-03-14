// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Studio settings. If these settings are applied on a user level, they take priority over the settings applied on a domain level.
    /// </summary>
    [OutputType]
    public sealed class DomainStudioWebPortalSettings
    {
        /// <summary>
        /// Applications supported in Studio that are hidden from the Studio left navigation pane.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.SageMaker.DomainAppType> HiddenAppTypes;
        /// <summary>
        /// The machine learning tools that are hidden from the Studio left navigation pane.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.SageMaker.DomainMlTools> HiddenMlTools;

        [OutputConstructor]
        private DomainStudioWebPortalSettings(
            ImmutableArray<Pulumi.AwsNative.SageMaker.DomainAppType> hiddenAppTypes,

            ImmutableArray<Pulumi.AwsNative.SageMaker.DomainMlTools> hiddenMlTools)
        {
            HiddenAppTypes = hiddenAppTypes;
            HiddenMlTools = hiddenMlTools;
        }
    }
}
