// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Information about a parameter used to provision a product.
    /// </summary>
    [OutputType]
    public sealed class ProjectProvisioningParameter
    {
        /// <summary>
        /// The parameter key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The parameter value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ProjectProvisioningParameter(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
