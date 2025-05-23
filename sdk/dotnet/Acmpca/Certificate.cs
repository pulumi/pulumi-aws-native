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
    /// The ``AWS::ACMPCA::Certificate`` resource is used to issue a certificate using your private certificate authority. For more information, see the [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html) action.
    /// </summary>
    [AwsNativeResourceType("aws-native:acmpca:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies X.509 certificate information to be included in the issued certificate. An ``APIPassthrough`` or ``APICSRPassthrough`` template variant must be selected, or else this parameter is ignored.
        /// </summary>
        [Output("apiPassthrough")]
        public Output<Outputs.CertificateApiPassthrough?> ApiPassthrough { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the issued certificate.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The issued Base64 PEM-encoded certificate.
        /// </summary>
        [Output("certificate")]
        public Output<string> CertificateValue { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA issues the certificate.
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// The certificate signing request (CSR) for the certificate.
        /// </summary>
        [Output("certificateSigningRequest")]
        public Output<string> CertificateSigningRequest { get; private set; } = null!;

        /// <summary>
        /// The name of the algorithm that will be used to sign the certificate to be issued. 
        ///  This parameter should not be confused with the ``SigningAlgorithm`` parameter used to sign a CSR in the ``CreateCertificateAuthority`` action.
        ///   The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
        /// </summary>
        [Output("signingAlgorithm")]
        public Output<string> SigningAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, PCAshort defaults to the ``EndEntityCertificate/V1`` template. For more information about PCAshort templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).
        /// </summary>
        [Output("templateArn")]
        public Output<string?> TemplateArn { get; private set; } = null!;

        /// <summary>
        /// The period of time during which the certificate will be valid.
        /// </summary>
        [Output("validity")]
        public Output<Outputs.CertificateValidity> Validity { get; private set; } = null!;

        /// <summary>
        /// Information describing the start of the validity period of the certificate. This parameter sets the "Not Before" date for the certificate.
        ///  By default, when issuing a certificate, PCAshort sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ``ValidityNotBefore`` parameter can be used to customize the "Not Before" value. 
        ///  Unlike the ``Validity`` parameter, the ``ValidityNotBefore`` parameter is optional.
        ///  The ``ValidityNotBefore`` value is expressed as an explicit date and time, using the ``Validity`` type value ``ABSOLUTE``.
        /// </summary>
        [Output("validityNotBefore")]
        public Output<Outputs.CertificateValidity?> ValidityNotBefore { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiPassthrough",
                    "certificateAuthorityArn",
                    "certificateSigningRequest",
                    "signingAlgorithm",
                    "templateArn",
                    "validity",
                    "validityNotBefore",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies X.509 certificate information to be included in the issued certificate. An ``APIPassthrough`` or ``APICSRPassthrough`` template variant must be selected, or else this parameter is ignored.
        /// </summary>
        [Input("apiPassthrough")]
        public Input<Inputs.CertificateApiPassthroughArgs>? ApiPassthrough { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA issues the certificate.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// The certificate signing request (CSR) for the certificate.
        /// </summary>
        [Input("certificateSigningRequest", required: true)]
        public Input<string> CertificateSigningRequest { get; set; } = null!;

        /// <summary>
        /// The name of the algorithm that will be used to sign the certificate to be issued. 
        ///  This parameter should not be confused with the ``SigningAlgorithm`` parameter used to sign a CSR in the ``CreateCertificateAuthority`` action.
        ///   The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        /// <summary>
        /// Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, PCAshort defaults to the ``EndEntityCertificate/V1`` template. For more information about PCAshort templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).
        /// </summary>
        [Input("templateArn")]
        public Input<string>? TemplateArn { get; set; }

        /// <summary>
        /// The period of time during which the certificate will be valid.
        /// </summary>
        [Input("validity", required: true)]
        public Input<Inputs.CertificateValidityArgs> Validity { get; set; } = null!;

        /// <summary>
        /// Information describing the start of the validity period of the certificate. This parameter sets the "Not Before" date for the certificate.
        ///  By default, when issuing a certificate, PCAshort sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ``ValidityNotBefore`` parameter can be used to customize the "Not Before" value. 
        ///  Unlike the ``Validity`` parameter, the ``ValidityNotBefore`` parameter is optional.
        ///  The ``ValidityNotBefore`` value is expressed as an explicit date and time, using the ``Validity`` type value ``ABSOLUTE``.
        /// </summary>
        [Input("validityNotBefore")]
        public Input<Inputs.CertificateValidityArgs>? ValidityNotBefore { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
