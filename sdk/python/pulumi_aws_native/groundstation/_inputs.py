# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ConfigAntennaDownlinkConfigArgs',
    'ConfigAntennaDownlinkDemodDecodeConfigArgs',
    'ConfigAntennaUplinkConfigArgs',
    'ConfigDataflowEndpointConfigArgs',
    'ConfigDataArgs',
    'ConfigDecodeConfigArgs',
    'ConfigDemodulationConfigArgs',
    'ConfigEirpArgs',
    'ConfigFrequencyBandwidthArgs',
    'ConfigFrequencyArgs',
    'ConfigS3RecordingConfigArgs',
    'ConfigSpectrumConfigArgs',
    'ConfigTagArgs',
    'ConfigTrackingConfigArgs',
    'ConfigUplinkEchoConfigArgs',
    'ConfigUplinkSpectrumConfigArgs',
    'DataflowEndpointGroupDataflowEndpointArgs',
    'DataflowEndpointGroupEndpointDetailsArgs',
    'DataflowEndpointGroupSecurityDetailsArgs',
    'DataflowEndpointGroupSocketAddressArgs',
    'DataflowEndpointGroupTagArgs',
    'MissionProfileDataflowEdgeArgs',
    'MissionProfileTagArgs',
]

@pulumi.input_type
class ConfigAntennaDownlinkConfigArgs:
    def __init__(__self__, *,
                 spectrum_config: Optional[pulumi.Input['ConfigSpectrumConfigArgs']] = None):
        if spectrum_config is not None:
            pulumi.set(__self__, "spectrum_config", spectrum_config)

    @property
    @pulumi.getter(name="spectrumConfig")
    def spectrum_config(self) -> Optional[pulumi.Input['ConfigSpectrumConfigArgs']]:
        return pulumi.get(self, "spectrum_config")

    @spectrum_config.setter
    def spectrum_config(self, value: Optional[pulumi.Input['ConfigSpectrumConfigArgs']]):
        pulumi.set(self, "spectrum_config", value)


@pulumi.input_type
class ConfigAntennaDownlinkDemodDecodeConfigArgs:
    def __init__(__self__, *,
                 decode_config: Optional[pulumi.Input['ConfigDecodeConfigArgs']] = None,
                 demodulation_config: Optional[pulumi.Input['ConfigDemodulationConfigArgs']] = None,
                 spectrum_config: Optional[pulumi.Input['ConfigSpectrumConfigArgs']] = None):
        if decode_config is not None:
            pulumi.set(__self__, "decode_config", decode_config)
        if demodulation_config is not None:
            pulumi.set(__self__, "demodulation_config", demodulation_config)
        if spectrum_config is not None:
            pulumi.set(__self__, "spectrum_config", spectrum_config)

    @property
    @pulumi.getter(name="decodeConfig")
    def decode_config(self) -> Optional[pulumi.Input['ConfigDecodeConfigArgs']]:
        return pulumi.get(self, "decode_config")

    @decode_config.setter
    def decode_config(self, value: Optional[pulumi.Input['ConfigDecodeConfigArgs']]):
        pulumi.set(self, "decode_config", value)

    @property
    @pulumi.getter(name="demodulationConfig")
    def demodulation_config(self) -> Optional[pulumi.Input['ConfigDemodulationConfigArgs']]:
        return pulumi.get(self, "demodulation_config")

    @demodulation_config.setter
    def demodulation_config(self, value: Optional[pulumi.Input['ConfigDemodulationConfigArgs']]):
        pulumi.set(self, "demodulation_config", value)

    @property
    @pulumi.getter(name="spectrumConfig")
    def spectrum_config(self) -> Optional[pulumi.Input['ConfigSpectrumConfigArgs']]:
        return pulumi.get(self, "spectrum_config")

    @spectrum_config.setter
    def spectrum_config(self, value: Optional[pulumi.Input['ConfigSpectrumConfigArgs']]):
        pulumi.set(self, "spectrum_config", value)


