// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Comprehend.Inputs
{

    public sealed class DocumentClassifierInputDataConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("augmentedManifests")]
        private InputList<Inputs.DocumentClassifierAugmentedManifestsListItemArgs>? _augmentedManifests;

        /// <summary>
        /// A list of augmented manifest files that provide training data for your custom model. An augmented manifest file is a labeled dataset that is produced by Amazon SageMaker Ground Truth.
        /// 
        /// This parameter is required if you set `DataFormat` to `AUGMENTED_MANIFEST` .
        /// </summary>
        public InputList<Inputs.DocumentClassifierAugmentedManifestsListItemArgs> AugmentedManifests
        {
            get => _augmentedManifests ?? (_augmentedManifests = new InputList<Inputs.DocumentClassifierAugmentedManifestsListItemArgs>());
            set => _augmentedManifests = value;
        }

        /// <summary>
        /// The format of your training data:
        /// 
        /// - `COMPREHEND_CSV` : A two-column CSV file, where labels are provided in the first column, and documents are provided in the second. If you use this value, you must provide the `S3Uri` parameter in your request.
        /// - `AUGMENTED_MANIFEST` : A labeled dataset that is produced by Amazon SageMaker Ground Truth. This file is in JSON lines format. Each line is a complete JSON object that contains a training document and its associated labels.
        /// 
        /// If you use this value, you must provide the `AugmentedManifests` parameter in your request.
        /// 
        /// If you don't specify a value, Amazon Comprehend uses `COMPREHEND_CSV` as the default.
        /// </summary>
        [Input("dataFormat")]
        public Input<Pulumi.AwsNative.Comprehend.DocumentClassifierInputDataConfigDataFormat>? DataFormat { get; set; }

        [Input("documentReaderConfig")]
        public Input<Inputs.DocumentClassifierDocumentReaderConfigArgs>? DocumentReaderConfig { get; set; }

        /// <summary>
        /// The type of input documents for training the model. Provide plain-text documents to create a plain-text model, and provide semi-structured documents to create a native document model.
        /// </summary>
        [Input("documentType")]
        public Input<Pulumi.AwsNative.Comprehend.DocumentClassifierInputDataConfigDocumentType>? DocumentType { get; set; }

        /// <summary>
        /// The S3 location of the training documents. This parameter is required in a request to create a native document model.
        /// </summary>
        [Input("documents")]
        public Input<Inputs.DocumentClassifierDocumentsArgs>? Documents { get; set; }

        /// <summary>
        /// Indicates the delimiter used to separate each label for training a multi-label classifier. The default delimiter between labels is a pipe (|). You can use a different character as a delimiter (if it's an allowed character) by specifying it under Delimiter for labels. If the training documents use a delimiter other than the default or the delimiter you specify, the labels on that line will be combined to make a single unique label, such as LABELLABELLABEL.
        /// </summary>
        [Input("labelDelimiter")]
        public Input<string>? LabelDelimiter { get; set; }

        /// <summary>
        /// The Amazon S3 URI for the input data. The S3 bucket must be in the same Region as the API endpoint that you are calling. The URI can point to a single input file or it can provide the prefix for a collection of input files.
        /// 
        /// For example, if you use the URI `S3://bucketName/prefix` , if the prefix is a single file, Amazon Comprehend uses that file as input. If more than one file begins with the prefix, Amazon Comprehend uses all of them as input.
        /// 
        /// This parameter is required if you set `DataFormat` to `COMPREHEND_CSV` .
        /// </summary>
        [Input("s3Uri")]
        public Input<string>? S3Uri { get; set; }

        /// <summary>
        /// This specifies the Amazon S3 location that contains the test annotations for the document classifier. The URI must be in the same AWS Region as the API endpoint that you are calling.
        /// </summary>
        [Input("testS3Uri")]
        public Input<string>? TestS3Uri { get; set; }

        public DocumentClassifierInputDataConfigArgs()
        {
        }
        public static new DocumentClassifierInputDataConfigArgs Empty => new DocumentClassifierInputDataConfigArgs();
    }
}
