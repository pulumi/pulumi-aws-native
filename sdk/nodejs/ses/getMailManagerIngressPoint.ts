// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::SES::MailManagerIngressPoint Resource Type
 */
export function getMailManagerIngressPoint(args: GetMailManagerIngressPointArgs, opts?: pulumi.InvokeOptions): Promise<GetMailManagerIngressPointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ses:getMailManagerIngressPoint", {
        "ingressPointId": args.ingressPointId,
    }, opts);
}

export interface GetMailManagerIngressPointArgs {
    /**
     * The identifier of the ingress endpoint resource.
     */
    ingressPointId: string;
}

export interface GetMailManagerIngressPointResult {
    /**
     * The DNS A Record that identifies your ingress endpoint. Configure your DNS Mail Exchange (MX) record with this value to route emails to Mail Manager.
     */
    readonly aRecord?: string;
    /**
     * The Amazon Resource Name (ARN) of the ingress endpoint resource.
     */
    readonly ingressPointArn?: string;
    /**
     * The identifier of the ingress endpoint resource.
     */
    readonly ingressPointId?: string;
    /**
     * A user friendly name for an ingress endpoint resource.
     */
    readonly ingressPointName?: string;
    /**
     * The identifier of an existing rule set that you attach to an ingress endpoint resource.
     */
    readonly ruleSetId?: string;
    /**
     * The status of the ingress endpoint resource.
     */
    readonly status?: enums.ses.MailManagerIngressPointIngressPointStatus;
    /**
     * The update status of an ingress endpoint.
     */
    readonly statusToUpdate?: enums.ses.MailManagerIngressPointIngressPointStatusToUpdate;
    /**
     * The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The identifier of an existing traffic policy that you attach to an ingress endpoint resource.
     */
    readonly trafficPolicyId?: string;
}
/**
 * Definition of AWS::SES::MailManagerIngressPoint Resource Type
 */
export function getMailManagerIngressPointOutput(args: GetMailManagerIngressPointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMailManagerIngressPointResult> {
    return pulumi.output(args).apply((a: any) => getMailManagerIngressPoint(a, opts))
}

export interface GetMailManagerIngressPointOutputArgs {
    /**
     * The identifier of the ingress endpoint resource.
     */
    ingressPointId: pulumi.Input<string>;
}