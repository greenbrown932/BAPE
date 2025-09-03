// internal/plugin/plugin.go
// DEPRECATED: This file is deprecated and will be removed. The plugin interface has been moved to pkg/plugin/plugin.go.
package plugin

import (
	"context"
	"time"
)

type AnalysisPlugin interface {
	Name() string
	Version() string
	Initialize() error
	Analyze(ctx context.Context, input AnalysisInput) (*AnalysisResult, error)
	Cleanup() error
	HealthCheck() error
}

type AnalysisInput struct {
	Binary   []byte
	Metadata map[string]interface{}
	Config   map[string]interface{}
	WorkDir  string
	Timeout  time.Duration
}

type AnalysisResult struct {
	Success   bool
	Data      map[string]interface{}
	Artifacts []string
	Metrics   AnalysisMetrics
	Error     error
}

type AnalysisMetrics struct {
	ExecutionTime time.Duration
	MemoryUsage   int64
}
