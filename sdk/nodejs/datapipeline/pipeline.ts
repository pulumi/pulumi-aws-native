// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DataPipeline::Pipeline
 *
 * @deprecated Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
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
        pulumi.log.warn("Pipeline is deprecated: Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
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

    public readonly activate!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parameterObjects!: pulumi.Output<outputs.datapipeline.PipelineParameterObject[]>;
    public readonly parameterValues!: pulumi.Output<outputs.datapipeline.PipelineParameterValue[] | undefined>;
    public readonly pipelineObjects!: pulumi.Output<outputs.datapipeline.PipelineObject[] | undefined>;
    public readonly pipelineTags!: pulumi.Output<outputs.datapipeline.PipelineTag[] | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Pipeline is deprecated: Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.parameterObjects === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameterObjects'");
            }
            resourceInputs["activate"] = args ? args.activate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameterObjects"] = args ? args.parameterObjects : undefined;
            resourceInputs["parameterValues"] = args ? args.parameterValues : undefined;
            resourceInputs["pipelineObjects"] = args ? args.pipelineObjects : undefined;
            resourceInputs["pipelineTags"] = args ? args.pipelineTags : undefined;
        } else {
            resourceInputs["activate"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameterObjects"] = undefined /*out*/;
            resourceInputs["parameterValues"] = undefined /*out*/;
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
    activate?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parameterObjects: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineParameterObjectArgs>[]>;
    parameterValues?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineParameterValueArgs>[]>;
    pipelineObjects?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineObjectArgs>[]>;
    pipelineTags?: pulumi.Input<pulumi.Input<inputs.datapipeline.PipelineTagArgs>[]>;
}