// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function importValue(args: ImportValueArgs, opts?: pulumi.InvokeOptions): Promise<ImportValueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:index:importValue", {
        "name": args.name,
    }, opts);
}

export interface ImportValueArgs {
    name: string;
}

export interface ImportValueResult {
    readonly value?: any;
}
export function importValueOutput(args: ImportValueOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ImportValueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:index:importValue", {
        "name": args.name,
    }, opts);
}

export interface ImportValueOutputArgs {
    name: pulumi.Input<string>;
}