@pulumi.input_type
class ConfigAntennaUplinkConfigArgs:
    def __init__(__self__, *,
                 spectrum_config: Optional[pulumi.Input['ConfigUplinkSpectrumConfigArgs']] = None,
                 target_eirp: Optional[pulumi.Input['ConfigEirpArgs']] = None,
                 transmit_disabled: Optional[pulumi.Input[bool]] = None):
        if spectrum_config is not None:
            pulumi.set(__self__, "spectrum_config", spectrum_config)
        if target_eirp is not None:
            pulumi.set(__self__, "target_eirp", target_eirp)
        if transmit_disabled is not None:
            pulumi.set(__self__, "transmit_disabled", transmit_disabled)

    @property
    @pulumi.getter(name="spectrumConfig")
    def spectrum_config(self) -> Optional[pulumi.Input['ConfigUplinkSpectrumConfigArgs']]:
        return pulumi.get(self, "spectrum_config")

    @spectrum_config.setter
    def spectrum_config(self, value: Optional[pulumi.Input['ConfigUplinkSpectrumConfigArgs']]):
        pulumi.set(self, "spectrum_config", value)

    @property
    @pulumi.getter(name="targetEirp")
    def target_eirp(self) -> Optional[pulumi.Input['ConfigEirpArgs']]:
        return pulumi.get(self, "target_eirp")

    @target_eirp.setter
    def target_eirp(self, value: Optional[pulumi.Input['ConfigEirpArgs']]):
        pulumi.set(self, "target_eirp", value)

    @property
    @pulumi.getter(name="transmitDisabled")
    def transmit_disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "transmit_disabled")

    @transmit_disabled.setter
    def transmit_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transmit_disabled", value)


@pulumi.input_type
class ConfigDataflowEndpointConfigArgs:
    def __init__(__self__, *,
                 dataflow_endpoint_name: Optional[pulumi.Input[str]] = None,
                 dataflow_endpoint_region: Optional[pulumi.Input[str]] = None):
        if dataflow_endpoint_name is not None:
            pulumi.set(__self__, "dataflow_endpoint_name", dataflow_endpoint_name)
        if dataflow_endpoint_region is not None:
            pulumi.set(__self__, "dataflow_endpoint_region", dataflow_endpoint_region)

    @property
    @pulumi.getter(name="dataflowEndpointName")
    def dataflow_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dataflow_endpoint_name")

    @dataflow_endpoint_name.setter
    def dataflow_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataflow_endpoint_name", value)

    @property
    @pulumi.getter(name="dataflowEndpointRegion")
    def dataflow_endpoint_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dataflow_endpoint_region")

    @dataflow_endpoint_region.setter
    def dataflow_endpoint_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataflow_endpoint_region", value)


@pulumi.input_type
class ConfigDataArgs:
    def __init__(__self__, *,
                 antenna_downlink_config: Optional[pulumi.Input['ConfigAntennaDownlinkConfigArgs']] = None,
                 antenna_downlink_demod_decode_config: Optional[pulumi.Input['ConfigAntennaDownlinkDemodDecodeConfigArgs']] = None,
                 antenna_uplink_config: Optional[pulumi.Input['ConfigAntennaUplinkConfigArgs']] = None,
                 dataflow_endpoint_config: Optional[pulumi.Input['ConfigDataflowEndpointConfigArgs']] = None,
                 s3_recording_config: Optional[pulumi.Input['ConfigS3RecordingConfigArgs']] = None,
                 tracking_config: Optional[pulumi.Input['ConfigTrackingConfigArgs']] = None,
                 uplink_echo_config: Optional[pulumi.Input['ConfigUplinkEchoConfigArgs']] = None):
        if antenna_downlink_config is not None:
            pulumi.set(__self__, "antenna_downlink_config", antenna_downlink_config)
        if antenna_downlink_demod_decode_config is not None:
            pulumi.set(__self__, "antenna_downlink_demod_decode_config", antenna_downlink_demod_decode_config)
        if antenna_uplink_config is not None:
            pulumi.set(__self__, "antenna_uplink_config", antenna_uplink_config)
        if dataflow_endpoint_config is not None:
            pulumi.set(__self__, "dataflow_endpoint_config", dataflow_endpoint_config)
        if s3_recording_config is not None:
            pulumi.set(__self__, "s3_recording_config", s3_recording_config)
        if tracking_config is not None:
            pulumi.set(__self__, "tracking_config", tracking_config)
        if uplink_echo_config is not None:
            pulumi.set(__self__, "uplink_echo_config", uplink_echo_config)

    @property
    @pulumi.getter(name="antennaDownlinkConfig")
    def antenna_downlink_config(self) -> Optional[pulumi.Input['ConfigAntennaDownlinkConfigArgs']]:
        return pulumi.get(self, "antenna_downlink_config")

    @antenna_downlink_config.setter
    def antenna_downlink_config(self, value: Optional[pulumi.Input['ConfigAntennaDownlinkConfigArgs']]):
        pulumi.set(self, "antenna_downlink_config", value)

    @property
    @pulumi.getter(name="antennaDownlinkDemodDecodeConfig")
    def antenna_downlink_demod_decode_config(self) -> Optional[pulumi.Input['ConfigAntennaDownlinkDemodDecodeConfigArgs']]:
        return pulumi.get(self, "antenna_downlink_demod_decode_config")

    @antenna_downlink_demod_decode_config.setter
    def antenna_downlink_demod_decode_config(self, value: Optional[pulumi.Input['ConfigAntennaDownlinkDemodDecodeConfigArgs']]):
        pulumi.set(self, "antenna_downlink_demod_decode_config", value)

    @property
    @pulumi.getter(name="antennaUplinkConfig")
    def antenna_uplink_config(self) -> Optional[pulumi.Input['ConfigAntennaUplinkConfigArgs']]:
        return pulumi.get(self, "antenna_uplink_config")

    @antenna_uplink_config.setter
    def antenna_uplink_config(self, value: Optional[pulumi.Input['ConfigAntennaUplinkConfigArgs']]):
        pulumi.set(self, "antenna_uplink_config", value)

    @property
    @pulumi.getter(name="dataflowEndpointConfig")
    def dataflow_endpoint_config(self) -> Optional[pulumi.Input['ConfigDataflowEndpointConfigArgs']]:
        return pulumi.get(self, "dataflow_endpoint_config")

    @dataflow_endpoint_config.setter
    def dataflow_endpoint_config(self, value: Optional[pulumi.Input['ConfigDataflowEndpointConfigArgs']]):
        pulumi.set(self, "dataflow_endpoint_config", value)

    @property
    @pulumi.getter(name="s3RecordingConfig")
    def s3_recording_config(self) -> Optional[pulumi.Input['ConfigS3RecordingConfigArgs']]:
        return pulumi.get(self, "s3_recording_config")

    @s3_recording_config.setter
    def s3_recording_config(self, value: Optional[pulumi.Input['ConfigS3RecordingConfigArgs']]):
        pulumi.set(self, "s3_recording_config", value)

    @property
    @pulumi.getter(name="trackingConfig")
    def tracking_config(self) -> Optional[pulumi.Input['ConfigTrackingConfigArgs']]:
        return pulumi.get(self, "tracking_config")

    @tracking_config.setter
    def tracking_config(self, value: Optional[pulumi.Input['ConfigTrackingConfigArgs']]):
        pulumi.set(self, "tracking_config", value)

    @property
    @pulumi.getter(name="uplinkEchoConfig")
    def uplink_echo_config(self) -> Optional[pulumi.Input['ConfigUplinkEchoConfigArgs']]:
        return pulumi.get(self, "uplink_echo_config")

    @uplink_echo_config.setter
    def uplink_echo_config(self, value: Optional[pulumi.Input['ConfigUplinkEchoConfigArgs']]):
        pulumi.set(self, "uplink_echo_config", value)


