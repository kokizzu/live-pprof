package metrics

import (
	"context"

	"github.com/moderato-app/live-pprof/api"
	"github.com/moderato-app/live-pprof/internal"
	"github.com/moderato-app/live-pprof/internal/logging"
)

// MockMetricsServer returns mock data instead of real data to save development time
type MockMetricsServer struct {
	api.UnimplementedMockMetricsServer

	mockResources *MockAssets
}

func NewMockMetricsServer() *MockMetricsServer {
	return &MockMetricsServer{
		mockResources: newMockAssets(),
	}
}

func (m *MockMetricsServer) HeapMetrics(_ context.Context, req *api.GoMetricsRequest) (*api.GoMetricsResponse, error) {
	logging.Sugar.Debug("HeapMetrics req:", req)
	return m.dispatch(req, internal.MetricsTypeHeap)
}

func (m *MockMetricsServer) CPUMetrics(_ context.Context, req *api.GoMetricsRequest) (*api.GoMetricsResponse, error) {
	logging.Sugar.Debug("CPUMetrics req:", req)
	return m.dispatch(req, internal.MetricsTypeCPU)
}

func (m *MockMetricsServer) AllocsMetrics(_ context.Context, req *api.GoMetricsRequest) (*api.GoMetricsResponse, error) {
	logging.Sugar.Debug("AllocsMetrics req:", req)
	return m.dispatch(req, internal.MetricsTypeAllocs)
}

func (m *MockMetricsServer) GoroutineMetrics(_ context.Context, req *api.GoMetricsRequest) (*api.GoMetricsResponse, error) {
	logging.Sugar.Debug("GoroutineMetrics req:", req)
	return m.dispatch(req, internal.MetricsTypeGoroutine)
}

func (m *MockMetricsServer) dispatch(req *api.GoMetricsRequest, mt internal.MetricsType) (*api.GoMetricsResponse, error) {
	_, err := internal.MetricsURL(req.Url, mt, false)
	if err != nil {
		return nil, err
	}

	mtr, err := m.mockResources.GetMetrics(mt)
	if err != nil {
		logging.Sugar.Error(err)
		return nil, err
	}
	resp := toResp(mtr)

	return resp, nil
}
