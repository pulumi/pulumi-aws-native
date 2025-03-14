// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Test and Publish a resource that has been registered in the CloudFormation Registry.
 */
export class PublicTypeVersion extends pulumi.CustomResource {
    /**
     * Get an existing PublicTypeVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublicTypeVersion {
        return new PublicTypeVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:PublicTypeVersion';

    /**
     * Returns true if the given object is an instance of PublicTypeVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicTypeVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicTypeVersion.__pulumiType;
    }

    /**
     * The Amazon Resource Number (ARN) of the extension.
     */
    public readonly arn!: pulumi.Output<string | undefined>;
    /**
     * A url to the S3 bucket where logs for the testType run will be available
     */
    public readonly logDeliveryBucket!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Number (ARN) assigned to the public extension upon publication
     */
    public /*out*/ readonly publicTypeArn!: pulumi.Output<string>;
    /**
     * The version number of a public third-party extension
     */
    public readonly publicVersionNumber!: pulumi.Output<string | undefined>;
    /**
     * The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
     */
    public /*out*/ readonly publisherId!: pulumi.Output<string>;
    /**
     * The kind of extension
     */
    public readonly type!: pulumi.Output<enums.cloudformation.PublicTypeVersionType | undefined>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    public readonly typeName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Number (ARN) of the extension with the versionId.
     */
    public /*out*/ readonly typeVersionArn!: pulumi.Output<string>;

    /**
     * Create a PublicTypeVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PublicTypeVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["logDeliveryBucket"] = args ? args.logDeliveryBucket : undefined;
            resourceInputs["publicVersionNumber"] = args ? args.publicVersionNumber : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["typeName"] = args ? args.typeName : undefined;
            resourceInputs["publicTypeArn"] = undefined /*out*/;
            resourceInputs["publisherId"] = undefined /*out*/;
            resourceInputs["typeVersionArn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["logDeliveryBucket"] = undefined /*out*/;
            resourceInputs["publicTypeArn"] = undefined /*out*/;
            resourceInputs["publicVersionNumber"] = undefined /*out*/;
            resourceInputs["publisherId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["typeName"] = undefined /*out*/;
            resourceInputs["typeVersionArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["arn", "logDeliveryBucket", "publicVersionNumber", "type", "typeName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PublicTypeVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublicTypeVersion resource.
 */
export interface PublicTypeVersionArgs {
    /**
     * The Amazon Resource Number (ARN) of the extension.
     */
    arn?: pulumi.Input<string>;
    /**
     * A url to the S3 bucket where logs for the testType run will be available
     */
    logDeliveryBucket?: pulumi.Input<string>;
    /**
     * The version number of a public third-party extension
     */
    publicVersionNumber?: pulumi.Input<string>;
    /**
     * The kind of extension
     */
    type?: pulumi.Input<enums.cloudformation.PublicTypeVersionType>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    typeName?: pulumi.Input<string>;
}
