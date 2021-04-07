package connectors

import (
	"fmt"

	"github.com/microlib/logger/pkg/multi"
)

// Mock all connections
type MockConnectors struct {
	Logger *multi.Logger
}

func (c *MockConnectors) Error(msg string, val ...interface{}) {
	c.Logger.Error(fmt.Sprintf(msg, val...))
}

func (c *MockConnectors) Info(msg string, val ...interface{}) {
	c.Logger.Info(fmt.Sprintf(msg, val...))
}

func (c *MockConnectors) Debug(msg string, val ...interface{}) {
	c.Logger.Debug(fmt.Sprintf(msg, val...))
}

func (c *MockConnectors) Trace(msg string, val ...interface{}) {
	c.Logger.Trace(fmt.Sprintf(msg, val...))
}

func NewTestConnectors(logger *multi.Logger) Clients {
	conns := &MockConnectors{Logger: logger}
	return conns
}
