// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Url
 */
export class Url extends pulumi.CustomResource {
    /**
     * Get an existing Url resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Url {
        return new Url(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:Url';

    /**
     * Returns true if the given object is an instance of Url.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Url {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Url.__pulumiType;
    }

    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    public readonly authType!: pulumi.Output<enums.lambda.UrlAuthType>;
    public readonly cors!: pulumi.Output<outputs.lambda.UrlCors | undefined>;
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    public /*out*/ readonly functionArn!: pulumi.Output<string>;
    /**
     * The generated url for this resource.
     */
    public /*out*/ readonly functionUrl!: pulumi.Output<string>;
    /**
     * The invocation mode for the function’s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.
     */
    public readonly invokeMode!: pulumi.Output<enums.lambda.UrlInvokeMode | undefined>;
    /**
     * The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    public readonly targetFunctionArn!: pulumi.Output<string>;

    /**
     * Create a Url resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UrlArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authType'");
            }
            if ((!args || args.targetFunctionArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetFunctionArn'");
            }
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["invokeMode"] = args ? args.invokeMode : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
            resourceInputs["targetFunctionArn"] = args ? args.targetFunctionArn : undefined;
            resourceInputs["functionArn"] = undefined /*out*/;
            resourceInputs["functionUrl"] = undefined /*out*/;
        } else {
            resourceInputs["authType"] = undefined /*out*/;
            resourceInputs["cors"] = undefined /*out*/;
            resourceInputs["functionArn"] = undefined /*out*/;
            resourceInputs["functionUrl"] = undefined /*out*/;
            resourceInputs["invokeMode"] = undefined /*out*/;
            resourceInputs["qualifier"] = undefined /*out*/;
            resourceInputs["targetFunctionArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Url.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Url resource.
 */
export interface UrlArgs {
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    authType: pulumi.Input<enums.lambda.UrlAuthType>;
    cors?: pulumi.Input<inputs.lambda.UrlCorsArgs>;
    /**
     * The invocation mode for the function’s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.
     */
    invokeMode?: pulumi.Input<enums.lambda.UrlInvokeMode>;
    /**
     * The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    targetFunctionArn: pulumi.Input<string>;
}