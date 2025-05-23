// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
    /// </summary>
    public sealed class ModelCardUserContextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain associated with the user.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the user's profile.
        /// </summary>
        [Input("userProfileArn")]
        public Input<string>? UserProfileArn { get; set; }

        /// <summary>
        /// The name of the user's profile.
        /// </summary>
        [Input("userProfileName")]
        public Input<string>? UserProfileName { get; set; }

        public ModelCardUserContextArgs()
        {
        }
        public static new ModelCardUserContextArgs Empty => new ModelCardUserContextArgs();
    }
}
