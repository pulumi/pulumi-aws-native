// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private certificate authority.
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/acmpca"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rootCA, err := acmpca.NewCertificateAuthority(ctx, "rootCA", &acmpca.CertificateAuthorityArgs{
//				Type:             pulumi.String("ROOT"),
//				KeyAlgorithm:     pulumi.String("RSA_2048"),
//				SigningAlgorithm: pulumi.String("SHA256WITHRSA"),
//				Subject: &acmpca.CertificateAuthoritySubjectArgs{
//					Country:                    pulumi.String("US"),
//					Organization:               pulumi.String("string"),
//					OrganizationalUnit:         pulumi.String("string"),
//					DistinguishedNameQualifier: pulumi.String("string"),
//					State:                      pulumi.String("string"),
//					CommonName:                 pulumi.String("123"),
//					SerialNumber:               pulumi.String("string"),
//					Locality:                   pulumi.String("string"),
//					Title:                      pulumi.String("string"),
//					Surname:                    pulumi.String("string"),
//					GivenName:                  pulumi.String("string"),
//					Initials:                   pulumi.String("DG"),
//					Pseudonym:                  pulumi.String("string"),
//					GenerationQualifier:        pulumi.String("DBG"),
//				},
//				RevocationConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationArgs{
//					CrlConfiguration: &acmpca.CertificateAuthorityCrlConfigurationArgs{
//						Enabled: pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			rootCACertificate, err := acmpca.NewCertificate(ctx, "rootCACertificate", &acmpca.CertificateArgs{
//				CertificateAuthorityArn:   rootCA.ID(),
//				CertificateSigningRequest: rootCA.CertificateSigningRequest,
//				SigningAlgorithm:          pulumi.String("SHA256WITHRSA"),
//				TemplateArn:               pulumi.String("arn:aws:acm-pca:::template/RootCACertificate/V1"),
//				Validity: &acmpca.CertificateValidityArgs{
//					Type:  pulumi.String("DAYS"),
//					Value: pulumi.Float64(100),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			rootCAActivation, err := acmpca.NewCertificateAuthorityActivation(ctx, "rootCAActivation", &acmpca.CertificateAuthorityActivationArgs{
//				CertificateAuthorityArn: rootCA.ID(),
//				Certificate:             rootCACertificate.Certificate,
//				Status:                  pulumi.String("ACTIVE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = acmpca.NewPermission(ctx, "rootCAPermission", &acmpca.PermissionArgs{
//				Actions: pulumi.StringArray{
//					pulumi.String("IssueCertificate"),
//					pulumi.String("GetCertificate"),
//					pulumi.String("ListPermissions"),
//				},
//				CertificateAuthorityArn: rootCA.ID(),
//				Principal:               pulumi.String("acm.amazonaws.com"),
//			})
//			if err != nil {
//				return err
//			}
//			subordinateCAOne, err := acmpca.NewCertificateAuthority(ctx, "subordinateCAOne", &acmpca.CertificateAuthorityArgs{
//				Type:             pulumi.String("SUBORDINATE"),
//				KeyAlgorithm:     pulumi.String("RSA_2048"),
//				SigningAlgorithm: pulumi.String("SHA256WITHRSA"),
//				Subject: &acmpca.CertificateAuthoritySubjectArgs{
//					Country:                    pulumi.String("US"),
//					Organization:               pulumi.String("string"),
//					OrganizationalUnit:         pulumi.String("string"),
//					DistinguishedNameQualifier: pulumi.String("string"),
//					State:                      pulumi.String("string"),
//					CommonName:                 pulumi.String("Sub1"),
//					SerialNumber:               pulumi.String("string"),
//					Locality:                   pulumi.String("string"),
//					Title:                      pulumi.String("string"),
//					Surname:                    pulumi.String("string"),
//					GivenName:                  pulumi.String("string"),
//					Initials:                   pulumi.String("DG"),
//					Pseudonym:                  pulumi.String("string"),
//					GenerationQualifier:        pulumi.String("DBG"),
//				},
//				RevocationConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationArgs{},
//				Tags:                    aws.TagArray{},
//			})
//			if err != nil {
//				return err
//			}
//			subordinateCAOneCACertificate, err := acmpca.NewCertificate(ctx, "subordinateCAOneCACertificate", &acmpca.CertificateArgs{
//				CertificateAuthorityArn:   rootCA.ID(),
//				CertificateSigningRequest: subordinateCAOne.CertificateSigningRequest,
//				SigningAlgorithm:          pulumi.String("SHA256WITHRSA"),
//				TemplateArn:               pulumi.String("arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen3/V1"),
//				Validity: &acmpca.CertificateValidityArgs{
//					Type:  pulumi.String("DAYS"),
//					Value: pulumi.Float64(90),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				rootCAActivation,
//			}))
//			if err != nil {
//				return err
//			}
//			subordinateCAOneActivation, err := acmpca.NewCertificateAuthorityActivation(ctx, "subordinateCAOneActivation", &acmpca.CertificateAuthorityActivationArgs{
//				CertificateAuthorityArn: subordinateCAOne.ID(),
//				Certificate:             subordinateCAOneCACertificate.Certificate,
//				CertificateChain:        rootCAActivation.CompleteCertificateChain,
//				Status:                  pulumi.String("ACTIVE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = acmpca.NewPermission(ctx, "subordinateCAOnePermission", &acmpca.PermissionArgs{
//				Actions: pulumi.StringArray{
//					pulumi.String("IssueCertificate"),
//					pulumi.String("GetCertificate"),
//					pulumi.String("ListPermissions"),
//				},
//				CertificateAuthorityArn: subordinateCAOne.ID(),
//				Principal:               pulumi.String("acm.amazonaws.com"),
//			})
//			if err != nil {
//				return err
//			}
//			subordinateCATwo, err := acmpca.NewCertificateAuthority(ctx, "subordinateCATwo", &acmpca.CertificateAuthorityArgs{
//				Type:             pulumi.String("SUBORDINATE"),
//				KeyAlgorithm:     pulumi.String("RSA_2048"),
//				SigningAlgorithm: pulumi.String("SHA256WITHRSA"),
//				Subject: &acmpca.CertificateAuthoritySubjectArgs{
//					Country:                    pulumi.String("US"),
//					Organization:               pulumi.String("string"),
//					OrganizationalUnit:         pulumi.String("string"),
//					DistinguishedNameQualifier: pulumi.String("string"),
//					State:                      pulumi.String("string"),
//					SerialNumber:               pulumi.String("string"),
//					Locality:                   pulumi.String("string"),
//					Title:                      pulumi.String("string"),
//					Surname:                    pulumi.String("string"),
//					GivenName:                  pulumi.String("string"),
//					Initials:                   pulumi.String("DG"),
//					Pseudonym:                  pulumi.String("string"),
//					GenerationQualifier:        pulumi.String("DBG"),
//				},
//				Tags: aws.TagArray{
//					&aws.TagArgs{
//						Key:   pulumi.String("Key1"),
//						Value: pulumi.String("Value1"),
//					},
//					&aws.TagArgs{
//						Key:   pulumi.String("Key2"),
//						Value: pulumi.String("Value2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			subordinateCATwoCACertificate, err := acmpca.NewCertificate(ctx, "subordinateCATwoCACertificate", &acmpca.CertificateArgs{
//				CertificateAuthorityArn:   subordinateCAOne.ID(),
//				CertificateSigningRequest: subordinateCATwo.CertificateSigningRequest,
//				SigningAlgorithm:          pulumi.String("SHA256WITHRSA"),
//				TemplateArn:               pulumi.String("arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen2/V1"),
//				Validity: &acmpca.CertificateValidityArgs{
//					Type:  pulumi.String("DAYS"),
//					Value: pulumi.Float64(80),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				subordinateCAOneActivation,
//			}))
//			if err != nil {
//				return err
//			}
//			subordinateCATwoActivation, err := acmpca.NewCertificateAuthorityActivation(ctx, "subordinateCATwoActivation", &acmpca.CertificateAuthorityActivationArgs{
//				CertificateAuthorityArn: subordinateCATwo.ID(),
//				Certificate:             subordinateCATwoCACertificate.Certificate,
//				CertificateChain:        subordinateCAOneActivation.CompleteCertificateChain,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = acmpca.NewPermission(ctx, "subordinateCATwoPermission", &acmpca.PermissionArgs{
//				Actions: pulumi.StringArray{
//					pulumi.String("IssueCertificate"),
//					pulumi.String("GetCertificate"),
//					pulumi.String("ListPermissions"),
//				},
//				CertificateAuthorityArn: subordinateCATwo.ID(),
//				Principal:               pulumi.String("acm.amazonaws.com"),
//			})
//			if err != nil {
//				return err
//			}
//			endEntityCertificate, err := acmpca.NewCertificate(ctx, "endEntityCertificate", &acmpca.CertificateArgs{
//				CertificateAuthorityArn: subordinateCATwo.ID(),
//				CertificateSigningRequest: pulumi.String(`-----BEGIN CERTIFICATE REQUEST-----
//
// MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
// BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
// aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
// 9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
// wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
// 1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
// WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
// wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
// BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
// KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
// hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
// Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
// ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
// 29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
// 97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
// -----END CERTIFICATE REQUEST-----`),
//
//				SigningAlgorithm: pulumi.String("SHA256WITHRSA"),
//				Validity: &acmpca.CertificateValidityArgs{
//					Type:  pulumi.String("DAYS"),
//					Value: pulumi.Float64(70),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				subordinateCATwoActivation,
//			}))
//			if err != nil {
//				return err
//			}
//			ctx.Export("completeCertificateChain", subordinateCATwoActivation.CompleteCertificateChain)
//			ctx.Export("certificateArn", endEntityCertificate.Arn)
//			return nil
//		})
//	}
//
// ```
type CertificateAuthority struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the certificate authority.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
	CertificateSigningRequest pulumi.StringOutput `pulumi:"certificateSigningRequest"`
	// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
	CsrExtensions CertificateAuthorityCsrExtensionsPtrOutput `pulumi:"csrExtensions"`
	// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
	KeyStorageSecurityStandard pulumi.StringPtrOutput `pulumi:"keyStorageSecurityStandard"`
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
	RevocationConfiguration CertificateAuthorityRevocationConfigurationPtrOutput `pulumi:"revocationConfiguration"`
	// Algorithm your CA uses to sign certificate requests.
	SigningAlgorithm pulumi.StringOutput `pulumi:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your CA.
	Subject CertificateAuthoritySubjectOutput `pulumi:"subject"`
	// Key-value pairs that will be attached to the new private CA. You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The type of the certificate authority.
	Type pulumi.StringOutput `pulumi:"type"`
	// Usage mode of the ceritificate authority.
	UsageMode pulumi.StringPtrOutput `pulumi:"usageMode"`
}

// NewCertificateAuthority registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthority(ctx *pulumi.Context,
	name string, args *CertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'KeyAlgorithm'")
	}
	if args.SigningAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'SigningAlgorithm'")
	}
	if args.Subject == nil {
		return nil, errors.New("invalid value for required argument 'Subject'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"csrExtensions",
		"keyAlgorithm",
		"keyStorageSecurityStandard",
		"signingAlgorithm",
		"subject",
		"type",
		"usageMode",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertificateAuthority
	err := ctx.RegisterResource("aws-native:acmpca:CertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthority gets an existing CertificateAuthority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityState, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	var resource CertificateAuthority
	err := ctx.ReadResource("aws-native:acmpca:CertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthority resources.
type certificateAuthorityState struct {
}

type CertificateAuthorityState struct {
}

func (CertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityState)(nil)).Elem()
}

type certificateAuthorityArgs struct {
	// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
	CsrExtensions *CertificateAuthorityCsrExtensions `pulumi:"csrExtensions"`
	// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	KeyAlgorithm string `pulumi:"keyAlgorithm"`
	// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
	KeyStorageSecurityStandard *string `pulumi:"keyStorageSecurityStandard"`
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
	RevocationConfiguration *CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	// Algorithm your CA uses to sign certificate requests.
	SigningAlgorithm string `pulumi:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your CA.
	Subject CertificateAuthoritySubject `pulumi:"subject"`
	// Key-value pairs that will be attached to the new private CA. You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags []aws.Tag `pulumi:"tags"`
	// The type of the certificate authority.
	Type string `pulumi:"type"`
	// Usage mode of the ceritificate authority.
	UsageMode *string `pulumi:"usageMode"`
}

