// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ConfigurationSetArgs } from "./configurationSet";
export type ConfigurationSet = import("./configurationSet").ConfigurationSet;
export const ConfigurationSet: typeof import("./configurationSet").ConfigurationSet = null as any;
utilities.lazyLoad(exports, ["ConfigurationSet"], () => require("./configurationSet"));

export { ConfigurationSetEventDestinationArgs } from "./configurationSetEventDestination";
export type ConfigurationSetEventDestination = import("./configurationSetEventDestination").ConfigurationSetEventDestination;
export const ConfigurationSetEventDestination: typeof import("./configurationSetEventDestination").ConfigurationSetEventDestination = null as any;
utilities.lazyLoad(exports, ["ConfigurationSetEventDestination"], () => require("./configurationSetEventDestination"));

export { ContactListArgs } from "./contactList";
export type ContactList = import("./contactList").ContactList;
export const ContactList: typeof import("./contactList").ContactList = null as any;
utilities.lazyLoad(exports, ["ContactList"], () => require("./contactList"));

export { DedicatedIpPoolArgs } from "./dedicatedIpPool";
export type DedicatedIpPool = import("./dedicatedIpPool").DedicatedIpPool;
export const DedicatedIpPool: typeof import("./dedicatedIpPool").DedicatedIpPool = null as any;
utilities.lazyLoad(exports, ["DedicatedIpPool"], () => require("./dedicatedIpPool"));

export { EmailIdentityArgs } from "./emailIdentity";
export type EmailIdentity = import("./emailIdentity").EmailIdentity;
export const EmailIdentity: typeof import("./emailIdentity").EmailIdentity = null as any;
utilities.lazyLoad(exports, ["EmailIdentity"], () => require("./emailIdentity"));

export { GetConfigurationSetArgs, GetConfigurationSetResult, GetConfigurationSetOutputArgs } from "./getConfigurationSet";
export const getConfigurationSet: typeof import("./getConfigurationSet").getConfigurationSet = null as any;
export const getConfigurationSetOutput: typeof import("./getConfigurationSet").getConfigurationSetOutput = null as any;
utilities.lazyLoad(exports, ["getConfigurationSet","getConfigurationSetOutput"], () => require("./getConfigurationSet"));

export { GetConfigurationSetEventDestinationArgs, GetConfigurationSetEventDestinationResult, GetConfigurationSetEventDestinationOutputArgs } from "./getConfigurationSetEventDestination";
export const getConfigurationSetEventDestination: typeof import("./getConfigurationSetEventDestination").getConfigurationSetEventDestination = null as any;
export const getConfigurationSetEventDestinationOutput: typeof import("./getConfigurationSetEventDestination").getConfigurationSetEventDestinationOutput = null as any;
utilities.lazyLoad(exports, ["getConfigurationSetEventDestination","getConfigurationSetEventDestinationOutput"], () => require("./getConfigurationSetEventDestination"));

export { GetContactListArgs, GetContactListResult, GetContactListOutputArgs } from "./getContactList";
export const getContactList: typeof import("./getContactList").getContactList = null as any;
export const getContactListOutput: typeof import("./getContactList").getContactListOutput = null as any;
utilities.lazyLoad(exports, ["getContactList","getContactListOutput"], () => require("./getContactList"));

export { GetDedicatedIpPoolArgs, GetDedicatedIpPoolResult, GetDedicatedIpPoolOutputArgs } from "./getDedicatedIpPool";
export const getDedicatedIpPool: typeof import("./getDedicatedIpPool").getDedicatedIpPool = null as any;
export const getDedicatedIpPoolOutput: typeof import("./getDedicatedIpPool").getDedicatedIpPoolOutput = null as any;
utilities.lazyLoad(exports, ["getDedicatedIpPool","getDedicatedIpPoolOutput"], () => require("./getDedicatedIpPool"));

export { GetEmailIdentityArgs, GetEmailIdentityResult, GetEmailIdentityOutputArgs } from "./getEmailIdentity";
export const getEmailIdentity: typeof import("./getEmailIdentity").getEmailIdentity = null as any;
export const getEmailIdentityOutput: typeof import("./getEmailIdentity").getEmailIdentityOutput = null as any;
utilities.lazyLoad(exports, ["getEmailIdentity","getEmailIdentityOutput"], () => require("./getEmailIdentity"));

export { GetMailManagerAddonInstanceArgs, GetMailManagerAddonInstanceResult, GetMailManagerAddonInstanceOutputArgs } from "./getMailManagerAddonInstance";
export const getMailManagerAddonInstance: typeof import("./getMailManagerAddonInstance").getMailManagerAddonInstance = null as any;
export const getMailManagerAddonInstanceOutput: typeof import("./getMailManagerAddonInstance").getMailManagerAddonInstanceOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerAddonInstance","getMailManagerAddonInstanceOutput"], () => require("./getMailManagerAddonInstance"));

