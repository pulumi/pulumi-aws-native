// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class UserProfileHiddenSageMakerImage
    {
        /// <summary>
        /// The SageMaker image name that you are hiding from the Studio user interface.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.UserProfileHiddenSageMakerImageSageMakerImageName? SageMakerImageName;
        /// <summary>
        /// The version aliases you are hiding from the Studio user interface.
        /// </summary>
        public readonly ImmutableArray<string> VersionAliases;

        [OutputConstructor]
        private UserProfileHiddenSageMakerImage(
            Pulumi.AwsNative.SageMaker.UserProfileHiddenSageMakerImageSageMakerImageName? sageMakerImageName,

            ImmutableArray<string> versionAliases)
        {
            SageMakerImageName = sageMakerImageName;
            VersionAliases = versionAliases;
        }
    }
}
