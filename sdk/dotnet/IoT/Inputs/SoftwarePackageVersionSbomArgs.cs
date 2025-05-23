// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The sbom zip archive location of the package version
    /// </summary>
    public sealed class SoftwarePackageVersionSbomArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3Location", required: true)]
        public Input<Inputs.SoftwarePackageVersionS3LocationArgs> S3Location { get; set; } = null!;

        public SoftwarePackageVersionSbomArgs()
        {
        }
        public static new SoftwarePackageVersionSbomArgs Empty => new SoftwarePackageVersionSbomArgs();
    }
}
