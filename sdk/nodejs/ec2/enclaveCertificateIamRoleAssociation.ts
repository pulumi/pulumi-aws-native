// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate. This association is based on Amazon Resource Names and it enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave.
 */
export class EnclaveCertificateIamRoleAssociation extends pulumi.CustomResource {
    /**
     * Get an existing EnclaveCertificateIamRoleAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnclaveCertificateIamRoleAssociation {
        return new EnclaveCertificateIamRoleAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:EnclaveCertificateIamRoleAssociation';

    /**
     * Returns true if the given object is an instance of EnclaveCertificateIamRoleAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnclaveCertificateIamRoleAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnclaveCertificateIamRoleAssociation.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.
     */
    public readonly certificateArn!: pulumi.Output<string>;
    /**
     * The name of the Amazon S3 bucket to which the certificate was uploaded.
     */
    public /*out*/ readonly certificateS3BucketName!: pulumi.Output<string>;
    /**
     * The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.
     */
    public /*out*/ readonly certificateS3ObjectKey!: pulumi.Output<string>;
    /**
     * The ID of the AWS KMS CMK used to encrypt the private key of the certificate.
     */
    public /*out*/ readonly encryptionKmsKeyId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.
     */
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a EnclaveCertificateIamRoleAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnclaveCertificateIamRoleAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificateArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateArn'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["certificateArn"] = args ? args.certificateArn : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["certificateS3BucketName"] = undefined /*out*/;
            resourceInputs["certificateS3ObjectKey"] = undefined /*out*/;
            resourceInputs["encryptionKmsKeyId"] = undefined /*out*/;
        } else {
            resourceInputs["certificateArn"] = undefined /*out*/;
            resourceInputs["certificateS3BucketName"] = undefined /*out*/;
            resourceInputs["certificateS3ObjectKey"] = undefined /*out*/;
            resourceInputs["encryptionKmsKeyId"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnclaveCertificateIamRoleAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnclaveCertificateIamRoleAssociation resource.
 */
export interface EnclaveCertificateIamRoleAssociationArgs {
    /**
     * The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.
     */
    certificateArn: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.
     */
    roleArn: pulumi.Input<string>;
}