from google.protobuf import timestamp_pb2 as _timestamp_pb2
import role_pb2 as _role_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
DISABLE: UserStatus
ENABLE: UserStatus
UNKNOWN: UserStatus

class AddUserRequest(_message.Message):
    __slots__ = ["customConfig", "hasSuper", "menuList", "name", "roleId", "status", "weight"]
    class CustomConfigEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    CUSTOMCONFIG_FIELD_NUMBER: _ClassVar[int]
    HASSUPER_FIELD_NUMBER: _ClassVar[int]
    MENULIST_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ROLEID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    customConfig: _containers.ScalarMap[str, str]
    hasSuper: bool
    menuList: _containers.RepeatedScalarFieldContainer[str]
    name: str
    roleId: int
    status: UserStatus
    weight: float
    def __init__(self, name: _Optional[str] = ..., weight: _Optional[float] = ..., hasSuper: bool = ..., status: _Optional[_Union[UserStatus, str]] = ..., roleId: _Optional[int] = ..., menuList: _Optional[_Iterable[str]] = ..., customConfig: _Optional[_Mapping[str, str]] = ...) -> None: ...

class AddUserResponse(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class EditUserRequest(_message.Message):
    __slots__ = ["customConfig", "hasSuper", "id", "menuList", "name", "roleId", "status", "weight"]
    class CustomConfigEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    CUSTOMCONFIG_FIELD_NUMBER: _ClassVar[int]
    HASSUPER_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    MENULIST_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ROLEID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    customConfig: _containers.ScalarMap[str, str]
    hasSuper: bool
    id: int
    menuList: _containers.RepeatedScalarFieldContainer[str]
    name: str
    roleId: int
    status: UserStatus
    weight: float
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., weight: _Optional[float] = ..., hasSuper: bool = ..., status: _Optional[_Union[UserStatus, str]] = ..., roleId: _Optional[int] = ..., menuList: _Optional[_Iterable[str]] = ..., customConfig: _Optional[_Mapping[str, str]] = ...) -> None: ...

class EditUserResponse(_message.Message):
    __slots__ = ["success"]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class GetUserRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class GetUserResponse(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: User
    def __init__(self, data: _Optional[_Union[User, _Mapping]] = ...) -> None: ...

class RemoveUserRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class RemoveUserResponse(_message.Message):
    __slots__ = ["success"]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class SearchUserRequest(_message.Message):
    __slots__ = ["createdTimeRange", "menuList", "name", "page", "size", "status"]
    class CreatedTimeRange(_message.Message):
        __slots__ = ["end", "start"]
        END_FIELD_NUMBER: _ClassVar[int]
        START_FIELD_NUMBER: _ClassVar[int]
        end: _timestamp_pb2.Timestamp
        start: _timestamp_pb2.Timestamp
        def __init__(self, start: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    CREATEDTIMERANGE_FIELD_NUMBER: _ClassVar[int]
    MENULIST_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PAGE_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    createdTimeRange: SearchUserRequest.CreatedTimeRange
    menuList: _containers.RepeatedScalarFieldContainer[str]
    name: str
    page: int
    size: int
    status: UserStatus
    def __init__(self, name: _Optional[str] = ..., status: _Optional[_Union[UserStatus, str]] = ..., createdTimeRange: _Optional[_Union[SearchUserRequest.CreatedTimeRange, _Mapping]] = ..., menuList: _Optional[_Iterable[str]] = ..., page: _Optional[int] = ..., size: _Optional[int] = ...) -> None: ...

class SearchUserResponse(_message.Message):
    __slots__ = ["list", "total"]
    LIST_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    list: _containers.RepeatedCompositeFieldContainer[User]
    total: int
    def __init__(self, total: _Optional[int] = ..., list: _Optional[_Iterable[_Union[User, _Mapping]]] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["createdBy", "createdTime", "customConfig", "hasSuper", "id", "menuList", "name", "role", "status", "weight"]
    class CustomConfigEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    CREATEDBY_FIELD_NUMBER: _ClassVar[int]
    CREATEDTIME_FIELD_NUMBER: _ClassVar[int]
    CUSTOMCONFIG_FIELD_NUMBER: _ClassVar[int]
    HASSUPER_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    MENULIST_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    createdBy: User
    createdTime: _timestamp_pb2.Timestamp
    customConfig: _containers.ScalarMap[str, str]
    hasSuper: bool
    id: int
    menuList: _containers.RepeatedScalarFieldContainer[str]
    name: str
    role: _role_pb2.Role
    status: UserStatus
    weight: float
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., weight: _Optional[float] = ..., hasSuper: bool = ..., status: _Optional[_Union[UserStatus, str]] = ..., role: _Optional[_Union[_role_pb2.Role, _Mapping]] = ..., createdBy: _Optional[_Union[User, _Mapping]] = ..., menuList: _Optional[_Iterable[str]] = ..., customConfig: _Optional[_Mapping[str, str]] = ..., createdTime: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class UserStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
