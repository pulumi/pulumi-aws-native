// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApplicationApiGatewayEndpointType = {
    Regional: "REGIONAL",
    Private: "PRIVATE",
} as const;

export type ApplicationApiGatewayEndpointType = (typeof ApplicationApiGatewayEndpointType)[keyof typeof ApplicationApiGatewayEndpointType];

export const ApplicationProxyType = {
    ApiGateway: "API_GATEWAY",
} as const;

export type ApplicationProxyType = (typeof ApplicationProxyType)[keyof typeof ApplicationProxyType];

export const EnvironmentNetworkFabricType = {
    TransitGateway: "TRANSIT_GATEWAY",
    None: "NONE",
} as const;

export type EnvironmentNetworkFabricType = (typeof EnvironmentNetworkFabricType)[keyof typeof EnvironmentNetworkFabricType];

export const RouteActivationState = {
    Inactive: "INACTIVE",
    Active: "ACTIVE",
} as const;

export type RouteActivationState = (typeof RouteActivationState)[keyof typeof RouteActivationState];

export const RouteMethod = {
    Delete: "DELETE",
    Get: "GET",
    Head: "HEAD",
    Options: "OPTIONS",
    Patch: "PATCH",
    Post: "POST",
    Put: "PUT",
} as const;

export type RouteMethod = (typeof RouteMethod)[keyof typeof RouteMethod];

export const RouteType = {
    Default: "DEFAULT",
    UriPath: "URI_PATH",
} as const;

export type RouteType = (typeof RouteType)[keyof typeof RouteType];

export const ServiceEndpointType = {
    Lambda: "LAMBDA",
    Url: "URL",
} as const;

export type ServiceEndpointType = (typeof ServiceEndpointType)[keyof typeof ServiceEndpointType];
