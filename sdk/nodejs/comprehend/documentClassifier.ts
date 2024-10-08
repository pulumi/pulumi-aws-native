// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Document Classifier enables training document classifier models.
 */
export class DocumentClassifier extends pulumi.CustomResource {
    /**
     * Get an existing DocumentClassifier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DocumentClassifier {
        return new DocumentClassifier(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:comprehend:DocumentClassifier';

    /**
     * Returns true if the given object is an instance of DocumentClassifier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentClassifier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentClassifier.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the document classifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.
     */
    public readonly dataAccessRoleArn!: pulumi.Output<string>;
    /**
     * The name of the document classifier.
     */
    public readonly documentClassifierName!: pulumi.Output<string>;
    /**
     * Specifies the format and location of the input data for the job.
     */
    public readonly inputDataConfig!: pulumi.Output<outputs.comprehend.DocumentClassifierInputDataConfig>;
    /**
     * The language of the input documents. You can specify any of the languages supported by Amazon Comprehend. All documents must be in the same language.
     */
    public readonly languageCode!: pulumi.Output<enums.comprehend.DocumentClassifierLanguageCode>;
    /**
     * Indicates the mode in which the classifier will be trained. The classifier can be trained in multi-class (single-label) mode or multi-label mode. Multi-class mode identifies a single class label for each document and multi-label mode identifies one or more class labels for each document. Multiple labels for an individual document are separated by a delimiter. The default delimiter between labels is a pipe (|).
     */
    public readonly mode!: pulumi.Output<enums.comprehend.DocumentClassifierMode | undefined>;
    /**
     * ID for the AWS KMS key that Amazon Comprehend uses to encrypt trained custom models. The ModelKmsKeyId can be either of the following formats:
     *
     * - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
     * - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`
     */
    public readonly modelKmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The resource-based policy to attach to your custom document classifier model. You can use this policy to allow another AWS account to import your custom model.
     *
     * Provide your policy as a JSON body that you enter as a UTF-8 encoded string without line breaks. To provide valid JSON, enclose the attribute names and values in double quotes. If the JSON body is also enclosed in double quotes, then you must escape the double quotes that are inside the policy:
     *
     * `"{\"attribute\": \"value\", \"attribute\": [\"value\"]}"`
     *
     * To avoid escaping quotes, you can use single quotes to enclose the policy and double quotes to enclose the JSON names and values:
     *
     * `'{"attribute": "value", "attribute": ["value"]}'`
     */
    public readonly modelPolicy!: pulumi.Output<string | undefined>;
    /**
     * Provides output results configuration parameters for custom classifier jobs.
     */
    public readonly outputDataConfig!: pulumi.Output<outputs.comprehend.DocumentClassifierOutputDataConfig | undefined>;
    /**
     * Tags to associate with the document classifier. A tag is a key-value pair that adds as a metadata to a resource used by Amazon Comprehend. For example, a tag with "Sales" as the key might be added to a resource to indicate its use by the sales department.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The version name given to the newly created classifier. Version names can have a maximum of 256 characters. Alphanumeric characters, hyphens (-) and underscores (_) are allowed. The version name must be unique among all models with the same classifier name in the AWS account / AWS Region .
     */
    public readonly versionName!: pulumi.Output<string | undefined>;
    /**
     * ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job. The VolumeKmsKeyId can be either of the following formats:
     *
     * - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
     * - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`
     */
    public readonly volumeKmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Configuration parameters for a private Virtual Private Cloud (VPC) containing the resources you are using for your custom classifier. For more information, see [Amazon VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html) .
     */
    public readonly vpcConfig!: pulumi.Output<outputs.comprehend.DocumentClassifierVpcConfig | undefined>;

