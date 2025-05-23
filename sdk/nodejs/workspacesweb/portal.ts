// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::WorkSpacesWeb::Portal Resource Type
 */
export class Portal extends pulumi.CustomResource {
    /**
     * Get an existing Portal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Portal {
        return new Portal(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:workspacesweb:Portal';

    /**
     * Returns true if the given object is an instance of Portal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Portal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Portal.__pulumiType;
    }

    /**
     * The additional encryption context of the portal.
     */
    public readonly additionalEncryptionContext!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of authentication integration points used when signing into the web portal. Defaults to `Standard` .
     *
     * `Standard` web portals are authenticated directly through your identity provider (IdP). User and group access to your web portal is controlled through your IdP. You need to include an IdP resource in your template to integrate your IdP with your web portal. Completing the configuration for your IdP requires exchanging WorkSpaces Secure Browser’s SP metadata with your IdP’s IdP metadata. If your IdP requires the SP metadata first before returning the IdP metadata, you should follow these steps:
     *
     * 1. Create and deploy a CloudFormation template with a `Standard` portal with no `IdentityProvider` resource.
     *
     * 2. Retrieve the SP metadata using `Fn:GetAtt` , the WorkSpaces Secure Browser console, or by the calling the `GetPortalServiceProviderMetadata` API.
     *
     * 3. Submit the data to your IdP.
     *
     * 4. Add an `IdentityProvider` resource to your CloudFormation template.
     *
     * `IAM Identity Center` web portals are authenticated through AWS IAM Identity Center . They provide additional features, such as IdP-initiated authentication. Identity sources (including external identity provider integration) and other identity provider information must be configured in IAM Identity Center . User and group assignment must be done through the WorkSpaces Secure Browser console. These cannot be configured in CloudFormation.
     */
    public readonly authenticationType!: pulumi.Output<enums.workspacesweb.PortalAuthenticationType | undefined>;
    /**
     * The ARN of the browser settings that is associated with this web portal.
     */
    public readonly browserSettingsArn!: pulumi.Output<string | undefined>;
    /**
     * The browser that users see when using a streaming session.
     */
    public /*out*/ readonly browserType!: pulumi.Output<enums.workspacesweb.PortalBrowserType>;
    /**
     * The creation date of the web portal.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The customer managed key of the web portal.
     *
     * *Pattern* : `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`
     */
    public readonly customerManagedKey!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the data protection settings.
     */
    public readonly dataProtectionSettingsArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the web portal.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The type and resources of the underlying instance.
     */
    public readonly instanceType!: pulumi.Output<enums.workspacesweb.PortalInstanceType | undefined>;
    /**
     * The ARN of the IP access settings that is associated with the web portal.
     */
    public readonly ipAccessSettingsArn!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of concurrent sessions for the portal.
     */
    public readonly maxConcurrentSessions!: pulumi.Output<number | undefined>;
    /**
     * The ARN of the network settings that is associated with the web portal.
     */
    public readonly networkSettingsArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the web portal.
     */
    public /*out*/ readonly portalArn!: pulumi.Output<string>;
    /**
     * The endpoint URL of the web portal that users access in order to start streaming sessions.
     */
    public /*out*/ readonly portalEndpoint!: pulumi.Output<string>;
    /**
     * The status of the web portal.
     */
    public /*out*/ readonly portalStatus!: pulumi.Output<enums.workspacesweb.PortalStatus>;
    /**
     * The renderer that is used in streaming sessions.
     */
    public /*out*/ readonly rendererType!: pulumi.Output<enums.workspacesweb.PortalRendererType>;
    /**
     * The SAML metadata of the service provider.
     */
    public /*out*/ readonly serviceProviderSamlMetadata!: pulumi.Output<string>;
    /**
     * A message that explains why the web portal is in its current status.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * The tags to add to the web portal. A tag is a key-value pair.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The ARN of the trust store that is associated with the web portal.
     */
    public readonly trustStoreArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the user access logging settings that is associated with the web portal.
     */
    public readonly userAccessLoggingSettingsArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the user settings that is associated with the web portal.
     */
    public readonly userSettingsArn!: pulumi.Output<string | undefined>;

