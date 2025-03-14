# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from ._enums import *

__all__ = [
    'AppMonitorConfigurationArgs',
    'AppMonitorConfigurationArgsDict',
    'AppMonitorCustomEventsArgs',
    'AppMonitorCustomEventsArgsDict',
    'AppMonitorMetricDefinitionArgs',
    'AppMonitorMetricDefinitionArgsDict',
    'AppMonitorMetricDestinationArgs',
    'AppMonitorMetricDestinationArgsDict',
    'AppMonitorResourcePolicyArgs',
    'AppMonitorResourcePolicyArgsDict',
]

MYPY = False

if not MYPY:
    class AppMonitorConfigurationArgsDict(TypedDict):
        """
        AppMonitor configuration
        """
        allow_cookies: NotRequired[pulumi.Input[bool]]
        """
        If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
        """
        enable_x_ray: NotRequired[pulumi.Input[bool]]
        """
        If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
        """
        excluded_pages: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        favorite_pages: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        A list of pages in the RUM console that are to be displayed with a favorite icon.
        """
        guest_role_arn: NotRequired[pulumi.Input[str]]
        """
        The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
        """
        identity_pool_id: NotRequired[pulumi.Input[str]]
        """
        The ID of the identity pool that is used to authorize the sending of data to RUM.
        """
        included_pages: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        metric_destinations: NotRequired[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDestinationArgsDict']]]]
        """
        An array of structures which define the destinations and the metrics that you want to send.
        """
        session_sample_rate: NotRequired[pulumi.Input[float]]
        """
        Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
        """
        telemetries: NotRequired[pulumi.Input[Sequence[pulumi.Input['AppMonitorTelemetry']]]]
        """
        An array that lists the types of telemetry data that this app monitor is to collect.
        """
elif False:
    AppMonitorConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppMonitorConfigurationArgs:
    def __init__(__self__, *,
                 allow_cookies: Optional[pulumi.Input[bool]] = None,
                 enable_x_ray: Optional[pulumi.Input[bool]] = None,
                 excluded_pages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 favorite_pages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 guest_role_arn: Optional[pulumi.Input[str]] = None,
                 identity_pool_id: Optional[pulumi.Input[str]] = None,
                 included_pages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metric_destinations: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDestinationArgs']]]] = None,
                 session_sample_rate: Optional[pulumi.Input[float]] = None,
                 telemetries: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorTelemetry']]]] = None):
        """
        AppMonitor configuration
        :param pulumi.Input[bool] allow_cookies: If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
        :param pulumi.Input[bool] enable_x_ray: If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_pages: A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] favorite_pages: A list of pages in the RUM console that are to be displayed with a favorite icon.
        :param pulumi.Input[str] guest_role_arn: The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
        :param pulumi.Input[str] identity_pool_id: The ID of the identity pool that is used to authorize the sending of data to RUM.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_pages: If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.
        :param pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDestinationArgs']]] metric_destinations: An array of structures which define the destinations and the metrics that you want to send.
        :param pulumi.Input[float] session_sample_rate: Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
        :param pulumi.Input[Sequence[pulumi.Input['AppMonitorTelemetry']]] telemetries: An array that lists the types of telemetry data that this app monitor is to collect.
        """
        if allow_cookies is not None:
            pulumi.set(__self__, "allow_cookies", allow_cookies)
        if enable_x_ray is not None:
            pulumi.set(__self__, "enable_x_ray", enable_x_ray)
        if excluded_pages is not None:
            pulumi.set(__self__, "excluded_pages", excluded_pages)
        if favorite_pages is not None:
            pulumi.set(__self__, "favorite_pages", favorite_pages)
        if guest_role_arn is not None:
            pulumi.set(__self__, "guest_role_arn", guest_role_arn)
        if identity_pool_id is not None:
            pulumi.set(__self__, "identity_pool_id", identity_pool_id)
        if included_pages is not None:
            pulumi.set(__self__, "included_pages", included_pages)
        if metric_destinations is not None:
            pulumi.set(__self__, "metric_destinations", metric_destinations)
        if session_sample_rate is not None:
            pulumi.set(__self__, "session_sample_rate", session_sample_rate)
        if telemetries is not None:
            pulumi.set(__self__, "telemetries", telemetries)

    @property
    @pulumi.getter(name="allowCookies")
    def allow_cookies(self) -> Optional[pulumi.Input[bool]]:
        """
        If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
        """
        return pulumi.get(self, "allow_cookies")

    @allow_cookies.setter
    def allow_cookies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_cookies", value)

    @property
    @pulumi.getter(name="enableXRay")
    def enable_x_ray(self) -> Optional[pulumi.Input[bool]]:
        """
        If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
        """
        return pulumi.get(self, "enable_x_ray")

    @enable_x_ray.setter
    def enable_x_ray(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_x_ray", value)

    @property
    @pulumi.getter(name="excludedPages")
    def excluded_pages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        return pulumi.get(self, "excluded_pages")

    @excluded_pages.setter
    def excluded_pages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_pages", value)

    @property
    @pulumi.getter(name="favoritePages")
    def favorite_pages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of pages in the RUM console that are to be displayed with a favorite icon.
        """
        return pulumi.get(self, "favorite_pages")

    @favorite_pages.setter
    def favorite_pages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "favorite_pages", value)

    @property
    @pulumi.getter(name="guestRoleArn")
    def guest_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
        """
        return pulumi.get(self, "guest_role_arn")

    @guest_role_arn.setter
    def guest_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "guest_role_arn", value)

    @property
    @pulumi.getter(name="identityPoolId")
    def identity_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the identity pool that is used to authorize the sending of data to RUM.
        """
        return pulumi.get(self, "identity_pool_id")

    @identity_pool_id.setter
    def identity_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_pool_id", value)

    @property
    @pulumi.getter(name="includedPages")
    def included_pages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        return pulumi.get(self, "included_pages")

    @included_pages.setter
    def included_pages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "included_pages", value)

    @property
    @pulumi.getter(name="metricDestinations")
    def metric_destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDestinationArgs']]]]:
        """
        An array of structures which define the destinations and the metrics that you want to send.
        """
        return pulumi.get(self, "metric_destinations")

    @metric_destinations.setter
    def metric_destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDestinationArgs']]]]):
        pulumi.set(self, "metric_destinations", value)

    @property
    @pulumi.getter(name="sessionSampleRate")
    def session_sample_rate(self) -> Optional[pulumi.Input[float]]:
        """
        Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
        """
        return pulumi.get(self, "session_sample_rate")

    @session_sample_rate.setter
    def session_sample_rate(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "session_sample_rate", value)

    @property
    @pulumi.getter
    def telemetries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorTelemetry']]]]:
        """
        An array that lists the types of telemetry data that this app monitor is to collect.
        """
        return pulumi.get(self, "telemetries")

    @telemetries.setter
    def telemetries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorTelemetry']]]]):
        pulumi.set(self, "telemetries", value)


