// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a runtime environment that can run migrated mainframe applications.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:m2:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * The description of the environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The target platform for the runtime environment.
     */
    public readonly engineType!: pulumi.Output<enums.m2.EnvironmentEngineType>;
    /**
     * The version of the runtime engine for the environment.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the runtime environment.
     */
    public /*out*/ readonly environmentArn!: pulumi.Output<string>;
    /**
     * The unique identifier of the environment.
     */
    public /*out*/ readonly environmentId!: pulumi.Output<string>;
    /**
     * Defines the details of a high availability configuration.
     */
    public readonly highAvailabilityConfig!: pulumi.Output<outputs.m2.EnvironmentHighAvailabilityConfig | undefined>;
    /**
     * The type of instance underlying the environment.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The name of the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network type supported by the runtime environment.
     */
    public readonly networkType!: pulumi.Output<enums.m2.EnvironmentNetworkType | undefined>;
    /**
     * Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the environment is publicly accessible.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * The list of security groups for the VPC associated with this environment.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The storage configurations defined for the runtime environment.
     */
    public readonly storageConfigurations!: pulumi.Output<outputs.m2.EnvironmentStorageConfiguration[] | undefined>;
    /**
     * The unique identifiers of the subnets assigned to this runtime environment.
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * Tags associated to this environment.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineType'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineType"] = args ? args.engineType : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["highAvailabilityConfig"] = args ? args.highAvailabilityConfig : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["storageConfigurations"] = args ? args.storageConfigurations : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["environmentArn"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["engineType"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["environmentArn"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["highAvailabilityConfig"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkType"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["publiclyAccessible"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["storageConfigurations"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "engineType", "kmsKeyId", "name", "networkType", "publiclyAccessible", "securityGroupIds[*]", "storageConfigurations[*]", "subnetIds[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * The description of the environment.
     */
    description?: pulumi.Input<string>;
    /**
     * The target platform for the runtime environment.
     */
    engineType: pulumi.Input<enums.m2.EnvironmentEngineType>;
    /**
     * The version of the runtime engine for the environment.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Defines the details of a high availability configuration.
     */
    highAvailabilityConfig?: pulumi.Input<inputs.m2.EnvironmentHighAvailabilityConfigArgs>;
    /**
     * The type of instance underlying the environment.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The network type supported by the runtime environment.
     */
    networkType?: pulumi.Input<enums.m2.EnvironmentNetworkType>;
    /**
     * Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Specifies whether the environment is publicly accessible.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * The list of security groups for the VPC associated with this environment.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage configurations defined for the runtime environment.
     */
    storageConfigurations?: pulumi.Input<pulumi.Input<inputs.m2.EnvironmentStorageConfigurationArgs>[]>;
    /**
     * The unique identifiers of the subnets assigned to this runtime environment.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags associated to this environment.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
