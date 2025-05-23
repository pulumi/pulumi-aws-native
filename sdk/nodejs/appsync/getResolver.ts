// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::AppSync::Resolver`` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).
 *   When you submit an update, CFNLong updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CFNshort template. Changing the S3 file content without changing a property value will not result in an update operation.
 *  See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *User Guide*.
 */
export function getResolver(args: GetResolverArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appsync:getResolver", {
        "resolverArn": args.resolverArn,
    }, opts);
}

export interface GetResolverArgs {
    /**
     * ARN of the resolver, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername` .
     */
    resolverArn: string;
}

export interface GetResolverResult {
    /**
     * The caching configuration for the resolver.
     */
    readonly cachingConfig?: outputs.appsync.ResolverCachingConfig;
    /**
     * The ``resolver`` code that contains the request and response functions. When code is used, the ``runtime`` is required. The runtime value must be ``APPSYNC_JS``.
     */
    readonly code?: string;
    /**
     * The resolver data source name.
     */
    readonly dataSourceName?: string;
    /**
     * The resolver type.
     *   +  *UNIT*: A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
     *   +  *PIPELINE*: A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of ``Function`` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
     */
    readonly kind?: string;
    /**
     * The maximum number of resolver request inputs that will be sent to a single LAMlong function in a ``BatchInvoke`` operation.
     */
    readonly maxBatchSize?: number;
    /**
     * Enables or disables enhanced resolver metrics for specified resolvers. Note that ``MetricsConfig`` won't be used unless the ``resolverLevelMetricsBehavior`` value is set to ``PER_RESOLVER_METRICS``. If the ``resolverLevelMetricsBehavior`` is set to ``FULL_REQUEST_RESOLVER_METRICS`` instead, ``MetricsConfig`` will be ignored. However, you can still set its value.
     */
    readonly metricsConfig?: enums.appsync.ResolverMetricsConfig;
    /**
     * Functions linked with the pipeline resolver.
     */
    readonly pipelineConfig?: outputs.appsync.ResolverPipelineConfig;
    /**
     * The request mapping template.
     *  Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
     */
    readonly requestMappingTemplate?: string;
    /**
     * ARN of the resolver, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername` .
     */
    readonly resolverArn?: string;
    /**
     * The response mapping template.
     */
    readonly responseMappingTemplate?: string;
    /**
     * Describes a runtime used by an APSYlong resolver or APSYlong function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
     */
    readonly runtime?: outputs.appsync.ResolverAppSyncRuntime;
    /**
     * The ``SyncConfig`` for a resolver attached to a versioned data source.
     */
    readonly syncConfig?: outputs.appsync.ResolverSyncConfig;
}
/**
 * The ``AWS::AppSync::Resolver`` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).
 *   When you submit an update, CFNLong updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CFNshort template. Changing the S3 file content without changing a property value will not result in an update operation.
 *  See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *User Guide*.
 */
export function getResolverOutput(args: GetResolverOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResolverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:appsync:getResolver", {
        "resolverArn": args.resolverArn,
    }, opts);
}

export interface GetResolverOutputArgs {
    /**
     * ARN of the resolver, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername` .
     */
    resolverArn: pulumi.Input<string>;
}
