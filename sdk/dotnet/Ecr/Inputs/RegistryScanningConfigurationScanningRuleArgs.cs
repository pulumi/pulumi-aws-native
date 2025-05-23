// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecr.Inputs
{

    /// <summary>
    /// The scanning rules associated with the registry.
    /// </summary>
    public sealed class RegistryScanningConfigurationScanningRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("repositoryFilters", required: true)]
        private InputList<Inputs.RegistryScanningConfigurationRepositoryFilterArgs>? _repositoryFilters;

        /// <summary>
        /// The details of a scanning repository filter. For more information on how to use filters, see [Using filters](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html#image-scanning-filters) in the *Amazon Elastic Container Registry User Guide*.
        /// </summary>
        public InputList<Inputs.RegistryScanningConfigurationRepositoryFilterArgs> RepositoryFilters
        {
            get => _repositoryFilters ?? (_repositoryFilters = new InputList<Inputs.RegistryScanningConfigurationRepositoryFilterArgs>());
            set => _repositoryFilters = value;
        }

        /// <summary>
        /// The frequency that scans are performed at for a private registry. When the ``ENHANCED`` scan type is specified, the supported scan frequencies are ``CONTINUOUS_SCAN`` and ``SCAN_ON_PUSH``. When the ``BASIC`` scan type is specified, the ``SCAN_ON_PUSH`` scan frequency is supported. If scan on push is not specified, then the ``MANUAL`` scan frequency is set by default.
        /// </summary>
        [Input("scanFrequency", required: true)]
        public Input<Pulumi.AwsNative.Ecr.RegistryScanningConfigurationScanFrequency> ScanFrequency { get; set; } = null!;

        public RegistryScanningConfigurationScanningRuleArgs()
        {
        }
        public static new RegistryScanningConfigurationScanningRuleArgs Empty => new RegistryScanningConfigurationScanningRuleArgs();
    }
}
