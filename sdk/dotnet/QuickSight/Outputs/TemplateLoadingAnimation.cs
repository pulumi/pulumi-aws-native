// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateLoadingAnimation
    {
        /// <summary>
        /// The visibility configuration of `LoadingAnimation` .
        /// </summary>
        public readonly object? Visibility;

        [OutputConstructor]
        private TemplateLoadingAnimation(object? visibility)
        {
            Visibility = visibility;
        }
    }
}
