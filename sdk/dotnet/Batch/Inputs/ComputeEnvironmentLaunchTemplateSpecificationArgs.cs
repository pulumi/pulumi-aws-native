// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class ComputeEnvironmentLaunchTemplateSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the launch template.
        /// </summary>
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        /// <summary>
        /// The name of the launch template.
        /// </summary>
        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        /// <summary>
        /// The version number of the launch template, `$Latest` , or `$Default` .
        /// 
        /// If the value is `$Latest` , the latest version of the launch template is used. If the value is `$Default` , the default version of the launch template is used.
        /// 
        /// &gt; If the AMI ID that's used in a compute environment is from the launch template, the AMI isn't changed when the compute environment is updated. It's only changed if the `updateToLatestImageVersion` parameter for the compute environment is set to `true` . During an infrastructure update, if either `$Latest` or `$Default` is specified, AWS Batch re-evaluates the launch template version, and it might use a different version of the launch template. This is the case even if the launch template isn't specified in the update. When updating a compute environment, changing the launch template requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* . 
        /// 
        /// Default: `$Default` .
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ComputeEnvironmentLaunchTemplateSpecificationArgs()
        {
        }
        public static new ComputeEnvironmentLaunchTemplateSpecificationArgs Empty => new ComputeEnvironmentLaunchTemplateSpecificationArgs();
    }
}
