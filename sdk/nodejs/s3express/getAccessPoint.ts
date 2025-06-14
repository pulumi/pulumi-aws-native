// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::S3Express::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.
 */
export function getAccessPoint(args: GetAccessPointArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:s3express:getAccessPoint", {
        "name": args.name,
    }, opts);
}

export interface GetAccessPointArgs {
    /**
     * The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. For directory buckets, the access point name must consist of a base name that you provide and suﬃx that includes the ZoneID (AWS Availability Zone or Local Zone) of your bucket location, followed by --xa-s3.
     */
    name: string;
}

export interface GetAccessPointResult {
    /**
     * The Amazon Resource Name (ARN) of the specified accesspoint.
     */
    readonly arn?: string;
    /**
     * Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
     */
    readonly networkOrigin?: enums.s3express.AccessPointNetworkOrigin;
    /**
     * The Access Point Policy you want to apply to this access point.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3Express::AccessPoint` for more information about the expected schema for this property.
     */
    readonly policy?: any;
    /**
     * The PublicAccessBlock configuration that you want to apply to this Access Point.
     */
    readonly publicAccessBlockConfiguration?: outputs.s3express.AccessPointPublicAccessBlockConfiguration;
    /**
     * For directory buckets, you can ﬁlter access control to speciﬁc preﬁxes, API operations, or a combination of both.
     */
    readonly scope?: outputs.s3express.AccessPointScope;
}
/**
 * The AWS::S3Express::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.
 */
export function getAccessPointOutput(args: GetAccessPointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccessPointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:s3express:getAccessPoint", {
        "name": args.name,
    }, opts);
}

export interface GetAccessPointOutputArgs {
    /**
     * The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. For directory buckets, the access point name must consist of a base name that you provide and suﬃx that includes the ZoneID (AWS Availability Zone or Local Zone) of your bucket location, followed by --xa-s3.
     */
    name: pulumi.Input<string>;
}
