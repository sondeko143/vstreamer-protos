# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from vstreamer_protos.commander import commander_pb2 as vstreamer__protos_dot_commander_dot_commander__pb2


class CommanderStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.process_command = channel.unary_unary(
                '/voicerecog.Commander/process_command',
                request_serializer=vstreamer__protos_dot_commander_dot_commander__pb2.Command.SerializeToString,
                response_deserializer=vstreamer__protos_dot_commander_dot_commander__pb2.Response.FromString,
                )
        self.sync_process_command = channel.unary_unary(
                '/voicerecog.Commander/sync_process_command',
                request_serializer=vstreamer__protos_dot_commander_dot_commander__pb2.Command.SerializeToString,
                response_deserializer=vstreamer__protos_dot_commander_dot_commander__pb2.Response.FromString,
                )


class CommanderServicer(object):
    """Missing associated documentation comment in .proto file."""

    def process_command(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def sync_process_command(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CommanderServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'process_command': grpc.unary_unary_rpc_method_handler(
                    servicer.process_command,
                    request_deserializer=vstreamer__protos_dot_commander_dot_commander__pb2.Command.FromString,
                    response_serializer=vstreamer__protos_dot_commander_dot_commander__pb2.Response.SerializeToString,
            ),
            'sync_process_command': grpc.unary_unary_rpc_method_handler(
                    servicer.sync_process_command,
                    request_deserializer=vstreamer__protos_dot_commander_dot_commander__pb2.Command.FromString,
                    response_serializer=vstreamer__protos_dot_commander_dot_commander__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'voicerecog.Commander', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Commander(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def process_command(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/voicerecog.Commander/process_command',
            vstreamer__protos_dot_commander_dot_commander__pb2.Command.SerializeToString,
            vstreamer__protos_dot_commander_dot_commander__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def sync_process_command(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/voicerecog.Commander/sync_process_command',
            vstreamer__protos_dot_commander_dot_commander__pb2.Command.SerializeToString,
            vstreamer__protos_dot_commander_dot_commander__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
