// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClusterArgs } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;

export { ClusterParameterGroupArgs } from "./clusterParameterGroup";
export type ClusterParameterGroup = import("./clusterParameterGroup").ClusterParameterGroup;
export const ClusterParameterGroup: typeof import("./clusterParameterGroup").ClusterParameterGroup = null as any;

export { ClusterSecurityGroupArgs } from "./clusterSecurityGroup";
export type ClusterSecurityGroup = import("./clusterSecurityGroup").ClusterSecurityGroup;
export const ClusterSecurityGroup: typeof import("./clusterSecurityGroup").ClusterSecurityGroup = null as any;

export { ClusterSecurityGroupIngressArgs } from "./clusterSecurityGroupIngress";
export type ClusterSecurityGroupIngress = import("./clusterSecurityGroupIngress").ClusterSecurityGroupIngress;
export const ClusterSecurityGroupIngress: typeof import("./clusterSecurityGroupIngress").ClusterSecurityGroupIngress = null as any;

export { ClusterSubnetGroupArgs } from "./clusterSubnetGroup";
export type ClusterSubnetGroup = import("./clusterSubnetGroup").ClusterSubnetGroup;
export const ClusterSubnetGroup: typeof import("./clusterSubnetGroup").ClusterSubnetGroup = null as any;

export { EndpointAccessArgs } from "./endpointAccess";
export type EndpointAccess = import("./endpointAccess").EndpointAccess;
export const EndpointAccess: typeof import("./endpointAccess").EndpointAccess = null as any;

export { EndpointAuthorizationArgs } from "./endpointAuthorization";
export type EndpointAuthorization = import("./endpointAuthorization").EndpointAuthorization;
export const EndpointAuthorization: typeof import("./endpointAuthorization").EndpointAuthorization = null as any;

export { EventSubscriptionArgs } from "./eventSubscription";
export type EventSubscription = import("./eventSubscription").EventSubscription;
export const EventSubscription: typeof import("./eventSubscription").EventSubscription = null as any;

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;

export { GetClusterParameterGroupArgs, GetClusterParameterGroupResult, GetClusterParameterGroupOutputArgs } from "./getClusterParameterGroup";
export const getClusterParameterGroup: typeof import("./getClusterParameterGroup").getClusterParameterGroup = null as any;
export const getClusterParameterGroupOutput: typeof import("./getClusterParameterGroup").getClusterParameterGroupOutput = null as any;

export { GetClusterSecurityGroupArgs, GetClusterSecurityGroupResult, GetClusterSecurityGroupOutputArgs } from "./getClusterSecurityGroup";
export const getClusterSecurityGroup: typeof import("./getClusterSecurityGroup").getClusterSecurityGroup = null as any;
export const getClusterSecurityGroupOutput: typeof import("./getClusterSecurityGroup").getClusterSecurityGroupOutput = null as any;

export { GetClusterSecurityGroupIngressArgs, GetClusterSecurityGroupIngressResult, GetClusterSecurityGroupIngressOutputArgs } from "./getClusterSecurityGroupIngress";
export const getClusterSecurityGroupIngress: typeof import("./getClusterSecurityGroupIngress").getClusterSecurityGroupIngress = null as any;
export const getClusterSecurityGroupIngressOutput: typeof import("./getClusterSecurityGroupIngress").getClusterSecurityGroupIngressOutput = null as any;

export { GetClusterSubnetGroupArgs, GetClusterSubnetGroupResult, GetClusterSubnetGroupOutputArgs } from "./getClusterSubnetGroup";
export const getClusterSubnetGroup: typeof import("./getClusterSubnetGroup").getClusterSubnetGroup = null as any;
export const getClusterSubnetGroupOutput: typeof import("./getClusterSubnetGroup").getClusterSubnetGroupOutput = null as any;

export { GetEndpointAccessArgs, GetEndpointAccessResult, GetEndpointAccessOutputArgs } from "./getEndpointAccess";
export const getEndpointAccess: typeof import("./getEndpointAccess").getEndpointAccess = null as any;
export const getEndpointAccessOutput: typeof import("./getEndpointAccess").getEndpointAccessOutput = null as any;

