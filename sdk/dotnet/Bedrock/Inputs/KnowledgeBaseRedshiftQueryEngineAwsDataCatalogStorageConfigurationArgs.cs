// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Configurations for Redshift query engine AWS Data Catalog backed storage
    /// </summary>
    public sealed class KnowledgeBaseRedshiftQueryEngineAwsDataCatalogStorageConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("tableNames", required: true)]
        private InputList<string>? _tableNames;
        public InputList<string> TableNames
        {
            get => _tableNames ?? (_tableNames = new InputList<string>());
            set => _tableNames = value;
        }

        public KnowledgeBaseRedshiftQueryEngineAwsDataCatalogStorageConfigurationArgs()
        {
        }
        public static new KnowledgeBaseRedshiftQueryEngineAwsDataCatalogStorageConfigurationArgs Empty => new KnowledgeBaseRedshiftQueryEngineAwsDataCatalogStorageConfigurationArgs();
    }
}
