// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datapipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
     */
    public readonly activate!: pulumi.Output<boolean | undefined>;
    /**
     * A description of the pipeline.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the pipeline.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parameter objects used with the pipeline.
     */
    public readonly parameterObjects!: pulumi.Output<outputs.datapipeline.PipelineParameterObject[] | undefined>;
    /**
     * The parameter values used with the pipeline.
     */
    public readonly parameterValues!: pulumi.Output<outputs.datapipeline.PipelineParameterValue[] | undefined>;
    public /*out*/ readonly pipelineId!: pulumi.Output<string>;
    /**
     * The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
     */
    public readonly pipelineObjects!: pulumi.Output<outputs.datapipeline.PipelineObject[] | undefined>;
    /**
     * A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
     */
    public readonly pipelineTags!: pulumi.Output<outputs.datapipeline.PipelineTag[] | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PipelineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["activate"] = args ? args.activate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameterObjects"] = args ? args.parameterObjects : undefined;
            resourceInputs["parameterValues"] = args ? args.parameterValues : undefined;
            resourceInputs["pipelineObjects"] = args ? args.pipelineObjects : undefined;
            resourceInputs["pipelineTags"] = args ? args.pipelineTags : undefined;
            resourceInputs["pipelineId"] = undefined /*out*/;
        } else {
            resourceInputs["activate"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameterObjects"] = undefined /*out*/;
            resourceInputs["parameterValues"] = undefined /*out*/;
            resourceInputs["pipelineId"] = undefined /*out*/;
            resourceInputs["pipelineObjects"] = undefined /*out*/;
            resourceInputs["pipelineTags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
     */
    activate?: pulumi.Input<boolean>;
    /**
     * A description of the pipeline.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the pipeline.
     */
    name?: pulumi.Input<string>;
    /**
     * The parameter objects used with the pipeline.
     */
    parameterObjects?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineParameterObjectArgs>[]>;
    /**
     * The parameter values used with the pipeline.
     */
    parameterValues?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineParameterValueArgs>[]>;
    /**
     * The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
     */
    pipelineObjects?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineObjectArgs>[]>;
    /**
     * A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
     */
    pipelineTags?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineTagArgs>[]>;
}