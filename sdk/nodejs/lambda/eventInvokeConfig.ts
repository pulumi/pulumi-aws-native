// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.
 */
export class EventInvokeConfig extends pulumi.CustomResource {
    /**
     * Get an existing EventInvokeConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventInvokeConfig {
        return new EventInvokeConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:EventInvokeConfig';

    /**
     * Returns true if the given object is an instance of EventInvokeConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventInvokeConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventInvokeConfig.__pulumiType;
    }

    /**
     * A destination for events after they have been sent to a function for processing.
     *
     * **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
     * - *Queue* - The ARN of a standard SQS queue.
     * - *Bucket* - The ARN of an Amazon S3 bucket.
     * - *Topic* - The ARN of a standard SNS topic.
     * - *Event Bus* - The ARN of an Amazon EventBridge event bus.
     *
     * > S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.
     */
    public readonly destinationConfig!: pulumi.Output<outputs.lambda.EventInvokeConfigDestinationConfig | undefined>;
    /**
     * The name of the Lambda function.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * The maximum age of a request that Lambda sends to a function for processing.
     */
    public readonly maximumEventAgeInSeconds!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of times to retry when the function returns an error.
     */
    public readonly maximumRetryAttempts!: pulumi.Output<number | undefined>;
    /**
     * The identifier of a version or alias.
     */
    public readonly qualifier!: pulumi.Output<string>;

    /**
     * Create a EventInvokeConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventInvokeConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            if ((!args || args.qualifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'qualifier'");
            }
            resourceInputs["destinationConfig"] = args ? args.destinationConfig : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["maximumEventAgeInSeconds"] = args ? args.maximumEventAgeInSeconds : undefined;
            resourceInputs["maximumRetryAttempts"] = args ? args.maximumRetryAttempts : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
        } else {
            resourceInputs["destinationConfig"] = undefined /*out*/;
            resourceInputs["functionName"] = undefined /*out*/;
            resourceInputs["maximumEventAgeInSeconds"] = undefined /*out*/;
            resourceInputs["maximumRetryAttempts"] = undefined /*out*/;
            resourceInputs["qualifier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["functionName", "qualifier"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EventInvokeConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventInvokeConfig resource.
 */
export interface EventInvokeConfigArgs {
    /**
     * A destination for events after they have been sent to a function for processing.
     *
     * **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
     * - *Queue* - The ARN of a standard SQS queue.
     * - *Bucket* - The ARN of an Amazon S3 bucket.
     * - *Topic* - The ARN of a standard SNS topic.
     * - *Event Bus* - The ARN of an Amazon EventBridge event bus.
     *
     * > S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.
     */
    destinationConfig?: pulumi.Input<inputs.lambda.EventInvokeConfigDestinationConfigArgs>;
    /**
     * The name of the Lambda function.
     */
    functionName: pulumi.Input<string>;
    /**
     * The maximum age of a request that Lambda sends to a function for processing.
     */
    maximumEventAgeInSeconds?: pulumi.Input<number>;
    /**
     * The maximum number of times to retry when the function returns an error.
     */
    maximumRetryAttempts?: pulumi.Input<number>;
    /**
     * The identifier of a version or alias.
     */
    qualifier: pulumi.Input<string>;
}
