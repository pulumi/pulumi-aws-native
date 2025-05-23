// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Outputs
{

    [OutputType]
    public sealed class ApplicationPersonalizationConfiguration
    {
        /// <summary>
        /// An option to allow Amazon Q Business to customize chat responses using user specific metadata—specifically, location and job information—in your IAM Identity Center instance.
        /// </summary>
        public readonly Pulumi.AwsNative.QBusiness.ApplicationPersonalizationControlMode PersonalizationControlMode;

        [OutputConstructor]
        private ApplicationPersonalizationConfiguration(Pulumi.AwsNative.QBusiness.ApplicationPersonalizationControlMode personalizationControlMode)
        {
            PersonalizationControlMode = personalizationControlMode;
        }
    }
}
