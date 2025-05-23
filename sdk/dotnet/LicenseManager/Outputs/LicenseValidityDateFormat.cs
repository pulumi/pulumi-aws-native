// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager.Outputs
{

    [OutputType]
    public sealed class LicenseValidityDateFormat
    {
        /// <summary>
        /// Validity begin date for the license.
        /// </summary>
        public readonly string Begin;
        /// <summary>
        /// Validity begin date for the license.
        /// </summary>
        public readonly string End;

        [OutputConstructor]
        private LicenseValidityDateFormat(
            string begin,

            string end)
        {
            Begin = begin;
            End = end;
        }
    }
}
