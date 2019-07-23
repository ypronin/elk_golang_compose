package handlers

import "github.com/sirupsen/logrus"

type Handlers struct {
	logger *logrus.Entry
}

func NewHandlers(logger *logrus.Entry) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
