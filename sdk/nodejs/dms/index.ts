// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CertificateArgs } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;

export { EndpointArgs } from "./endpoint";
export type Endpoint = import("./endpoint").Endpoint;
export const Endpoint: typeof import("./endpoint").Endpoint = null as any;

export { EventSubscriptionArgs } from "./eventSubscription";
export type EventSubscription = import("./eventSubscription").EventSubscription;
export const EventSubscription: typeof import("./eventSubscription").EventSubscription = null as any;

export { GetCertificateArgs, GetCertificateResult, GetCertificateOutputArgs } from "./getCertificate";
export const getCertificate: typeof import("./getCertificate").getCertificate = null as any;
export const getCertificateOutput: typeof import("./getCertificate").getCertificateOutput = null as any;

export { GetEndpointArgs, GetEndpointResult, GetEndpointOutputArgs } from "./getEndpoint";
export const getEndpoint: typeof import("./getEndpoint").getEndpoint = null as any;
export const getEndpointOutput: typeof import("./getEndpoint").getEndpointOutput = null as any;

export { GetEventSubscriptionArgs, GetEventSubscriptionResult, GetEventSubscriptionOutputArgs } from "./getEventSubscription";
export const getEventSubscription: typeof import("./getEventSubscription").getEventSubscription = null as any;
export const getEventSubscriptionOutput: typeof import("./getEventSubscription").getEventSubscriptionOutput = null as any;

export { GetReplicationInstanceArgs, GetReplicationInstanceResult, GetReplicationInstanceOutputArgs } from "./getReplicationInstance";
export const getReplicationInstance: typeof import("./getReplicationInstance").getReplicationInstance = null as any;
export const getReplicationInstanceOutput: typeof import("./getReplicationInstance").getReplicationInstanceOutput = null as any;

export { GetReplicationSubnetGroupArgs, GetReplicationSubnetGroupResult, GetReplicationSubnetGroupOutputArgs } from "./getReplicationSubnetGroup";
export const getReplicationSubnetGroup: typeof import("./getReplicationSubnetGroup").getReplicationSubnetGroup = null as any;
export const getReplicationSubnetGroupOutput: typeof import("./getReplicationSubnetGroup").getReplicationSubnetGroupOutput = null as any;

export { GetReplicationTaskArgs, GetReplicationTaskResult, GetReplicationTaskOutputArgs } from "./getReplicationTask";
export const getReplicationTask: typeof import("./getReplicationTask").getReplicationTask = null as any;
export const getReplicationTaskOutput: typeof import("./getReplicationTask").getReplicationTaskOutput = null as any;

export { ReplicationInstanceArgs } from "./replicationInstance";
export type ReplicationInstance = import("./replicationInstance").ReplicationInstance;
export const ReplicationInstance: typeof import("./replicationInstance").ReplicationInstance = null as any;

export { ReplicationSubnetGroupArgs } from "./replicationSubnetGroup";
export type ReplicationSubnetGroup = import("./replicationSubnetGroup").ReplicationSubnetGroup;
export const ReplicationSubnetGroup: typeof import("./replicationSubnetGroup").ReplicationSubnetGroup = null as any;

export { ReplicationTaskArgs } from "./replicationTask";
export type ReplicationTask = import("./replicationTask").ReplicationTask;
export const ReplicationTask: typeof import("./replicationTask").ReplicationTask = null as any;

utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));
utilities.lazyLoad(exports, ["Endpoint"], () => require("./endpoint"));
utilities.lazyLoad(exports, ["EventSubscription"], () => require("./eventSubscription"));
utilities.lazyLoad(exports, ["getCertificate","getCertificateOutput"], () => require("./getCertificate"));
utilities.lazyLoad(exports, ["getEndpoint","getEndpointOutput"], () => require("./getEndpoint"));
utilities.lazyLoad(exports, ["getEventSubscription","getEventSubscriptionOutput"], () => require("./getEventSubscription"));
utilities.lazyLoad(exports, ["getReplicationInstance","getReplicationInstanceOutput"], () => require("./getReplicationInstance"));
utilities.lazyLoad(exports, ["getReplicationSubnetGroup","getReplicationSubnetGroupOutput"], () => require("./getReplicationSubnetGroup"));
utilities.lazyLoad(exports, ["getReplicationTask","getReplicationTaskOutput"], () => require("./getReplicationTask"));
utilities.lazyLoad(exports, ["ReplicationInstance"], () => require("./replicationInstance"));
utilities.lazyLoad(exports, ["ReplicationSubnetGroup"], () => require("./replicationSubnetGroup"));
utilities.lazyLoad(exports, ["ReplicationTask"], () => require("./replicationTask"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:dms:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws-native:dms:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "aws-native:dms:EventSubscription":
                return new EventSubscription(name, <any>undefined, { urn })
            case "aws-native:dms:ReplicationInstance":
                return new ReplicationInstance(name, <any>undefined, { urn })
            case "aws-native:dms:ReplicationSubnetGroup":
                return new ReplicationSubnetGroup(name, <any>undefined, { urn })
            case "aws-native:dms:ReplicationTask":
                return new ReplicationTask(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "dms", _module)