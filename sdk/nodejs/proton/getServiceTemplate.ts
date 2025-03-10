// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Proton::ServiceTemplate Resource Type
 */
export function getServiceTemplate(args: GetServiceTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:proton:getServiceTemplate", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceTemplateArgs {
    /**
     * <p>The Amazon Resource Name (ARN) of the service template.</p>
     */
    arn: string;
}

export interface GetServiceTemplateResult {
    /**
     * <p>The Amazon Resource Name (ARN) of the service template.</p>
     */
    readonly arn?: string;
    /**
     * <p>A description of the service template.</p>
     */
    readonly description?: string;
    /**
     * <p>The name of the service template as displayed in the developer interface.</p>
     */
    readonly displayName?: string;
    /**
     * <p>An optional list of metadata items that you can associate with the Proton service template. A tag is a key-value pair.</p>
     *          <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
     *         <i>Proton User Guide</i>.</p>
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Definition of AWS::Proton::ServiceTemplate Resource Type
 */
export function getServiceTemplateOutput(args: GetServiceTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:proton:getServiceTemplate", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceTemplateOutputArgs {
    /**
     * <p>The Amazon Resource Name (ARN) of the service template.</p>
     */
    arn: pulumi.Input<string>;
}
