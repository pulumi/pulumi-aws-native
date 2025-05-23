// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecretsManager::SecretTargetAttachment
 */
export class SecretTargetAttachment extends pulumi.CustomResource {
    /**
     * Get an existing SecretTargetAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecretTargetAttachment {
        return new SecretTargetAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:secretsmanager:SecretTargetAttachment';

    /**
     * Returns true if the given object is an instance of SecretTargetAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretTargetAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretTargetAttachment.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The ARN or name of the secret. To reference a secret also created in this template, use the see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID. This field is unique for each target attachment definition.
     */
    public readonly secretId!: pulumi.Output<string>;
    /**
     * The ID of the database or cluster.
     */
    public readonly targetId!: pulumi.Output<string>;
    /**
     * A string that defines the type of service or database associated with the secret. This value instructs Secrets Manager how to update the secret with the details of the service or database. This value must be one of the following:
     *
     * - AWS::RDS::DBInstance
     * - AWS::RDS::DBCluster
     * - AWS::Redshift::Cluster
     * - AWS::RedshiftServerless::Namespace
     * - AWS::DocDB::DBInstance
     * - AWS::DocDB::DBCluster
     * - AWS::DocDBElastic::Cluster
     */
    public readonly targetType!: pulumi.Output<string>;

    /**
     * Create a SecretTargetAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretTargetAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.secretId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretId'");
            }
            if ((!args || args.targetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetId'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            resourceInputs["secretId"] = args ? args.secretId : undefined;
            resourceInputs["targetId"] = args ? args.targetId : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["secretId"] = undefined /*out*/;
            resourceInputs["targetId"] = undefined /*out*/;
            resourceInputs["targetType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["secretId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SecretTargetAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecretTargetAttachment resource.
 */
export interface SecretTargetAttachmentArgs {
    /**
     * The ARN or name of the secret. To reference a secret also created in this template, use the see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID. This field is unique for each target attachment definition.
     */
    secretId: pulumi.Input<string>;
    /**
     * The ID of the database or cluster.
     */
    targetId: pulumi.Input<string>;
    /**
     * A string that defines the type of service or database associated with the secret. This value instructs Secrets Manager how to update the secret with the details of the service or database. This value must be one of the following:
     *
     * - AWS::RDS::DBInstance
     * - AWS::RDS::DBCluster
     * - AWS::Redshift::Cluster
     * - AWS::RedshiftServerless::Namespace
     * - AWS::DocDB::DBInstance
     * - AWS::DocDB::DBCluster
     * - AWS::DocDBElastic::Cluster
     */
    targetType: pulumi.Input<string>;
}
