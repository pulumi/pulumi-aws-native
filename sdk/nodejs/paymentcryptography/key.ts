// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::PaymentCryptography::Key Resource Type
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:paymentcryptography:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    /**
     * Specifies whether the key is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether the key is exportable. This data is immutable after the key is created.
     */
    public readonly exportable!: pulumi.Output<boolean>;
    /**
     * The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
     */
    public readonly keyAttributes!: pulumi.Output<outputs.paymentcryptography.KeyAttributes>;
    /**
     * The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.
     *
     * For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
     */
    public readonly keyCheckValueAlgorithm!: pulumi.Output<enums.paymentcryptography.KeyCheckValueAlgorithm | undefined>;
    public /*out*/ readonly keyIdentifier!: pulumi.Output<string>;
    /**
     * The source of the key material. For keys created within AWS Payment Cryptography, the value is `AWS_PAYMENT_CRYPTOGRAPHY` . For keys imported into AWS Payment Cryptography, the value is `EXTERNAL` .
     */
    public /*out*/ readonly keyOrigin!: pulumi.Output<enums.paymentcryptography.KeyOrigin>;
    /**
     * The state of key that is being created or deleted.
     */
    public /*out*/ readonly keyState!: pulumi.Output<enums.paymentcryptography.KeyState>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.exportable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exportable'");
            }
            if ((!args || args.keyAttributes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyAttributes'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["exportable"] = args ? args.exportable : undefined;
            resourceInputs["keyAttributes"] = args ? args.keyAttributes : undefined;
            resourceInputs["keyCheckValueAlgorithm"] = args ? args.keyCheckValueAlgorithm : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["keyIdentifier"] = undefined /*out*/;
            resourceInputs["keyOrigin"] = undefined /*out*/;
            resourceInputs["keyState"] = undefined /*out*/;
        } else {
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["exportable"] = undefined /*out*/;
            resourceInputs["keyAttributes"] = undefined /*out*/;
            resourceInputs["keyCheckValueAlgorithm"] = undefined /*out*/;
            resourceInputs["keyIdentifier"] = undefined /*out*/;
            resourceInputs["keyOrigin"] = undefined /*out*/;
            resourceInputs["keyState"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Key.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * Specifies whether the key is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether the key is exportable. This data is immutable after the key is created.
     */
    exportable: pulumi.Input<boolean>;
    /**
     * The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
     */
    keyAttributes: pulumi.Input<inputs.paymentcryptography.KeyAttributesArgs>;
    /**
     * The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.
     *
     * For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
     */
    keyCheckValueAlgorithm?: pulumi.Input<enums.paymentcryptography.KeyCheckValueAlgorithm>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}