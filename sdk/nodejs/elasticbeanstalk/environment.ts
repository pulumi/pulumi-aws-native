// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticBeanstalk::Environment
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticbeanstalk:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * The name of the application that is associated with this environment.
     */
    public readonly applicationName!: pulumi.Output<string>;
    /**
     * If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
     */
    public readonly cnamePrefix!: pulumi.Output<string | undefined>;
    /**
     * Your description for this environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * For load-balanced, autoscaling environments, the URL to the load balancer. For single-instance environments, the IP address of the instance.
     *
     * Example load balancer URL:
     *
     * Example instance IP address:
     *
     * `192.0.2.0`
     */
    public /*out*/ readonly endpointUrl!: pulumi.Output<string>;
    /**
     * A unique name for the environment.
     */
    public readonly environmentName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
     */
    public readonly operationsRole!: pulumi.Output<string | undefined>;
    /**
     * Key-value pairs defining configuration options for this environment, such as the instance type.
     */
    public readonly optionSettings!: pulumi.Output<outputs.elasticbeanstalk.EnvironmentOptionSetting[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the custom platform to use with the environment.
     */
    public readonly platformArn!: pulumi.Output<string | undefined>;
    /**
     * The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
     */
    public readonly solutionStackName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the tags applied to resources in the environment.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The name of the Elastic Beanstalk configuration template to use with the environment.
     */
    public readonly templateName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
     */
    public readonly tier!: pulumi.Output<outputs.elasticbeanstalk.EnvironmentTier | undefined>;
    /**
     * The name of the application version to deploy.
     */
    public readonly versionLabel!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationName'");
            }
            resourceInputs["applicationName"] = args ? args.applicationName : undefined;
            resourceInputs["cnamePrefix"] = args ? args.cnamePrefix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["operationsRole"] = args ? args.operationsRole : undefined;
            resourceInputs["optionSettings"] = args ? args.optionSettings : undefined;
            resourceInputs["platformArn"] = args ? args.platformArn : undefined;
            resourceInputs["solutionStackName"] = args ? args.solutionStackName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["versionLabel"] = args ? args.versionLabel : undefined;
            resourceInputs["endpointUrl"] = undefined /*out*/;
        } else {
            resourceInputs["applicationName"] = undefined /*out*/;
            resourceInputs["cnamePrefix"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endpointUrl"] = undefined /*out*/;
            resourceInputs["environmentName"] = undefined /*out*/;
            resourceInputs["operationsRole"] = undefined /*out*/;
            resourceInputs["optionSettings"] = undefined /*out*/;
            resourceInputs["platformArn"] = undefined /*out*/;
            resourceInputs["solutionStackName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["templateName"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
            resourceInputs["versionLabel"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationName", "cnamePrefix", "environmentName", "solutionStackName", "tier.name", "tier.type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * The name of the application that is associated with this environment.
     */
    applicationName: pulumi.Input<string>;
    /**
     * If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
     */
    cnamePrefix?: pulumi.Input<string>;
    /**
     * Your description for this environment.
     */
    description?: pulumi.Input<string>;
    /**
     * A unique name for the environment.
     */
    environmentName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
     */
    operationsRole?: pulumi.Input<string>;
    /**
     * Key-value pairs defining configuration options for this environment, such as the instance type.
     */
    optionSettings?: pulumi.Input<pulumi.Input<inputs.elasticbeanstalk.EnvironmentOptionSettingArgs>[]>;
    /**
     * The Amazon Resource Name (ARN) of the custom platform to use with the environment.
     */
    platformArn?: pulumi.Input<string>;
    /**
     * The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
     */
    solutionStackName?: pulumi.Input<string>;
    /**
     * Specifies the tags applied to resources in the environment.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The name of the Elastic Beanstalk configuration template to use with the environment.
     */
    templateName?: pulumi.Input<string>;
    /**
     * Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
     */
    tier?: pulumi.Input<inputs.elasticbeanstalk.EnvironmentTierArgs>;
    /**
     * The name of the application version to deploy.
     */
    versionLabel?: pulumi.Input<string>;
}
