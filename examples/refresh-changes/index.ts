import * as aws from "@pulumi/aws-native";

const ssmParam = new aws.ssm.Parameter('param', {
  type: aws.ssm.ParameterType.String,
  value: 'old-value',
});

export const paramName = ssmParam.name;
export const paramValue = ssmParam.value;
