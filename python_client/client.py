import grpc

from python_client.pb.user_pb2 import SearchUserRequest, GetUserRequest
from python_client.pb.user_pb2_grpc import UserServiceStub


def run():
    with grpc.insecure_channel('localhost:50001') as channel:
        svc = UserServiceStub(channel)

        # 列表中所有名称
        response = svc.Search(SearchUserRequest())
        print([i.name for i in response.list])

        response = svc.Get(GetUserRequest(id=1))
        print(response.data.menuList)


if __name__ == '__main__':
    run()
