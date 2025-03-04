// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.
 */
export class VpcIngressConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpcIngressConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpcIngressConnection {
        return new VpcIngressConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apprunner:VpcIngressConnection';

    /**
     * Returns true if the given object is an instance of VpcIngressConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcIngressConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcIngressConnection.__pulumiType;
    }

    /**
     * The Domain name associated with the VPC Ingress Connection.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource.
     */
    public readonly ingressVpcConfiguration!: pulumi.Output<outputs.apprunner.VpcIngressConnectionIngressVpcConfiguration>;
    /**
     * The Amazon Resource Name (ARN) of the service.
     */
    public readonly serviceArn!: pulumi.Output<string>;
    /**
     * The current status of the VpcIngressConnection.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.apprunner.VpcIngressConnectionStatus>;
    /**
     * An optional list of metadata items that you can associate with the VPC Ingress Connection resource. A tag is a key-value pair.
     */
    public readonly tags!: pulumi.Output<outputs.CreateOnlyTag[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the VpcIngressConnection.
     */
    public /*out*/ readonly vpcIngressConnectionArn!: pulumi.Output<string>;
    /**
     * The customer-provided Vpc Ingress Connection name.
     */
    public readonly vpcIngressConnectionName!: pulumi.Output<string | undefined>;

    /**
     * Create a VpcIngressConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcIngressConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ingressVpcConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ingressVpcConfiguration'");
            }
            if ((!args || args.serviceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceArn'");
            }
            resourceInputs["ingressVpcConfiguration"] = args ? args.ingressVpcConfiguration : undefined;
            resourceInputs["serviceArn"] = args ? args.serviceArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcIngressConnectionName"] = args ? args.vpcIngressConnectionName : undefined;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vpcIngressConnectionArn"] = undefined /*out*/;
        } else {
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["ingressVpcConfiguration"] = undefined /*out*/;
            resourceInputs["serviceArn"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcIngressConnectionArn"] = undefined /*out*/;
            resourceInputs["vpcIngressConnectionName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["serviceArn", "tags[*]", "vpcIngressConnectionName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(VpcIngressConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpcIngressConnection resource.
 */
export interface VpcIngressConnectionArgs {
    /**
     * Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource.
     */
    ingressVpcConfiguration: pulumi.Input<inputs.apprunner.VpcIngressConnectionIngressVpcConfigurationArgs>;
    /**
     * The Amazon Resource Name (ARN) of the service.
     */
    serviceArn: pulumi.Input<string>;
    /**
     * An optional list of metadata items that you can associate with the VPC Ingress Connection resource. A tag is a key-value pair.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.CreateOnlyTagArgs>[]>;
    /**
     * The customer-provided Vpc Ingress Connection name.
     */
    vpcIngressConnectionName?: pulumi.Input<string>;
}
