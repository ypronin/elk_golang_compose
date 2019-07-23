package handlers

import (
	"github.com/elk_golang_compose/swagger/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func init() {
	RegisterHandlerSetterFabric(GetHealthV1HandlerSetter)
}

func GetHealthV1HandlerSetter(h *Handlers) HandlerSetter {
	return func(api *operations.RelaxAPI) {
		api.GetHealthV1Handler = operations.GetHealthV1HandlerFunc(h.GetHealthV1)
	}
}

func (h *Handlers) GetHealthV1(
	params operations.GetHealthV1Params,
) middleware.Responder {

	h.logger.Info("GetHealthV1")
	return operations.NewGetHealthV1OK().WithPayload("ok")
}
