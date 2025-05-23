// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Chatbot.Inputs
{

    public sealed class CustomActionAttachmentCriteriaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation to perform on the named variable.
        /// </summary>
        [Input("operator", required: true)]
        public Input<Pulumi.AwsNative.Chatbot.CustomActionAttachmentCriteriaOperator> Operator { get; set; } = null!;

        /// <summary>
        /// A value that is compared with the actual value of the variable based on the behavior of the operator.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The name of the variable to operate on.
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        public CustomActionAttachmentCriteriaArgs()
        {
        }
        public static new CustomActionAttachmentCriteriaArgs Empty => new CustomActionAttachmentCriteriaArgs();
    }
}