    /**
     * Create a Portal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PortalArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["additionalEncryptionContext"] = args ? args.additionalEncryptionContext : undefined;
            resourceInputs["authenticationType"] = args ? args.authenticationType : undefined;
            resourceInputs["browserSettingsArn"] = args ? args.browserSettingsArn : undefined;
            resourceInputs["customerManagedKey"] = args ? args.customerManagedKey : undefined;
            resourceInputs["dataProtectionSettingsArn"] = args ? args.dataProtectionSettingsArn : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["ipAccessSettingsArn"] = args ? args.ipAccessSettingsArn : undefined;
            resourceInputs["maxConcurrentSessions"] = args ? args.maxConcurrentSessions : undefined;
            resourceInputs["networkSettingsArn"] = args ? args.networkSettingsArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trustStoreArn"] = args ? args.trustStoreArn : undefined;
            resourceInputs["userAccessLoggingSettingsArn"] = args ? args.userAccessLoggingSettingsArn : undefined;
            resourceInputs["userSettingsArn"] = args ? args.userSettingsArn : undefined;
            resourceInputs["browserType"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["portalArn"] = undefined /*out*/;
            resourceInputs["portalEndpoint"] = undefined /*out*/;
            resourceInputs["portalStatus"] = undefined /*out*/;
            resourceInputs["rendererType"] = undefined /*out*/;
            resourceInputs["serviceProviderSamlMetadata"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
        } else {
            resourceInputs["additionalEncryptionContext"] = undefined /*out*/;
            resourceInputs["authenticationType"] = undefined /*out*/;
            resourceInputs["browserSettingsArn"] = undefined /*out*/;
            resourceInputs["browserType"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["customerManagedKey"] = undefined /*out*/;
            resourceInputs["dataProtectionSettingsArn"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["ipAccessSettingsArn"] = undefined /*out*/;
            resourceInputs["maxConcurrentSessions"] = undefined /*out*/;
            resourceInputs["networkSettingsArn"] = undefined /*out*/;
            resourceInputs["portalArn"] = undefined /*out*/;
            resourceInputs["portalEndpoint"] = undefined /*out*/;
            resourceInputs["portalStatus"] = undefined /*out*/;
            resourceInputs["rendererType"] = undefined /*out*/;
            resourceInputs["serviceProviderSamlMetadata"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["trustStoreArn"] = undefined /*out*/;
            resourceInputs["userAccessLoggingSettingsArn"] = undefined /*out*/;
            resourceInputs["userSettingsArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["additionalEncryptionContext.*", "customerManagedKey"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Portal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Portal resource.
 */
export interface PortalArgs {
    /**
     * The additional encryption context of the portal.
     */
    additionalEncryptionContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of authentication integration points used when signing into the web portal. Defaults to `Standard` .
     *
     * `Standard` web portals are authenticated directly through your identity provider (IdP). User and group access to your web portal is controlled through your IdP. You need to include an IdP resource in your template to integrate your IdP with your web portal. Completing the configuration for your IdP requires exchanging WorkSpaces Secure Browser’s SP metadata with your IdP’s IdP metadata. If your IdP requires the SP metadata first before returning the IdP metadata, you should follow these steps:
     *
     * 1. Create and deploy a CloudFormation template with a `Standard` portal with no `IdentityProvider` resource.
     *
     * 2. Retrieve the SP metadata using `Fn:GetAtt` , the WorkSpaces Secure Browser console, or by the calling the `GetPortalServiceProviderMetadata` API.
     *
     * 3. Submit the data to your IdP.
     *
     * 4. Add an `IdentityProvider` resource to your CloudFormation template.
     *
     * `IAM Identity Center` web portals are authenticated through AWS IAM Identity Center . They provide additional features, such as IdP-initiated authentication. Identity sources (including external identity provider integration) and other identity provider information must be configured in IAM Identity Center . User and group assignment must be done through the WorkSpaces Secure Browser console. These cannot be configured in CloudFormation.
     */
    authenticationType?: pulumi.Input<enums.workspacesweb.PortalAuthenticationType>;
    /**
     * The ARN of the browser settings that is associated with this web portal.
     */
    browserSettingsArn?: pulumi.Input<string>;
    /**
     * The customer managed key of the web portal.
     *
     * *Pattern* : `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`
     */
    customerManagedKey?: pulumi.Input<string>;
    /**
     * The ARN of the data protection settings.
     */
    dataProtectionSettingsArn?: pulumi.Input<string>;
    /**
     * The name of the web portal.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The type and resources of the underlying instance.
     */
    instanceType?: pulumi.Input<enums.workspacesweb.PortalInstanceType>;
    /**
     * The ARN of the IP access settings that is associated with the web portal.
     */
    ipAccessSettingsArn?: pulumi.Input<string>;
    /**
     * The maximum number of concurrent sessions for the portal.
     */
    maxConcurrentSessions?: pulumi.Input<number>;
    /**
     * The ARN of the network settings that is associated with the web portal.
     */
    networkSettingsArn?: pulumi.Input<string>;
    /**
     * The tags to add to the web portal. A tag is a key-value pair.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The ARN of the trust store that is associated with the web portal.
     */
    trustStoreArn?: pulumi.Input<string>;
    /**
     * The ARN of the user access logging settings that is associated with the web portal.
     */
    userAccessLoggingSettingsArn?: pulumi.Input<string>;
    /**
     * The ARN of the user settings that is associated with the web portal.
     */
    userSettingsArn?: pulumi.Input<string>;
}
