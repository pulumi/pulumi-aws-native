// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelCardModelPackageCreatorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the user's profile in Studio
        /// </summary>
        [Input("userProfileName")]
        public Input<string>? UserProfileName { get; set; }

        public ModelCardModelPackageCreatorArgs()
        {
        }
        public static new ModelCardModelPackageCreatorArgs Empty => new ModelCardModelPackageCreatorArgs();
    }
}
