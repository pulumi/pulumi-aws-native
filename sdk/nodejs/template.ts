import * as pulumi from "@pulumi/pulumi";
import * as cloudformation from "./cloudformation";
import * as fs from "fs";

// TODO(pdg): hook up first-class providers, elaborate nested stacks, etc.

const region = cloudformation.region!;
const stack = cloudformation.stack!;
const accountId = cloudformation.getAccountId();
const stackId = cloudformation.getStackId();

async function apply<T, U>(output: pulumi.Output<T>, fn: (t: T) => U): Promise<pulumi.Input<U>> {
    if (await (<any>output).isKnown) {
        return fn(<T>(await (<any>output).promise()));
    }
    return output.apply(fn);
}

// Generic container for resources.
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("cloudformation:template:component", name, {}, opts);
    }
}

// Generic CloudFormation resource to avoid mapping from CloudFormation type -> Pulumi type.
class CloudFormationResource extends pulumi.CustomResource {
    public readonly attributes!: pulumi.Output<any>;

    constructor(tok: string, name: string, inputs: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super(tok, name, inputs, opts);
    }
}

// functions records the set of standard CloudFormation intrinsic functions.
const functions = new Set<string>([
    "Fn::And",
    "Fn::Base64",
    "Fn::Cidr",
    "Fn::Equals",
    "Fn::FindInMap",
    "Fn::GetAtt",
    "Fn::GetAZs",
    "Fn::If",
    "Fn::ImportValue",
    "Fn::Join",
    "Fn::Not",
    "Fn::Or",
    "Fn::Select",
    "Fn::Split",
    "Fn::Sub",
    "Fn::Transform",
    "Ref",
]);

// EvalContext holds the information necessary to evaluate a CloudFormation expression
interface EvalContext {
    // The name of the context. Used in error messages.
    name: string;
    // The set of mappings for the stack.
    mappings: any;
    // The evaluated conditions for the stack.
    conditions: Record<string, pulumi.Input<boolean>>;
    // The evaluated parameter values for the stack.
    parameterValues: Record<string, any>;
    // The evaluated resources for the stack.
    resources: Record<string, CloudFormationResource | null>;
}

// subPattern is used by evalSub and findSubDeps to find the interpolations in a `Fn::Sub` template string.
const subPattern = /\${([^}]*)}/;

// evalSub evaluates a call to `Fn::Sub` by replacing all interpolations as appropriate.
async function evalSub(args: any[], context: EvalContext): Promise<any> {
    let [template, vars] = args;

    // Find all interpolations in the template.
    const result = [];

    let match;
    while (match = subPattern.exec(template)) {
        // If there's literal string data that precedes the match, stuff it into the result.
        if (match.index > 0) {
            result.push(template.slice(0, match.index));
        }
        const interpolation = match[1];
        if (interpolation.length == 0) {
            continue;
        } else if (interpolation[0] == '!') {
            // Literal value.
            result.push("${" + interpolation.slice(1, interpolation.length) + "}");
        } else if (vars && (interpolation in vars)) {
            // Interpolation variable.
            result.push(vars[interpolation]);
        } else if (interpolation in context.parameterValues) {
            // Parameter value.
            result.push(context.parameterValues[interpolation]);
        } else {
            // Resource reference or attribute, depending on whether or not a '.' is present.
            const args = interpolation.split(".", 2);
            if (args.length === 1) {
                result.push(await evalExpr({ "Ref": args }, context, false));
            } else {
                result.push(await evalExpr({ "Fn::GetAtt": args }, context, false));
            }
        }

        // Slice off the match.
        template = template.slice(match.index + match[0].length, template.length);
    }
    // If there's literal string data after the last interpolation, stuff it into the result.
    if (template.length > 0) {
        result.push(template);
    }

    return apply(pulumi.all(result), arr => arr.join(""));
}

// findSubDeps finds all resource references inside of the template string in a call to `Fn::Sub`.
function findSubDeps(args: any[], parameterValues: Record<string, any>): string[] {
    var [template, vars] = args;

    // Find all interpolations in the template.
    const result = [];

    let match;
    while (match = subPattern.exec(template)) {
        const interpolation = match[1];
        if (interpolation.length == 0) {
            continue;
        } else if (interpolation[0] == '!') {
            // Ignore: literal values do not refer to resources.
        } else if (vars && (interpolation in vars)) {
            // Ignore: any references in the vars will have been found by the caller.
        } else if (interpolation in parameterValues) {
            // Ignore: parameters do not refer to resources.
        } else {
            // Resource reference.
            const [resource] = interpolation.split(".", 2);
            result.push(resource);
        }
        template = template.slice(match.index + match[0].length, template.length);
    }

    return result;
}