// The set of arguments for constructing a CertificateAuthority resource.
type CertificateAuthorityArgs struct {
	// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
	CsrExtensions CertificateAuthorityCsrExtensionsPtrInput
	// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	KeyAlgorithm pulumi.StringInput
	// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
	KeyStorageSecurityStandard pulumi.StringPtrInput
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
	RevocationConfiguration CertificateAuthorityRevocationConfigurationPtrInput
	// Algorithm your CA uses to sign certificate requests.
	SigningAlgorithm pulumi.StringInput
	// Structure that contains X.500 distinguished name information for your CA.
	Subject CertificateAuthoritySubjectInput
	// Key-value pairs that will be attached to the new private CA. You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags aws.TagArrayInput
	// The type of the certificate authority.
	Type pulumi.StringInput
	// Usage mode of the ceritificate authority.
	UsageMode pulumi.StringPtrInput
}

func (CertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityArgs)(nil)).Elem()
}

type CertificateAuthorityInput interface {
	pulumi.Input

	ToCertificateAuthorityOutput() CertificateAuthorityOutput
	ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput
}

func (*CertificateAuthority) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (i *CertificateAuthority) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return i.ToCertificateAuthorityOutputWithContext(context.Background())
}

func (i *CertificateAuthority) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityOutput)
}

type CertificateAuthorityOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return o
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return o
}

// The Amazon Resource Name (ARN) of the certificate authority.
func (o CertificateAuthorityOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
func (o CertificateAuthorityOutput) CertificateSigningRequest() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CertificateSigningRequest }).(pulumi.StringOutput)
}

// Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
func (o CertificateAuthorityOutput) CsrExtensions() CertificateAuthorityCsrExtensionsPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateAuthorityCsrExtensionsPtrOutput { return v.CsrExtensions }).(CertificateAuthorityCsrExtensionsPtrOutput)
}

// Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
func (o CertificateAuthorityOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
func (o CertificateAuthorityOutput) KeyStorageSecurityStandard() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringPtrOutput { return v.KeyStorageSecurityStandard }).(pulumi.StringPtrOutput)
}

// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
func (o CertificateAuthorityOutput) RevocationConfiguration() CertificateAuthorityRevocationConfigurationPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateAuthorityRevocationConfigurationPtrOutput {
		return v.RevocationConfiguration
	}).(CertificateAuthorityRevocationConfigurationPtrOutput)
}

// Algorithm your CA uses to sign certificate requests.
func (o CertificateAuthorityOutput) SigningAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.SigningAlgorithm }).(pulumi.StringOutput)
}

// Structure that contains X.500 distinguished name information for your CA.
func (o CertificateAuthorityOutput) Subject() CertificateAuthoritySubjectOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateAuthoritySubjectOutput { return v.Subject }).(CertificateAuthoritySubjectOutput)
}

// Key-value pairs that will be attached to the new private CA. You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
func (o CertificateAuthorityOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *CertificateAuthority) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The type of the certificate authority.
func (o CertificateAuthorityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Usage mode of the ceritificate authority.
func (o CertificateAuthorityOutput) UsageMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringPtrOutput { return v.UsageMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityInput)(nil)).Elem(), &CertificateAuthority{})
	pulumi.RegisterOutputType(CertificateAuthorityOutput{})
}
