// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A calculated attribute definition for Customer Profiles
 */
export function getCalculatedAttributeDefinition(args: GetCalculatedAttributeDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetCalculatedAttributeDefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:customerprofiles:getCalculatedAttributeDefinition", {
        "calculatedAttributeName": args.calculatedAttributeName,
        "domainName": args.domainName,
    }, opts);
}

export interface GetCalculatedAttributeDefinitionArgs {
    calculatedAttributeName: string;
    domainName: string;
}

export interface GetCalculatedAttributeDefinitionResult {
    readonly attributeDetails?: outputs.customerprofiles.CalculatedAttributeDefinitionAttributeDetails;
    readonly conditions?: outputs.customerprofiles.CalculatedAttributeDefinitionConditions;
    /**
     * The timestamp of when the calculated attribute definition was created.
     */
    readonly createdAt?: string;
    readonly description?: string;
    readonly displayName?: string;
    /**
     * The timestamp of when the calculated attribute definition was most recently edited.
     */
    readonly lastUpdatedAt?: string;
    readonly statistic?: enums.customerprofiles.CalculatedAttributeDefinitionStatistic;
    readonly tags?: outputs.customerprofiles.CalculatedAttributeDefinitionTag[];
}
/**
 * A calculated attribute definition for Customer Profiles
 */
export function getCalculatedAttributeDefinitionOutput(args: GetCalculatedAttributeDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCalculatedAttributeDefinitionResult> {
    return pulumi.output(args).apply((a: any) => getCalculatedAttributeDefinition(a, opts))
}

export interface GetCalculatedAttributeDefinitionOutputArgs {
    calculatedAttributeName: pulumi.Input<string>;
    domainName: pulumi.Input<string>;
}