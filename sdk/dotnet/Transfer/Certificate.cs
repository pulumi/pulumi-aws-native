// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer
{
    /// <summary>
    /// Resource Type definition for AWS::Transfer::Certificate
    /// </summary>
    [AwsNativeResourceType("aws-native:transfer:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the active date for the certificate.
        /// </summary>
        [Output("activeDate")]
        public Output<string?> ActiveDate { get; private set; } = null!;

        /// <summary>
        /// Specifies the unique Amazon Resource Name (ARN) for the agreement.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies the certificate body to be imported.
        /// </summary>
        [Output("certificate")]
        public Output<string> CertificateValue { get; private set; } = null!;

        /// <summary>
        /// Specifies the certificate chain to be imported.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the certificate.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// A textual description for the certificate.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the inactive date for the certificate.
        /// </summary>
        [Output("inactiveDate")]
        public Output<string?> InactiveDate { get; private set; } = null!;

        /// <summary>
        /// Specifies the not after date for the certificate.
        /// </summary>
        [Output("notAfterDate")]
        public Output<string> NotAfterDate { get; private set; } = null!;

        /// <summary>
        /// Specifies the not before date for the certificate.
        /// </summary>
        [Output("notBeforeDate")]
        public Output<string> NotBeforeDate { get; private set; } = null!;

        /// <summary>
        /// Specifies the private key for the certificate.
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Specifies Certificate's serial.
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// A status description for the certificate.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Transfer.CertificateStatus> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CertificateTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Describing the type of certificate. With or without a private key.
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.Transfer.CertificateType> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the usage type for the certificate.
        /// </summary>
        [Output("usage")]
        public Output<Pulumi.AwsNative.Transfer.CertificateUsage> Usage { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:transfer:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:transfer:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
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
        /// Specifies the active date for the certificate.
        /// </summary>
        [Input("activeDate")]
        public Input<string>? ActiveDate { get; set; }

        /// <summary>
        /// Specifies the certificate body to be imported.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<string> CertificateValue { get; set; } = null!;

        /// <summary>
        /// Specifies the certificate chain to be imported.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// A textual description for the certificate.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the inactive date for the certificate.
        /// </summary>
        [Input("inactiveDate")]
        public Input<string>? InactiveDate { get; set; }

        /// <summary>
        /// Specifies the private key for the certificate.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("tags")]
        private InputList<Inputs.CertificateTagArgs>? _tags;

        /// <summary>
        /// Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.
        /// </summary>
        public InputList<Inputs.CertificateTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CertificateTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the usage type for the certificate.
        /// </summary>
        [Input("usage", required: true)]
        public Input<Pulumi.AwsNative.Transfer.CertificateUsage> Usage { get; set; } = null!;

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}