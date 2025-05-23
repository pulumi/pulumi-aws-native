// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleOpenSearchAction
    {
        /// <summary>
        /// The endpoint of your OpenSearch domain.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The unique identifier for the document you are storing.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The OpenSearch index where you want to store your data.
        /// </summary>
        public readonly string Index;
        /// <summary>
        /// The IAM role ARN that has access to OpenSearch.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The type of document you are storing.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TopicRuleOpenSearchAction(
            string endpoint,

            string id,

            string index,

            string roleArn,

            string type)
        {
            Endpoint = endpoint;
            Id = id;
            Index = index;
            RoleArn = roleArn;
            Type = type;
        }
    }
}
