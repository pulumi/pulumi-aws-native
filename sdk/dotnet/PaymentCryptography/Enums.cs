// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.PaymentCryptography
{
    [EnumType]
    public readonly struct KeyAlgorithm : IEquatable<KeyAlgorithm>
    {
        private readonly string _value;

        private KeyAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyAlgorithm Tdes2key { get; } = new KeyAlgorithm("TDES_2KEY");
        public static KeyAlgorithm Tdes3key { get; } = new KeyAlgorithm("TDES_3KEY");
        public static KeyAlgorithm Aes128 { get; } = new KeyAlgorithm("AES_128");
        public static KeyAlgorithm Aes192 { get; } = new KeyAlgorithm("AES_192");
        public static KeyAlgorithm Aes256 { get; } = new KeyAlgorithm("AES_256");
        public static KeyAlgorithm HmacSha256 { get; } = new KeyAlgorithm("HMAC_SHA256");
        public static KeyAlgorithm HmacSha384 { get; } = new KeyAlgorithm("HMAC_SHA384");
        public static KeyAlgorithm HmacSha512 { get; } = new KeyAlgorithm("HMAC_SHA512");
        public static KeyAlgorithm HmacSha224 { get; } = new KeyAlgorithm("HMAC_SHA224");
        public static KeyAlgorithm Rsa2048 { get; } = new KeyAlgorithm("RSA_2048");
        public static KeyAlgorithm Rsa3072 { get; } = new KeyAlgorithm("RSA_3072");
        public static KeyAlgorithm Rsa4096 { get; } = new KeyAlgorithm("RSA_4096");
        public static KeyAlgorithm EccNistP256 { get; } = new KeyAlgorithm("ECC_NIST_P256");
        public static KeyAlgorithm EccNistP384 { get; } = new KeyAlgorithm("ECC_NIST_P384");
        public static KeyAlgorithm EccNistP521 { get; } = new KeyAlgorithm("ECC_NIST_P521");

        public static bool operator ==(KeyAlgorithm left, KeyAlgorithm right) => left.Equals(right);
        public static bool operator !=(KeyAlgorithm left, KeyAlgorithm right) => !left.Equals(right);

        public static explicit operator string(KeyAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyAlgorithm other && Equals(other);
        public bool Equals(KeyAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyCheckValueAlgorithm : IEquatable<KeyCheckValueAlgorithm>
    {
        private readonly string _value;

        private KeyCheckValueAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyCheckValueAlgorithm Cmac { get; } = new KeyCheckValueAlgorithm("CMAC");
        public static KeyCheckValueAlgorithm AnsiX924 { get; } = new KeyCheckValueAlgorithm("ANSI_X9_24");
        public static KeyCheckValueAlgorithm Hmac { get; } = new KeyCheckValueAlgorithm("HMAC");
        public static KeyCheckValueAlgorithm Sha1 { get; } = new KeyCheckValueAlgorithm("SHA_1");

        public static bool operator ==(KeyCheckValueAlgorithm left, KeyCheckValueAlgorithm right) => left.Equals(right);
        public static bool operator !=(KeyCheckValueAlgorithm left, KeyCheckValueAlgorithm right) => !left.Equals(right);

        public static explicit operator string(KeyCheckValueAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyCheckValueAlgorithm other && Equals(other);
        public bool Equals(KeyCheckValueAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyClass : IEquatable<KeyClass>
    {
        private readonly string _value;

        private KeyClass(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyClass SymmetricKey { get; } = new KeyClass("SYMMETRIC_KEY");
        public static KeyClass AsymmetricKeyPair { get; } = new KeyClass("ASYMMETRIC_KEY_PAIR");
        public static KeyClass PrivateKey { get; } = new KeyClass("PRIVATE_KEY");
        public static KeyClass PublicKey { get; } = new KeyClass("PUBLIC_KEY");

        public static bool operator ==(KeyClass left, KeyClass right) => left.Equals(right);
        public static bool operator !=(KeyClass left, KeyClass right) => !left.Equals(right);

        public static explicit operator string(KeyClass value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyClass other && Equals(other);
        public bool Equals(KeyClass other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyDeriveKeyUsage : IEquatable<KeyDeriveKeyUsage>
    {
        private readonly string _value;

        private KeyDeriveKeyUsage(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyDeriveKeyUsage Tr31b0BaseDerivationKey { get; } = new KeyDeriveKeyUsage("TR31_B0_BASE_DERIVATION_KEY");
        public static KeyDeriveKeyUsage Tr31c0CardVerificationKey { get; } = new KeyDeriveKeyUsage("TR31_C0_CARD_VERIFICATION_KEY");
        public static KeyDeriveKeyUsage Tr31d0SymmetricDataEncryptionKey { get; } = new KeyDeriveKeyUsage("TR31_D0_SYMMETRIC_DATA_ENCRYPTION_KEY");
        public static KeyDeriveKeyUsage Tr31e0EmvMkeyAppCryptograms { get; } = new KeyDeriveKeyUsage("TR31_E0_EMV_MKEY_APP_CRYPTOGRAMS");
        public static KeyDeriveKeyUsage Tr31e1EmvMkeyConfidentiality { get; } = new KeyDeriveKeyUsage("TR31_E1_EMV_MKEY_CONFIDENTIALITY");
        public static KeyDeriveKeyUsage Tr31e2EmvMkeyIntegrity { get; } = new KeyDeriveKeyUsage("TR31_E2_EMV_MKEY_INTEGRITY");
        public static KeyDeriveKeyUsage Tr31e4EmvMkeyDynamicNumbers { get; } = new KeyDeriveKeyUsage("TR31_E4_EMV_MKEY_DYNAMIC_NUMBERS");
        public static KeyDeriveKeyUsage Tr31e5EmvMkeyCardPersonalization { get; } = new KeyDeriveKeyUsage("TR31_E5_EMV_MKEY_CARD_PERSONALIZATION");
        public static KeyDeriveKeyUsage Tr31e6EmvMkeyOther { get; } = new KeyDeriveKeyUsage("TR31_E6_EMV_MKEY_OTHER");
        public static KeyDeriveKeyUsage Tr31k0KeyEncryptionKey { get; } = new KeyDeriveKeyUsage("TR31_K0_KEY_ENCRYPTION_KEY");
        public static KeyDeriveKeyUsage Tr31k1KeyBlockProtectionKey { get; } = new KeyDeriveKeyUsage("TR31_K1_KEY_BLOCK_PROTECTION_KEY");
        public static KeyDeriveKeyUsage Tr31m3Iso97973MacKey { get; } = new KeyDeriveKeyUsage("TR31_M3_ISO_9797_3_MAC_KEY");
        public static KeyDeriveKeyUsage Tr31m1Iso97971MacKey { get; } = new KeyDeriveKeyUsage("TR31_M1_ISO_9797_1_MAC_KEY");
        public static KeyDeriveKeyUsage Tr31m6Iso97975CmacKey { get; } = new KeyDeriveKeyUsage("TR31_M6_ISO_9797_5_CMAC_KEY");
        public static KeyDeriveKeyUsage Tr31m7HmacKey { get; } = new KeyDeriveKeyUsage("TR31_M7_HMAC_KEY");
        public static KeyDeriveKeyUsage Tr31p0PinEncryptionKey { get; } = new KeyDeriveKeyUsage("TR31_P0_PIN_ENCRYPTION_KEY");
        public static KeyDeriveKeyUsage Tr31p1PinGenerationKey { get; } = new KeyDeriveKeyUsage("TR31_P1_PIN_GENERATION_KEY");
        public static KeyDeriveKeyUsage Tr31v1Ibm3624PinVerificationKey { get; } = new KeyDeriveKeyUsage("TR31_V1_IBM3624_PIN_VERIFICATION_KEY");
        public static KeyDeriveKeyUsage Tr31v2VisaPinVerificationKey { get; } = new KeyDeriveKeyUsage("TR31_V2_VISA_PIN_VERIFICATION_KEY");

        public static bool operator ==(KeyDeriveKeyUsage left, KeyDeriveKeyUsage right) => left.Equals(right);
        public static bool operator !=(KeyDeriveKeyUsage left, KeyDeriveKeyUsage right) => !left.Equals(right);

        public static explicit operator string(KeyDeriveKeyUsage value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyDeriveKeyUsage other && Equals(other);
        public bool Equals(KeyDeriveKeyUsage other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines the source of a key
    /// </summary>
    [EnumType]
    public readonly struct KeyOrigin : IEquatable<KeyOrigin>
    {
        private readonly string _value;

        private KeyOrigin(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyOrigin External { get; } = new KeyOrigin("EXTERNAL");
        public static KeyOrigin AwsPaymentCryptography { get; } = new KeyOrigin("AWS_PAYMENT_CRYPTOGRAPHY");

        public static bool operator ==(KeyOrigin left, KeyOrigin right) => left.Equals(right);
        public static bool operator !=(KeyOrigin left, KeyOrigin right) => !left.Equals(right);

        public static explicit operator string(KeyOrigin value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyOrigin other && Equals(other);
        public bool Equals(KeyOrigin other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines the state of a key
    /// </summary>
    [EnumType]
    public readonly struct KeyState : IEquatable<KeyState>
    {
        private readonly string _value;

        private KeyState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyState CreateInProgress { get; } = new KeyState("CREATE_IN_PROGRESS");
        public static KeyState CreateComplete { get; } = new KeyState("CREATE_COMPLETE");
        public static KeyState DeletePending { get; } = new KeyState("DELETE_PENDING");
        public static KeyState DeleteComplete { get; } = new KeyState("DELETE_COMPLETE");

        public static bool operator ==(KeyState left, KeyState right) => left.Equals(right);
        public static bool operator !=(KeyState left, KeyState right) => !left.Equals(right);

        public static explicit operator string(KeyState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyState other && Equals(other);
        public bool Equals(KeyState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyUsage : IEquatable<KeyUsage>
    {
        private readonly string _value;

        private KeyUsage(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyUsage Tr31b0BaseDerivationKey { get; } = new KeyUsage("TR31_B0_BASE_DERIVATION_KEY");
        public static KeyUsage Tr31c0CardVerificationKey { get; } = new KeyUsage("TR31_C0_CARD_VERIFICATION_KEY");
        public static KeyUsage Tr31d0SymmetricDataEncryptionKey { get; } = new KeyUsage("TR31_D0_SYMMETRIC_DATA_ENCRYPTION_KEY");
        public static KeyUsage Tr31d1AsymmetricKeyForDataEncryption { get; } = new KeyUsage("TR31_D1_ASYMMETRIC_KEY_FOR_DATA_ENCRYPTION");
        public static KeyUsage Tr31e0EmvMkeyAppCryptograms { get; } = new KeyUsage("TR31_E0_EMV_MKEY_APP_CRYPTOGRAMS");
        public static KeyUsage Tr31e1EmvMkeyConfidentiality { get; } = new KeyUsage("TR31_E1_EMV_MKEY_CONFIDENTIALITY");
        public static KeyUsage Tr31e2EmvMkeyIntegrity { get; } = new KeyUsage("TR31_E2_EMV_MKEY_INTEGRITY");
        public static KeyUsage Tr31e4EmvMkeyDynamicNumbers { get; } = new KeyUsage("TR31_E4_EMV_MKEY_DYNAMIC_NUMBERS");
        public static KeyUsage Tr31e5EmvMkeyCardPersonalization { get; } = new KeyUsage("TR31_E5_EMV_MKEY_CARD_PERSONALIZATION");
        public static KeyUsage Tr31e6EmvMkeyOther { get; } = new KeyUsage("TR31_E6_EMV_MKEY_OTHER");
        public static KeyUsage Tr31k0KeyEncryptionKey { get; } = new KeyUsage("TR31_K0_KEY_ENCRYPTION_KEY");
        public static KeyUsage Tr31k1KeyBlockProtectionKey { get; } = new KeyUsage("TR31_K1_KEY_BLOCK_PROTECTION_KEY");
        public static KeyUsage Tr31k3AsymmetricKeyForKeyAgreement { get; } = new KeyUsage("TR31_K3_ASYMMETRIC_KEY_FOR_KEY_AGREEMENT");
        public static KeyUsage Tr31m3Iso97973MacKey { get; } = new KeyUsage("TR31_M3_ISO_9797_3_MAC_KEY");
        public static KeyUsage Tr31m1Iso97971MacKey { get; } = new KeyUsage("TR31_M1_ISO_9797_1_MAC_KEY");
        public static KeyUsage Tr31m6Iso97975CmacKey { get; } = new KeyUsage("TR31_M6_ISO_9797_5_CMAC_KEY");
        public static KeyUsage Tr31m7HmacKey { get; } = new KeyUsage("TR31_M7_HMAC_KEY");
        public static KeyUsage Tr31p0PinEncryptionKey { get; } = new KeyUsage("TR31_P0_PIN_ENCRYPTION_KEY");
        public static KeyUsage Tr31p1PinGenerationKey { get; } = new KeyUsage("TR31_P1_PIN_GENERATION_KEY");
        public static KeyUsage Tr31s0AsymmetricKeyForDigitalSignature { get; } = new KeyUsage("TR31_S0_ASYMMETRIC_KEY_FOR_DIGITAL_SIGNATURE");
        public static KeyUsage Tr31v1Ibm3624PinVerificationKey { get; } = new KeyUsage("TR31_V1_IBM3624_PIN_VERIFICATION_KEY");
        public static KeyUsage Tr31v2VisaPinVerificationKey { get; } = new KeyUsage("TR31_V2_VISA_PIN_VERIFICATION_KEY");
        public static KeyUsage Tr31k2Tr34AsymmetricKey { get; } = new KeyUsage("TR31_K2_TR34_ASYMMETRIC_KEY");

        public static bool operator ==(KeyUsage left, KeyUsage right) => left.Equals(right);
        public static bool operator !=(KeyUsage left, KeyUsage right) => !left.Equals(right);

        public static explicit operator string(KeyUsage value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyUsage other && Equals(other);
        public bool Equals(KeyUsage other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
