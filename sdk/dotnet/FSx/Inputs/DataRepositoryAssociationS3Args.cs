// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    /// <summary>
    /// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
    /// </summary>
    public sealed class DataRepositoryAssociationS3Args : global::Pulumi.ResourceArgs
    {
        [Input("autoExportPolicy")]
        public Input<Inputs.DataRepositoryAssociationAutoExportPolicyArgs>? AutoExportPolicy { get; set; }

        [Input("autoImportPolicy")]
        public Input<Inputs.DataRepositoryAssociationAutoImportPolicyArgs>? AutoImportPolicy { get; set; }

        public DataRepositoryAssociationS3Args()
        {
        }
        public static new DataRepositoryAssociationS3Args Empty => new DataRepositoryAssociationS3Args();
    }
}