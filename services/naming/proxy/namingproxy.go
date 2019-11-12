package proxy

import (
	"RAMid/aux"
	"RAMid/distribution/clientproxy"
	"RAMid/plugins"
	"RAMid/repository"
	"RAMid/util"
	"plugin"
)

type NamingProxy struct{}

func (NamingProxy) Register(p1 string, proxy interface{}) bool {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = proxy
	namingproxy := clientproxy.ClientProxy{Host: util.NAMING_HOST, Port: util.NAMING_PORT, Id: 0}
	request := aux.Request{Op: "Register", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	//Carrega o arquivo do componente
	manager := plugins.Manager{}
	componente, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_REQUESTOR))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Transmitir")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Transmitir := funcao.(func(chan interface{}))

	// invoke requestor
	ch := make(chan interface{})
	go Transmitir(ch)
	ch <- inv

	retorno := <-ch
	ter := retorno.([]interface{})

	return ter[0].(bool)
}

func (NamingProxy) Lookup(p1 string) interface{} {

	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	namingproxy := clientproxy.ClientProxy{Host: util.NAMING_HOST, Port: util.NAMING_PORT, Id: 0}
	request := aux.Request{Op: "Lookup", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	//Carrega o arquivo do componente
	manager := plugins.Manager{}
	componente, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_REQUESTOR))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Transmitir")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Transmitir := funcao.(func(chan interface{}))

	// invoke requestor
	ch := make(chan interface{})
	go Transmitir(ch)
	ch <- inv

	retorno := <-ch
	ter := retorno.([]interface{})

	// process reply
	proxyTemp := ter[0].(map[string]interface{})
	clientProxyTemp := clientproxy.ClientProxy{TypeName: proxyTemp["TypeName"].(string), Host: proxyTemp["Host"].(string), Port: int(proxyTemp["Port"].(float64))}
	clientProxy := repository.CheckRepository(clientProxyTemp)

	return clientProxy
}
