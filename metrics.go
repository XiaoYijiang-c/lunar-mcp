package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Metrics holds server metrics
type Metrics struct {
	TotalRequests   int64
	SuccessRequests int64
	FailedRequests int64
	ActiveConns    int64
	ToolsCalled    map[string]*int64
	mu            struct{ sync.RWMutex }
}

// ServerMetrics global metrics
var ServerMetrics = &Metrics{}

func init() {
	ServerMetrics.ToolsCalled = make(map[string]*int64)
}

var startTime = time.Now()

// IncrementRequest increments total requests
func (m *Metrics) IncrementRequest() {
	atomic.AddInt64(&m.TotalRequests, 1)
}

// IncrementSuccess increments success requests
func (m *Metrics) IncrementSuccess() {
	atomic.AddInt64(&m.SuccessRequests, 1)
}

// IncrementFailed increments failed requests
func (m *Metrics) IncrementFailed() {
	atomic.AddInt64(&m.FailedRequests, 1)
}

// IncrementTool increments tool call count
func (m *Metrics) IncrementTool(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if _, ok := m.ToolsCalled[name]; !ok {
		m.ToolsCalled[name] = new(int64)
	}
	atomic.AddInt64(m.ToolsCalled[name], 1)
}

// GetMetrics returns current metrics
func (m *Metrics) GetMetrics() map[string]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	tools := make(map[string]int64)
	for k, v := range m.ToolsCalled {
		tools[k] = atomic.LoadInt64(v)
	}
	
	return map[string]interface{}{
		"totalRequests":   atomic.LoadInt64(&m.TotalRequests),
		"successRequests": atomic.LoadInt64(&m.SuccessRequests),
		"failedRequests":  atomic.LoadInt64(&m.FailedRequests),
		"activeConns":    atomic.LoadInt64(&m.ActiveConns),
		"toolsCalled":    tools,
		"uptime":         time.Since(startTime).Seconds(),
	}
}

// SimpleLogger logs messages
type SimpleLogger struct{}

func (l *SimpleLogger) Log(event string, data interface{}) {
	log.Printf("[%s] %v", event, data)
}
