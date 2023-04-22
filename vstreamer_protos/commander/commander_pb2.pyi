from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
PAUSE: Operation
PLAYBACK: Operation
RELOAD: Operation
RESUME: Operation
SET_FILTERS: Operation
SPEECH: Operation
SUBTITLE: Operation
SUBTITLE_TRANSLATED: Operation
TRANSCRIBE: Operation
TRANSLATE: Operation

class Command(_message.Message):
    __slots__ = ["file_path", "filters", "operations", "sound", "text"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    OPERATIONS_FIELD_NUMBER: _ClassVar[int]
    SOUND_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    filters: _containers.RepeatedScalarFieldContainer[str]
    operations: _containers.RepeatedScalarFieldContainer[Operation]
    sound: Sound
    text: str
    def __init__(self, operations: _Optional[_Iterable[_Union[Operation, str]]] = ..., sound: _Optional[_Union[Sound, _Mapping]] = ..., text: _Optional[str] = ..., file_path: _Optional[str] = ..., filters: _Optional[_Iterable[str]] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ["result"]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: bool
    def __init__(self, result: bool = ...) -> None: ...

class Sound(_message.Message):
    __slots__ = ["data", "rate"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    RATE_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    rate: int
    def __init__(self, data: _Optional[bytes] = ..., rate: _Optional[int] = ...) -> None: ...

class Operation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
