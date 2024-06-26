// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolLambdaConfig
    {
        /// <summary>
        /// Creates an authentication challenge.
        /// </summary>
        public readonly string? CreateAuthChallenge;
        /// <summary>
        /// A custom email sender AWS Lambda trigger.
        /// </summary>
        public readonly Outputs.UserPoolCustomEmailSender? CustomEmailSender;
        /// <summary>
        /// A custom Message AWS Lambda trigger.
        /// </summary>
        public readonly string? CustomMessage;
        /// <summary>
        /// A custom SMS sender AWS Lambda trigger.
        /// </summary>
        public readonly Outputs.UserPoolCustomSmsSender? CustomSmsSender;
        /// <summary>
        /// Defines the authentication challenge.
        /// </summary>
        public readonly string? DefineAuthChallenge;
        /// <summary>
        /// The Amazon Resource Name of a AWS Key Management Service ( AWS KMS ) key. Amazon Cognito uses the key to encrypt codes and temporary passwords sent to `CustomEmailSender` and `CustomSMSSender` .
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// A post-authentication AWS Lambda trigger.
        /// </summary>
        public readonly string? PostAuthentication;
        /// <summary>
        /// A post-confirmation AWS Lambda trigger.
        /// </summary>
        public readonly string? PostConfirmation;
        /// <summary>
        /// A pre-authentication AWS Lambda trigger.
        /// </summary>
        public readonly string? PreAuthentication;
        /// <summary>
        /// A pre-registration AWS Lambda trigger.
        /// </summary>
        public readonly string? PreSignUp;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the function that you want to assign to your Lambda trigger.
        /// 
        /// Set this parameter for legacy purposes. If you also set an ARN in `PreTokenGenerationConfig` , its value must be identical to `PreTokenGeneration` . For new instances of pre token generation triggers, set the `LambdaArn` of `PreTokenGenerationConfig` .
        /// 
        /// You can set ``
        /// </summary>
        public readonly string? PreTokenGeneration;
        /// <summary>
        /// The detailed configuration of a pre token generation trigger. If you also set an ARN in `PreTokenGeneration` , its value must be identical to `PreTokenGenerationConfig` .
        /// </summary>
        public readonly Outputs.UserPoolPreTokenGenerationConfig? PreTokenGenerationConfig;
        /// <summary>
        /// The user migration Lambda config type.
        /// </summary>
        public readonly string? UserMigration;
        /// <summary>
        /// Verifies the authentication challenge response.
        /// </summary>
        public readonly string? VerifyAuthChallengeResponse;

        [OutputConstructor]
        private UserPoolLambdaConfig(
            string? createAuthChallenge,

            Outputs.UserPoolCustomEmailSender? customEmailSender,

            string? customMessage,

            Outputs.UserPoolCustomSmsSender? customSmsSender,

            string? defineAuthChallenge,

            string? kmsKeyId,

            string? postAuthentication,

            string? postConfirmation,

            string? preAuthentication,

            string? preSignUp,

            string? preTokenGeneration,

            Outputs.UserPoolPreTokenGenerationConfig? preTokenGenerationConfig,

            string? userMigration,

            string? verifyAuthChallengeResponse)
        {
            CreateAuthChallenge = createAuthChallenge;
            CustomEmailSender = customEmailSender;
            CustomMessage = customMessage;
            CustomSmsSender = customSmsSender;
            DefineAuthChallenge = defineAuthChallenge;
            KmsKeyId = kmsKeyId;
            PostAuthentication = postAuthentication;
            PostConfirmation = postConfirmation;
            PreAuthentication = preAuthentication;
            PreSignUp = preSignUp;
            PreTokenGeneration = preTokenGeneration;
            PreTokenGenerationConfig = preTokenGenerationConfig;
            UserMigration = userMigration;
            VerifyAuthChallengeResponse = verifyAuthChallengeResponse;
        }
    }
}
