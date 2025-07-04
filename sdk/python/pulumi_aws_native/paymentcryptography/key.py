# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 exportable: pulumi.Input[builtins.bool],
                 key_attributes: pulumi.Input['KeyAttributesArgs'],
                 derive_key_usage: Optional[pulumi.Input['KeyDeriveKeyUsage']] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 key_check_value_algorithm: Optional[pulumi.Input['KeyCheckValueAlgorithm']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[builtins.bool] exportable: Specifies whether the key is exportable. This data is immutable after the key is created.
        :param pulumi.Input['KeyAttributesArgs'] key_attributes: The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
        :param pulumi.Input['KeyDeriveKeyUsage'] derive_key_usage: The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether the key is enabled.
        :param pulumi.Input['KeyCheckValueAlgorithm'] key_check_value_algorithm: The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.
               
               For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
        """
        pulumi.set(__self__, "exportable", exportable)
        pulumi.set(__self__, "key_attributes", key_attributes)
        if derive_key_usage is not None:
            pulumi.set(__self__, "derive_key_usage", derive_key_usage)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if key_check_value_algorithm is not None:
            pulumi.set(__self__, "key_check_value_algorithm", key_check_value_algorithm)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def exportable(self) -> pulumi.Input[builtins.bool]:
        """
        Specifies whether the key is exportable. This data is immutable after the key is created.
        """
        return pulumi.get(self, "exportable")

    @exportable.setter
    def exportable(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "exportable", value)

    @property
    @pulumi.getter(name="keyAttributes")
    def key_attributes(self) -> pulumi.Input['KeyAttributesArgs']:
        """
        The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
        """
        return pulumi.get(self, "key_attributes")

    @key_attributes.setter
    def key_attributes(self, value: pulumi.Input['KeyAttributesArgs']):
        pulumi.set(self, "key_attributes", value)

    @property
    @pulumi.getter(name="deriveKeyUsage")
    def derive_key_usage(self) -> Optional[pulumi.Input['KeyDeriveKeyUsage']]:
        """
        The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
        """
        return pulumi.get(self, "derive_key_usage")

    @derive_key_usage.setter
    def derive_key_usage(self, value: Optional[pulumi.Input['KeyDeriveKeyUsage']]):
        pulumi.set(self, "derive_key_usage", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether the key is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="keyCheckValueAlgorithm")
    def key_check_value_algorithm(self) -> Optional[pulumi.Input['KeyCheckValueAlgorithm']]:
        """
        The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.

        For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
        """
        return pulumi.get(self, "key_check_value_algorithm")

    @key_check_value_algorithm.setter
    def key_check_value_algorithm(self, value: Optional[pulumi.Input['KeyCheckValueAlgorithm']]):
        pulumi.set(self, "key_check_value_algorithm", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:paymentcryptography:Key")
class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 derive_key_usage: Optional[pulumi.Input['KeyDeriveKeyUsage']] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 exportable: Optional[pulumi.Input[builtins.bool]] = None,
                 key_attributes: Optional[pulumi.Input[Union['KeyAttributesArgs', 'KeyAttributesArgsDict']]] = None,
                 key_check_value_algorithm: Optional[pulumi.Input['KeyCheckValueAlgorithm']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::PaymentCryptography::Key Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['KeyDeriveKeyUsage'] derive_key_usage: The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether the key is enabled.
        :param pulumi.Input[builtins.bool] exportable: Specifies whether the key is exportable. This data is immutable after the key is created.
        :param pulumi.Input[Union['KeyAttributesArgs', 'KeyAttributesArgsDict']] key_attributes: The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
        :param pulumi.Input['KeyCheckValueAlgorithm'] key_check_value_algorithm: The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.
               
               For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::PaymentCryptography::Key Resource Type

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 derive_key_usage: Optional[pulumi.Input['KeyDeriveKeyUsage']] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 exportable: Optional[pulumi.Input[builtins.bool]] = None,
                 key_attributes: Optional[pulumi.Input[Union['KeyAttributesArgs', 'KeyAttributesArgsDict']]] = None,
                 key_check_value_algorithm: Optional[pulumi.Input['KeyCheckValueAlgorithm']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["derive_key_usage"] = derive_key_usage
            __props__.__dict__["enabled"] = enabled
            if exportable is None and not opts.urn:
                raise TypeError("Missing required property 'exportable'")
            __props__.__dict__["exportable"] = exportable
            if key_attributes is None and not opts.urn:
                raise TypeError("Missing required property 'key_attributes'")
            __props__.__dict__["key_attributes"] = key_attributes
            __props__.__dict__["key_check_value_algorithm"] = key_check_value_algorithm
            __props__.__dict__["tags"] = tags
            __props__.__dict__["key_identifier"] = None
            __props__.__dict__["key_origin"] = None
            __props__.__dict__["key_state"] = None
        super(Key, __self__).__init__(
            'aws-native:paymentcryptography:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["derive_key_usage"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["exportable"] = None
        __props__.__dict__["key_attributes"] = None
        __props__.__dict__["key_check_value_algorithm"] = None
        __props__.__dict__["key_identifier"] = None
        __props__.__dict__["key_origin"] = None
        __props__.__dict__["key_state"] = None
        __props__.__dict__["tags"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deriveKeyUsage")
    def derive_key_usage(self) -> pulumi.Output[Optional['KeyDeriveKeyUsage']]:
        """
        The cryptographic usage of an ECDH derived key as deﬁned in section A.5.2 of the TR-31 spec.
        """
        return pulumi.get(self, "derive_key_usage")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether the key is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def exportable(self) -> pulumi.Output[builtins.bool]:
        """
        Specifies whether the key is exportable. This data is immutable after the key is created.
        """
        return pulumi.get(self, "exportable")

    @property
    @pulumi.getter(name="keyAttributes")
    def key_attributes(self) -> pulumi.Output['outputs.KeyAttributes']:
        """
        The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key. This data is immutable after the key is created.
        """
        return pulumi.get(self, "key_attributes")

    @property
    @pulumi.getter(name="keyCheckValueAlgorithm")
    def key_check_value_algorithm(self) -> pulumi.Output[Optional['KeyCheckValueAlgorithm']]:
        """
        The algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV). It is used to validate the key integrity.

        For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of zero, with the key to be checked and retaining the 3 highest order bytes of the encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where the input data is 16 bytes of zero and retaining the 3 highest order bytes of the encrypted result.
        """
        return pulumi.get(self, "key_check_value_algorithm")

    @property
    @pulumi.getter(name="keyIdentifier")
    def key_identifier(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "key_identifier")

    @property
    @pulumi.getter(name="keyOrigin")
    def key_origin(self) -> pulumi.Output['KeyOrigin']:
        """
        The source of the key material. For keys created within AWS Payment Cryptography, the value is `AWS_PAYMENT_CRYPTOGRAPHY` . For keys imported into AWS Payment Cryptography, the value is `EXTERNAL` .
        """
        return pulumi.get(self, "key_origin")

    @property
    @pulumi.getter(name="keyState")
    def key_state(self) -> pulumi.Output['KeyState']:
        """
        The state of key that is being created or deleted.
        """
        return pulumi.get(self, "key_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

