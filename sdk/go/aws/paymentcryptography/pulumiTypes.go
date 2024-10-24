// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package paymentcryptography

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type KeyAttributes struct {
	// The key algorithm to be use during creation of an AWS Payment Cryptography key.
	//
	// For symmetric keys, AWS Payment Cryptography supports `AES` and `TDES` algorithms. For asymmetric keys, AWS Payment Cryptography supports `RSA` and `ECC_NIST` algorithms.
	KeyAlgorithm KeyAlgorithm `pulumi:"keyAlgorithm"`
	// The type of AWS Payment Cryptography key to create, which determines the classiﬁcation of the cryptographic method and whether AWS Payment Cryptography key contains a symmetric key or an asymmetric key pair.
	KeyClass KeyClass `pulumi:"keyClass"`
	// The list of cryptographic operations that you can perform using the key.
	KeyModesOfUse KeyModesOfUse `pulumi:"keyModesOfUse"`
	// The cryptographic usage of an AWS Payment Cryptography key as deﬁned in section A.5.2 of the TR-31 spec.
	KeyUsage KeyUsage `pulumi:"keyUsage"`
}

// KeyAttributesInput is an input type that accepts KeyAttributesArgs and KeyAttributesOutput values.
// You can construct a concrete instance of `KeyAttributesInput` via:
//
//	KeyAttributesArgs{...}
type KeyAttributesInput interface {
	pulumi.Input

	ToKeyAttributesOutput() KeyAttributesOutput
	ToKeyAttributesOutputWithContext(context.Context) KeyAttributesOutput
}

type KeyAttributesArgs struct {
	// The key algorithm to be use during creation of an AWS Payment Cryptography key.
	//
	// For symmetric keys, AWS Payment Cryptography supports `AES` and `TDES` algorithms. For asymmetric keys, AWS Payment Cryptography supports `RSA` and `ECC_NIST` algorithms.
	KeyAlgorithm KeyAlgorithmInput `pulumi:"keyAlgorithm"`
	// The type of AWS Payment Cryptography key to create, which determines the classiﬁcation of the cryptographic method and whether AWS Payment Cryptography key contains a symmetric key or an asymmetric key pair.
	KeyClass KeyClassInput `pulumi:"keyClass"`
	// The list of cryptographic operations that you can perform using the key.
	KeyModesOfUse KeyModesOfUseInput `pulumi:"keyModesOfUse"`
	// The cryptographic usage of an AWS Payment Cryptography key as deﬁned in section A.5.2 of the TR-31 spec.
	KeyUsage KeyUsageInput `pulumi:"keyUsage"`
}

func (KeyAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (i KeyAttributesArgs) ToKeyAttributesOutput() KeyAttributesOutput {
	return i.ToKeyAttributesOutputWithContext(context.Background())
}

func (i KeyAttributesArgs) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesOutput)
}

type KeyAttributesOutput struct{ *pulumi.OutputState }

func (KeyAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesOutput) ToKeyAttributesOutput() KeyAttributesOutput {
	return o
}

func (o KeyAttributesOutput) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return o
}

// The key algorithm to be use during creation of an AWS Payment Cryptography key.
//
// For symmetric keys, AWS Payment Cryptography supports `AES` and `TDES` algorithms. For asymmetric keys, AWS Payment Cryptography supports `RSA` and `ECC_NIST` algorithms.
func (o KeyAttributesOutput) KeyAlgorithm() KeyAlgorithmOutput {
	return o.ApplyT(func(v KeyAttributes) KeyAlgorithm { return v.KeyAlgorithm }).(KeyAlgorithmOutput)
}

// The type of AWS Payment Cryptography key to create, which determines the classiﬁcation of the cryptographic method and whether AWS Payment Cryptography key contains a symmetric key or an asymmetric key pair.
func (o KeyAttributesOutput) KeyClass() KeyClassOutput {
	return o.ApplyT(func(v KeyAttributes) KeyClass { return v.KeyClass }).(KeyClassOutput)
}

// The list of cryptographic operations that you can perform using the key.
func (o KeyAttributesOutput) KeyModesOfUse() KeyModesOfUseOutput {
	return o.ApplyT(func(v KeyAttributes) KeyModesOfUse { return v.KeyModesOfUse }).(KeyModesOfUseOutput)
}

