// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
 */
export function getWebACL(args: GetWebACLArgs, opts?: pulumi.InvokeOptions): Promise<GetWebACLResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:wafv2:getWebACL", {
        "id": args.id,
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

export interface GetWebACLArgs {
    id: string;
    name: string;
    scope: enums.wafv2.WebACLScope;
}

export interface GetWebACLResult {
    readonly arn?: string;
    readonly capacity?: number;
    readonly captchaConfig?: outputs.wafv2.WebACLCaptchaConfig;
    readonly customResponseBodies?: outputs.wafv2.WebACLCustomResponseBodies;
    readonly defaultAction?: outputs.wafv2.WebACLDefaultAction;
    readonly description?: string;
    readonly id?: string;
    readonly labelNamespace?: string;
    /**
     * Collection of Rules.
     */
    readonly rules?: outputs.wafv2.WebACLRule[];
    readonly tags?: outputs.wafv2.WebACLTag[];
    readonly visibilityConfig?: outputs.wafv2.WebACLVisibilityConfig;
}

export function getWebACLOutput(args: GetWebACLOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWebACLResult> {
    return pulumi.output(args).apply(a => getWebACL(a, opts))
}

export interface GetWebACLOutputArgs {
    id: pulumi.Input<string>;
    name: pulumi.Input<string>;
    scope: pulumi.Input<enums.wafv2.WebACLScope>;
}