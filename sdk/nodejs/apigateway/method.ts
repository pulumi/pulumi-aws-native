// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Method
 */
export class Method extends pulumi.CustomResource {
    /**
     * Get an existing Method resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Method {
        return new Method(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:Method';

    /**
     * Returns true if the given object is an instance of Method.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Method {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Method.__pulumiType;
    }

    /**
     * Indicates whether the method requires clients to submit a valid API key.
     */
    public readonly apiKeyRequired!: pulumi.Output<boolean | undefined>;
    /**
     * A list of authorization scopes configured on the method.
     */
    public readonly authorizationScopes!: pulumi.Output<string[] | undefined>;
    /**
     * The method's authorization type.
     */
    public readonly authorizationType!: pulumi.Output<enums.apigateway.MethodAuthorizationType | undefined>;
    /**
     * The identifier of the authorizer to use on this method.
     */
    public readonly authorizerId!: pulumi.Output<string | undefined>;
    /**
     * The backend system that the method calls when it receives a request.
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * The backend system that the method calls when it receives a request.
     */
    public readonly integration!: pulumi.Output<outputs.apigateway.MethodIntegration | undefined>;
    /**
     * The responses that can be sent to the client who calls the method.
     */
    public readonly methodResponses!: pulumi.Output<outputs.apigateway.MethodResponse[] | undefined>;
    /**
     * A friendly operation name for the method.
     */
    public readonly operationName!: pulumi.Output<string | undefined>;
    /**
     * The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
     */
    public readonly requestModels!: pulumi.Output<any | undefined>;
    /**
     * The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
     */
    public readonly requestParameters!: pulumi.Output<any | undefined>;
    /**
     * The ID of the associated request validator.
     */
    public readonly requestValidatorId!: pulumi.Output<string | undefined>;
    /**
     * The ID of an API Gateway resource.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The ID of the RestApi resource in which API Gateway creates the method.
     */
    public readonly restApiId!: pulumi.Output<string>;

    /**
     * Create a Method resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MethodArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.httpMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            resourceInputs["apiKeyRequired"] = args ? args.apiKeyRequired : undefined;
            resourceInputs["authorizationScopes"] = args ? args.authorizationScopes : undefined;
            resourceInputs["authorizationType"] = args ? args.authorizationType : undefined;
            resourceInputs["authorizerId"] = args ? args.authorizerId : undefined;
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["integration"] = args ? args.integration : undefined;
            resourceInputs["methodResponses"] = args ? args.methodResponses : undefined;
            resourceInputs["operationName"] = args ? args.operationName : undefined;
            resourceInputs["requestModels"] = args ? args.requestModels : undefined;
            resourceInputs["requestParameters"] = args ? args.requestParameters : undefined;
            resourceInputs["requestValidatorId"] = args ? args.requestValidatorId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["restApiId"] = args ? args.restApiId : undefined;
        } else {
            resourceInputs["apiKeyRequired"] = undefined /*out*/;
            resourceInputs["authorizationScopes"] = undefined /*out*/;
            resourceInputs["authorizationType"] = undefined /*out*/;
            resourceInputs["authorizerId"] = undefined /*out*/;
            resourceInputs["httpMethod"] = undefined /*out*/;
            resourceInputs["integration"] = undefined /*out*/;
            resourceInputs["methodResponses"] = undefined /*out*/;
            resourceInputs["operationName"] = undefined /*out*/;
            resourceInputs["requestModels"] = undefined /*out*/;
            resourceInputs["requestParameters"] = undefined /*out*/;
            resourceInputs["requestValidatorId"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["restApiId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Method.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Method resource.
 */
export interface MethodArgs {
    /**
     * Indicates whether the method requires clients to submit a valid API key.
     */
    apiKeyRequired?: pulumi.Input<boolean>;
    /**
     * A list of authorization scopes configured on the method.
     */
    authorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The method's authorization type.
     */
    authorizationType?: pulumi.Input<enums.apigateway.MethodAuthorizationType>;
    /**
     * The identifier of the authorizer to use on this method.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * The backend system that the method calls when it receives a request.
     */
    httpMethod: pulumi.Input<string>;
    /**
     * The backend system that the method calls when it receives a request.
     */
    integration?: pulumi.Input<inputs.apigateway.MethodIntegrationArgs>;
    /**
     * The responses that can be sent to the client who calls the method.
     */
    methodResponses?: pulumi.Input<pulumi.Input<inputs.apigateway.MethodResponseArgs>[]>;
    /**
     * A friendly operation name for the method.
     */
    operationName?: pulumi.Input<string>;
    /**
     * The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
     */
    requestModels?: any;
    /**
     * The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
     */
    requestParameters?: any;
    /**
     * The ID of the associated request validator.
     */
    requestValidatorId?: pulumi.Input<string>;
    /**
     * The ID of an API Gateway resource.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The ID of the RestApi resource in which API Gateway creates the method.
     */
    restApiId: pulumi.Input<string>;
}