// The cryptographic usage of an AWS Payment Cryptography key as deﬁned in section A.5.2 of the TR-31 spec.
func (o KeyAttributesOutput) KeyUsage() KeyUsageOutput {
	return o.ApplyT(func(v KeyAttributes) KeyUsage { return v.KeyUsage }).(KeyUsageOutput)
}

type KeyAttributesPtrOutput struct{ *pulumi.OutputState }

func (KeyAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) Elem() KeyAttributesOutput {
	return o.ApplyT(func(v *KeyAttributes) KeyAttributes {
		if v != nil {
			return *v
		}
		var ret KeyAttributes
		return ret
	}).(KeyAttributesOutput)
}

// The key algorithm to be use during creation of an AWS Payment Cryptography key.
//
// For symmetric keys, AWS Payment Cryptography supports `AES` and `TDES` algorithms. For asymmetric keys, AWS Payment Cryptography supports `RSA` and `ECC_NIST` algorithms.
func (o KeyAttributesPtrOutput) KeyAlgorithm() KeyAlgorithmPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *KeyAlgorithm {
		if v == nil {
			return nil
		}
		return &v.KeyAlgorithm
	}).(KeyAlgorithmPtrOutput)
}

// The type of AWS Payment Cryptography key to create, which determines the classiﬁcation of the cryptographic method and whether AWS Payment Cryptography key contains a symmetric key or an asymmetric key pair.
func (o KeyAttributesPtrOutput) KeyClass() KeyClassPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *KeyClass {
		if v == nil {
			return nil
		}
		return &v.KeyClass
	}).(KeyClassPtrOutput)
}

// The list of cryptographic operations that you can perform using the key.
func (o KeyAttributesPtrOutput) KeyModesOfUse() KeyModesOfUsePtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *KeyModesOfUse {
		if v == nil {
			return nil
		}
		return &v.KeyModesOfUse
	}).(KeyModesOfUsePtrOutput)
}

// The cryptographic usage of an AWS Payment Cryptography key as deﬁned in section A.5.2 of the TR-31 spec.
func (o KeyAttributesPtrOutput) KeyUsage() KeyUsagePtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *KeyUsage {
		if v == nil {
			return nil
		}
		return &v.KeyUsage
	}).(KeyUsagePtrOutput)
}

type KeyModesOfUse struct {
	// Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.
	Decrypt *bool `pulumi:"decrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.
	DeriveKey *bool `pulumi:"deriveKey"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.
	Encrypt *bool `pulumi:"encrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
	Generate *bool `pulumi:"generate"`
	// Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by `KeyUsage` .
	NoRestrictions *bool `pulumi:"noRestrictions"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used for signing.
	Sign   *bool `pulumi:"sign"`
	Unwrap *bool `pulumi:"unwrap"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.
	Verify *bool `pulumi:"verify"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.
	Wrap *bool `pulumi:"wrap"`
}

// KeyModesOfUseInput is an input type that accepts KeyModesOfUseArgs and KeyModesOfUseOutput values.
// You can construct a concrete instance of `KeyModesOfUseInput` via:
//
//	KeyModesOfUseArgs{...}
type KeyModesOfUseInput interface {
	pulumi.Input

	ToKeyModesOfUseOutput() KeyModesOfUseOutput
	ToKeyModesOfUseOutputWithContext(context.Context) KeyModesOfUseOutput
}

