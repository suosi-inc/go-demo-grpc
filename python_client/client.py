import time

import grpc

from python_client.pb.user_pb2 import SearchUserRequest, GetUserRequest
from python_client.pb.user_pb2_grpc import UserServiceStub

# 配置
HOST = '127.0.0.1'
PORT = '50001'
APP_ID = 'appId'
APP_KEY = 'appKey'


# 时间拦截器
class Prof(grpc.UnaryUnaryClientInterceptor):
    def intercept_unary_unary(self, continuation, client_call_details, request):
        print("时间拦截器开始")

        start = time.time_ns()
        # 没有这个调用，后面的拦截器以及真正的rpc函数都不会执行了
        response = continuation(client_call_details, request)
        end = time.time_ns()

        print("耗时 %d 微秒" % ((end - start) / 1e3))
        print("时间拦截器结束")
        return response


# 其他拦截器
class Other(grpc.UnaryUnaryClientInterceptor):
    def intercept_unary_unary(self, continuation, client_call_details, request):
        print("其他拦截器开始")

        response = continuation(client_call_details, request)

        print("其他拦截器结束")
        return response


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
        # 认证授权信息
        grpc.metadata_call_credentials(Auth(APP_ID, APP_KEY)),
    )
    channel = grpc.secure_channel('%s:%s' % (HOST, PORT), composite_credentials)
    channel = grpc.intercept_channel(channel, Prof(), Other())

    svc = UserServiceStub(channel)

    # 打印列表中所有名称
    response = svc.Search(SearchUserRequest())
    print([i.name for i in response.list])

    # 打印用户拥有的菜单
    response = svc.Get(GetUserRequest(id=1))
    print(response.data.menuList)


if __name__ == '__main__':
    run()