    /**
     * Create a DocumentClassifier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentClassifierArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataAccessRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataAccessRoleArn'");
            }
            if ((!args || args.inputDataConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputDataConfig'");
            }
            if ((!args || args.languageCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'languageCode'");
            }
            resourceInputs["dataAccessRoleArn"] = args ? args.dataAccessRoleArn : undefined;
            resourceInputs["documentClassifierName"] = args ? args.documentClassifierName : undefined;
            resourceInputs["inputDataConfig"] = args ? args.inputDataConfig : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["modelKmsKeyId"] = args ? args.modelKmsKeyId : undefined;
            resourceInputs["modelPolicy"] = args ? args.modelPolicy : undefined;
            resourceInputs["outputDataConfig"] = args ? args.outputDataConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versionName"] = args ? args.versionName : undefined;
            resourceInputs["volumeKmsKeyId"] = args ? args.volumeKmsKeyId : undefined;
            resourceInputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dataAccessRoleArn"] = undefined /*out*/;
            resourceInputs["documentClassifierName"] = undefined /*out*/;
            resourceInputs["inputDataConfig"] = undefined /*out*/;
            resourceInputs["languageCode"] = undefined /*out*/;
            resourceInputs["mode"] = undefined /*out*/;
            resourceInputs["modelKmsKeyId"] = undefined /*out*/;
            resourceInputs["modelPolicy"] = undefined /*out*/;
            resourceInputs["outputDataConfig"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["versionName"] = undefined /*out*/;
            resourceInputs["volumeKmsKeyId"] = undefined /*out*/;
            resourceInputs["vpcConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["dataAccessRoleArn", "documentClassifierName", "inputDataConfig", "languageCode", "mode", "modelKmsKeyId", "outputDataConfig", "versionName", "volumeKmsKeyId", "vpcConfig"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DocumentClassifier.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DocumentClassifier resource.
 */
export interface DocumentClassifierArgs {
    /**
     * The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.
     */
    dataAccessRoleArn: pulumi.Input<string>;
    /**
     * The name of the document classifier.
     */
    documentClassifierName?: pulumi.Input<string>;
    /**
     * Specifies the format and location of the input data for the job.
     */
    inputDataConfig: pulumi.Input<inputs.comprehend.DocumentClassifierInputDataConfigArgs>;
    /**
     * The language of the input documents. You can specify any of the languages supported by Amazon Comprehend. All documents must be in the same language.
     */
    languageCode: pulumi.Input<enums.comprehend.DocumentClassifierLanguageCode>;
    /**
     * Indicates the mode in which the classifier will be trained. The classifier can be trained in multi-class (single-label) mode or multi-label mode. Multi-class mode identifies a single class label for each document and multi-label mode identifies one or more class labels for each document. Multiple labels for an individual document are separated by a delimiter. The default delimiter between labels is a pipe (|).
     */
    mode?: pulumi.Input<enums.comprehend.DocumentClassifierMode>;
    /**
     * ID for the AWS KMS key that Amazon Comprehend uses to encrypt trained custom models. The ModelKmsKeyId can be either of the following formats:
     *
     * - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
     * - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`
     */
    modelKmsKeyId?: pulumi.Input<string>;
    /**
     * The resource-based policy to attach to your custom document classifier model. You can use this policy to allow another AWS account to import your custom model.
     *
     * Provide your policy as a JSON body that you enter as a UTF-8 encoded string without line breaks. To provide valid JSON, enclose the attribute names and values in double quotes. If the JSON body is also enclosed in double quotes, then you must escape the double quotes that are inside the policy:
     *
     * `"{\"attribute\": \"value\", \"attribute\": [\"value\"]}"`
     *
     * To avoid escaping quotes, you can use single quotes to enclose the policy and double quotes to enclose the JSON names and values:
     *
     * `'{"attribute": "value", "attribute": ["value"]}'`
     */
    modelPolicy?: pulumi.Input<string>;
    /**
     * Provides output results configuration parameters for custom classifier jobs.
     */
    outputDataConfig?: pulumi.Input<inputs.comprehend.DocumentClassifierOutputDataConfigArgs>;
    /**
     * Tags to associate with the document classifier. A tag is a key-value pair that adds as a metadata to a resource used by Amazon Comprehend. For example, a tag with "Sales" as the key might be added to a resource to indicate its use by the sales department.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The version name given to the newly created classifier. Version names can have a maximum of 256 characters. Alphanumeric characters, hyphens (-) and underscores (_) are allowed. The version name must be unique among all models with the same classifier name in the AWS account / AWS Region .
     */
    versionName?: pulumi.Input<string>;
    /**
     * ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job. The VolumeKmsKeyId can be either of the following formats:
     *
     * - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
     * - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`
     */
    volumeKmsKeyId?: pulumi.Input<string>;
    /**
     * Configuration parameters for a private Virtual Private Cloud (VPC) containing the resources you are using for your custom classifier. For more information, see [Amazon VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html) .
     */
    vpcConfig?: pulumi.Input<inputs.comprehend.DocumentClassifierVpcConfigArgs>;
}
