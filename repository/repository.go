package repository

import (
	"ConversorMoedasApp/adaptativemiddleware/cliente/proxies"
	"RAMid/distribution/clientproxy"
	"reflect"
)

func CheckRepository(proxy clientproxy.ClientProxy) interface{} {

	var clientProxy interface{}

	switch proxy.TypeName {

	case reflect.TypeOf(proxies.ConversorProxy{}).String():
		conversorProxy := proxies.NewConversorProxy()
		conversorProxy.Proxy.TypeName = proxy.TypeName
		conversorProxy.Proxy.Host = proxy.Host
		conversorProxy.Proxy.Port = proxy.Port
		clientProxy = conversorProxy
	}

	return clientProxy
}
