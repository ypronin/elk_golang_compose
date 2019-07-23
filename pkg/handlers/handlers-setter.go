package handlers

import (
	"github.com/elk_golang_compose/swagger/restapi/operations"
)

type (
	HandlerSetter               func(*operations.RelaxAPI)
	HandlerSetterFabric         func(*Handlers) HandlerSetter
	HandlerSetterFabricRegister []HandlerSetterFabric
)

func (s HandlerSetter) BindTo(api *operations.RelaxAPI) {
	s(api)
}

var handlerSetterFabricRegister []HandlerSetterFabric

func RegisterHandlerSetterFabric(f HandlerSetterFabric) {
	handlerSetterFabricRegister = append(handlerSetterFabricRegister, f)
}

func RegisteredHandlerSetter(h *Handlers) HandlerSetter {
	return func(api *operations.RelaxAPI) {
		for _, handlerSetterFabric := range handlerSetterFabricRegister {
			handlerSetterFabric(h).BindTo(api)
		}
	}
}
