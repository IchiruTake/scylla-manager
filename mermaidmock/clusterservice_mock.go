// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package mermaidmock is a generated GoMock package.
package mermaidmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/scylladb/mermaid/cluster"
	uuid "github.com/scylladb/mermaid/uuid"
	reflect "reflect"
)

// MockClusterService is a mock of ClusterService interface
type MockClusterService struct {
	ctrl     *gomock.Controller
	recorder *MockClusterServiceMockRecorder
}

// MockClusterServiceMockRecorder is the mock recorder for MockClusterService
type MockClusterServiceMockRecorder struct {
	mock *MockClusterService
}

// NewMockClusterService creates a new mock instance
func NewMockClusterService(ctrl *gomock.Controller) *MockClusterService {
	mock := &MockClusterService{ctrl: ctrl}
	mock.recorder = &MockClusterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterService) EXPECT() *MockClusterServiceMockRecorder {
	return m.recorder
}

// ListClusters mocks base method
func (m *MockClusterService) ListClusters(ctx context.Context, f *cluster.Filter) ([]*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "ListClusters", ctx, f)
	ret0, _ := ret[0].([]*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters
func (mr *MockClusterServiceMockRecorder) ListClusters(ctx, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockClusterService)(nil).ListClusters), ctx, f)
}

// GetCluster mocks base method
func (m *MockClusterService) GetCluster(ctx context.Context, idOrName string) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "GetCluster", ctx, idOrName)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockClusterServiceMockRecorder) GetCluster(ctx, idOrName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterService)(nil).GetCluster), ctx, idOrName)
}

// PutCluster mocks base method
func (m *MockClusterService) PutCluster(ctx context.Context, c *cluster.Cluster) error {
	ret := m.ctrl.Call(m, "PutCluster", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCluster indicates an expected call of PutCluster
func (mr *MockClusterServiceMockRecorder) PutCluster(ctx, c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCluster", reflect.TypeOf((*MockClusterService)(nil).PutCluster), ctx, c)
}

// DeleteCluster mocks base method
func (m *MockClusterService) DeleteCluster(ctx context.Context, id uuid.UUID) error {
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockClusterServiceMockRecorder) DeleteCluster(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterService)(nil).DeleteCluster), ctx, id)
}
