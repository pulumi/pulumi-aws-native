// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Specifies one of the following actions to receive notifications when the alarm state changes.
    /// </summary>
    public sealed class AlarmModelAlarmActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
        ///  You must use expressions for all parameters in ``DynamoDBv2Action``. The expressions accept literals, operators, functions, references, and substitution templates.
        ///   **Examples**
        ///  +  For literal values, the expressions must contain single quotes. For example, the value for the ``tableName`` parameter can be ``'GreenhouseTemperatureTable'``.
        ///   +  For references, you must specify either variables or input values. For example, the value for the ``tableName`` parameter can be ``$variable.ddbtableName``.
        ///   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
        ///  In the following example, the value for the ``contentExpression`` parameter in ``Payload`` uses a substitution template. 
        ///   ``'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`` 
        ///   +  For a string concatenation, you must use ``+``. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
        ///  In the following example, the value for the ``tableName`` parameter uses a string concatenation. 
        ///   ``'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`` 
        ///   
        ///  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
        ///  The value for the ``type`` parameter in ``Payload`` must be ``JSON``.
        /// </summary>
        [Input("dynamoDBv2")]
        public Input<Inputs.AlarmModelDynamoDBv2Args>? DynamoDBv2 { get; set; }

        /// <summary>
        /// Defines an action to write to the Amazon DynamoDB table that you created. The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
        ///  You must use expressions for all parameters in ``DynamoDBAction``. The expressions accept literals, operators, functions, references, and substitution templates.
        ///   **Examples**
        ///  +  For literal values, the expressions must contain single quotes. For example, the value for the ``hashKeyType`` parameter can be ``'STRING'``.
        ///   +  For references, you must specify either variables or input values. For example, the value for the ``hashKeyField`` parameter can be ``$input.GreenhouseInput.name``.
        ///   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
        ///  In the following example, the value for the ``hashKeyValue`` parameter uses a substitution template. 
        ///   ``'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`` 
        ///   +  For a string concatenation, you must use ``+``. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
        ///  In the following example, the value for the ``tableName`` parameter uses a string concatenation. 
        ///   ``'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`` 
        ///   
        ///  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
        ///  If the defined payload type is a string, ``DynamoDBAction`` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the ``payloadField`` parameter is ``&lt;payload-field&gt;_raw``.
        /// </summary>
        [Input("dynamoDb")]
        public Input<Inputs.AlarmModelDynamoDbArgs>? DynamoDb { get; set; }

        /// <summary>
        /// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
        /// </summary>
        [Input("firehose")]
        public Input<Inputs.AlarmModelFirehoseArgs>? Firehose { get; set; }

        /// <summary>
        /// Sends an ITE input, passing in information about the detector model instance and the event that triggered the action.
        /// </summary>
        [Input("iotEvents")]
        public Input<Inputs.AlarmModelIotEventsArgs>? IotEvents { get; set; }

        /// <summary>
        /// Sends information about the detector model instance and the event that triggered the action to a specified asset property in ITSW.
        ///  You must use expressions for all parameters in ``IotSiteWiseAction``. The expressions accept literals, operators, functions, references, and substitutions templates.
        ///   **Examples**
        ///  +  For literal values, the expressions must contain single quotes. For example, the value for the ``propertyAlias`` parameter can be ``'/company/windfarm/3/turbine/7/temperature'``.
        ///   +  For references, you must specify either variables or input values. For example, the value for the ``assetId`` parameter can be ``$input.TurbineInput.assetId1``.
        ///   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
        ///  In the following example, the value for the ``propertyAlias`` parameter uses a substitution template. 
        ///   ``'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`` 
        ///   
        ///  You must specify either ``propertyAlias`` or both ``assetId`` and ``propertyId`` to identify the target asset property in ITSW.
        ///  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
        /// </summary>
        [Input("iotSiteWise")]
        public Input<Inputs.AlarmModelIotSiteWiseArgs>? IotSiteWise { get; set; }

        /// <summary>
        /// Information required to publish the MQTT message through the IoT message broker.
        /// </summary>
        [Input("iotTopicPublish")]
        public Input<Inputs.AlarmModelIotTopicPublishArgs>? IotTopicPublish { get; set; }

        /// <summary>
        /// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
        /// </summary>
        [Input("lambda")]
        public Input<Inputs.AlarmModelLambdaArgs>? Lambda { get; set; }

        /// <summary>
        /// Information required to publish the Amazon SNS message.
        /// </summary>
        [Input("sns")]
        public Input<Inputs.AlarmModelSnsArgs>? Sns { get; set; }

        /// <summary>
        /// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
        /// </summary>
        [Input("sqs")]
        public Input<Inputs.AlarmModelSqsArgs>? Sqs { get; set; }

        public AlarmModelAlarmActionArgs()
        {
        }
        public static new AlarmModelAlarmActionArgs Empty => new AlarmModelAlarmActionArgs();
    }
}
