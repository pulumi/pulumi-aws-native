// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.
 */
export function getDomainUnit(args: GetDomainUnitArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainUnitResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:datazone:getDomainUnit", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetDomainUnitArgs {
    /**
     * The ID of the domain where the domain unit was created.
     */
    domainId: string;
    /**
     * The ID of the domain unit.
     */
    id: string;
}

export interface GetDomainUnitResult {
    /**
     * The timestamp at which the domain unit was created.
     */
    readonly createdAt?: string;
    /**
     * The description of the domain unit.
     */
    readonly description?: string;
    /**
     * The ID of the domain where the domain unit was created.
     */
    readonly domainId?: string;
    /**
     * The ID of the domain unit.
     */
    readonly id?: string;
    /**
     * The identifier of the domain unit that you want to get.
     */
    readonly identifier?: string;
    /**
     * The timestamp at which the domain unit was last updated.
     */
    readonly lastUpdatedAt?: string;
    /**
     * The name of the domain unit.
     */
    readonly name?: string;
    /**
     * The ID of the parent domain unit.
     */
    readonly parentDomainUnitId?: string;
}
/**
 * A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.
 */
export function getDomainUnitOutput(args: GetDomainUnitOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDomainUnitResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:datazone:getDomainUnit", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetDomainUnitOutputArgs {
    /**
     * The ID of the domain where the domain unit was created.
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the domain unit.
     */
    id: pulumi.Input<string>;
}
