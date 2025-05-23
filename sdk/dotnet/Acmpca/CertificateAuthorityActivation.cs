// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca
{
    /// <summary>
    /// Used to install the certificate authority certificate and update the certificate authority status.
    /// </summary>
    [AwsNativeResourceType("aws-native:acmpca:CertificateAuthorityActivation")]
    public partial class CertificateAuthorityActivation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate Authority certificate that will be installed in the Certificate Authority.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Arn of the Certificate Authority.
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// Certificate chain for the Certificate Authority certificate.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// The complete certificate chain, including the Certificate Authority certificate.
        /// </summary>
        [Output("completeCertificateChain")]
        public Output<string> CompleteCertificateChain { get; private set; } = null!;

        /// <summary>
        /// The status of the Certificate Authority.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthorityActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthorityActivation(string name, CertificateAuthorityActivationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:CertificateAuthorityActivation", name, args ?? new CertificateAuthorityActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthorityActivation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:CertificateAuthorityActivation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateAuthorityArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateAuthorityActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthorityActivation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateAuthorityActivation(name, id, options);
        }
    }

    public sealed class CertificateAuthorityActivationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate Authority certificate that will be installed in the Certificate Authority.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// Arn of the Certificate Authority.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// Certificate chain for the Certificate Authority certificate.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The status of the Certificate Authority.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CertificateAuthorityActivationArgs()
        {
        }
        public static new CertificateAuthorityActivationArgs Empty => new CertificateAuthorityActivationArgs();
    }
}