// evalEquals returns true if the given evaluated expressions are equal.
function evalEquals(a: any, b: any): boolean {
    if (a === null) {
        return b === null;
    } else if (typeof a !== "object") {
        return typeof b !== "object" && `${a}` == `${b}`;
    } else if (a instanceof Array) {
        return b instanceof Array && a.every((e, i) => evalEquals(a[i], b[i]));
    }

    return Object.keys(a).every(k => evalEquals(a[k], b[k])) && Object.keys(b).every(k => evalEquals(b[k], a[k]));
}

// evalFn evaluates a CloudFormation intrinsic function.
async function evalFn(fn: string, args: any[], context: EvalContext): Promise<any> {
    switch (fn) {
        case "Condition": {
            const [cond] = args;
            return context.conditions[<string>cond] || false;
        }

        case "Fn::And":
            return args.every(arg => !!arg);

        case "Fn::Base64":
            return new Buffer(args[0]).toString("base64");

        case "Fn::Cidr":
            // Leave these in as they are and let CloudFormation handle them. Note that this defeats diffing if the
            // CIDR values actually change.
            return { "Fn::Cidr": args };

        case "Fn::Equals":
            return evalEquals(args[0], args[1]);

        case "Fn::FindInMap": {
            const [mapName, topLevelKey, secondLevelKey] = args;
            return context.mappings[mapName][topLevelKey][secondLevelKey];
        }

        case "Fn::GetAtt": {
            const [resourceName, attributeName] = args;
            const resource = context.resources[resourceName];
            if (resource === undefined) {
                throw new Error(`${context.name} refers to missing resource ${resourceName}`);
            } else if (resource === null) {
                // The referenced resource is defined but is not being created. Tolerating these references makes the
                // evaluator simpler, as it can evaluate both branches of a `Fn::If` without needing to worry about
                // references to unknown resources.
                return undefined;
            }
            return apply(resource.attributes, (attrs: any) => attrs[attributeName]);
        }

        case "Fn::GetAZs": {
            const [azRegion] = args;
            return cloudformation.getAzs(azRegion || region);
        }

        case "Fn::If": {
            const [cond, ifv, elsev] = args;
            return apply(pulumi.output(context.conditions[cond]), v => v ? ifv : elsev);
        }

        case "Fn::ImportValue":
            // Leave these in as they are and let CloudFormation handle them. Note that this defeats diffing if the
            // imported values actually change.
            return { "Fn::ImportValue": args[0] };
            
        case "Fn::Join": {
            const [delimiter, list] = args;
            return list.join(delimiter);
        }

        case "Fn::Not":
            return !args[0];

        case "Fn::Or":
            return args.some(arg => !!arg);

        case "Fn::Select": {
            const [index, list] = args;
            return list[index];
        }

        case "Fn::Split": {
            const [delimiter, str] = args;
            return str.split(delimiter);
        }

        case "Fn::Sub":
            return evalSub(args, context);

        case "Fn::Transform":
            // Leave these in as they are and let CloudFormation handle them. Note that this defeats diffing if the
            // transformed values actually change.
            return { "Fn::Transform": args[0] };

        case "Ref": {
            const entityName = args[0];
            if (entityName in context.parameterValues) {
                return context.parameterValues[entityName];
            }
            const resource = context.resources[entityName];
            if (resource === undefined) {
                throw new Error(`${context.name} refers to missing resource ${entityName}`);
            } else if (resource === null) {
                // The referenced resource is defined but is not being created. Tolerating these references makes the
                // evaluator simpler, as it can evaluate both branches of a `Fn::If` without needing to worry about
                // references to unknown resources.
                return undefined;
            }
            return resource.id;
        }

        default:
            throw new Error(`unexpected function ${fn} passed to evalFn`);
    }
}

