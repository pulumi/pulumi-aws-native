import * as pulumi from "@pulumi/pulumi";

const __config = new pulumi.Config("cloudformation");
export const region = __config.get("region");
export const stack = __config.get("stack");

export interface StackArgs {
    stack?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}

export class Stack extends pulumi.ProviderResource {
    constructor(name: string, args: StackArgs, opts?: pulumi.ResourceOptions) {
        const inputs: pulumi.Inputs = {
            stack: args.stack || name,
            region: args.region,
        };

        super("cloudformation", name, inputs, opts);
    }
}

export function getAccountId(opts?: pulumi.InvokeOptions): Promise<string> {
    return pulumi.runtime.invoke("cloudformation:index:getAccountId", {}, opts).then((res: any) => res.accountId);
}

export function getAzs(opts?: pulumi.InvokeOptions): Promise<string[]> {
    return pulumi.runtime.invoke("cloudformation:index:getAzs", {}, opts).then((res: any) => res.azs);
}

export function getStackId(opts?: pulumi.InvokeOptions): Promise<string> {
    return pulumi.runtime.invoke("cloudformation:index:getStackId", {}, opts).then((res: any) => res.stackId);
}

export interface ResourceOptions {
    logicalId?: pulumi.Input<string>;
    metadata?: pulumi.Inputs;
    creationPolicy?: pulumi.Inputs;
    deletionPolicy?: pulumi.Inputs;
    updatePolicy?: pulumi.Inputs;
    updateReplacePolicy?: pulumi.Inputs;
}
{{range $_, $type := .RootTypes}}
export interface {{$type.Name}} {
    {{- range $_, $member := $type.Members}}
    {{$member.Name}}{{if $member.Optional}}?{{end}}: {{$member.Type}};
    {{- end}}
}
{{end}}
{{range $_, $module := .Modules}}
export namespace {{$module.Namespace}} {
    {{range $_, $type := $module.NestedTypes}}
    export interface {{$type.Name}} {
        {{- range $_, $member := $type.Members}}
        {{$member.Name}}{{if $member.Optional}}?{{end}}: {{$member.Type}};
        {{- end}}
    }
    {{end}}
    {{range $_, $type := $module.Resources}}
    export class {{$type.Name}} extends pulumi.CustomResource {
        public readonly attributes: pulumi.Output<{{$type.AttributesType}}>;

        constructor(name: string, properties{{if $type.PropertiesOptional}}?{{end}}: {{$type.PropertiesType}}, cfnOpts?: ResourceOptions, opts?: pulumi.CustomResourceOptions) {
            const inputs: pulumi.Inputs = {
                ...(cfnOpts || {}),
                properties: properties,
                attributes: undefined,
            };
            super("{{$type.Token}}", name, inputs, opts);
        }
    }
    {{end}}
}
{{end}}
