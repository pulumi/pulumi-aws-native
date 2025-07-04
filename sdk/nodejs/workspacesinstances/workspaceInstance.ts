// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WorkspacesInstances::WorkspaceInstance
 */
export class WorkspaceInstance extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkspaceInstance {
        return new WorkspaceInstance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:workspacesinstances:WorkspaceInstance';

    /**
     * Returns true if the given object is an instance of WorkspaceInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceInstance.__pulumiType;
    }

    public /*out*/ readonly ec2ManagedInstance!: pulumi.Output<outputs.workspacesinstances.WorkspaceInstanceEc2ManagedInstance>;
    public readonly managedInstance!: pulumi.Output<outputs.workspacesinstances.ManagedInstanceProperties | undefined>;
    /**
     * The current state of the workspace instance
     */
    public /*out*/ readonly provisionState!: pulumi.Output<enums.workspacesinstances.WorkspaceInstanceProvisionState>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Unique identifier for the workspace instance
     */
    public /*out*/ readonly workspaceInstanceId!: pulumi.Output<string>;

    /**
     * Create a WorkspaceInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkspaceInstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["managedInstance"] = args ? args.managedInstance : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ec2ManagedInstance"] = undefined /*out*/;
            resourceInputs["provisionState"] = undefined /*out*/;
            resourceInputs["workspaceInstanceId"] = undefined /*out*/;
        } else {
            resourceInputs["ec2ManagedInstance"] = undefined /*out*/;
            resourceInputs["managedInstance"] = undefined /*out*/;
            resourceInputs["provisionState"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["workspaceInstanceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["managedInstance"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(WorkspaceInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkspaceInstance resource.
 */
export interface WorkspaceInstanceArgs {
    managedInstance?: pulumi.Input<inputs.workspacesinstances.ManagedInstancePropertiesArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
