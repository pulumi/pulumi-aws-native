// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::UsagePlanKey
 */
export class UsagePlanKey extends pulumi.CustomResource {
    /**
     * Get an existing UsagePlanKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UsagePlanKey {
        return new UsagePlanKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:UsagePlanKey';

    /**
     * Returns true if the given object is an instance of UsagePlanKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsagePlanKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsagePlanKey.__pulumiType;
    }

    /**
     * The ID of the usage plan key.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The type of usage plan key. Currently, the only valid key type is API_KEY.
     */
    public readonly keyType!: pulumi.Output<enums.apigateway.UsagePlanKeyKeyType>;
    /**
     * The ID of the usage plan.
     */
    public readonly usagePlanId!: pulumi.Output<string>;

    /**
     * Create a UsagePlanKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UsagePlanKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            if ((!args || args.keyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyType'");
            }
            if ((!args || args.usagePlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'usagePlanId'");
            }
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["usagePlanId"] = args ? args.usagePlanId : undefined;
        } else {
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["keyType"] = undefined /*out*/;
            resourceInputs["usagePlanId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UsagePlanKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UsagePlanKey resource.
 */
export interface UsagePlanKeyArgs {
    /**
     * The ID of the usage plan key.
     */
    keyId: pulumi.Input<string>;
    /**
     * The type of usage plan key. Currently, the only valid key type is API_KEY.
     */
    keyType: pulumi.Input<enums.apigateway.UsagePlanKeyKeyType>;
    /**
     * The ID of the usage plan.
     */
    usagePlanId: pulumi.Input<string>;
}