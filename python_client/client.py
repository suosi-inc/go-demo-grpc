import grpc

from python_client.pb.user_pb2 import SearchUserRequest, GetUserRequest
from python_client.pb.user_pb2_grpc import UserServiceStub


class Auth(grpc.AuthMetadataPlugin):
    def __init__(self, app_id, app_key):
        self.app_id = app_id
        self.app_key = app_key

    def __call__(self, context, callback):
        callback((('app-id', self.app_id), ('app-key', self.app_key)), None)


def run():
    composite_credentials = grpc.composite_channel_credentials(
        # 此方式可免去安全证书校验, 但服务端客户端需要保持一致
        grpc.local_channel_credentials(),
        grpc.metadata_call_credentials(Auth("appId", "appKey")),
    )
    channel = grpc.secure_channel('127.0.0.1:50001', composite_credentials)

    svc = UserServiceStub(channel)

    # 列表中所有名称
    response = svc.Search(SearchUserRequest())
    print([i.name for i in response.list])

    response = svc.Get(GetUserRequest(id=1))
    print(response.data.menuList)


if __name__ == '__main__':
    run()
