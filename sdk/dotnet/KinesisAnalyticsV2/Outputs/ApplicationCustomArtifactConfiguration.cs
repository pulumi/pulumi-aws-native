// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// The configuration of connectors and user-defined functions.
    /// </summary>
    [OutputType]
    public sealed class ApplicationCustomArtifactConfiguration
    {
        /// <summary>
        /// Set this to either `UDF` or `DEPENDENCY_JAR`. `UDF` stands for user-defined functions. This type of artifact must be in an S3 bucket. A `DEPENDENCY_JAR` can be in either Maven or an S3 bucket.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationCustomArtifactConfigurationArtifactType ArtifactType;
        /// <summary>
        /// The parameters required to fully specify a Maven reference.
        /// </summary>
        public readonly Outputs.ApplicationMavenReference? MavenReference;
        /// <summary>
        /// The location of the custom artifacts.
        /// </summary>
        public readonly Outputs.ApplicationS3ContentLocation? S3ContentLocation;

        [OutputConstructor]
        private ApplicationCustomArtifactConfiguration(
            Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationCustomArtifactConfigurationArtifactType artifactType,

            Outputs.ApplicationMavenReference? mavenReference,

            Outputs.ApplicationS3ContentLocation? s3ContentLocation)
        {
            ArtifactType = artifactType;
            MavenReference = mavenReference;
            S3ContentLocation = s3ContentLocation;
        }
    }
}
