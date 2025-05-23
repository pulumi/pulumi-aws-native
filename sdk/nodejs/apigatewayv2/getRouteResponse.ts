// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGatewayV2::RouteResponse`` resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.
 */
export function getRouteResponse(args: GetRouteResponseArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteResponseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigatewayv2:getRouteResponse", {
        "apiId": args.apiId,
        "routeId": args.routeId,
        "routeResponseId": args.routeResponseId,
    }, opts);
}

export interface GetRouteResponseArgs {
    /**
     * The API identifier.
     */
    apiId: string;
    /**
     * The route ID.
     */
    routeId: string;
    /**
     * The route response ID.
     */
    routeResponseId: string;
}

export interface GetRouteResponseResult {
    /**
     * The model selection expression for the route response. Supported only for WebSocket APIs.
     */
    readonly modelSelectionExpression?: string;
    /**
     * The response models for the route response.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
     */
    readonly responseModels?: any;
    /**
     * The route response parameters.
     */
    readonly responseParameters?: {[key: string]: outputs.apigatewayv2.RouteResponseParameterConstraints};
    /**
     * The route response ID.
     */
    readonly routeResponseId?: string;
    /**
     * The route response key.
     */
    readonly routeResponseKey?: string;
}
/**
 * The ``AWS::ApiGatewayV2::RouteResponse`` resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.
 */
export function getRouteResponseOutput(args: GetRouteResponseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRouteResponseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:apigatewayv2:getRouteResponse", {
        "apiId": args.apiId,
        "routeId": args.routeId,
        "routeResponseId": args.routeResponseId,
    }, opts);
}

export interface GetRouteResponseOutputArgs {
    /**
     * The API identifier.
     */
    apiId: pulumi.Input<string>;
    /**
     * The route ID.
     */
    routeId: pulumi.Input<string>;
    /**
     * The route response ID.
     */
    routeResponseId: pulumi.Input<string>;
}