@pulumi.input_type
class ConfigDecodeConfigArgs:
    def __init__(__self__, *,
                 unvalidated_json: Optional[pulumi.Input[str]] = None):
        if unvalidated_json is not None:
            pulumi.set(__self__, "unvalidated_json", unvalidated_json)

    @property
    @pulumi.getter(name="unvalidatedJSON")
    def unvalidated_json(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unvalidated_json")

    @unvalidated_json.setter
    def unvalidated_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unvalidated_json", value)


@pulumi.input_type
class ConfigDemodulationConfigArgs:
    def __init__(__self__, *,
                 unvalidated_json: Optional[pulumi.Input[str]] = None):
        if unvalidated_json is not None:
            pulumi.set(__self__, "unvalidated_json", unvalidated_json)

    @property
    @pulumi.getter(name="unvalidatedJSON")
    def unvalidated_json(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unvalidated_json")

    @unvalidated_json.setter
    def unvalidated_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unvalidated_json", value)


@pulumi.input_type
class ConfigEirpArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input['ConfigEirpUnits']] = None,
                 value: Optional[pulumi.Input[float]] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input['ConfigEirpUnits']]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input['ConfigEirpUnits']]):
        pulumi.set(self, "units", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigFrequencyBandwidthArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input['ConfigBandwidthUnits']] = None,
                 value: Optional[pulumi.Input[float]] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input['ConfigBandwidthUnits']]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input['ConfigBandwidthUnits']]):
        pulumi.set(self, "units", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigFrequencyArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input['ConfigFrequencyUnits']] = None,
                 value: Optional[pulumi.Input[float]] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input['ConfigFrequencyUnits']]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input['ConfigFrequencyUnits']]):
        pulumi.set(self, "units", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigS3RecordingConfigArgs:
    def __init__(__self__, *,
                 bucket_arn: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        if bucket_arn is not None:
            pulumi.set(__self__, "bucket_arn", bucket_arn)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="bucketArn")
    def bucket_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket_arn")

    @bucket_arn.setter
    def bucket_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_arn", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class ConfigSpectrumConfigArgs:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input['ConfigFrequencyBandwidthArgs']] = None,
                 center_frequency: Optional[pulumi.Input['ConfigFrequencyArgs']] = None,
                 polarization: Optional[pulumi.Input['ConfigPolarization']] = None):
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if center_frequency is not None:
            pulumi.set(__self__, "center_frequency", center_frequency)
        if polarization is not None:
            pulumi.set(__self__, "polarization", polarization)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input['ConfigFrequencyBandwidthArgs']]:
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input['ConfigFrequencyBandwidthArgs']]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="centerFrequency")
    def center_frequency(self) -> Optional[pulumi.Input['ConfigFrequencyArgs']]:
        return pulumi.get(self, "center_frequency")

    @center_frequency.setter
    def center_frequency(self, value: Optional[pulumi.Input['ConfigFrequencyArgs']]):
        pulumi.set(self, "center_frequency", value)

    @property
    @pulumi.getter
    def polarization(self) -> Optional[pulumi.Input['ConfigPolarization']]:
        return pulumi.get(self, "polarization")

    @polarization.setter
    def polarization(self, value: Optional[pulumi.Input['ConfigPolarization']]):
        pulumi.set(self, "polarization", value)