if not MYPY:
    class AppMonitorCustomEventsArgsDict(TypedDict):
        """
        AppMonitor custom events configuration
        """
        status: NotRequired[pulumi.Input['AppMonitorCustomEventsStatus']]
        """
        Indicates whether AppMonitor accepts custom events.
        """
elif False:
    AppMonitorCustomEventsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppMonitorCustomEventsArgs:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input['AppMonitorCustomEventsStatus']] = None):
        """
        AppMonitor custom events configuration
        :param pulumi.Input['AppMonitorCustomEventsStatus'] status: Indicates whether AppMonitor accepts custom events.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['AppMonitorCustomEventsStatus']]:
        """
        Indicates whether AppMonitor accepts custom events.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['AppMonitorCustomEventsStatus']]):
        pulumi.set(self, "status", value)


if not MYPY:
    class AppMonitorMetricDefinitionArgsDict(TypedDict):
        """
        A single metric definition
        """
        name: pulumi.Input[str]
        """
        The name for the metric that is defined in this structure. For extended metrics, valid values are the following:

        PerformanceNavigationDuration

        PerformanceResourceDuration

        NavigationSatisfiedTransaction

        NavigationToleratedTransaction

        NavigationFrustratedTransaction

        WebVitalsCumulativeLayoutShift

        WebVitalsFirstInputDelay

        WebVitalsLargestContentfulPaint

        JsErrorCount

        HttpErrorCount

        SessionCount
        """
        dimension_keys: NotRequired[pulumi.Input[Mapping[str, pulumi.Input[str]]]]
        """
        Use this field only if you are sending the metric to CloudWatch.

        This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:

        "metadata.pageId": "PageId"

        "metadata.browserName": "BrowserName"

        "metadata.deviceType": "DeviceType"

        "metadata.osName": "OSName"

        "metadata.countryCode": "CountryCode"

        "event_details.fileType": "FileType"

        All dimensions listed in this field must also be included in EventPattern.
        """
        event_pattern: NotRequired[pulumi.Input[str]]
        """
        The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.

        When you define extended metrics, the metric definition is not valid if EventPattern is omitted.

        Example event patterns:

        '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'

        If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        """
        namespace: NotRequired[pulumi.Input[str]]
        """
        The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        """
        unit_label: NotRequired[pulumi.Input[str]]
        """
        The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        """
        value_key: NotRequired[pulumi.Input[str]]
        """
        The field within the event object that the metric value is sourced from.

        If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.

        If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        """