export { GetEndpointAuthorizationArgs, GetEndpointAuthorizationResult, GetEndpointAuthorizationOutputArgs } from "./getEndpointAuthorization";
export const getEndpointAuthorization: typeof import("./getEndpointAuthorization").getEndpointAuthorization = null as any;
export const getEndpointAuthorizationOutput: typeof import("./getEndpointAuthorization").getEndpointAuthorizationOutput = null as any;

export { GetEventSubscriptionArgs, GetEventSubscriptionResult, GetEventSubscriptionOutputArgs } from "./getEventSubscription";
export const getEventSubscription: typeof import("./getEventSubscription").getEventSubscription = null as any;
export const getEventSubscriptionOutput: typeof import("./getEventSubscription").getEventSubscriptionOutput = null as any;

export { GetScheduledActionArgs, GetScheduledActionResult, GetScheduledActionOutputArgs } from "./getScheduledAction";
export const getScheduledAction: typeof import("./getScheduledAction").getScheduledAction = null as any;
export const getScheduledActionOutput: typeof import("./getScheduledAction").getScheduledActionOutput = null as any;

export { ScheduledActionArgs } from "./scheduledAction";
export type ScheduledAction = import("./scheduledAction").ScheduledAction;
export const ScheduledAction: typeof import("./scheduledAction").ScheduledAction = null as any;

utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));
utilities.lazyLoad(exports, ["ClusterParameterGroup"], () => require("./clusterParameterGroup"));
utilities.lazyLoad(exports, ["ClusterSecurityGroup"], () => require("./clusterSecurityGroup"));
utilities.lazyLoad(exports, ["ClusterSecurityGroupIngress"], () => require("./clusterSecurityGroupIngress"));
utilities.lazyLoad(exports, ["ClusterSubnetGroup"], () => require("./clusterSubnetGroup"));
utilities.lazyLoad(exports, ["EndpointAccess"], () => require("./endpointAccess"));
utilities.lazyLoad(exports, ["EndpointAuthorization"], () => require("./endpointAuthorization"));
utilities.lazyLoad(exports, ["EventSubscription"], () => require("./eventSubscription"));
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));
utilities.lazyLoad(exports, ["getClusterParameterGroup","getClusterParameterGroupOutput"], () => require("./getClusterParameterGroup"));
utilities.lazyLoad(exports, ["getClusterSecurityGroup","getClusterSecurityGroupOutput"], () => require("./getClusterSecurityGroup"));
utilities.lazyLoad(exports, ["getClusterSecurityGroupIngress","getClusterSecurityGroupIngressOutput"], () => require("./getClusterSecurityGroupIngress"));
utilities.lazyLoad(exports, ["getClusterSubnetGroup","getClusterSubnetGroupOutput"], () => require("./getClusterSubnetGroup"));
utilities.lazyLoad(exports, ["getEndpointAccess","getEndpointAccessOutput"], () => require("./getEndpointAccess"));
utilities.lazyLoad(exports, ["getEndpointAuthorization","getEndpointAuthorizationOutput"], () => require("./getEndpointAuthorization"));
utilities.lazyLoad(exports, ["getEventSubscription","getEventSubscriptionOutput"], () => require("./getEventSubscription"));
utilities.lazyLoad(exports, ["getScheduledAction","getScheduledActionOutput"], () => require("./getScheduledAction"));
utilities.lazyLoad(exports, ["ScheduledAction"], () => require("./scheduledAction"));

// Export enums:
export * from "../types/enums/redshift";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:redshift:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:redshift:ClusterParameterGroup":
                return new ClusterParameterGroup(name, <any>undefined, { urn })
            case "aws-native:redshift:ClusterSecurityGroup":
                return new ClusterSecurityGroup(name, <any>undefined, { urn })
            case "aws-native:redshift:ClusterSecurityGroupIngress":
                return new ClusterSecurityGroupIngress(name, <any>undefined, { urn })
            case "aws-native:redshift:ClusterSubnetGroup":
                return new ClusterSubnetGroup(name, <any>undefined, { urn })
            case "aws-native:redshift:EndpointAccess":
                return new EndpointAccess(name, <any>undefined, { urn })
            case "aws-native:redshift:EndpointAuthorization":
                return new EndpointAuthorization(name, <any>undefined, { urn })
            case "aws-native:redshift:EventSubscription":
                return new EventSubscription(name, <any>undefined, { urn })
            case "aws-native:redshift:ScheduledAction":
                return new ScheduledAction(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "redshift", _module)