export { GetMailManagerAddonSubscriptionArgs, GetMailManagerAddonSubscriptionResult, GetMailManagerAddonSubscriptionOutputArgs } from "./getMailManagerAddonSubscription";
export const getMailManagerAddonSubscription: typeof import("./getMailManagerAddonSubscription").getMailManagerAddonSubscription = null as any;
export const getMailManagerAddonSubscriptionOutput: typeof import("./getMailManagerAddonSubscription").getMailManagerAddonSubscriptionOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerAddonSubscription","getMailManagerAddonSubscriptionOutput"], () => require("./getMailManagerAddonSubscription"));

export { GetMailManagerAddressListArgs, GetMailManagerAddressListResult, GetMailManagerAddressListOutputArgs } from "./getMailManagerAddressList";
export const getMailManagerAddressList: typeof import("./getMailManagerAddressList").getMailManagerAddressList = null as any;
export const getMailManagerAddressListOutput: typeof import("./getMailManagerAddressList").getMailManagerAddressListOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerAddressList","getMailManagerAddressListOutput"], () => require("./getMailManagerAddressList"));

export { GetMailManagerArchiveArgs, GetMailManagerArchiveResult, GetMailManagerArchiveOutputArgs } from "./getMailManagerArchive";
export const getMailManagerArchive: typeof import("./getMailManagerArchive").getMailManagerArchive = null as any;
export const getMailManagerArchiveOutput: typeof import("./getMailManagerArchive").getMailManagerArchiveOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerArchive","getMailManagerArchiveOutput"], () => require("./getMailManagerArchive"));

export { GetMailManagerIngressPointArgs, GetMailManagerIngressPointResult, GetMailManagerIngressPointOutputArgs } from "./getMailManagerIngressPoint";
export const getMailManagerIngressPoint: typeof import("./getMailManagerIngressPoint").getMailManagerIngressPoint = null as any;
export const getMailManagerIngressPointOutput: typeof import("./getMailManagerIngressPoint").getMailManagerIngressPointOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerIngressPoint","getMailManagerIngressPointOutput"], () => require("./getMailManagerIngressPoint"));

export { GetMailManagerRelayArgs, GetMailManagerRelayResult, GetMailManagerRelayOutputArgs } from "./getMailManagerRelay";
export const getMailManagerRelay: typeof import("./getMailManagerRelay").getMailManagerRelay = null as any;
export const getMailManagerRelayOutput: typeof import("./getMailManagerRelay").getMailManagerRelayOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerRelay","getMailManagerRelayOutput"], () => require("./getMailManagerRelay"));

export { GetMailManagerRuleSetArgs, GetMailManagerRuleSetResult, GetMailManagerRuleSetOutputArgs } from "./getMailManagerRuleSet";
export const getMailManagerRuleSet: typeof import("./getMailManagerRuleSet").getMailManagerRuleSet = null as any;
export const getMailManagerRuleSetOutput: typeof import("./getMailManagerRuleSet").getMailManagerRuleSetOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerRuleSet","getMailManagerRuleSetOutput"], () => require("./getMailManagerRuleSet"));

export { GetMailManagerTrafficPolicyArgs, GetMailManagerTrafficPolicyResult, GetMailManagerTrafficPolicyOutputArgs } from "./getMailManagerTrafficPolicy";
export const getMailManagerTrafficPolicy: typeof import("./getMailManagerTrafficPolicy").getMailManagerTrafficPolicy = null as any;
export const getMailManagerTrafficPolicyOutput: typeof import("./getMailManagerTrafficPolicy").getMailManagerTrafficPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getMailManagerTrafficPolicy","getMailManagerTrafficPolicyOutput"], () => require("./getMailManagerTrafficPolicy"));

export { GetTemplateArgs, GetTemplateResult, GetTemplateOutputArgs } from "./getTemplate";
export const getTemplate: typeof import("./getTemplate").getTemplate = null as any;
export const getTemplateOutput: typeof import("./getTemplate").getTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getTemplate","getTemplateOutput"], () => require("./getTemplate"));

export { GetVdmAttributesArgs, GetVdmAttributesResult, GetVdmAttributesOutputArgs } from "./getVdmAttributes";
export const getVdmAttributes: typeof import("./getVdmAttributes").getVdmAttributes = null as any;
export const getVdmAttributesOutput: typeof import("./getVdmAttributes").getVdmAttributesOutput = null as any;
utilities.lazyLoad(exports, ["getVdmAttributes","getVdmAttributesOutput"], () => require("./getVdmAttributes"));

export { MailManagerAddonInstanceArgs } from "./mailManagerAddonInstance";
export type MailManagerAddonInstance = import("./mailManagerAddonInstance").MailManagerAddonInstance;
export const MailManagerAddonInstance: typeof import("./mailManagerAddonInstance").MailManagerAddonInstance = null as any;
utilities.lazyLoad(exports, ["MailManagerAddonInstance"], () => require("./mailManagerAddonInstance"));

