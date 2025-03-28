// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda
 */
export function getAccessPointPolicy(args: GetAccessPointPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPointPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:s3objectlambda:getAccessPointPolicy", {
        "objectLambdaAccessPoint": args.objectLambdaAccessPoint,
    }, opts);
}

export interface GetAccessPointPolicyArgs {
    /**
     * The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
     */
    objectLambdaAccessPoint: string;
}

export interface GetAccessPointPolicyResult {
    /**
     * A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3ObjectLambda::AccessPointPolicy` for more information about the expected schema for this property.
     */
    readonly policyDocument?: any;
}
/**
 * AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda
 */
export function getAccessPointPolicyOutput(args: GetAccessPointPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccessPointPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:s3objectlambda:getAccessPointPolicy", {
        "objectLambdaAccessPoint": args.objectLambdaAccessPoint,
    }, opts);
}

export interface GetAccessPointPolicyOutputArgs {
    /**
     * The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
     */
    objectLambdaAccessPoint: pulumi.Input<string>;
}
