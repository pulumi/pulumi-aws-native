// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::Flow Resource Type
 */
export class Flow extends pulumi.CustomResource {
    /**
     * Get an existing Flow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Flow {
        return new Flow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:bedrock:Flow';

    /**
     * Returns true if the given object is an instance of Flow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Flow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Flow.__pulumiType;
    }

    /**
     * Arn representation of the Flow
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Identifier for a Flow
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A KMS key ARN
     */
    public readonly customerEncryptionKeyArn!: pulumi.Output<string | undefined>;
    /**
     * The definition of the nodes and connections between the nodes in the flow.
     */
    public readonly definition!: pulumi.Output<outputs.bedrock.FlowDefinition | undefined>;
    /**
     * The Amazon S3 location of the flow definition.
     */
    public readonly definitionS3Location!: pulumi.Output<outputs.bedrock.FlowS3Location | undefined>;
    /**
     * A JSON string containing a Definition with the same schema as the Definition property of this resource
     */
    public readonly definitionString!: pulumi.Output<string | undefined>;
    /**
     * A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
     *
     * Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
     */
    public readonly definitionSubstitutions!: pulumi.Output<{[key: string]: string | number | boolean} | undefined>;
    /**
     * Description of the flow
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ARN of a IAM role
     */
    public readonly executionRoleArn!: pulumi.Output<string>;
    /**
     * Name for the flow
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The status of the flow. The following statuses are possible:
     *
     * - NotPrepared – The flow has been created or updated, but hasn't been prepared. If you just created the flow, you can't test it. If you updated the flow, the `DRAFT` version won't contain the latest changes for testing. Send a [PrepareFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PrepareFlow.html) request to package the latest changes into the `DRAFT` version.
     * - Preparing – The flow is being prepared so that the `DRAFT` version contains the latest changes for testing.
     * - Prepared – The flow is prepared and the `DRAFT` version contains the latest changes for testing.
     * - Failed – The last API operation that you invoked on the flow failed. Send a [GetFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlow.html) request and check the error message in the `validations` field.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.bedrock.FlowStatus>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly testAliasTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Draft Version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Flow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.executionRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            resourceInputs["customerEncryptionKeyArn"] = args ? args.customerEncryptionKeyArn : undefined;
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["definitionS3Location"] = args ? args.definitionS3Location : undefined;
            resourceInputs["definitionString"] = args ? args.definitionString : undefined;
            resourceInputs["definitionSubstitutions"] = args ? args.definitionSubstitutions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["testAliasTags"] = args ? args.testAliasTags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["customerEncryptionKeyArn"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["definitionS3Location"] = undefined /*out*/;
            resourceInputs["definitionString"] = undefined /*out*/;
            resourceInputs["definitionSubstitutions"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["executionRoleArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["testAliasTags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Flow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Flow resource.
 */
export interface FlowArgs {
    /**
     * A KMS key ARN
     */
    customerEncryptionKeyArn?: pulumi.Input<string>;
    /**
     * The definition of the nodes and connections between the nodes in the flow.
     */
    definition?: pulumi.Input<inputs.bedrock.FlowDefinitionArgs>;
    /**
     * The Amazon S3 location of the flow definition.
     */
    definitionS3Location?: pulumi.Input<inputs.bedrock.FlowS3LocationArgs>;
    /**
     * A JSON string containing a Definition with the same schema as the Definition property of this resource
     */
    definitionString?: pulumi.Input<string>;
    /**
     * A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
     *
     * Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
     */
    definitionSubstitutions?: pulumi.Input<{[key: string]: pulumi.Input<string | number | boolean>}>;
    /**
     * Description of the flow
     */
    description?: pulumi.Input<string>;
    /**
     * ARN of a IAM role
     */
    executionRoleArn: pulumi.Input<string>;
    /**
     * Name for the flow
     */
    name?: pulumi.Input<string>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    testAliasTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}