// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::OIDCProvider
 */
export class OidcProvider extends pulumi.CustomResource {
    /**
     * Get an existing OidcProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OidcProvider {
        return new OidcProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:OidcProvider';

    /**
     * Returns true if the given object is an instance of OidcProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcProvider.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the OIDC provider
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly clientIdList!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.iam.OidcProviderTag[] | undefined>;
    public readonly thumbprintList!: pulumi.Output<string[]>;
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a OidcProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OidcProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.thumbprintList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'thumbprintList'");
            }
            resourceInputs["clientIdList"] = args ? args.clientIdList : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thumbprintList"] = args ? args.thumbprintList : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["clientIdList"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["thumbprintList"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["url"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(OidcProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OidcProvider resource.
 */
export interface OidcProviderArgs {
    clientIdList?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.iam.OidcProviderTagArgs>[]>;
    thumbprintList: pulumi.Input<pulumi.Input<string>[]>;
    url?: pulumi.Input<string>;
}