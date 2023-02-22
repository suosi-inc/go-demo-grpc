# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import python_client.pb.user_pb2 as user__pb2


class UserServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Get = channel.unary_unary(
                '/protobuf.UserService/Get',
                request_serializer=user__pb2.GetUserRequest.SerializeToString,
                response_deserializer=user__pb2.GetUserResponse.FromString,
                )
        self.Search = channel.unary_unary(
                '/protobuf.UserService/Search',
                request_serializer=user__pb2.SearchUserRequest.SerializeToString,
                response_deserializer=user__pb2.SearchUserResponse.FromString,
                )
        self.Add = channel.unary_unary(
                '/protobuf.UserService/Add',
                request_serializer=user__pb2.AddUserRequest.SerializeToString,
                response_deserializer=user__pb2.AddUserResponse.FromString,
                )
        self.Edit = channel.unary_unary(
                '/protobuf.UserService/Edit',
                request_serializer=user__pb2.EditUserRequest.SerializeToString,
                response_deserializer=user__pb2.EditUserResponse.FromString,
                )
        self.Remove = channel.unary_unary(
                '/protobuf.UserService/Remove',
                request_serializer=user__pb2.RemoveUserRequest.SerializeToString,
                response_deserializer=user__pb2.RemoveUserResponse.FromString,
                )


class UserServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Search(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Add(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Edit(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Remove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=user__pb2.GetUserRequest.FromString,
                    response_serializer=user__pb2.GetUserResponse.SerializeToString,
            ),
            'Search': grpc.unary_unary_rpc_method_handler(
                    servicer.Search,
                    request_deserializer=user__pb2.SearchUserRequest.FromString,
                    response_serializer=user__pb2.SearchUserResponse.SerializeToString,
            ),
            'Add': grpc.unary_unary_rpc_method_handler(
                    servicer.Add,
                    request_deserializer=user__pb2.AddUserRequest.FromString,
                    response_serializer=user__pb2.AddUserResponse.SerializeToString,
            ),
            'Edit': grpc.unary_unary_rpc_method_handler(
                    servicer.Edit,
                    request_deserializer=user__pb2.EditUserRequest.FromString,
                    response_serializer=user__pb2.EditUserResponse.SerializeToString,
            ),
            'Remove': grpc.unary_unary_rpc_method_handler(
                    servicer.Remove,
                    request_deserializer=user__pb2.RemoveUserRequest.FromString,
                    response_serializer=user__pb2.RemoveUserResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'protobuf.UserService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UserService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/protobuf.UserService/Get',
            user__pb2.GetUserRequest.SerializeToString,
            user__pb2.GetUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Search(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/protobuf.UserService/Search',
            user__pb2.SearchUserRequest.SerializeToString,
            user__pb2.SearchUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Add(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/protobuf.UserService/Add',
            user__pb2.AddUserRequest.SerializeToString,
            user__pb2.AddUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Edit(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/protobuf.UserService/Edit',
            user__pb2.EditUserRequest.SerializeToString,
            user__pb2.EditUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Remove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/protobuf.UserService/Remove',
            user__pb2.RemoveUserRequest.SerializeToString,
            user__pb2.RemoveUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
