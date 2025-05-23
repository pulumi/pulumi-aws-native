// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    /// <summary>
    /// The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
    /// </summary>
    public sealed class TemplateSsmDocumentDetailsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or Amazon Resource Name (ARN) of the SSM document to use to create a conformance pack. If you use the document name, AWS Config checks only your account and AWS Region for the SSM document.
        /// </summary>
        [Input("documentName")]
        public Input<string>? DocumentName { get; set; }

        /// <summary>
        /// The version of the SSM document to use to create a conformance pack. By default, AWS Config uses the latest version.
        /// 
        /// &gt; This field is optional.
        /// </summary>
        [Input("documentVersion")]
        public Input<string>? DocumentVersion { get; set; }

        public TemplateSsmDocumentDetailsPropertiesArgs()
        {
        }
        public static new TemplateSsmDocumentDetailsPropertiesArgs Empty => new TemplateSsmDocumentDetailsPropertiesArgs();
    }
}
