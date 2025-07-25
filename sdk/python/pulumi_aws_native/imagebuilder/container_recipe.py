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
from ._enums import *
from ._inputs import *

__all__ = ['ContainerRecipeArgs', 'ContainerRecipe']

@pulumi.input_type
class ContainerRecipeArgs:
    def __init__(__self__, *,
                 components: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerRecipeComponentConfigurationArgs']]]] = None,
                 container_type: Optional[pulumi.Input['ContainerRecipeContainerType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_data: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_uri: Optional[pulumi.Input[builtins.str]] = None,
                 image_os_version_override: Optional[pulumi.Input[builtins.str]] = None,
                 instance_configuration: Optional[pulumi.Input['ContainerRecipeInstanceConfigurationArgs']] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_image: Optional[pulumi.Input[builtins.str]] = None,
                 platform_override: Optional[pulumi.Input['ContainerRecipePlatformOverride']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_repository: Optional[pulumi.Input['ContainerRecipeTargetContainerRepositoryArgs']] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 working_directory: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ContainerRecipe resource.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerRecipeComponentConfigurationArgs']]] components: Components for build and test that are included in the container recipe.
        :param pulumi.Input['ContainerRecipeContainerType'] container_type: Specifies the type of container, such as Docker.
        :param pulumi.Input[builtins.str] description: The description of the container recipe.
        :param pulumi.Input[builtins.str] dockerfile_template_data: Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
        :param pulumi.Input[builtins.str] dockerfile_template_uri: The S3 URI for the Dockerfile that will be used to build your container image.
        :param pulumi.Input[builtins.str] image_os_version_override: Specifies the operating system version for the source image.
        :param pulumi.Input['ContainerRecipeInstanceConfigurationArgs'] instance_configuration: A group of options that can be used to configure an instance for building and testing container images.
        :param pulumi.Input[builtins.str] kms_key_id: Identifies which KMS key is used to encrypt the container image.
        :param pulumi.Input[builtins.str] name: The name of the container recipe.
        :param pulumi.Input[builtins.str] parent_image: The source image for the container recipe.
        :param pulumi.Input['ContainerRecipePlatformOverride'] platform_override: Specifies the operating system platform when you use a custom source image.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags that are attached to the container recipe.
        :param pulumi.Input['ContainerRecipeTargetContainerRepositoryArgs'] target_repository: The destination repository for the container image.
        :param pulumi.Input[builtins.str] version: The semantic version of the container recipe (<major>.<minor>.<patch>).
        :param pulumi.Input[builtins.str] working_directory: The working directory to be used during build and test workflows.
        """
        if components is not None:
            pulumi.set(__self__, "components", components)
        if container_type is not None:
            pulumi.set(__self__, "container_type", container_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dockerfile_template_data is not None:
            pulumi.set(__self__, "dockerfile_template_data", dockerfile_template_data)
        if dockerfile_template_uri is not None:
            pulumi.set(__self__, "dockerfile_template_uri", dockerfile_template_uri)
        if image_os_version_override is not None:
            pulumi.set(__self__, "image_os_version_override", image_os_version_override)
        if instance_configuration is not None:
            pulumi.set(__self__, "instance_configuration", instance_configuration)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_image is not None:
            pulumi.set(__self__, "parent_image", parent_image)
        if platform_override is not None:
            pulumi.set(__self__, "platform_override", platform_override)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_repository is not None:
            pulumi.set(__self__, "target_repository", target_repository)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if working_directory is not None:
            pulumi.set(__self__, "working_directory", working_directory)

    @property
    @pulumi.getter
    def components(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerRecipeComponentConfigurationArgs']]]]:
        """
        Components for build and test that are included in the container recipe.
        """
        return pulumi.get(self, "components")

    @components.setter
    def components(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerRecipeComponentConfigurationArgs']]]]):
        pulumi.set(self, "components", value)

    @property
    @pulumi.getter(name="containerType")
    def container_type(self) -> Optional[pulumi.Input['ContainerRecipeContainerType']]:
        """
        Specifies the type of container, such as Docker.
        """
        return pulumi.get(self, "container_type")

    @container_type.setter
    def container_type(self, value: Optional[pulumi.Input['ContainerRecipeContainerType']]):
        pulumi.set(self, "container_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the container recipe.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dockerfileTemplateData")
    def dockerfile_template_data(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
        """
        return pulumi.get(self, "dockerfile_template_data")

    @dockerfile_template_data.setter
    def dockerfile_template_data(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dockerfile_template_data", value)

    @property
    @pulumi.getter(name="dockerfileTemplateUri")
    def dockerfile_template_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The S3 URI for the Dockerfile that will be used to build your container image.
        """
        return pulumi.get(self, "dockerfile_template_uri")

    @dockerfile_template_uri.setter
    def dockerfile_template_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dockerfile_template_uri", value)

    @property
    @pulumi.getter(name="imageOsVersionOverride")
    def image_os_version_override(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the operating system version for the source image.
        """
        return pulumi.get(self, "image_os_version_override")

    @image_os_version_override.setter
    def image_os_version_override(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_os_version_override", value)

    @property
    @pulumi.getter(name="instanceConfiguration")
    def instance_configuration(self) -> Optional[pulumi.Input['ContainerRecipeInstanceConfigurationArgs']]:
        """
        A group of options that can be used to configure an instance for building and testing container images.
        """
        return pulumi.get(self, "instance_configuration")

    @instance_configuration.setter
    def instance_configuration(self, value: Optional[pulumi.Input['ContainerRecipeInstanceConfigurationArgs']]):
        pulumi.set(self, "instance_configuration", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Identifies which KMS key is used to encrypt the container image.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the container recipe.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentImage")
    def parent_image(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The source image for the container recipe.
        """
        return pulumi.get(self, "parent_image")

    @parent_image.setter
    def parent_image(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "parent_image", value)

    @property
    @pulumi.getter(name="platformOverride")
    def platform_override(self) -> Optional[pulumi.Input['ContainerRecipePlatformOverride']]:
        """
        Specifies the operating system platform when you use a custom source image.
        """
        return pulumi.get(self, "platform_override")

    @platform_override.setter
    def platform_override(self, value: Optional[pulumi.Input['ContainerRecipePlatformOverride']]):
        pulumi.set(self, "platform_override", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Tags that are attached to the container recipe.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetRepository")
    def target_repository(self) -> Optional[pulumi.Input['ContainerRecipeTargetContainerRepositoryArgs']]:
        """
        The destination repository for the container image.
        """
        return pulumi.get(self, "target_repository")

    @target_repository.setter
    def target_repository(self, value: Optional[pulumi.Input['ContainerRecipeTargetContainerRepositoryArgs']]):
        pulumi.set(self, "target_repository", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The semantic version of the container recipe (<major>.<minor>.<patch>).
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The working directory to be used during build and test workflows.
        """
        return pulumi.get(self, "working_directory")

    @working_directory.setter
    def working_directory(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "working_directory", value)


@pulumi.type_token("aws-native:imagebuilder:ContainerRecipe")
class ContainerRecipe(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerRecipeComponentConfigurationArgs', 'ContainerRecipeComponentConfigurationArgsDict']]]]] = None,
                 container_type: Optional[pulumi.Input['ContainerRecipeContainerType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_data: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_uri: Optional[pulumi.Input[builtins.str]] = None,
                 image_os_version_override: Optional[pulumi.Input[builtins.str]] = None,
                 instance_configuration: Optional[pulumi.Input[Union['ContainerRecipeInstanceConfigurationArgs', 'ContainerRecipeInstanceConfigurationArgsDict']]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_image: Optional[pulumi.Input[builtins.str]] = None,
                 platform_override: Optional[pulumi.Input['ContainerRecipePlatformOverride']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_repository: Optional[pulumi.Input[Union['ContainerRecipeTargetContainerRepositoryArgs', 'ContainerRecipeTargetContainerRepositoryArgsDict']]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 working_directory: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::ImageBuilder::ContainerRecipe

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ContainerRecipeComponentConfigurationArgs', 'ContainerRecipeComponentConfigurationArgsDict']]]] components: Components for build and test that are included in the container recipe.
        :param pulumi.Input['ContainerRecipeContainerType'] container_type: Specifies the type of container, such as Docker.
        :param pulumi.Input[builtins.str] description: The description of the container recipe.
        :param pulumi.Input[builtins.str] dockerfile_template_data: Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
        :param pulumi.Input[builtins.str] dockerfile_template_uri: The S3 URI for the Dockerfile that will be used to build your container image.
        :param pulumi.Input[builtins.str] image_os_version_override: Specifies the operating system version for the source image.
        :param pulumi.Input[Union['ContainerRecipeInstanceConfigurationArgs', 'ContainerRecipeInstanceConfigurationArgsDict']] instance_configuration: A group of options that can be used to configure an instance for building and testing container images.
        :param pulumi.Input[builtins.str] kms_key_id: Identifies which KMS key is used to encrypt the container image.
        :param pulumi.Input[builtins.str] name: The name of the container recipe.
        :param pulumi.Input[builtins.str] parent_image: The source image for the container recipe.
        :param pulumi.Input['ContainerRecipePlatformOverride'] platform_override: Specifies the operating system platform when you use a custom source image.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags that are attached to the container recipe.
        :param pulumi.Input[Union['ContainerRecipeTargetContainerRepositoryArgs', 'ContainerRecipeTargetContainerRepositoryArgsDict']] target_repository: The destination repository for the container image.
        :param pulumi.Input[builtins.str] version: The semantic version of the container recipe (<major>.<minor>.<patch>).
        :param pulumi.Input[builtins.str] working_directory: The working directory to be used during build and test workflows.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContainerRecipeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::ImageBuilder::ContainerRecipe

        :param str resource_name: The name of the resource.
        :param ContainerRecipeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerRecipeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerRecipeComponentConfigurationArgs', 'ContainerRecipeComponentConfigurationArgsDict']]]]] = None,
                 container_type: Optional[pulumi.Input['ContainerRecipeContainerType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_data: Optional[pulumi.Input[builtins.str]] = None,
                 dockerfile_template_uri: Optional[pulumi.Input[builtins.str]] = None,
                 image_os_version_override: Optional[pulumi.Input[builtins.str]] = None,
                 instance_configuration: Optional[pulumi.Input[Union['ContainerRecipeInstanceConfigurationArgs', 'ContainerRecipeInstanceConfigurationArgsDict']]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_image: Optional[pulumi.Input[builtins.str]] = None,
                 platform_override: Optional[pulumi.Input['ContainerRecipePlatformOverride']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_repository: Optional[pulumi.Input[Union['ContainerRecipeTargetContainerRepositoryArgs', 'ContainerRecipeTargetContainerRepositoryArgsDict']]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 working_directory: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerRecipeArgs.__new__(ContainerRecipeArgs)

            __props__.__dict__["components"] = components
            __props__.__dict__["container_type"] = container_type
            __props__.__dict__["description"] = description
            __props__.__dict__["dockerfile_template_data"] = dockerfile_template_data
            __props__.__dict__["dockerfile_template_uri"] = dockerfile_template_uri
            __props__.__dict__["image_os_version_override"] = image_os_version_override
            __props__.__dict__["instance_configuration"] = instance_configuration
            __props__.__dict__["kms_key_id"] = kms_key_id
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_image"] = parent_image
            __props__.__dict__["platform_override"] = platform_override
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_repository"] = target_repository
            __props__.__dict__["version"] = version
            __props__.__dict__["working_directory"] = working_directory
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["components[*]", "containerType", "description", "dockerfileTemplateData", "dockerfileTemplateUri", "imageOsVersionOverride", "instanceConfiguration", "kmsKeyId", "name", "parentImage", "platformOverride", "targetRepository", "version", "workingDirectory"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ContainerRecipe, __self__).__init__(
            'aws-native:imagebuilder:ContainerRecipe',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ContainerRecipe':
        """
        Get an existing ContainerRecipe resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ContainerRecipeArgs.__new__(ContainerRecipeArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["components"] = None
        __props__.__dict__["container_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["dockerfile_template_data"] = None
        __props__.__dict__["dockerfile_template_uri"] = None
        __props__.__dict__["image_os_version_override"] = None
        __props__.__dict__["instance_configuration"] = None
        __props__.__dict__["kms_key_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parent_image"] = None
        __props__.__dict__["platform_override"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_repository"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["working_directory"] = None
        return ContainerRecipe(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the container recipe.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def components(self) -> pulumi.Output[Optional[Sequence['outputs.ContainerRecipeComponentConfiguration']]]:
        """
        Components for build and test that are included in the container recipe.
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="containerType")
    def container_type(self) -> pulumi.Output[Optional['ContainerRecipeContainerType']]:
        """
        Specifies the type of container, such as Docker.
        """
        return pulumi.get(self, "container_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the container recipe.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dockerfileTemplateData")
    def dockerfile_template_data(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
        """
        return pulumi.get(self, "dockerfile_template_data")

    @property
    @pulumi.getter(name="dockerfileTemplateUri")
    def dockerfile_template_uri(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The S3 URI for the Dockerfile that will be used to build your container image.
        """
        return pulumi.get(self, "dockerfile_template_uri")

    @property
    @pulumi.getter(name="imageOsVersionOverride")
    def image_os_version_override(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the operating system version for the source image.
        """
        return pulumi.get(self, "image_os_version_override")

    @property
    @pulumi.getter(name="instanceConfiguration")
    def instance_configuration(self) -> pulumi.Output[Optional['outputs.ContainerRecipeInstanceConfiguration']]:
        """
        A group of options that can be used to configure an instance for building and testing container images.
        """
        return pulumi.get(self, "instance_configuration")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Identifies which KMS key is used to encrypt the container image.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the container recipe.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentImage")
    def parent_image(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The source image for the container recipe.
        """
        return pulumi.get(self, "parent_image")

    @property
    @pulumi.getter(name="platformOverride")
    def platform_override(self) -> pulumi.Output[Optional['ContainerRecipePlatformOverride']]:
        """
        Specifies the operating system platform when you use a custom source image.
        """
        return pulumi.get(self, "platform_override")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Tags that are attached to the container recipe.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetRepository")
    def target_repository(self) -> pulumi.Output[Optional['outputs.ContainerRecipeTargetContainerRepository']]:
        """
        The destination repository for the container image.
        """
        return pulumi.get(self, "target_repository")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The semantic version of the container recipe (<major>.<minor>.<patch>).
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The working directory to be used during build and test workflows.
        """
        return pulumi.get(self, "working_directory")