type KeyModesOfUseArgs struct {
	// Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.
	Decrypt pulumi.BoolPtrInput `pulumi:"decrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.
	DeriveKey pulumi.BoolPtrInput `pulumi:"deriveKey"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.
	Encrypt pulumi.BoolPtrInput `pulumi:"encrypt"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
	Generate pulumi.BoolPtrInput `pulumi:"generate"`
	// Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by `KeyUsage` .
	NoRestrictions pulumi.BoolPtrInput `pulumi:"noRestrictions"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used for signing.
	Sign   pulumi.BoolPtrInput `pulumi:"sign"`
	Unwrap pulumi.BoolPtrInput `pulumi:"unwrap"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.
	Verify pulumi.BoolPtrInput `pulumi:"verify"`
	// Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.
	Wrap pulumi.BoolPtrInput `pulumi:"wrap"`
}

func (KeyModesOfUseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyModesOfUse)(nil)).Elem()
}

func (i KeyModesOfUseArgs) ToKeyModesOfUseOutput() KeyModesOfUseOutput {
	return i.ToKeyModesOfUseOutputWithContext(context.Background())
}

func (i KeyModesOfUseArgs) ToKeyModesOfUseOutputWithContext(ctx context.Context) KeyModesOfUseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyModesOfUseOutput)
}

type KeyModesOfUseOutput struct{ *pulumi.OutputState }

func (KeyModesOfUseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyModesOfUse)(nil)).Elem()
}

func (o KeyModesOfUseOutput) ToKeyModesOfUseOutput() KeyModesOfUseOutput {
	return o
}

func (o KeyModesOfUseOutput) ToKeyModesOfUseOutputWithContext(ctx context.Context) KeyModesOfUseOutput {
	return o
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.
func (o KeyModesOfUseOutput) Decrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Decrypt }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.
func (o KeyModesOfUseOutput) DeriveKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.DeriveKey }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.
func (o KeyModesOfUseOutput) Encrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Encrypt }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
func (o KeyModesOfUseOutput) Generate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Generate }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by `KeyUsage` .
func (o KeyModesOfUseOutput) NoRestrictions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.NoRestrictions }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used for signing.
func (o KeyModesOfUseOutput) Sign() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Sign }).(pulumi.BoolPtrOutput)
}

func (o KeyModesOfUseOutput) Unwrap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Unwrap }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.
func (o KeyModesOfUseOutput) Verify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Verify }).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.
func (o KeyModesOfUseOutput) Wrap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyModesOfUse) *bool { return v.Wrap }).(pulumi.BoolPtrOutput)
}

type KeyModesOfUsePtrOutput struct{ *pulumi.OutputState }

func (KeyModesOfUsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyModesOfUse)(nil)).Elem()
}

func (o KeyModesOfUsePtrOutput) ToKeyModesOfUsePtrOutput() KeyModesOfUsePtrOutput {
	return o
}

func (o KeyModesOfUsePtrOutput) ToKeyModesOfUsePtrOutputWithContext(ctx context.Context) KeyModesOfUsePtrOutput {
	return o
}

func (o KeyModesOfUsePtrOutput) Elem() KeyModesOfUseOutput {
	return o.ApplyT(func(v *KeyModesOfUse) KeyModesOfUse {
		if v != nil {
			return *v
		}
		var ret KeyModesOfUse
		return ret
	}).(KeyModesOfUseOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.
func (o KeyModesOfUsePtrOutput) Decrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Decrypt
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.
func (o KeyModesOfUsePtrOutput) DeriveKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.DeriveKey
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.
func (o KeyModesOfUsePtrOutput) Encrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Encrypt
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
func (o KeyModesOfUsePtrOutput) Generate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Generate
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by `KeyUsage` .
func (o KeyModesOfUsePtrOutput) NoRestrictions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.NoRestrictions
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used for signing.
func (o KeyModesOfUsePtrOutput) Sign() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Sign
	}).(pulumi.BoolPtrOutput)
}

func (o KeyModesOfUsePtrOutput) Unwrap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Unwrap
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.
func (o KeyModesOfUsePtrOutput) Verify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Verify
	}).(pulumi.BoolPtrOutput)
}

// Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.
func (o KeyModesOfUsePtrOutput) Wrap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyModesOfUse) *bool {
		if v == nil {
			return nil
		}
		return v.Wrap
	}).(pulumi.BoolPtrOutput)
}

type KeyTag struct {
	// The key of the tag.
	Key string `pulumi:"key"`
	// The value of the tag.
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAttributesInput)(nil)).Elem(), KeyAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyModesOfUseInput)(nil)).Elem(), KeyModesOfUseArgs{})
	pulumi.RegisterOutputType(KeyAttributesOutput{})
	pulumi.RegisterOutputType(KeyAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyModesOfUseOutput{})
	pulumi.RegisterOutputType(KeyModesOfUsePtrOutput{})
}