export { MailManagerAddonSubscriptionArgs } from "./mailManagerAddonSubscription";
export type MailManagerAddonSubscription = import("./mailManagerAddonSubscription").MailManagerAddonSubscription;
export const MailManagerAddonSubscription: typeof import("./mailManagerAddonSubscription").MailManagerAddonSubscription = null as any;
utilities.lazyLoad(exports, ["MailManagerAddonSubscription"], () => require("./mailManagerAddonSubscription"));

export { MailManagerAddressListArgs } from "./mailManagerAddressList";
export type MailManagerAddressList = import("./mailManagerAddressList").MailManagerAddressList;
export const MailManagerAddressList: typeof import("./mailManagerAddressList").MailManagerAddressList = null as any;
utilities.lazyLoad(exports, ["MailManagerAddressList"], () => require("./mailManagerAddressList"));

export { MailManagerArchiveArgs } from "./mailManagerArchive";
export type MailManagerArchive = import("./mailManagerArchive").MailManagerArchive;
export const MailManagerArchive: typeof import("./mailManagerArchive").MailManagerArchive = null as any;
utilities.lazyLoad(exports, ["MailManagerArchive"], () => require("./mailManagerArchive"));

export { MailManagerIngressPointArgs } from "./mailManagerIngressPoint";
export type MailManagerIngressPoint = import("./mailManagerIngressPoint").MailManagerIngressPoint;
export const MailManagerIngressPoint: typeof import("./mailManagerIngressPoint").MailManagerIngressPoint = null as any;
utilities.lazyLoad(exports, ["MailManagerIngressPoint"], () => require("./mailManagerIngressPoint"));

export { MailManagerRelayArgs } from "./mailManagerRelay";
export type MailManagerRelay = import("./mailManagerRelay").MailManagerRelay;
export const MailManagerRelay: typeof import("./mailManagerRelay").MailManagerRelay = null as any;
utilities.lazyLoad(exports, ["MailManagerRelay"], () => require("./mailManagerRelay"));

export { MailManagerRuleSetArgs } from "./mailManagerRuleSet";
export type MailManagerRuleSet = import("./mailManagerRuleSet").MailManagerRuleSet;
export const MailManagerRuleSet: typeof import("./mailManagerRuleSet").MailManagerRuleSet = null as any;
utilities.lazyLoad(exports, ["MailManagerRuleSet"], () => require("./mailManagerRuleSet"));

export { MailManagerTrafficPolicyArgs } from "./mailManagerTrafficPolicy";
export type MailManagerTrafficPolicy = import("./mailManagerTrafficPolicy").MailManagerTrafficPolicy;
export const MailManagerTrafficPolicy: typeof import("./mailManagerTrafficPolicy").MailManagerTrafficPolicy = null as any;
utilities.lazyLoad(exports, ["MailManagerTrafficPolicy"], () => require("./mailManagerTrafficPolicy"));

export { TemplateArgs } from "./template";
export type Template = import("./template").Template;
export const Template: typeof import("./template").Template = null as any;
utilities.lazyLoad(exports, ["Template"], () => require("./template"));

export { VdmAttributesArgs } from "./vdmAttributes";
export type VdmAttributes = import("./vdmAttributes").VdmAttributes;
export const VdmAttributes: typeof import("./vdmAttributes").VdmAttributes = null as any;
utilities.lazyLoad(exports, ["VdmAttributes"], () => require("./vdmAttributes"));


// Export enums:
export * from "../types/enums/ses";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ses:ConfigurationSet":
                return new ConfigurationSet(name, <any>undefined, { urn })
            case "aws-native:ses:ConfigurationSetEventDestination":
                return new ConfigurationSetEventDestination(name, <any>undefined, { urn })
            case "aws-native:ses:ContactList":
                return new ContactList(name, <any>undefined, { urn })
            case "aws-native:ses:DedicatedIpPool":
                return new DedicatedIpPool(name, <any>undefined, { urn })
            case "aws-native:ses:EmailIdentity":
                return new EmailIdentity(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerAddonInstance":
                return new MailManagerAddonInstance(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerAddonSubscription":
                return new MailManagerAddonSubscription(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerAddressList":
                return new MailManagerAddressList(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerArchive":
                return new MailManagerArchive(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerIngressPoint":
                return new MailManagerIngressPoint(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerRelay":
                return new MailManagerRelay(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerRuleSet":
                return new MailManagerRuleSet(name, <any>undefined, { urn })
            case "aws-native:ses:MailManagerTrafficPolicy":
                return new MailManagerTrafficPolicy(name, <any>undefined, { urn })
            case "aws-native:ses:Template":
                return new Template(name, <any>undefined, { urn })
            case "aws-native:ses:VdmAttributes":
                return new VdmAttributes(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ses", _module)
