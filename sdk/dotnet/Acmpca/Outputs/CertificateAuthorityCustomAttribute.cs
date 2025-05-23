// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Outputs
{

    /// <summary>
    /// Structure that contains X.500 attribute type and value.
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityCustomAttribute
    {
        public readonly string ObjectIdentifier;
        public readonly string Value;

        [OutputConstructor]
        private CertificateAuthorityCustomAttribute(
            string objectIdentifier,

            string value)
        {
            ObjectIdentifier = objectIdentifier;
            Value = value;
        }
    }
}
