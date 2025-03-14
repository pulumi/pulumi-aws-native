# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DataProtectionSettingsRedactionPlaceHolderType',
    'IdentityProviderType',
    'PortalAuthenticationType',
    'PortalBrowserType',
    'PortalInstanceType',
    'PortalRendererType',
    'PortalStatus',
    'UserSettingsEnabledType',
    'UserSettingsMaxDisplayResolution',
    'UserSettingsToolbarItem',
    'UserSettingsToolbarType',
    'UserSettingsVisualMode',
]


class DataProtectionSettingsRedactionPlaceHolderType(str, Enum):
    CUSTOM_TEXT = "CustomText"


class IdentityProviderType(str, Enum):
    SAML = "SAML"
    FACEBOOK = "Facebook"
    GOOGLE = "Google"
    LOGIN_WITH_AMAZON = "LoginWithAmazon"
    SIGN_IN_WITH_APPLE = "SignInWithApple"
    OIDC = "OIDC"


class PortalAuthenticationType(str, Enum):
    STANDARD = "Standard"
    IAM_IDENTITY_CENTER = "IAM_Identity_Center"


class PortalBrowserType(str, Enum):
    CHROME = "Chrome"


class PortalInstanceType(str, Enum):
    STANDARD_REGULAR = "standard.regular"
    STANDARD_LARGE = "standard.large"
    STANDARD_XLARGE = "standard.xlarge"


class PortalRendererType(str, Enum):
    APP_STREAM = "AppStream"


class PortalStatus(str, Enum):
    INCOMPLETE = "Incomplete"
    PENDING = "Pending"
    ACTIVE = "Active"


class UserSettingsEnabledType(str, Enum):
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class UserSettingsMaxDisplayResolution(str, Enum):
    SIZE4096X2160 = "size4096X2160"
    SIZE3840X2160 = "size3840X2160"
    SIZE3440X1440 = "size3440X1440"
    SIZE2560X1440 = "size2560X1440"
    SIZE1920X1080 = "size1920X1080"
    SIZE1280X720 = "size1280X720"
    SIZE1024X768 = "size1024X768"
    SIZE800X600 = "size800X600"


class UserSettingsToolbarItem(str, Enum):
    WINDOWS = "Windows"
    DUAL_MONITOR = "DualMonitor"
    FULL_SCREEN = "FullScreen"
    WEBCAM = "Webcam"
    MICROPHONE = "Microphone"


class UserSettingsToolbarType(str, Enum):
    FLOATING = "Floating"
    DOCKED = "Docked"


class UserSettingsVisualMode(str, Enum):
    DARK = "Dark"
    LIGHT = "Light"
