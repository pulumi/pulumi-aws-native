// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Associates a service with a service network.
 */
export class ServiceNetworkServiceAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ServiceNetworkServiceAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceNetworkServiceAssociation {
        return new ServiceNetworkServiceAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:vpclattice:ServiceNetworkServiceAssociation';

    /**
     * Returns true if the given object is an instance of ServiceNetworkServiceAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceNetworkServiceAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceNetworkServiceAssociation.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the association between the service network and the service.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the of the association between the service network and the service.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The date and time that the association was created, specified in ISO-8601 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The DNS information of the service.
     */
    public readonly dnsEntry!: pulumi.Output<outputs.vpclattice.ServiceNetworkServiceAssociationDnsEntry | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the service.
     */
    public /*out*/ readonly serviceArn!: pulumi.Output<string>;
    /**
     * The ID of the service.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    /**
     * The ID or ARN of the service.
     */
    public readonly serviceIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The name of the service.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the service network
     */
    public /*out*/ readonly serviceNetworkArn!: pulumi.Output<string>;
    /**
     * The ID of the service network.
     */
    public /*out*/ readonly serviceNetworkId!: pulumi.Output<string>;
    /**
     * The ID or ARN of the service network. You must use an ARN if the resources are in different accounts.
     */
    public readonly serviceNetworkIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The name of the service network.
     */
    public /*out*/ readonly serviceNetworkName!: pulumi.Output<string>;
    /**
     * The status of the association between the service network and the service.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.vpclattice.ServiceNetworkServiceAssociationStatus>;
    /**
     * The tags for the association.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a ServiceNetworkServiceAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceNetworkServiceAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dnsEntry"] = args ? args.dnsEntry : undefined;
            resourceInputs["serviceIdentifier"] = args ? args.serviceIdentifier : undefined;
            resourceInputs["serviceNetworkIdentifier"] = args ? args.serviceNetworkIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["serviceArn"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceNetworkArn"] = undefined /*out*/;
            resourceInputs["serviceNetworkId"] = undefined /*out*/;
            resourceInputs["serviceNetworkName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dnsEntry"] = undefined /*out*/;
            resourceInputs["serviceArn"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["serviceIdentifier"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceNetworkArn"] = undefined /*out*/;
            resourceInputs["serviceNetworkId"] = undefined /*out*/;
            resourceInputs["serviceNetworkIdentifier"] = undefined /*out*/;
            resourceInputs["serviceNetworkName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["serviceIdentifier", "serviceNetworkIdentifier"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ServiceNetworkServiceAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceNetworkServiceAssociation resource.
 */
export interface ServiceNetworkServiceAssociationArgs {
    /**
     * The DNS information of the service.
     */
    dnsEntry?: pulumi.Input<inputs.vpclattice.ServiceNetworkServiceAssociationDnsEntryArgs>;
    /**
     * The ID or ARN of the service.
     */
    serviceIdentifier?: pulumi.Input<string>;
    /**
     * The ID or ARN of the service network. You must use an ARN if the resources are in different accounts.
     */
    serviceNetworkIdentifier?: pulumi.Input<string>;
    /**
     * The tags for the association.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
