// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolGroup
 *
 * @deprecated UserPoolGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class UserPoolGroup extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPoolGroup {
        pulumi.log.warn("UserPoolGroup is deprecated: UserPoolGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new UserPoolGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPoolGroup';

    /**
     * Returns true if the given object is an instance of UserPoolGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolGroup.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly groupName!: pulumi.Output<string | undefined>;
    public readonly precedence!: pulumi.Output<number | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a UserPoolGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated UserPoolGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: UserPoolGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("UserPoolGroup is deprecated: UserPoolGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["precedence"] = args ? args.precedence : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["groupName"] = undefined /*out*/;
            resourceInputs["precedence"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPoolGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPoolGroup resource.
 */
export interface UserPoolGroupArgs {
    description?: pulumi.Input<string>;
    groupName?: pulumi.Input<string>;
    precedence?: pulumi.Input<number>;
    roleArn?: pulumi.Input<string>;
    userPoolId: pulumi.Input<string>;
}