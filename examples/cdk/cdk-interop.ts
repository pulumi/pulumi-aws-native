import * as cdk from '@aws-cdk/core';
import * as pulumi from "@pulumi/pulumi";

function firstToLower(str: string) {
    return str.replace(
      /\w\S*/g,
      function(txt) {
        return txt.charAt(0).toLowerCase() + txt.substr(1);
      }
    );
}

function normalize(value: any): any {
    if (!value) return value;

    if (Array.isArray(value)) {
        const result: any[] = [];
        for(let i=0; i<value.length; i++){
            result[i] = normalize(value[i]);
        }
        return result
    }

    if (typeof value !== 'object') return value;

    const result :any = {};
    Object.entries(value).forEach(([k, v]) => {
        result[firstToLower(k)] = normalize(v);
    });
    return result;
}

class CdkResource extends pulumi.CustomResource {
    constructor(name: string, type: string, args: any, opts?: pulumi.CustomResourceOptions) {
        const [_, mod, res] = type.split("::");
        super(`aws-native:${mod}:${res}`, name, args, opts);
    }
}

export class CdkComponent extends pulumi.ComponentResource {
    constructor(name: string, args: (stack: cdk.Stack) => void, opts?: pulumi.CustomResourceOptions) {
        super("aws-native:cdk:Component", name, args, opts);

        const app = new cdk.App();
        const stack = new cdk.Stack(app);
        args(stack);

        //debugger;
        const template = app.synth().getStackByName(stack.stackName).template;
        console.debug(template.Resources.appsvc.Properties);
        const resources = template.Resources;

        Object.entries(resources).forEach(
            ([key, value]) => {
                const typeName = (value as any).Type;
                const sourceProps = (value as any).Properties;
                opts = opts || { parent: this };
                new CdkResource(key, typeName, normalize(sourceProps), opts);
            });
    }
}