// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("documentDataFieldName", required: true)]
        public Input<string> DocumentDataFieldName { get; set; } = null!;

        [Input("documentTitleFieldName")]
        public Input<string>? DocumentTitleFieldName { get; set; }

        [Input("fieldMappings")]
        private InputList<Inputs.DataSourceToIndexFieldMappingArgs>? _fieldMappings;
        public InputList<Inputs.DataSourceToIndexFieldMappingArgs> FieldMappings
        {
            get => _fieldMappings ?? (_fieldMappings = new InputList<Inputs.DataSourceToIndexFieldMappingArgs>());
            set => _fieldMappings = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs()
        {
        }
        public static new DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs Empty => new DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs();
    }
}