// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::VerifiedPermissions::PolicyStore Resource Type
 */
export function getPolicyStore(args: GetPolicyStoreArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyStoreResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:verifiedpermissions:getPolicyStore", {
        "policyStoreId": args.policyStoreId,
    }, opts);
}

export interface GetPolicyStoreArgs {
    policyStoreId: string;
}

export interface GetPolicyStoreResult {
    readonly arn?: string;
    readonly policyStoreId?: string;
    readonly schema?: outputs.verifiedpermissions.PolicyStoreSchemaDefinition;
    readonly validationSettings?: outputs.verifiedpermissions.PolicyStoreValidationSettings;
}
/**
 * Definition of AWS::VerifiedPermissions::PolicyStore Resource Type
 */
export function getPolicyStoreOutput(args: GetPolicyStoreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyStoreResult> {
    return pulumi.output(args).apply((a: any) => getPolicyStore(a, opts))
}

export interface GetPolicyStoreOutputArgs {
    policyStoreId: pulumi.Input<string>;
}