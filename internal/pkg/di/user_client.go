package di

import (
	"sync"

	"github.com/suosi-inc/go-demo/grpc/protobuf"
)

const (
	DefaultUserClientName = "defaultUserClient"
)

var UserClientDi = &userClientDi{
	store: make(map[string]protobuf.UserServiceClient),
	mutex: sync.RWMutex{},
}

type userClientDi struct {
	store map[string]protobuf.UserServiceClient
	mutex sync.RWMutex
}

func (d *userClientDi) SetUserClientDi(name string, v protobuf.UserServiceClient) {
	d.mutex.Lock()
	d.store[name] = v
	d.mutex.Unlock()
}

func (d *userClientDi) GetUserClientDi(name string) protobuf.UserServiceClient {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.store[name]
}

func GetUserClientWithName(name string) protobuf.UserServiceClient {
	return UserClientDi.GetUserClientDi(name)
}

func SetUserClientWithName(name string, v protobuf.UserServiceClient) {
	UserClientDi.SetUserClientDi(name, v)
}

func GetUserClient() protobuf.UserServiceClient {
	return UserClientDi.GetUserClientDi(DefaultUserClientName)
}

func SetUserClient(v protobuf.UserServiceClient) {
	UserClientDi.SetUserClientDi(DefaultUserClientName, v)
}
