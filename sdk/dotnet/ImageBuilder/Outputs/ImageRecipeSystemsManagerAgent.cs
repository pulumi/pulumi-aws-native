// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// Contains settings for the SSM agent on your build instance.
    /// </summary>
    [OutputType]
    public sealed class ImageRecipeSystemsManagerAgent
    {
        /// <summary>
        /// Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.
        /// </summary>
        public readonly bool? UninstallAfterBuild;

        [OutputConstructor]
        private ImageRecipeSystemsManagerAgent(bool? uninstallAfterBuild)
        {
            UninstallAfterBuild = uninstallAfterBuild;
        }
    }
}