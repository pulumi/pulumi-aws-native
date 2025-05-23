// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetListenerArgs, GetListenerResult, GetListenerOutputArgs } from "./getListener";
export const getListener: typeof import("./getListener").getListener = null as any;
export const getListenerOutput: typeof import("./getListener").getListenerOutput = null as any;
utilities.lazyLoad(exports, ["getListener","getListenerOutput"], () => require("./getListener"));

export { GetListenerRuleArgs, GetListenerRuleResult, GetListenerRuleOutputArgs } from "./getListenerRule";
export const getListenerRule: typeof import("./getListenerRule").getListenerRule = null as any;
export const getListenerRuleOutput: typeof import("./getListenerRule").getListenerRuleOutput = null as any;
utilities.lazyLoad(exports, ["getListenerRule","getListenerRuleOutput"], () => require("./getListenerRule"));

export { GetLoadBalancerArgs, GetLoadBalancerResult, GetLoadBalancerOutputArgs } from "./getLoadBalancer";
export const getLoadBalancer: typeof import("./getLoadBalancer").getLoadBalancer = null as any;
export const getLoadBalancerOutput: typeof import("./getLoadBalancer").getLoadBalancerOutput = null as any;
utilities.lazyLoad(exports, ["getLoadBalancer","getLoadBalancerOutput"], () => require("./getLoadBalancer"));

export { GetTargetGroupArgs, GetTargetGroupResult, GetTargetGroupOutputArgs } from "./getTargetGroup";
export const getTargetGroup: typeof import("./getTargetGroup").getTargetGroup = null as any;
export const getTargetGroupOutput: typeof import("./getTargetGroup").getTargetGroupOutput = null as any;
utilities.lazyLoad(exports, ["getTargetGroup","getTargetGroupOutput"], () => require("./getTargetGroup"));

export { GetTrustStoreArgs, GetTrustStoreResult, GetTrustStoreOutputArgs } from "./getTrustStore";
export const getTrustStore: typeof import("./getTrustStore").getTrustStore = null as any;
export const getTrustStoreOutput: typeof import("./getTrustStore").getTrustStoreOutput = null as any;
utilities.lazyLoad(exports, ["getTrustStore","getTrustStoreOutput"], () => require("./getTrustStore"));

export { GetTrustStoreRevocationArgs, GetTrustStoreRevocationResult, GetTrustStoreRevocationOutputArgs } from "./getTrustStoreRevocation";
export const getTrustStoreRevocation: typeof import("./getTrustStoreRevocation").getTrustStoreRevocation = null as any;
export const getTrustStoreRevocationOutput: typeof import("./getTrustStoreRevocation").getTrustStoreRevocationOutput = null as any;
utilities.lazyLoad(exports, ["getTrustStoreRevocation","getTrustStoreRevocationOutput"], () => require("./getTrustStoreRevocation"));

export { ListenerArgs } from "./listener";
export type Listener = import("./listener").Listener;
export const Listener: typeof import("./listener").Listener = null as any;
utilities.lazyLoad(exports, ["Listener"], () => require("./listener"));

export { ListenerRuleArgs } from "./listenerRule";
export type ListenerRule = import("./listenerRule").ListenerRule;
export const ListenerRule: typeof import("./listenerRule").ListenerRule = null as any;
utilities.lazyLoad(exports, ["ListenerRule"], () => require("./listenerRule"));

export { LoadBalancerArgs } from "./loadBalancer";
export type LoadBalancer = import("./loadBalancer").LoadBalancer;
export const LoadBalancer: typeof import("./loadBalancer").LoadBalancer = null as any;
utilities.lazyLoad(exports, ["LoadBalancer"], () => require("./loadBalancer"));

export { TargetGroupArgs } from "./targetGroup";
export type TargetGroup = import("./targetGroup").TargetGroup;
export const TargetGroup: typeof import("./targetGroup").TargetGroup = null as any;
utilities.lazyLoad(exports, ["TargetGroup"], () => require("./targetGroup"));

export { TrustStoreArgs } from "./trustStore";
export type TrustStore = import("./trustStore").TrustStore;
export const TrustStore: typeof import("./trustStore").TrustStore = null as any;
utilities.lazyLoad(exports, ["TrustStore"], () => require("./trustStore"));

export { TrustStoreRevocationArgs } from "./trustStoreRevocation";
export type TrustStoreRevocation = import("./trustStoreRevocation").TrustStoreRevocation;
export const TrustStoreRevocation: typeof import("./trustStoreRevocation").TrustStoreRevocation = null as any;
utilities.lazyLoad(exports, ["TrustStoreRevocation"], () => require("./trustStoreRevocation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:elasticloadbalancingv2:Listener":
                return new Listener(name, <any>undefined, { urn })
            case "aws-native:elasticloadbalancingv2:ListenerRule":
                return new ListenerRule(name, <any>undefined, { urn })
            case "aws-native:elasticloadbalancingv2:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "aws-native:elasticloadbalancingv2:TargetGroup":
                return new TargetGroup(name, <any>undefined, { urn })
            case "aws-native:elasticloadbalancingv2:TrustStore":
                return new TrustStore(name, <any>undefined, { urn })
            case "aws-native:elasticloadbalancingv2:TrustStoreRevocation":
                return new TrustStoreRevocation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "elasticloadbalancingv2", _module)
