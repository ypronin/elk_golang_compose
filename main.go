package main

import (
	"io"
	"os"

	"github.com/elk_golang_compose/swagger/restapi/operations"

	"github.com/go-openapi/loads"

	apphandlers "github.com/elk_golang_compose/pkg/handlers"
	"github.com/elk_golang_compose/pkg/logging"
	"github.com/elk_golang_compose/pkg/utils"
	"github.com/elk_golang_compose/swagger/restapi"
	"github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	elogrus "gopkg.in/sohlich/elogrus.v3"
)

type Fields = logrus.Fields

const (
	serviceName = "app"
)

type ContextHook struct {
	Data Fields
}

func main() {

	logger := logging.Init(serviceName)

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		logger.Panic(err)
	}
	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "pokerlog")
	if err != nil {
		logger.Panic(err)
	}
	logger.Logger.Hooks.Add(hook)

	f, err := os.OpenFile(utils.LOG_PATH, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	mw := io.MultiWriter(os.Stdout, f)
	logger.Logger.SetOutput(mw)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Info("Hello world")

	api := operations.NewRelaxAPI(swaggerSpec)
	handlers := apphandlers.NewHandlers()
	apphandlers.RegisteredHandlerSetter(handlers).BindTo(api)

	server := restapi.NewServer(api)
	server.Port = 8000
	server.EnabledListeners = []string{"http"}

	if err := server.Serve(); err != nil {
		f.Close()
		logger.Fatalln(err)
	}

}
