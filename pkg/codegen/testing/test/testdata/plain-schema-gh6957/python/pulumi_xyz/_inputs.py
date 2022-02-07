# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'FooArgs',
]

@pulumi.input_type
class FooArgs:
    def __init__(__self__, *,
                 a: Optional[pulumi.Input[bool]] = None):
        if a is not None:
            pulumi.set(__self__, "a", a)

    @property
    @pulumi.getter
    def a(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "a")

    @a.setter
    def a(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "a", value)

