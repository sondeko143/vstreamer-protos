from typing import Any, AnyStr, Callable, Optional, Sequence, Tuple, Union
import grpc

from vstreamer_protos.commander import commander_pb2


class CommanderStub(object):
    def __init__(self, channel: grpc.Channel) -> None: ...


class CommanderServicer(object):
    def process_command(self, request: commander_pb2.Command, context: grpc.ServicerContext) -> commander_pb2.Response: ...


def add_CommanderServicer_to_server(servicer: CommanderServicer, server: grpc.Server) -> None: ...


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
    metadata: Optional[Sequence[Tuple[str, Union[str, bytes]]]] = None) -> commander_pb2.Response: ...