elif False:
    AppMonitorMetricDefinitionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppMonitorMetricDefinitionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 dimension_keys: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 unit_label: Optional[pulumi.Input[str]] = None,
                 value_key: Optional[pulumi.Input[str]] = None):
        """
        A single metric definition
        :param pulumi.Input[str] name: The name for the metric that is defined in this structure. For extended metrics, valid values are the following:
               
               PerformanceNavigationDuration
               
               PerformanceResourceDuration
               
               NavigationSatisfiedTransaction
               
               NavigationToleratedTransaction
               
               NavigationFrustratedTransaction
               
               WebVitalsCumulativeLayoutShift
               
               WebVitalsFirstInputDelay
               
               WebVitalsLargestContentfulPaint
               
               JsErrorCount
               
               HttpErrorCount
               
               SessionCount
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] dimension_keys: Use this field only if you are sending the metric to CloudWatch.
               
               This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:
               
               "metadata.pageId": "PageId"
               
               "metadata.browserName": "BrowserName"
               
               "metadata.deviceType": "DeviceType"
               
               "metadata.osName": "OSName"
               
               "metadata.countryCode": "CountryCode"
               
               "event_details.fileType": "FileType"
               
               All dimensions listed in this field must also be included in EventPattern.
        :param pulumi.Input[str] event_pattern: The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
               
               When you define extended metrics, the metric definition is not valid if EventPattern is omitted.
               
               Example event patterns:
               
               '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'
               
               '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'
               
               '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'
               
               If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        :param pulumi.Input[str] namespace: The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        :param pulumi.Input[str] unit_label: The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        :param pulumi.Input[str] value_key: The field within the event object that the metric value is sourced from.
               
               If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.
               
               If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        """
        pulumi.set(__self__, "name", name)
        if dimension_keys is not None:
            pulumi.set(__self__, "dimension_keys", dimension_keys)
        if event_pattern is not None:
            pulumi.set(__self__, "event_pattern", event_pattern)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if unit_label is not None:
            pulumi.set(__self__, "unit_label", unit_label)
        if value_key is not None:
            pulumi.set(__self__, "value_key", value_key)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name for the metric that is defined in this structure. For extended metrics, valid values are the following:

        PerformanceNavigationDuration

        PerformanceResourceDuration

        NavigationSatisfiedTransaction

        NavigationToleratedTransaction

        NavigationFrustratedTransaction

        WebVitalsCumulativeLayoutShift

        WebVitalsFirstInputDelay

        WebVitalsLargestContentfulPaint

        JsErrorCount

        HttpErrorCount

        SessionCount
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="dimensionKeys")
    def dimension_keys(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Use this field only if you are sending the metric to CloudWatch.

        This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:

        "metadata.pageId": "PageId"

        "metadata.browserName": "BrowserName"

        "metadata.deviceType": "DeviceType"

        "metadata.osName": "OSName"

        "metadata.countryCode": "CountryCode"

        "event_details.fileType": "FileType"

        All dimensions listed in this field must also be included in EventPattern.
        """
        return pulumi.get(self, "dimension_keys")

    @dimension_keys.setter
    def dimension_keys(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "dimension_keys", value)

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.

        When you define extended metrics, the metric definition is not valid if EventPattern is omitted.

        Example event patterns:

        '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'

        If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        """
        return pulumi.get(self, "event_pattern")

    @event_pattern.setter
    def event_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_pattern", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="unitLabel")
    def unit_label(self) -> Optional[pulumi.Input[str]]:
        """
        The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        """
        return pulumi.get(self, "unit_label")

    @unit_label.setter
    def unit_label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit_label", value)

    @property
    @pulumi.getter(name="valueKey")
    def value_key(self) -> Optional[pulumi.Input[str]]:
        """
        The field within the event object that the metric value is sourced from.

        If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.

        If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        """
        return pulumi.get(self, "value_key")

    @value_key.setter
    def value_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_key", value)


if not MYPY:
    class AppMonitorMetricDestinationArgsDict(TypedDict):
        """
        An structure which defines the destination and the metrics that you want to send.
        """
        destination: pulumi.Input['AppMonitorMetricDestinationDestination']
        """
        Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
        """
        destination_arn: NotRequired[pulumi.Input[str]]
        """
        Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
        """
        iam_role_arn: NotRequired[pulumi.Input[str]]
        """
        This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.

        This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
        """
        metric_definitions: NotRequired[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDefinitionArgsDict']]]]
        """
        An array of structures which define the metrics that you want to send.
        """
elif False:
    AppMonitorMetricDestinationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppMonitorMetricDestinationArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input['AppMonitorMetricDestinationDestination'],
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 metric_definitions: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDefinitionArgs']]]] = None):
        """
        An structure which defines the destination and the metrics that you want to send.
        :param pulumi.Input['AppMonitorMetricDestinationDestination'] destination: Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
        :param pulumi.Input[str] destination_arn: Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
        :param pulumi.Input[str] iam_role_arn: This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
               
               This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
        :param pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDefinitionArgs']]] metric_definitions: An array of structures which define the metrics that you want to send.
        """
        pulumi.set(__self__, "destination", destination)
        if destination_arn is not None:
            pulumi.set(__self__, "destination_arn", destination_arn)
        if iam_role_arn is not None:
            pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        if metric_definitions is not None:
            pulumi.set(__self__, "metric_definitions", metric_definitions)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input['AppMonitorMetricDestinationDestination']:
        """
        Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input['AppMonitorMetricDestinationDestination']):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
        """
        return pulumi.get(self, "destination_arn")

    @destination_arn.setter
    def destination_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_arn", value)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.

        This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
        """
        return pulumi.get(self, "iam_role_arn")

    @iam_role_arn.setter
    def iam_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_role_arn", value)

    @property
    @pulumi.getter(name="metricDefinitions")
    def metric_definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDefinitionArgs']]]]:
        """
        An array of structures which define the metrics that you want to send.
        """
        return pulumi.get(self, "metric_definitions")

    @metric_definitions.setter
    def metric_definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AppMonitorMetricDefinitionArgs']]]]):
        pulumi.set(self, "metric_definitions", value)


if not MYPY:
    class AppMonitorResourcePolicyArgsDict(TypedDict):
        """
        A structure that defines resource policy attached to your app monitor.
        """
        policy_document: pulumi.Input[str]
        """
        The JSON to use as the resource policy. The document can be up to 4 KB in size. 
        """
        policy_revision_id: NotRequired[pulumi.Input[str]]
        """
        A string value that you can use to conditionally update your policy. You can provide the revision ID of your existing policy to make mutating requests against that policy. 

         When you assign a policy revision ID, then later requests about that policy will be rejected with an InvalidPolicyRevisionIdException error if they don't provide the correct current revision ID.
        """
elif False:
    AppMonitorResourcePolicyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppMonitorResourcePolicyArgs:
    def __init__(__self__, *,
                 policy_document: pulumi.Input[str],
                 policy_revision_id: Optional[pulumi.Input[str]] = None):
        """
        A structure that defines resource policy attached to your app monitor.
        :param pulumi.Input[str] policy_document: The JSON to use as the resource policy. The document can be up to 4 KB in size. 
        :param pulumi.Input[str] policy_revision_id: A string value that you can use to conditionally update your policy. You can provide the revision ID of your existing policy to make mutating requests against that policy. 
               
                When you assign a policy revision ID, then later requests about that policy will be rejected with an InvalidPolicyRevisionIdException error if they don't provide the correct current revision ID.
        """
        pulumi.set(__self__, "policy_document", policy_document)
        if policy_revision_id is not None:
            pulumi.set(__self__, "policy_revision_id", policy_revision_id)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Input[str]:
        """
        The JSON to use as the resource policy. The document can be up to 4 KB in size. 
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_document", value)

    @property
    @pulumi.getter(name="policyRevisionId")
    def policy_revision_id(self) -> Optional[pulumi.Input[str]]:
        """
        A string value that you can use to conditionally update your policy. You can provide the revision ID of your existing policy to make mutating requests against that policy. 

         When you assign a policy revision ID, then later requests about that policy will be rejected with an InvalidPolicyRevisionIdException error if they don't provide the correct current revision ID.
        """
        return pulumi.get(self, "policy_revision_id")

    @policy_revision_id.setter
    def policy_revision_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_revision_id", value)


