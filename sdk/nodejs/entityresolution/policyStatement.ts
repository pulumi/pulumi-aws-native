// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Policy Statement defined in AWS Entity Resolution Service
 */
export class PolicyStatement extends pulumi.CustomResource {
    /**
     * Get an existing PolicyStatement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyStatement {
        return new PolicyStatement(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:entityresolution:PolicyStatement';

    /**
     * Returns true if the given object is an instance of PolicyStatement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyStatement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyStatement.__pulumiType;
    }

    public readonly action!: pulumi.Output<string[] | undefined>;
    public readonly arn!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<string | undefined>;
    public readonly effect!: pulumi.Output<enums.entityresolution.PolicyStatementStatementEffect | undefined>;
    public readonly principal!: pulumi.Output<string[] | undefined>;
    public readonly statementId!: pulumi.Output<string>;

    /**
     * Create a PolicyStatement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyStatementArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.arn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'arn'");
            }
            if ((!args || args.statementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statementId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["effect"] = args ? args.effect : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["statementId"] = args ? args.statementId : undefined;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["condition"] = undefined /*out*/;
            resourceInputs["effect"] = undefined /*out*/;
            resourceInputs["principal"] = undefined /*out*/;
            resourceInputs["statementId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["arn", "statementId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PolicyStatement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyStatement resource.
 */
export interface PolicyStatementArgs {
    action?: pulumi.Input<pulumi.Input<string>[]>;
    arn: pulumi.Input<string>;
    condition?: pulumi.Input<string>;
    effect?: pulumi.Input<enums.entityresolution.PolicyStatementStatementEffect>;
    principal?: pulumi.Input<pulumi.Input<string>[]>;
    statementId: pulumi.Input<string>;
}