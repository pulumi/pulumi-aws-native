// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::OpsWorks::Stack
 *
 * @deprecated Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Stack {
        pulumi.log.warn("Stack is deprecated: Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Stack(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:opsworks:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    public readonly agentVersion!: pulumi.Output<string | undefined>;
    public readonly attributes!: pulumi.Output<any | undefined>;
    public readonly chefConfiguration!: pulumi.Output<outputs.opsworks.StackChefConfiguration | undefined>;
    public readonly cloneAppIds!: pulumi.Output<string[] | undefined>;
    public readonly clonePermissions!: pulumi.Output<boolean | undefined>;
    public readonly configurationManager!: pulumi.Output<outputs.opsworks.StackConfigurationManager | undefined>;
    public readonly customCookbooksSource!: pulumi.Output<outputs.opsworks.StackSource | undefined>;
    public readonly customJson!: pulumi.Output<any | undefined>;
    public readonly defaultAvailabilityZone!: pulumi.Output<string | undefined>;
    public readonly defaultInstanceProfileArn!: pulumi.Output<string>;
    public readonly defaultOs!: pulumi.Output<string | undefined>;
    public readonly defaultRootDeviceType!: pulumi.Output<string | undefined>;
    public readonly defaultSshKeyName!: pulumi.Output<string | undefined>;
    public readonly defaultSubnetId!: pulumi.Output<string | undefined>;
    public readonly ecsClusterArn!: pulumi.Output<string | undefined>;
    public readonly elasticIps!: pulumi.Output<outputs.opsworks.StackElasticIp[] | undefined>;
    public readonly hostnameTheme!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly rdsDbInstances!: pulumi.Output<outputs.opsworks.StackRdsDbInstance[] | undefined>;
    public readonly serviceRoleArn!: pulumi.Output<string>;
    public readonly sourceStackId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.opsworks.StackTag[] | undefined>;
    public readonly useCustomCookbooks!: pulumi.Output<boolean | undefined>;
    public readonly useOpsworksSecurityGroups!: pulumi.Output<boolean | undefined>;
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: StackArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Stack is deprecated: Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.defaultInstanceProfileArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultInstanceProfileArn'");
            }
            if ((!args || args.serviceRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            resourceInputs["agentVersion"] = args ? args.agentVersion : undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["chefConfiguration"] = args ? args.chefConfiguration : undefined;
            resourceInputs["cloneAppIds"] = args ? args.cloneAppIds : undefined;
            resourceInputs["clonePermissions"] = args ? args.clonePermissions : undefined;
            resourceInputs["configurationManager"] = args ? args.configurationManager : undefined;
            resourceInputs["customCookbooksSource"] = args ? args.customCookbooksSource : undefined;
            resourceInputs["customJson"] = args ? args.customJson : undefined;
            resourceInputs["defaultAvailabilityZone"] = args ? args.defaultAvailabilityZone : undefined;
            resourceInputs["defaultInstanceProfileArn"] = args ? args.defaultInstanceProfileArn : undefined;
            resourceInputs["defaultOs"] = args ? args.defaultOs : undefined;
            resourceInputs["defaultRootDeviceType"] = args ? args.defaultRootDeviceType : undefined;
            resourceInputs["defaultSshKeyName"] = args ? args.defaultSshKeyName : undefined;
            resourceInputs["defaultSubnetId"] = args ? args.defaultSubnetId : undefined;
            resourceInputs["ecsClusterArn"] = args ? args.ecsClusterArn : undefined;
            resourceInputs["elasticIps"] = args ? args.elasticIps : undefined;
            resourceInputs["hostnameTheme"] = args ? args.hostnameTheme : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rdsDbInstances"] = args ? args.rdsDbInstances : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["sourceStackId"] = args ? args.sourceStackId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useCustomCookbooks"] = args ? args.useCustomCookbooks : undefined;
            resourceInputs["useOpsworksSecurityGroups"] = args ? args.useOpsworksSecurityGroups : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        } else {
            resourceInputs["agentVersion"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["chefConfiguration"] = undefined /*out*/;
            resourceInputs["cloneAppIds"] = undefined /*out*/;
            resourceInputs["clonePermissions"] = undefined /*out*/;
            resourceInputs["configurationManager"] = undefined /*out*/;
            resourceInputs["customCookbooksSource"] = undefined /*out*/;
            resourceInputs["customJson"] = undefined /*out*/;
            resourceInputs["defaultAvailabilityZone"] = undefined /*out*/;
            resourceInputs["defaultInstanceProfileArn"] = undefined /*out*/;
            resourceInputs["defaultOs"] = undefined /*out*/;
            resourceInputs["defaultRootDeviceType"] = undefined /*out*/;
            resourceInputs["defaultSshKeyName"] = undefined /*out*/;
            resourceInputs["defaultSubnetId"] = undefined /*out*/;
            resourceInputs["ecsClusterArn"] = undefined /*out*/;
            resourceInputs["elasticIps"] = undefined /*out*/;
            resourceInputs["hostnameTheme"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["rdsDbInstances"] = undefined /*out*/;
            resourceInputs["serviceRoleArn"] = undefined /*out*/;
            resourceInputs["sourceStackId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["useCustomCookbooks"] = undefined /*out*/;
            resourceInputs["useOpsworksSecurityGroups"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    agentVersion?: pulumi.Input<string>;
    attributes?: any;
    chefConfiguration?: pulumi.Input<inputs.opsworks.StackChefConfigurationArgs>;
    cloneAppIds?: pulumi.Input<pulumi.Input<string>[]>;
    clonePermissions?: pulumi.Input<boolean>;
    configurationManager?: pulumi.Input<inputs.opsworks.StackConfigurationManagerArgs>;
    customCookbooksSource?: pulumi.Input<inputs.opsworks.StackSourceArgs>;
    customJson?: any;
    defaultAvailabilityZone?: pulumi.Input<string>;
    defaultInstanceProfileArn: pulumi.Input<string>;
    defaultOs?: pulumi.Input<string>;
    defaultRootDeviceType?: pulumi.Input<string>;
    defaultSshKeyName?: pulumi.Input<string>;
    defaultSubnetId?: pulumi.Input<string>;
    ecsClusterArn?: pulumi.Input<string>;
    elasticIps?: pulumi.Input<pulumi.Input<inputs.opsworks.StackElasticIpArgs>[]>;
    hostnameTheme?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rdsDbInstances?: pulumi.Input<pulumi.Input<inputs.opsworks.StackRdsDbInstanceArgs>[]>;
    serviceRoleArn: pulumi.Input<string>;
    sourceStackId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.opsworks.StackTagArgs>[]>;
    useCustomCookbooks?: pulumi.Input<boolean>;
    useOpsworksSecurityGroups?: pulumi.Input<boolean>;
    vpcId?: pulumi.Input<string>;
}