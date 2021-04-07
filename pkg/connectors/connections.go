package connectors

import "github.com/microlib/logger/pkg/multi"

// Connections struct - all backend connections in a common object
type Connectors struct {
	Logger *multi.Logger
}

func NewClientConnections(logger *multi.Logger) Clients {
	return &Connectors{Logger: logger}
}

func (c *Connectors) Error(msg string, val ...interface{}) {
	c.Logger.Errorf(msg, val...)
}

func (c *Connectors) Info(msg string, val ...interface{}) {
	c.Logger.Infof(msg, val...)
}

func (c *Connectors) Debug(msg string, val ...interface{}) {
	c.Logger.Debugf(msg, val...)
}

func (c *Connectors) Trace(msg string, val ...interface{}) {
	c.Logger.Tracef(msg, val...)
}
