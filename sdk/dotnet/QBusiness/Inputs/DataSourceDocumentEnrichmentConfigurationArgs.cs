// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class DataSourceDocumentEnrichmentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("inlineConfigurations")]
        private InputList<Inputs.DataSourceInlineDocumentEnrichmentConfigurationArgs>? _inlineConfigurations;

        /// <summary>
        /// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Q Business.
        /// </summary>
        public InputList<Inputs.DataSourceInlineDocumentEnrichmentConfigurationArgs> InlineConfigurations
        {
            get => _inlineConfigurations ?? (_inlineConfigurations = new InputList<Inputs.DataSourceInlineDocumentEnrichmentConfigurationArgs>());
            set => _inlineConfigurations = value;
        }

        /// <summary>
        /// Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted. You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html) .
        /// </summary>
        [Input("postExtractionHookConfiguration")]
        public Input<Inputs.DataSourceHookConfigurationArgs>? PostExtractionHookConfiguration { get; set; }

        /// <summary>
        /// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text. You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html) .
        /// </summary>
        [Input("preExtractionHookConfiguration")]
        public Input<Inputs.DataSourceHookConfigurationArgs>? PreExtractionHookConfiguration { get; set; }

        public DataSourceDocumentEnrichmentConfigurationArgs()
        {
        }
        public static new DataSourceDocumentEnrichmentConfigurationArgs Empty => new DataSourceDocumentEnrichmentConfigurationArgs();
    }
}