// evalExpr evaluates a CloudFormation expression.
async function evalExpr(expr: any, context: EvalContext, inCondition: boolean): Promise<any> {
    // If this is a primitive or null, leave it be.
    if (expr === null || typeof expr !== "object") {
        return expr;
    }

    // If this is an array, evaluate each element.
    if (expr instanceof Array) {
        return expr.map(e => evalExpr(e, context, inCondition));
    }

    // If this is an output, evaluate the eventual result. If the value is known promptly, evaluate it now.
    if (pulumi.Output.isInstance(expr)) {
        return apply(expr, v => evalExpr(v, context, inCondition));
    }

    // Otherwise, check for a function call.
    let fn = "";
    const exprKeys = Object.keys(expr);
    if (exprKeys.length === 1) {
        fn = exprKeys[0];
    }

    // If this is a function call, evaluate its args and then evaluate it.
    if (functions.has(fn) || (inCondition && fn == "Condition")) {
        let fnArgs = expr[fn];
        if (!(fnArgs instanceof Array)) {
            fnArgs = [fnArgs];
        }

        const args = pulumi.all(<any[]>fnArgs.map((arg: any) => evalExpr(arg, context, inCondition)));
        return apply(args, args => evalFn(fn, args, context));
    }

    // Otherwise, this is a normal object. Evaluate the value of each of its properties.
    return exprKeys.reduce((o: any, k: string) => {
        o[k] = evalExpr(expr[k], context, inCondition);
        return o;
    }, {});
}

// findDeps finds the resource or condition dependencies in a CloudFormation expression.
function findDeps(expr: any, parameterValues: Record<string, any>, isCondition: boolean): string[] {
    // If this is a primitive or null, it has no dependencies.
    if (expr === null || typeof expr !== "object") {
        return [];
    }

    // If this is an array, find the dependencies of each element.
    if (expr instanceof Array) {
        return expr.reduce((deps, e) => {
            deps.push(...findDeps(e, parameterValues, isCondition));
            return deps;
        }, []);
    }

    // Otherwise, check for a function call.
    let fn = "";
    const exprKeys = Object.keys(expr);
    if (exprKeys.length === 1) {
        fn = exprKeys[0];
    }

    let args = expr[fn];
    if (!(args instanceof Array)) {
        args = [args];
    }

    // If we're in a condition definition, the only call we need to look for is "Condition". If we're in a normal
    // expression, look for `Fn::GetAtt`, `Fn::Sub`, and `Fn::Ref`.
    if (isCondition) {
        if (fn == "Condition") {
            const [condition] = args;
            return [condition];
        }
    } else {
        switch (fn) {
            case "Fn::GetAtt": {
                const [resource] = args;
                return [resource];
            }

            case "Fn::Sub":
                return findSubDeps(args, parameterValues);

            case "Ref": {
                const [resource] = args;
                return resource in parameterValues ? [] : [resource];
            }
        }
    }

    // If this is not a function call, it is a normal object. Find the dependencies in each of its properties.
    return exprKeys.reduce((deps: string[], k: string) => {
        deps.push(...findDeps(<any>expr[k], parameterValues, isCondition));
        return deps;
    }, []);
}

// topologicalSort performs a topological sort of the definitions in the given map.
function topologicalSort(defs: any, parameterValues: Record<string, any>, isCondition: boolean): string[] {
    const order: string[] = [];
    const done = new Set<string>();

    function sort(defName: string) {
        if (done.has(defName)) {
            return;
        }
        done.add(defName);

        const def = defs[defName];
        if (def === undefined) {
            throw new Error(`missing def for ${defName}`);
        }
        const deps = findDeps(def, parameterValues, isCondition);

        const dependsOn = def.DependsOn;
        if (typeof dependsOn === "string") {
            deps.push(dependsOn);
        } else if (dependsOn instanceof Array) {
            deps.push(...dependsOn);
        }

        for (const d of new Set<string>(deps)) {
            sort(d);
        }
        order.push(defName);
    }

    for (const defName of Object.keys(defs)) {
        sort(defName);
    }

    return order;
}


interface RegionInfo {
    partition: string;
    urlSuffix: string;
}

// getRegionInfo returns the partition and URL suffix for the given region.
function getRegionInfo(region: string): RegionInfo {
    switch (region) {
        case "cn-north-1":
        case "cn-northwest-1":
            return {partition: "aws-cn", urlSuffix: "amazonaws.com.cn"};
        default:
            return {partition: "aws", urlSuffix: "amazonaws.com"};
    }
}

// Tree records the logical structure of a CloudFormation template.
export interface Tree {
    root: any;
    parents: Record<string, string>;
}

function createTree(path: string, node: any, mapping: Record<string, Component>, parent?: Component) {
    const component = new Component(path, { parent });
    mapping[path] = component;

    for (const child of Object.keys(node)) {
        createTree(`${path}/${child}`, node[child], mapping, component);
    }
}

