// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceConfluenceAttachmentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("attachmentFieldMappings")]
        private InputList<Inputs.DataSourceConfluenceAttachmentToIndexFieldMappingArgs>? _attachmentFieldMappings;
        public InputList<Inputs.DataSourceConfluenceAttachmentToIndexFieldMappingArgs> AttachmentFieldMappings
        {
            get => _attachmentFieldMappings ?? (_attachmentFieldMappings = new InputList<Inputs.DataSourceConfluenceAttachmentToIndexFieldMappingArgs>());
            set => _attachmentFieldMappings = value;
        }

        [Input("crawlAttachments")]
        public Input<bool>? CrawlAttachments { get; set; }

        public DataSourceConfluenceAttachmentConfigurationArgs()
        {
        }
        public static new DataSourceConfluenceAttachmentConfigurationArgs Empty => new DataSourceConfluenceAttachmentConfigurationArgs();
    }
}