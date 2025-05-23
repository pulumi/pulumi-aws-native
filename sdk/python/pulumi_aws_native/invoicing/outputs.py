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

__all__ = [
    'InvoiceUnitRule',
]

@pulumi.output_type
class InvoiceUnitRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "linkedAccounts":
            suggest = "linked_accounts"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InvoiceUnitRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InvoiceUnitRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InvoiceUnitRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 linked_accounts: Sequence[builtins.str]):
        """
        :param Sequence[builtins.str] linked_accounts: The list of `LINKED_ACCOUNT` IDs where charges are included within the invoice unit.
        """
        pulumi.set(__self__, "linked_accounts", linked_accounts)

    @property
    @pulumi.getter(name="linkedAccounts")
    def linked_accounts(self) -> Sequence[builtins.str]:
        """
        The list of `LINKED_ACCOUNT` IDs where charges are included within the invoice unit.
        """
        return pulumi.get(self, "linked_accounts")


