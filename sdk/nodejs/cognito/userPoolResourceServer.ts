// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolResourceServer
 */
export class UserPoolResourceServer extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolResourceServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPoolResourceServer {
        return new UserPoolResourceServer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPoolResourceServer';

    /**
     * Returns true if the given object is an instance of UserPoolResourceServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolResourceServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolResourceServer.__pulumiType;
    }

    /**
     * A unique resource server identifier for the resource server. The identifier can be an API friendly name like `solar-system-data` . You can also set an API URL like `https://solar-system-data-api.example.com` as your identifier.
     *
     * Amazon Cognito represents scopes in the access token in the format `$resource-server-identifier/$scope` . Longer scope-identifier strings increase the size of your access tokens.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * A friendly name for the resource server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of scopes. Each scope is a map with keys `ScopeName` and `ScopeDescription` .
     */
    public readonly scopes!: pulumi.Output<outputs.cognito.UserPoolResourceServerResourceServerScopeType[] | undefined>;
    /**
     * The ID of the user pool where you want to create a resource server.
     */
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a UserPoolResourceServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPoolResourceServerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
        } else {
            resourceInputs["identifier"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["scopes"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["identifier", "userPoolId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(UserPoolResourceServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPoolResourceServer resource.
 */
export interface UserPoolResourceServerArgs {
    /**
     * A unique resource server identifier for the resource server. The identifier can be an API friendly name like `solar-system-data` . You can also set an API URL like `https://solar-system-data-api.example.com` as your identifier.
     *
     * Amazon Cognito represents scopes in the access token in the format `$resource-server-identifier/$scope` . Longer scope-identifier strings increase the size of your access tokens.
     */
    identifier: pulumi.Input<string>;
    /**
     * A friendly name for the resource server.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of scopes. Each scope is a map with keys `ScopeName` and `ScopeDescription` .
     */
    scopes?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolResourceServerResourceServerScopeTypeArgs>[]>;
    /**
     * The ID of the user pool where you want to create a resource server.
     */
    userPoolId: pulumi.Input<string>;
}
