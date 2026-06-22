from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Operation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TRANSCRIBE: _ClassVar[Operation]
    TRANSLATE: _ClassVar[Operation]
    SUBTITLE: _ClassVar[Operation]
    TTS: _ClassVar[Operation]
    VC: _ClassVar[Operation]
    PLAYBACK: _ClassVar[Operation]
    PAUSE: _ClassVar[Operation]
    RESUME: _ClassVar[Operation]
    RELOAD: _ClassVar[Operation]
    SET_FILTERS: _ClassVar[Operation]
    PING: _ClassVar[Operation]
    FORWARD: _ClassVar[Operation]
TRANSCRIBE: Operation
TRANSLATE: Operation
SUBTITLE: Operation
TTS: Operation
VC: Operation
PLAYBACK: Operation
PAUSE: Operation
RESUME: Operation
RELOAD: Operation
SET_FILTERS: Operation
PING: Operation
FORWARD: Operation

class Command(_message.Message):
    __slots__ = ("chains", "operand")
    CHAINS_FIELD_NUMBER: _ClassVar[int]
    OPERAND_FIELD_NUMBER: _ClassVar[int]
    chains: _containers.RepeatedCompositeFieldContainer[OperationChain]
    operand: Operand
    def __init__(self, chains: _Optional[_Iterable[_Union[OperationChain, _Mapping]]] = ..., operand: _Optional[_Union[Operand, _Mapping]] = ...) -> None: ...

class Operand(_message.Message):
    __slots__ = ("sound", "text", "file_path", "filters", "trace_id", "origin_ts")
    SOUND_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TRACE_ID_FIELD_NUMBER: _ClassVar[int]
    ORIGIN_TS_FIELD_NUMBER: _ClassVar[int]
    sound: Sound
    text: str
    file_path: str
    filters: _containers.RepeatedScalarFieldContainer[str]
    trace_id: str
    origin_ts: float
    def __init__(self, sound: _Optional[_Union[Sound, _Mapping]] = ..., text: _Optional[str] = ..., file_path: _Optional[str] = ..., filters: _Optional[_Iterable[str]] = ..., trace_id: _Optional[str] = ..., origin_ts: _Optional[float] = ...) -> None: ...

class Sound(_message.Message):
    __slots__ = ("data", "rate", "format", "channels")
    DATA_FIELD_NUMBER: _ClassVar[int]
    RATE_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    CHANNELS_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    rate: int
    format: int
    channels: int
    def __init__(self, data: _Optional[bytes] = ..., rate: _Optional[int] = ..., format: _Optional[int] = ..., channels: _Optional[int] = ...) -> None: ...

class OperationChain(_message.Message):
    __slots__ = ("operations",)
    OPERATIONS_FIELD_NUMBER: _ClassVar[int]
    operations: _containers.RepeatedCompositeFieldContainer[OperationRoute]
    def __init__(self, operations: _Optional[_Iterable[_Union[OperationRoute, _Mapping]]] = ...) -> None: ...

class OperationRoute(_message.Message):
    __slots__ = ("operation", "remote", "queries")
    class QueriesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    OPERATION_FIELD_NUMBER: _ClassVar[int]
    REMOTE_FIELD_NUMBER: _ClassVar[int]
    QUERIES_FIELD_NUMBER: _ClassVar[int]
    operation: Operation
    remote: str
    queries: _containers.ScalarMap[str, str]
    def __init__(self, operation: _Optional[_Union[Operation, str]] = ..., remote: _Optional[str] = ..., queries: _Optional[_Mapping[str, str]] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ("result",)
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: bool
    def __init__(self, result: bool = ...) -> None: ...
