from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
PAUSE: Operation
PING: Operation
PLAYBACK: Operation
RELOAD: Operation
RESUME: Operation
SET_FILTERS: Operation
SUBTITLE: Operation
TRANSCRIBE: Operation
TRANSLATE: Operation
TTS: Operation
VC: Operation

class Command(_message.Message):
    __slots__ = ["chains", "operand"]
    CHAINS_FIELD_NUMBER: _ClassVar[int]
    OPERAND_FIELD_NUMBER: _ClassVar[int]
    chains: _containers.RepeatedCompositeFieldContainer[OperationChain]
    operand: Operand
    def __init__(self, chains: _Optional[_Iterable[_Union[OperationChain, _Mapping]]] = ..., operand: _Optional[_Union[Operand, _Mapping]] = ...) -> None: ...

class Operand(_message.Message):
    __slots__ = ["file_path", "filters", "sound", "text"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    SOUND_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    filters: _containers.RepeatedScalarFieldContainer[str]
    sound: Sound
    text: str
    def __init__(self, sound: _Optional[_Union[Sound, _Mapping]] = ..., text: _Optional[str] = ..., file_path: _Optional[str] = ..., filters: _Optional[_Iterable[str]] = ...) -> None: ...

class OperationChain(_message.Message):
    __slots__ = ["operations"]
    OPERATIONS_FIELD_NUMBER: _ClassVar[int]
    operations: _containers.RepeatedCompositeFieldContainer[OperationRoute]
    def __init__(self, operations: _Optional[_Iterable[_Union[OperationRoute, _Mapping]]] = ...) -> None: ...

class OperationRoute(_message.Message):
    __slots__ = ["operation", "queries", "remote"]
    class QueriesEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    OPERATION_FIELD_NUMBER: _ClassVar[int]
    QUERIES_FIELD_NUMBER: _ClassVar[int]
    REMOTE_FIELD_NUMBER: _ClassVar[int]
    operation: Operation
    queries: _containers.ScalarMap[str, str]
    remote: str
    def __init__(self, operation: _Optional[_Union[Operation, str]] = ..., remote: _Optional[str] = ..., queries: _Optional[_Mapping[str, str]] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ["result"]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: bool
    def __init__(self, result: bool = ...) -> None: ...

class Sound(_message.Message):
    __slots__ = ["channels", "data", "format", "rate"]
    CHANNELS_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    RATE_FIELD_NUMBER: _ClassVar[int]
    channels: int
    data: bytes
    format: int
    rate: int
    def __init__(self, data: _Optional[bytes] = ..., rate: _Optional[int] = ..., format: _Optional[int] = ..., channels: _Optional[int] = ...) -> None: ...

class Operation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
