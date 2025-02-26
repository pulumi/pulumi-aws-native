// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::UserHierarchyGroup
 */
export class UserHierarchyGroup extends pulumi.CustomResource {
    /**
     * Get an existing UserHierarchyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserHierarchyGroup {
        return new UserHierarchyGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:connect:UserHierarchyGroup';

    /**
     * Returns true if the given object is an instance of UserHierarchyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserHierarchyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserHierarchyGroup.__pulumiType;
    }

    /**
     * The identifier of the Amazon Connect instance.
     */
    public readonly instanceArn!: pulumi.Output<string>;
    /**
     * The name of the user hierarchy group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the parent user hierarchy group.
     */
    public readonly parentGroupArn!: pulumi.Output<string | undefined>;
    /**
     * One or more tags.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) for the user hierarchy group.
     */
    public /*out*/ readonly userHierarchyGroupArn!: pulumi.Output<string>;

    /**
     * Create a UserHierarchyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserHierarchyGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceArn'");
            }
            resourceInputs["instanceArn"] = args ? args.instanceArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentGroupArn"] = args ? args.parentGroupArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userHierarchyGroupArn"] = undefined /*out*/;
        } else {
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentGroupArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["userHierarchyGroupArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["parentGroupArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(UserHierarchyGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserHierarchyGroup resource.
 */
export interface UserHierarchyGroupArgs {
    /**
     * The identifier of the Amazon Connect instance.
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The name of the user hierarchy group.
     */
    name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the parent user hierarchy group.
     */
    parentGroupArn?: pulumi.Input<string>;
    /**
     * One or more tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
