// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::LaunchWizard::Deployment Resource Type
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:launchwizard:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * ARN of the LaunchWizard deployment
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Timestamp of LaunchWizard deployment creation
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Timestamp of LaunchWizard deployment deletion
     */
    public /*out*/ readonly deletedAt!: pulumi.Output<string>;
    /**
     * Deployment ID of the LaunchWizard deployment
     */
    public /*out*/ readonly deploymentId!: pulumi.Output<string>;
    /**
     * Workload deployment pattern name
     */
    public readonly deploymentPatternName!: pulumi.Output<string>;
    /**
     * Name of LaunchWizard deployment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource Group Name created for LaunchWizard deployment
     */
    public /*out*/ readonly resourceGroup!: pulumi.Output<string>;
    /**
     * LaunchWizard deployment specifications
     */
    public readonly specifications!: pulumi.Output<{[key: string]: string}>;
    /**
     * Status of LaunchWizard deployment
     */
    public /*out*/ readonly status!: pulumi.Output<enums.launchwizard.DeploymentStatus>;
    /**
     * Tags for LaunchWizard deployment
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Workload Name for LaunchWizard deployment
     */
    public readonly workloadName!: pulumi.Output<string>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deploymentPatternName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentPatternName'");
            }
            if ((!args || args.specifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specifications'");
            }
            if ((!args || args.workloadName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workloadName'");
            }
            resourceInputs["deploymentPatternName"] = args ? args.deploymentPatternName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["specifications"] = args ? args.specifications : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workloadName"] = args ? args.workloadName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deletedAt"] = undefined /*out*/;
            resourceInputs["deploymentId"] = undefined /*out*/;
            resourceInputs["resourceGroup"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deletedAt"] = undefined /*out*/;
            resourceInputs["deploymentId"] = undefined /*out*/;
            resourceInputs["deploymentPatternName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["resourceGroup"] = undefined /*out*/;
            resourceInputs["specifications"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["workloadName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["deploymentPatternName", "name", "workloadName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Deployment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * Workload deployment pattern name
     */
    deploymentPatternName: pulumi.Input<string>;
    /**
     * Name of LaunchWizard deployment
     */
    name?: pulumi.Input<string>;
    /**
     * LaunchWizard deployment specifications
     */
    specifications: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for LaunchWizard deployment
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Workload Name for LaunchWizard deployment
     */
    workloadName: pulumi.Input<string>;
}