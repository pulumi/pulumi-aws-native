// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
 */
export class EnvironmentBlueprintConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentBlueprintConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnvironmentBlueprintConfiguration {
        return new EnvironmentBlueprintConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datazone:EnvironmentBlueprintConfiguration';

    /**
     * Returns true if the given object is an instance of EnvironmentBlueprintConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentBlueprintConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentBlueprintConfiguration.__pulumiType;
    }

    /**
     * The timestamp of when an environment blueprint was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The identifier of the Amazon DataZone domain in which an environment blueprint exists.
     */
    public /*out*/ readonly domainId!: pulumi.Output<string>;
    /**
     * The identifier of the Amazon DataZone domain in which an environment blueprint exists.
     */
    public readonly domainIdentifier!: pulumi.Output<string>;
    /**
     * The enabled AWS Regions specified in a blueprint configuration.
     */
    public readonly enabledRegions!: pulumi.Output<string[]>;
    /**
     * The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
     */
    public /*out*/ readonly environmentBlueprintId!: pulumi.Output<string>;
    /**
     * The identifier of the environment blueprint.
     *
     * In the current release, only the following values are supported: `DefaultDataLake` and `DefaultDataWarehouse` .
     */
    public readonly environmentBlueprintIdentifier!: pulumi.Output<string>;
    /**
     * The environment role permission boundary.
     */
    public readonly environmentRolePermissionBoundary!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the manage access role.
     */
    public readonly manageAccessRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The provisioning configuration of a blueprint.
     */
    public readonly provisioningConfigurations!: pulumi.Output<outputs.datazone.EnvironmentBlueprintConfigurationProvisioningConfigurationProperties[] | undefined>;
    /**
     * The ARN of the provisioning role.
     */
    public readonly provisioningRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The regional parameters of the environment blueprint.
     */
    public readonly regionalParameters!: pulumi.Output<outputs.datazone.EnvironmentBlueprintConfigurationRegionalParameter[] | undefined>;
    /**
     * The timestamp of when the environment blueprint was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a EnvironmentBlueprintConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentBlueprintConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domainIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainIdentifier'");
            }
            if ((!args || args.enabledRegions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabledRegions'");
            }
            if ((!args || args.environmentBlueprintIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentBlueprintIdentifier'");
            }
            resourceInputs["domainIdentifier"] = args ? args.domainIdentifier : undefined;
            resourceInputs["enabledRegions"] = args ? args.enabledRegions : undefined;
            resourceInputs["environmentBlueprintIdentifier"] = args ? args.environmentBlueprintIdentifier : undefined;
            resourceInputs["environmentRolePermissionBoundary"] = args ? args.environmentRolePermissionBoundary : undefined;
            resourceInputs["manageAccessRoleArn"] = args ? args.manageAccessRoleArn : undefined;
            resourceInputs["provisioningConfigurations"] = args ? args.provisioningConfigurations : undefined;
            resourceInputs["provisioningRoleArn"] = args ? args.provisioningRoleArn : undefined;
            resourceInputs["regionalParameters"] = args ? args.regionalParameters : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domainId"] = undefined /*out*/;
            resourceInputs["environmentBlueprintId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domainId"] = undefined /*out*/;
            resourceInputs["domainIdentifier"] = undefined /*out*/;
            resourceInputs["enabledRegions"] = undefined /*out*/;
            resourceInputs["environmentBlueprintId"] = undefined /*out*/;
            resourceInputs["environmentBlueprintIdentifier"] = undefined /*out*/;
            resourceInputs["environmentRolePermissionBoundary"] = undefined /*out*/;
            resourceInputs["manageAccessRoleArn"] = undefined /*out*/;
            resourceInputs["provisioningConfigurations"] = undefined /*out*/;
            resourceInputs["provisioningRoleArn"] = undefined /*out*/;
            resourceInputs["regionalParameters"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["domainIdentifier", "environmentBlueprintIdentifier"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EnvironmentBlueprintConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnvironmentBlueprintConfiguration resource.
 */
export interface EnvironmentBlueprintConfigurationArgs {
    /**
     * The identifier of the Amazon DataZone domain in which an environment blueprint exists.
     */
    domainIdentifier: pulumi.Input<string>;
    /**
     * The enabled AWS Regions specified in a blueprint configuration.
     */
    enabledRegions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The identifier of the environment blueprint.
     *
     * In the current release, only the following values are supported: `DefaultDataLake` and `DefaultDataWarehouse` .
     */
    environmentBlueprintIdentifier: pulumi.Input<string>;
    /**
     * The environment role permission boundary.
     */
    environmentRolePermissionBoundary?: pulumi.Input<string>;
    /**
     * The ARN of the manage access role.
     */
    manageAccessRoleArn?: pulumi.Input<string>;
    /**
     * The provisioning configuration of a blueprint.
     */
    provisioningConfigurations?: pulumi.Input<pulumi.Input<inputs.datazone.EnvironmentBlueprintConfigurationProvisioningConfigurationPropertiesArgs>[]>;
    /**
     * The ARN of the provisioning role.
     */
    provisioningRoleArn?: pulumi.Input<string>;
    /**
     * The regional parameters of the environment blueprint.
     */
    regionalParameters?: pulumi.Input<pulumi.Input<inputs.datazone.EnvironmentBlueprintConfigurationRegionalParameterArgs>[]>;
}