@pulumi.input_type
class ConfigTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigTrackingConfigArgs:
    def __init__(__self__, *,
                 autotrack: Optional[pulumi.Input['ConfigTrackingConfigAutotrack']] = None):
        if autotrack is not None:
            pulumi.set(__self__, "autotrack", autotrack)

    @property
    @pulumi.getter
    def autotrack(self) -> Optional[pulumi.Input['ConfigTrackingConfigAutotrack']]:
        return pulumi.get(self, "autotrack")

    @autotrack.setter
    def autotrack(self, value: Optional[pulumi.Input['ConfigTrackingConfigAutotrack']]):
        pulumi.set(self, "autotrack", value)


@pulumi.input_type
class ConfigUplinkEchoConfigArgs:
    def __init__(__self__, *,
                 antenna_uplink_config_arn: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        if antenna_uplink_config_arn is not None:
            pulumi.set(__self__, "antenna_uplink_config_arn", antenna_uplink_config_arn)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="antennaUplinkConfigArn")
    def antenna_uplink_config_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "antenna_uplink_config_arn")

    @antenna_uplink_config_arn.setter
    def antenna_uplink_config_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "antenna_uplink_config_arn", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ConfigUplinkSpectrumConfigArgs:
    def __init__(__self__, *,
                 center_frequency: Optional[pulumi.Input['ConfigFrequencyArgs']] = None,
                 polarization: Optional[pulumi.Input['ConfigPolarization']] = None):
        if center_frequency is not None:
            pulumi.set(__self__, "center_frequency", center_frequency)
        if polarization is not None:
            pulumi.set(__self__, "polarization", polarization)

    @property
    @pulumi.getter(name="centerFrequency")
    def center_frequency(self) -> Optional[pulumi.Input['ConfigFrequencyArgs']]:
        return pulumi.get(self, "center_frequency")

    @center_frequency.setter
    def center_frequency(self, value: Optional[pulumi.Input['ConfigFrequencyArgs']]):
        pulumi.set(self, "center_frequency", value)

    @property
    @pulumi.getter
    def polarization(self) -> Optional[pulumi.Input['ConfigPolarization']]:
        return pulumi.get(self, "polarization")

    @polarization.setter
    def polarization(self, value: Optional[pulumi.Input['ConfigPolarization']]):
        pulumi.set(self, "polarization", value)


@pulumi.input_type
class DataflowEndpointGroupDataflowEndpointArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input['DataflowEndpointGroupSocketAddressArgs']] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        if address is not None:
            pulumi.set(__self__, "address", address)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input['DataflowEndpointGroupSocketAddressArgs']]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input['DataflowEndpointGroupSocketAddressArgs']]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DataflowEndpointGroupEndpointDetailsArgs:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input['DataflowEndpointGroupDataflowEndpointArgs']] = None,
                 security_details: Optional[pulumi.Input['DataflowEndpointGroupSecurityDetailsArgs']] = None):
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if security_details is not None:
            pulumi.set(__self__, "security_details", security_details)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input['DataflowEndpointGroupDataflowEndpointArgs']]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input['DataflowEndpointGroupDataflowEndpointArgs']]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="securityDetails")
    def security_details(self) -> Optional[pulumi.Input['DataflowEndpointGroupSecurityDetailsArgs']]:
        return pulumi.get(self, "security_details")

    @security_details.setter
    def security_details(self, value: Optional[pulumi.Input['DataflowEndpointGroupSecurityDetailsArgs']]):
        pulumi.set(self, "security_details", value)


@pulumi.input_type
class DataflowEndpointGroupSecurityDetailsArgs:
    def __init__(__self__, *,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)


@pulumi.input_type
class DataflowEndpointGroupSocketAddressArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class DataflowEndpointGroupTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class MissionProfileDataflowEdgeArgs:
    def __init__(__self__, *,
                 destination: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


@pulumi.input_type
class MissionProfileTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

