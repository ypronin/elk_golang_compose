package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/elk_golang_compose/pkg/utils"
	"github.com/elk_golang_compose/swagger/models"
	"github.com/elk_golang_compose/swagger/restapi/operations"
)

func init() {
	RegisterHandlerSetterFabric(GetJSONSchemeHandlerSetter)
}

func GetJSONSchemeHandlerSetter(h *Handlers) HandlerSetter {
	return func(api *operations.RelaxAPI) {
		api.GetJSONSchemaJSONHandler = operations.GetJSONSchemaJSONHandlerFunc(h.GetJSONSchemaHandler)
	}
}

func (h *Handlers) GetJSONSchemaHandler(
	params operations.GetJSONSchemaJSONParams,
) middleware.Responder {
	return operations.NewGetJSONSchemaJSONOK().WithPayload(&models.PokerTransaction{
		TransactionID:  utils.GenerateRandomString(10),
		PlayerID:       utils.GenerateRandomExternalPlayerID(),
		PlayerIP:       utils.GenerateRandomString(10),
		CurrencyID:     "EUR",
		PlayableBefore: 10000,
		Playable:       90000,
	})
}
