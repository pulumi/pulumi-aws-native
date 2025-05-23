// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    /// <summary>
    /// Resource Type definition for AWS::Cognito::UserPoolUser
    /// </summary>
    [AwsNativeResourceType("aws-native:cognito:UserPoolUser")]
    public partial class UserPoolUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.
        /// 
        /// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you use the AdminCreateUser API action, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `ClientMetadata` attribute, which provides the data that you assigned to the ClientMetadata parameter in your AdminCreateUser request. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
        /// 
        /// For more information, see [Using Lambda triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
        /// 
        /// &gt; When you use the `ClientMetadata` parameter, note that Amazon Cognito won't do the following:
        /// &gt; 
        /// &gt; - Store the `ClientMetadata` value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration doesn't include triggers, the `ClientMetadata` parameter serves no purpose.
        /// &gt; - Validate the `ClientMetadata` value.
        /// &gt; - Encrypt the `ClientMetadata` value. Don't send sensitive information in this parameter.
        /// </summary>
        [Output("clientMetadata")]
        public Output<ImmutableDictionary<string, string>?> ClientMetadata { get; private set; } = null!;

        /// <summary>
        /// Specify `EMAIL` if email will be used to send the welcome message. Specify `SMS` if the phone number will be used. The default value is `SMS` . You can specify more than one value.
        /// </summary>
        [Output("desiredDeliveryMediums")]
        public Output<ImmutableArray<string>> DesiredDeliveryMediums { get; private set; } = null!;

        /// <summary>
        /// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` . Otherwise, it is ignored.
        /// 
        /// If this parameter is set to `True` and the phone number or email address specified in the `UserAttributes` parameter already exists as an alias with a different user, this request migrates the alias from the previous user to the newly-created user. The previous user will no longer be able to log in using that alias.
        /// 
        /// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
        /// </summary>
        [Output("forceAliasCreation")]
        public Output<bool?> ForceAliasCreation { get; private set; } = null!;

        /// <summary>
        /// Set to `RESEND` to resend the invitation message to a user that already exists, and to reset the temporary-password duration with a new temporary password. Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
        /// </summary>
        [Output("messageAction")]
        public Output<string?> MessageAction { get; private set; } = null!;

        /// <summary>
        /// An array of name-value pairs that contain user attributes and attribute values to be set for the user to be created. You can create a user without specifying any attributes other than `Username` . However, any attributes that you specify as required (when creating a user pool or in the *Attributes* tab of the console) either you should supply (in your call to `AdminCreateUser` ) or the user should supply (when they sign up in response to your welcome message).
        /// 
        /// For custom attributes, you must prepend the `custom:` prefix to the attribute name.
        /// 
        /// To send a message inviting the user to sign up, you must specify the user's email address or phone number. You can do this in your call to AdminCreateUser or in the *Users* tab of the Amazon Cognito console for managing your user pools.
        /// 
        /// You must also provide an email address or phone number when you expect the user to do passwordless sign-in with an email or SMS OTP. These attributes must be provided when passwordless options are the only available, or when you don't submit a `TemporaryPassword` .
        /// 
        /// In your call to `AdminCreateUser` , you can set the `email_verified` attribute to `True` , and you can set the `phone_number_verified` attribute to `True` .
        /// 
        /// - *email* : The email address of the user to whom the message that contains the code and username will be sent. Required if the `email_verified` attribute is set to `True` , or if `"EMAIL"` is specified in the `DesiredDeliveryMediums` parameter.
        /// - *phone_number* : The phone number of the user to whom the message that contains the code and username will be sent. Required if the `phone_number_verified` attribute is set to `True` , or if `"SMS"` is specified in the `DesiredDeliveryMediums` parameter.
        /// </summary>
        [Output("userAttributes")]
        public Output<ImmutableArray<Outputs.UserPoolUserAttributeType>> UserAttributes { get; private set; } = null!;

        /// <summary>
        /// The ID of the user pool where you want to create a user.
        /// </summary>
        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;

        /// <summary>
        /// The value that you want to set as the username sign-in attribute. The following conditions apply to the username parameter.
        /// 
        /// - The username can't be a duplicate of another username in the same user pool.
        /// - You can't change the value of a username after you create it.
        /// - You can only provide a value if usernames are a valid sign-in attribute for your user pool. If your user pool only supports phone numbers or email addresses as sign-in attributes, Amazon Cognito automatically generates a username value. For more information, see [Customizing sign-in attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases) .
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda trigger. This set of key-value pairs are for custom validation of information that you collect from your users but don't need to retain.
        /// 
        /// Your Lambda function can analyze this additional data and act on it. Your function can automatically confirm and verify select users or perform external API operations like logging user attributes and validation data to Amazon CloudWatch Logs.
        /// 
        /// For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) .
        /// </summary>
        [Output("validationData")]
        public Output<ImmutableArray<Outputs.UserPoolUserAttributeType>> ValidationData { get; private set; } = null!;


        /// <summary>
        /// Create a UserPoolUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPoolUser(string name, UserPoolUserArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolUser", name, args ?? new UserPoolUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPoolUser(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolUser", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clientMetadata.*",
                    "desiredDeliveryMediums[*]",
                    "forceAliasCreation",
                    "messageAction",
                    "userAttributes[*]",
                    "userPoolId",
                    "username",
                    "validationData[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserPoolUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPoolUser Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserPoolUser(name, id, options);
        }
    }

    public sealed class UserPoolUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientMetadata")]
        private InputMap<string>? _clientMetadata;

        /// <summary>
        /// A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.
        /// 
        /// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you use the AdminCreateUser API action, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `ClientMetadata` attribute, which provides the data that you assigned to the ClientMetadata parameter in your AdminCreateUser request. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
        /// 
        /// For more information, see [Using Lambda triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
        /// 
        /// &gt; When you use the `ClientMetadata` parameter, note that Amazon Cognito won't do the following:
        /// &gt; 
        /// &gt; - Store the `ClientMetadata` value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration doesn't include triggers, the `ClientMetadata` parameter serves no purpose.
        /// &gt; - Validate the `ClientMetadata` value.
        /// &gt; - Encrypt the `ClientMetadata` value. Don't send sensitive information in this parameter.
        /// </summary>
        public InputMap<string> ClientMetadata
        {
            get => _clientMetadata ?? (_clientMetadata = new InputMap<string>());
            set => _clientMetadata = value;
        }

        [Input("desiredDeliveryMediums")]
        private InputList<string>? _desiredDeliveryMediums;

        /// <summary>
        /// Specify `EMAIL` if email will be used to send the welcome message. Specify `SMS` if the phone number will be used. The default value is `SMS` . You can specify more than one value.
        /// </summary>
        public InputList<string> DesiredDeliveryMediums
        {
            get => _desiredDeliveryMediums ?? (_desiredDeliveryMediums = new InputList<string>());
            set => _desiredDeliveryMediums = value;
        }

        /// <summary>
        /// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` . Otherwise, it is ignored.
        /// 
        /// If this parameter is set to `True` and the phone number or email address specified in the `UserAttributes` parameter already exists as an alias with a different user, this request migrates the alias from the previous user to the newly-created user. The previous user will no longer be able to log in using that alias.
        /// 
        /// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
        /// </summary>
        [Input("forceAliasCreation")]
        public Input<bool>? ForceAliasCreation { get; set; }

        /// <summary>
        /// Set to `RESEND` to resend the invitation message to a user that already exists, and to reset the temporary-password duration with a new temporary password. Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
        /// </summary>
        [Input("messageAction")]
        public Input<string>? MessageAction { get; set; }

        [Input("userAttributes")]
        private InputList<Inputs.UserPoolUserAttributeTypeArgs>? _userAttributes;

        /// <summary>
        /// An array of name-value pairs that contain user attributes and attribute values to be set for the user to be created. You can create a user without specifying any attributes other than `Username` . However, any attributes that you specify as required (when creating a user pool or in the *Attributes* tab of the console) either you should supply (in your call to `AdminCreateUser` ) or the user should supply (when they sign up in response to your welcome message).
        /// 
        /// For custom attributes, you must prepend the `custom:` prefix to the attribute name.
        /// 
        /// To send a message inviting the user to sign up, you must specify the user's email address or phone number. You can do this in your call to AdminCreateUser or in the *Users* tab of the Amazon Cognito console for managing your user pools.
        /// 
        /// You must also provide an email address or phone number when you expect the user to do passwordless sign-in with an email or SMS OTP. These attributes must be provided when passwordless options are the only available, or when you don't submit a `TemporaryPassword` .
        /// 
        /// In your call to `AdminCreateUser` , you can set the `email_verified` attribute to `True` , and you can set the `phone_number_verified` attribute to `True` .
        /// 
        /// - *email* : The email address of the user to whom the message that contains the code and username will be sent. Required if the `email_verified` attribute is set to `True` , or if `"EMAIL"` is specified in the `DesiredDeliveryMediums` parameter.
        /// - *phone_number* : The phone number of the user to whom the message that contains the code and username will be sent. Required if the `phone_number_verified` attribute is set to `True` , or if `"SMS"` is specified in the `DesiredDeliveryMediums` parameter.
        /// </summary>
        public InputList<Inputs.UserPoolUserAttributeTypeArgs> UserAttributes
        {
            get => _userAttributes ?? (_userAttributes = new InputList<Inputs.UserPoolUserAttributeTypeArgs>());
            set => _userAttributes = value;
        }

        /// <summary>
        /// The ID of the user pool where you want to create a user.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        /// <summary>
        /// The value that you want to set as the username sign-in attribute. The following conditions apply to the username parameter.
        /// 
        /// - The username can't be a duplicate of another username in the same user pool.
        /// - You can't change the value of a username after you create it.
        /// - You can only provide a value if usernames are a valid sign-in attribute for your user pool. If your user pool only supports phone numbers or email addresses as sign-in attributes, Amazon Cognito automatically generates a username value. For more information, see [Customizing sign-in attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases) .
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        [Input("validationData")]
        private InputList<Inputs.UserPoolUserAttributeTypeArgs>? _validationData;

        /// <summary>
        /// Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda trigger. This set of key-value pairs are for custom validation of information that you collect from your users but don't need to retain.
        /// 
        /// Your Lambda function can analyze this additional data and act on it. Your function can automatically confirm and verify select users or perform external API operations like logging user attributes and validation data to Amazon CloudWatch Logs.
        /// 
        /// For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) .
        /// </summary>
        public InputList<Inputs.UserPoolUserAttributeTypeArgs> ValidationData
        {
            get => _validationData ?? (_validationData = new InputList<Inputs.UserPoolUserAttributeTypeArgs>());
            set => _validationData = value;
        }

        public UserPoolUserArgs()
        {
        }
        public static new UserPoolUserArgs Empty => new UserPoolUserArgs();
    }
}