// load loads the given CloudFormation template using the specified parameters.
export async function load(template: any, parameters: any, tree?: Tree): Promise<any> {
    // Pull the pieces we need out of the template.
    const parameterDefs = template["Parameters"] || {};
    const mappings = template["Mappings"] || {};
    const conditionDefs = template["Conditions"] || {};
    const resourceDefs = template["Resources"] || {};
    const outputDefs = template["Outputs"] || {};

    // Set up the containers we'll use to hold our evaluated state.
    const conditions: Record<string, pulumi.Input<boolean>> = {};
    const resources: Record<string, CloudFormationResource | null> = {};
    const outputs: Record<string, any> = {};

    // Get the region info.
    const regionInfo = getRegionInfo(region);

    // Evaluate parameter values.
    const parameterValues: Record<string, any> = {
        "AWS::NoValue": null,
        "AWS::AccountId": accountId,
        "AWS::Partition": regionInfo.partition,
        "AWS::Region": region,
        "AWS::StackName": stack,
        "AWS::StackId": stackId,
        "AWS::URLSuffix": regionInfo.urlSuffix,
    };
    for (const k of Object.keys(parameterDefs)) {
        const def = parameterDefs[k];
        const value = parameters[k] || def.Default;
        if (value === undefined) {
            throw new Error(`missing required parameter ${k}`);
        }
        parameterValues[k] = value;
    }

    // Set up our evaluation context and evaluation function.
    const evalContext = {conditions, parameterValues, mappings, resources};
    const evaluate = (name: string, e: any, inCondition?: boolean) => evalExpr(e, { ...evalContext, name }, !!inCondition);

    // Evaluate conditions in topological order.
    const conditionOrder = topologicalSort(conditionDefs, parameterValues, true);
    for (const k of conditionOrder) {
        conditions[k] = await apply(pulumi.output(evaluate(k, conditionDefs[k], true)), v => !!v);
    }

    // If we were passed a tree, create its container nodes.
    tree = tree || {root: {}, parents: {}};
    const componentMap: Record<string, Component> = {};
    for (const k of Object.keys(tree.root)) {
        createTree(k, tree.root[k], componentMap);
    }

    // Create resources in topological order.
    const resourceOrder = topologicalSort(resourceDefs, parameterValues, false);
    for (const resourceName of resourceOrder) {
        const resourceDef = resourceDefs[resourceName];

        // If this is a resource that we're not going to create because its condition is false, record the fact that we
        // are not going to create it. This simplifies the evaluator, as it can evaluate both branches of a `Fn::If`
        // without needing to worry about references to unknown resources.
        //
        // Note that if the condition is not promptly known, we will consider it to be true.
        if (resourceDef.Condition && !conditions[<string>resourceDef.Condition]) {
            resources[resourceName] = null;
            continue;
        }

        const parent = componentMap[tree.parents[resourceName]];

        const dependsOnDef = resourceDef.DependsOn;
        const dependsOnNames = [];
        if (typeof dependsOnDef === "string") {
            dependsOnNames.push(dependsOnDef);
        } else if (dependsOnDef instanceof Array) {
            dependsOnNames.push(...dependsOnDef);
        }
        const dependsOn = <pulumi.Resource[]>dependsOnNames.map(n => resources[n]).filter(r => r !== null);

        let [_, resourceModule, resourceType] = resourceDef.Type.split("::");
        resources[resourceName] = new CloudFormationResource(`cloudformation:${resourceModule}:${resourceType}`, resourceName, {
            metadata: evaluate(resourceName, resourceDef.Metadata),
            creationPolicy: evaluate(resourceName, resourceDef.CreationPolicy),
            deletionPolicy: evaluate(resourceName, resourceDef.DeletionPolicy),
            updatePolicy: evaluate(resourceName, resourceDef.UpdatePolicy),
            updateReplacePolicy: evaluate(resourceName, resourceDef.UpdateReplacePolicy),
            properties: evaluate(resourceName, resourceDef.Properties),
            attributes: undefined,
        }, { parent, dependsOn });
    }

    // Evaluate and return the stack's outputs.
    for (const k of Object.keys(outputDefs)) {
        outputs[k] = await evaluate(k, outputDefs[k].Value);
    }
    return outputs;
}

// loadJSON loads and evaluates the JSON CloudFormation template at the given path using the specified parameters.
export function loadJSON(path: string, parameters: any): Promise<any> {
    return load(JSON.parse(fs.readFileSync(path, "utf8")), parameters);
}
