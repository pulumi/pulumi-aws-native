// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// The information required to specify a Maven reference. You can use Maven references to specify dependency JAR files.
    /// </summary>
    public sealed class ApplicationMavenReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The artifact ID of the Maven reference.
        /// </summary>
        [Input("artifactId", required: true)]
        public Input<string> ArtifactId { get; set; } = null!;

        /// <summary>
        /// The group ID of the Maven reference.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The version of the Maven reference.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ApplicationMavenReferenceArgs()
        {
        }
        public static new ApplicationMavenReferenceArgs Empty => new ApplicationMavenReferenceArgs();
    }
}
