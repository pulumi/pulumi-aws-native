// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class UserProfileHiddenSageMakerImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SageMaker image name that you are hiding from the Studio user interface.
        /// </summary>
        [Input("sageMakerImageName")]
        public Input<Pulumi.AwsNative.SageMaker.UserProfileHiddenSageMakerImageSageMakerImageName>? SageMakerImageName { get; set; }

        [Input("versionAliases")]
        private InputList<string>? _versionAliases;

        /// <summary>
        /// The version aliases you are hiding from the Studio user interface.
        /// </summary>
        public InputList<string> VersionAliases
        {
            get => _versionAliases ?? (_versionAliases = new InputList<string>());
            set => _versionAliases = value;
        }

        public UserProfileHiddenSageMakerImageArgs()
        {
        }
        public static new UserProfileHiddenSageMakerImageArgs Empty => new UserProfileHiddenSageMakerImageArgs();
    }
}
