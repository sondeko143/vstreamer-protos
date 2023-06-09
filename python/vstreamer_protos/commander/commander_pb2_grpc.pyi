from typing import Any, AnyStr, Callable, Coroutine, Optional, Sequence, Tuple, Union
import grpc

from vstreamer_protos.commander import commander_pb2


class CommanderStub(object):
    process_command: grpc.UnaryUnaryMultiCallable
    sync_process_command: grpc.UnaryUnaryMultiCallable
    def __init__(self, channel: Union[grpc.Channel, grpc.aio.Channel]) -> None: ...



class CommanderServicer(object):
    def process_command(self, request: commander_pb2.Command, context: grpc.ServicerContext) -> Union[Coroutine[Any, Any, commander_pb2.Response], commander_pb2.Response]: ...
    def sync_process_command(self, request: commander_pb2.Command, context: grpc.ServicerContext) -> Union[Coroutine[Any, Any, commander_pb2.Response], commander_pb2.Response]: ...


def add_CommanderServicer_to_server(servicer: CommanderServicer, server: Union[grpc.Server, grpc.aio.Server]) -> None: ...


 # This class is part of an EXPERIMENTAL API.
class Commander(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def process_command(
        request: commander_pb2.Command,
        target: str,
        method: str,
        request_serializer: Optional[Callable[[Any], bytes]] = None,
        response_deserializer: Optional[Callable[[bytes], Any]] = None,
        options: Sequence[Tuple[AnyStr, AnyStr]] = (),
        channel_credentials: Optional[grpc.ChannelCredentials] = None,
        insecure: bool = False,
        call_credentials: Optional[grpc.CallCredentials] = None,
        compression: Optional[grpc.Compression] = None,
        wait_for_ready: Optional[bool] = None,
        timeout: Optional[float] = 60,
        metadata: Optional[Sequence[Tuple[str, Union[str, bytes]]]] = None) -> Coroutine[Any, Any, commander_pb2.Response]: ...

    @staticmethod
    def sync_process_command(
        request: commander_pb2.Command,
        target: str,
        method: str,
        request_serializer: Optional[Callable[[Any], bytes]] = None,
        response_deserializer: Optional[Callable[[bytes], Any]] = None,
        options: Sequence[Tuple[AnyStr, AnyStr]] = (),
        channel_credentials: Optional[grpc.ChannelCredentials] = None,
        insecure: bool = False,
        call_credentials: Optional[grpc.CallCredentials] = None,
        compression: Optional[grpc.Compression] = None,
        wait_for_ready: Optional[bool] = None,
        timeout: Optional[float] = 60,
        metadata: Optional[Sequence[Tuple[str, Union[str, bytes]]]] = None) -> Coroutine[Any, Any, commander_pb2.Response]: ...
