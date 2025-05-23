// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainEncryptionAtRestOptions
    {
        /// <summary>
        /// Specify `true` to enable encryption at rest. Required if you enable fine-grained access control in [AdvancedSecurityOptionsInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html) .
        /// 
        /// If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template, the domain is deleted and recreated in order to modify the property.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The KMS key ID. Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a` . Required if you enable encryption at rest.
        /// 
        /// You can also use `keyAlias` as a value.
        /// 
        /// If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template, the domain is deleted and recreated in order to modify the property.
        /// </summary>
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private DomainEncryptionAtRestOptions(
            bool? enabled,

            string? kmsKeyId)
        {
            Enabled = enabled;
            KmsKeyId = kmsKeyId;
        }
    }
}
