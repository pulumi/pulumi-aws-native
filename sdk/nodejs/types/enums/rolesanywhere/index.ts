// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const TrustAnchorType = {
    AwsAcmPca: "AWS_ACM_PCA",
    CertificateBundle: "CERTIFICATE_BUNDLE",
    SelfSignedRepository: "SELF_SIGNED_REPOSITORY",
} as const;

export type TrustAnchorType = (typeof TrustAnchorType)[keyof typeof TrustAnchorType];