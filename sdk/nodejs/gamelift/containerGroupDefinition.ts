// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::GameLift::ContainerGroupDefinition resource creates an Amazon GameLift container group definition.
 */
export class ContainerGroupDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ContainerGroupDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContainerGroupDefinition {
        return new ContainerGroupDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:gamelift:ContainerGroupDefinition';

    /**
     * Returns true if the given object is an instance of ContainerGroupDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerGroupDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerGroupDefinition.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions.
     */
    public /*out*/ readonly containerGroupDefinitionArn!: pulumi.Output<string>;
    /**
     * The scope of the container group
     */
    public readonly containerGroupType!: pulumi.Output<enums.gamelift.ContainerGroupDefinitionContainerGroupType | undefined>;
    /**
     * A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057").
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The definition for the game server container in this group. This property is used only when the container group type is `GAME_SERVER` . This container definition specifies a container image with the game server build.
     */
    public readonly gameServerContainerDefinition!: pulumi.Output<outputs.gamelift.ContainerGroupDefinitionGameServerContainerDefinition | undefined>;
    /**
     * A descriptive label for the container group definition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The operating system of the container group
     */
    public readonly operatingSystem!: pulumi.Output<enums.gamelift.ContainerGroupDefinitionOperatingSystem>;
    /**
     * A specific ContainerGroupDefinition version to be updated
     */
    public readonly sourceVersionNumber!: pulumi.Output<number | undefined>;
    /**
     * A string indicating ContainerGroupDefinition status.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.gamelift.ContainerGroupDefinitionStatus>;
    /**
     * A string indicating the reason for ContainerGroupDefinition status.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * A collection of support container definitions that define the containers in this group.
     */
    public readonly supportContainerDefinitions!: pulumi.Output<outputs.gamelift.ContainerGroupDefinitionSupportContainerDefinition[] | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The total memory limit of container groups following this definition in MiB
     */
    public readonly totalMemoryLimitMebibytes!: pulumi.Output<number>;
    /**
     * The total amount of virtual CPUs on the container group definition
     */
    public readonly totalVcpuLimit!: pulumi.Output<number>;
    /**
     * The description of this version
     */
    public readonly versionDescription!: pulumi.Output<string | undefined>;
    /**
     * The version of this ContainerGroupDefinition
     */
    public /*out*/ readonly versionNumber!: pulumi.Output<number>;

    /**
     * Create a ContainerGroupDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerGroupDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.operatingSystem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operatingSystem'");
            }
            if ((!args || args.totalMemoryLimitMebibytes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'totalMemoryLimitMebibytes'");
            }
            if ((!args || args.totalVcpuLimit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'totalVcpuLimit'");
            }
            resourceInputs["containerGroupType"] = args ? args.containerGroupType : undefined;
            resourceInputs["gameServerContainerDefinition"] = args ? args.gameServerContainerDefinition : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operatingSystem"] = args ? args.operatingSystem : undefined;
            resourceInputs["sourceVersionNumber"] = args ? args.sourceVersionNumber : undefined;
            resourceInputs["supportContainerDefinitions"] = args ? args.supportContainerDefinitions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["totalMemoryLimitMebibytes"] = args ? args.totalMemoryLimitMebibytes : undefined;
            resourceInputs["totalVcpuLimit"] = args ? args.totalVcpuLimit : undefined;
            resourceInputs["versionDescription"] = args ? args.versionDescription : undefined;
            resourceInputs["containerGroupDefinitionArn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        } else {
            resourceInputs["containerGroupDefinitionArn"] = undefined /*out*/;
            resourceInputs["containerGroupType"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["gameServerContainerDefinition"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operatingSystem"] = undefined /*out*/;
            resourceInputs["sourceVersionNumber"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
            resourceInputs["supportContainerDefinitions"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["totalMemoryLimitMebibytes"] = undefined /*out*/;
            resourceInputs["totalVcpuLimit"] = undefined /*out*/;
            resourceInputs["versionDescription"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["containerGroupType", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ContainerGroupDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContainerGroupDefinition resource.
 */
export interface ContainerGroupDefinitionArgs {
    /**
     * The scope of the container group
     */
    containerGroupType?: pulumi.Input<enums.gamelift.ContainerGroupDefinitionContainerGroupType>;
    /**
     * The definition for the game server container in this group. This property is used only when the container group type is `GAME_SERVER` . This container definition specifies a container image with the game server build.
     */
    gameServerContainerDefinition?: pulumi.Input<inputs.gamelift.ContainerGroupDefinitionGameServerContainerDefinitionArgs>;
    /**
     * A descriptive label for the container group definition.
     */
    name?: pulumi.Input<string>;
    /**
     * The operating system of the container group
     */
    operatingSystem: pulumi.Input<enums.gamelift.ContainerGroupDefinitionOperatingSystem>;
    /**
     * A specific ContainerGroupDefinition version to be updated
     */
    sourceVersionNumber?: pulumi.Input<number>;
    /**
     * A collection of support container definitions that define the containers in this group.
     */
    supportContainerDefinitions?: pulumi.Input<pulumi.Input<inputs.gamelift.ContainerGroupDefinitionSupportContainerDefinitionArgs>[]>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The total memory limit of container groups following this definition in MiB
     */
    totalMemoryLimitMebibytes: pulumi.Input<number>;
    /**
     * The total amount of virtual CPUs on the container group definition
     */
    totalVcpuLimit: pulumi.Input<number>;
    /**
     * The description of this version
     */
    versionDescription?: pulumi.Input<string>;
}
