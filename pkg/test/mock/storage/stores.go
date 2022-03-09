// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/storage/stores.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	core "github.com/rancher/opni-monitoring/pkg/core"
	keyring "github.com/rancher/opni-monitoring/pkg/keyring"
	storage "github.com/rancher/opni-monitoring/pkg/storage"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method.
func (m *MockBackend) CreateCluster(ctx context.Context, cluster *core.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster.
func (mr *MockBackendMockRecorder) CreateCluster(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockBackend)(nil).CreateCluster), ctx, cluster)
}

// CreateRole mocks base method.
func (m *MockBackend) CreateRole(arg0 context.Context, arg1 *core.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockBackendMockRecorder) CreateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockBackend)(nil).CreateRole), arg0, arg1)
}

// CreateRoleBinding mocks base method.
func (m *MockBackend) CreateRoleBinding(arg0 context.Context, arg1 *core.RoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoleBinding indicates an expected call of CreateRoleBinding.
func (mr *MockBackendMockRecorder) CreateRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoleBinding", reflect.TypeOf((*MockBackend)(nil).CreateRoleBinding), arg0, arg1)
}

// CreateToken mocks base method.
func (m *MockBackend) CreateToken(ctx context.Context, ttl time.Duration, opts ...storage.TokenCreateOption) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ttl}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateToken", varargs...)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockBackendMockRecorder) CreateToken(ctx, ttl interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ttl}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockBackend)(nil).CreateToken), varargs...)
}

// DeleteCluster mocks base method.
func (m *MockBackend) DeleteCluster(ctx context.Context, ref *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockBackendMockRecorder) DeleteCluster(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockBackend)(nil).DeleteCluster), ctx, ref)
}

// DeleteRole mocks base method.
func (m *MockBackend) DeleteRole(arg0 context.Context, arg1 *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockBackendMockRecorder) DeleteRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockBackend)(nil).DeleteRole), arg0, arg1)
}

// DeleteRoleBinding mocks base method.
func (m *MockBackend) DeleteRoleBinding(arg0 context.Context, arg1 *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoleBinding indicates an expected call of DeleteRoleBinding.
func (mr *MockBackendMockRecorder) DeleteRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoleBinding", reflect.TypeOf((*MockBackend)(nil).DeleteRoleBinding), arg0, arg1)
}

// DeleteToken mocks base method.
func (m *MockBackend) DeleteToken(ctx context.Context, ref *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", ctx, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteToken indicates an expected call of DeleteToken.
func (mr *MockBackendMockRecorder) DeleteToken(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockBackend)(nil).DeleteToken), ctx, ref)
}

// GetCluster mocks base method.
func (m *MockBackend) GetCluster(ctx context.Context, ref *core.Reference) (*core.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, ref)
	ret0, _ := ret[0].(*core.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockBackendMockRecorder) GetCluster(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockBackend)(nil).GetCluster), ctx, ref)
}

// GetRole mocks base method.
func (m *MockBackend) GetRole(arg0 context.Context, arg1 *core.Reference) (*core.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(*core.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockBackendMockRecorder) GetRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockBackend)(nil).GetRole), arg0, arg1)
}

// GetRoleBinding mocks base method.
func (m *MockBackend) GetRoleBinding(arg0 context.Context, arg1 *core.Reference) (*core.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(*core.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleBinding indicates an expected call of GetRoleBinding.
func (mr *MockBackendMockRecorder) GetRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleBinding", reflect.TypeOf((*MockBackend)(nil).GetRoleBinding), arg0, arg1)
}

// GetToken mocks base method.
func (m *MockBackend) GetToken(ctx context.Context, ref *core.Reference) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, ref)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockBackendMockRecorder) GetToken(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockBackend)(nil).GetToken), ctx, ref)
}

// KeyValueStore mocks base method.
func (m *MockBackend) KeyValueStore(namespace string) (storage.KeyValueStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyValueStore", namespace)
	ret0, _ := ret[0].(storage.KeyValueStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyValueStore indicates an expected call of KeyValueStore.
func (mr *MockBackendMockRecorder) KeyValueStore(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyValueStore", reflect.TypeOf((*MockBackend)(nil).KeyValueStore), namespace)
}

// KeyringStore mocks base method.
func (m *MockBackend) KeyringStore(ctx context.Context, namespace string, ref *core.Reference) (storage.KeyringStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyringStore", ctx, namespace, ref)
	ret0, _ := ret[0].(storage.KeyringStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyringStore indicates an expected call of KeyringStore.
func (mr *MockBackendMockRecorder) KeyringStore(ctx, namespace, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyringStore", reflect.TypeOf((*MockBackend)(nil).KeyringStore), ctx, namespace, ref)
}

// ListClusters mocks base method.
func (m *MockBackend) ListClusters(ctx context.Context, matchLabels *core.LabelSelector, matchOptions core.MatchOptions) (*core.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", ctx, matchLabels, matchOptions)
	ret0, _ := ret[0].(*core.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockBackendMockRecorder) ListClusters(ctx, matchLabels, matchOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockBackend)(nil).ListClusters), ctx, matchLabels, matchOptions)
}

// ListRoleBindings mocks base method.
func (m *MockBackend) ListRoleBindings(arg0 context.Context) (*core.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleBindings", arg0)
	ret0, _ := ret[0].(*core.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleBindings indicates an expected call of ListRoleBindings.
func (mr *MockBackendMockRecorder) ListRoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleBindings", reflect.TypeOf((*MockBackend)(nil).ListRoleBindings), arg0)
}

// ListRoles mocks base method.
func (m *MockBackend) ListRoles(arg0 context.Context) (*core.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", arg0)
	ret0, _ := ret[0].(*core.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockBackendMockRecorder) ListRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockBackend)(nil).ListRoles), arg0)
}

// ListTokens mocks base method.
func (m *MockBackend) ListTokens(ctx context.Context) ([]*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTokens", ctx)
	ret0, _ := ret[0].([]*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTokens indicates an expected call of ListTokens.
func (mr *MockBackendMockRecorder) ListTokens(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTokens", reflect.TypeOf((*MockBackend)(nil).ListTokens), ctx)
}

// UpdateCluster mocks base method.
func (m *MockBackend) UpdateCluster(ctx context.Context, ref *core.Reference, mutator storage.ClusterMutator) (*core.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", ctx, ref, mutator)
	ret0, _ := ret[0].(*core.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster.
func (mr *MockBackendMockRecorder) UpdateCluster(ctx, ref, mutator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockBackend)(nil).UpdateCluster), ctx, ref, mutator)
}

// UpdateToken mocks base method.
func (m *MockBackend) UpdateToken(ctx context.Context, ref *core.Reference, mutator storage.TokenMutator) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", ctx, ref, mutator)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockBackendMockRecorder) UpdateToken(ctx, ref, mutator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockBackend)(nil).UpdateToken), ctx, ref, mutator)
}

// MockTokenStore is a mock of TokenStore interface.
type MockTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenStoreMockRecorder
}

// MockTokenStoreMockRecorder is the mock recorder for MockTokenStore.
type MockTokenStoreMockRecorder struct {
	mock *MockTokenStore
}

// NewMockTokenStore creates a new mock instance.
func NewMockTokenStore(ctrl *gomock.Controller) *MockTokenStore {
	mock := &MockTokenStore{ctrl: ctrl}
	mock.recorder = &MockTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenStore) EXPECT() *MockTokenStoreMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MockTokenStore) CreateToken(ctx context.Context, ttl time.Duration, opts ...storage.TokenCreateOption) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ttl}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateToken", varargs...)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockTokenStoreMockRecorder) CreateToken(ctx, ttl interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ttl}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockTokenStore)(nil).CreateToken), varargs...)
}

// DeleteToken mocks base method.
func (m *MockTokenStore) DeleteToken(ctx context.Context, ref *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", ctx, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteToken indicates an expected call of DeleteToken.
func (mr *MockTokenStoreMockRecorder) DeleteToken(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockTokenStore)(nil).DeleteToken), ctx, ref)
}

// GetToken mocks base method.
func (m *MockTokenStore) GetToken(ctx context.Context, ref *core.Reference) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, ref)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenStoreMockRecorder) GetToken(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenStore)(nil).GetToken), ctx, ref)
}

// ListTokens mocks base method.
func (m *MockTokenStore) ListTokens(ctx context.Context) ([]*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTokens", ctx)
	ret0, _ := ret[0].([]*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTokens indicates an expected call of ListTokens.
func (mr *MockTokenStoreMockRecorder) ListTokens(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTokens", reflect.TypeOf((*MockTokenStore)(nil).ListTokens), ctx)
}

// UpdateToken mocks base method.
func (m *MockTokenStore) UpdateToken(ctx context.Context, ref *core.Reference, mutator storage.TokenMutator) (*core.BootstrapToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", ctx, ref, mutator)
	ret0, _ := ret[0].(*core.BootstrapToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockTokenStoreMockRecorder) UpdateToken(ctx, ref, mutator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockTokenStore)(nil).UpdateToken), ctx, ref, mutator)
}

// MockClusterStore is a mock of ClusterStore interface.
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore.
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance.
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method.
func (m *MockClusterStore) CreateCluster(ctx context.Context, cluster *core.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster.
func (mr *MockClusterStoreMockRecorder) CreateCluster(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterStore)(nil).CreateCluster), ctx, cluster)
}

// DeleteCluster mocks base method.
func (m *MockClusterStore) DeleteCluster(ctx context.Context, ref *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockClusterStoreMockRecorder) DeleteCluster(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterStore)(nil).DeleteCluster), ctx, ref)
}

// GetCluster mocks base method.
func (m *MockClusterStore) GetCluster(ctx context.Context, ref *core.Reference) (*core.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, ref)
	ret0, _ := ret[0].(*core.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockClusterStoreMockRecorder) GetCluster(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterStore)(nil).GetCluster), ctx, ref)
}

// ListClusters mocks base method.
func (m *MockClusterStore) ListClusters(ctx context.Context, matchLabels *core.LabelSelector, matchOptions core.MatchOptions) (*core.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", ctx, matchLabels, matchOptions)
	ret0, _ := ret[0].(*core.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockClusterStoreMockRecorder) ListClusters(ctx, matchLabels, matchOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockClusterStore)(nil).ListClusters), ctx, matchLabels, matchOptions)
}

// UpdateCluster mocks base method.
func (m *MockClusterStore) UpdateCluster(ctx context.Context, ref *core.Reference, mutator storage.ClusterMutator) (*core.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", ctx, ref, mutator)
	ret0, _ := ret[0].(*core.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster.
func (mr *MockClusterStoreMockRecorder) UpdateCluster(ctx, ref, mutator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterStore)(nil).UpdateCluster), ctx, ref, mutator)
}

// MockRBACStore is a mock of RBACStore interface.
type MockRBACStore struct {
	ctrl     *gomock.Controller
	recorder *MockRBACStoreMockRecorder
}

// MockRBACStoreMockRecorder is the mock recorder for MockRBACStore.
type MockRBACStoreMockRecorder struct {
	mock *MockRBACStore
}

// NewMockRBACStore creates a new mock instance.
func NewMockRBACStore(ctrl *gomock.Controller) *MockRBACStore {
	mock := &MockRBACStore{ctrl: ctrl}
	mock.recorder = &MockRBACStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACStore) EXPECT() *MockRBACStoreMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRBACStore) CreateRole(arg0 context.Context, arg1 *core.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRBACStoreMockRecorder) CreateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRBACStore)(nil).CreateRole), arg0, arg1)
}

// CreateRoleBinding mocks base method.
func (m *MockRBACStore) CreateRoleBinding(arg0 context.Context, arg1 *core.RoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoleBinding indicates an expected call of CreateRoleBinding.
func (mr *MockRBACStoreMockRecorder) CreateRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoleBinding", reflect.TypeOf((*MockRBACStore)(nil).CreateRoleBinding), arg0, arg1)
}

// DeleteRole mocks base method.
func (m *MockRBACStore) DeleteRole(arg0 context.Context, arg1 *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRBACStoreMockRecorder) DeleteRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRBACStore)(nil).DeleteRole), arg0, arg1)
}

// DeleteRoleBinding mocks base method.
func (m *MockRBACStore) DeleteRoleBinding(arg0 context.Context, arg1 *core.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoleBinding indicates an expected call of DeleteRoleBinding.
func (mr *MockRBACStoreMockRecorder) DeleteRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoleBinding", reflect.TypeOf((*MockRBACStore)(nil).DeleteRoleBinding), arg0, arg1)
}

// GetRole mocks base method.
func (m *MockRBACStore) GetRole(arg0 context.Context, arg1 *core.Reference) (*core.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(*core.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockRBACStoreMockRecorder) GetRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockRBACStore)(nil).GetRole), arg0, arg1)
}

// GetRoleBinding mocks base method.
func (m *MockRBACStore) GetRoleBinding(arg0 context.Context, arg1 *core.Reference) (*core.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(*core.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleBinding indicates an expected call of GetRoleBinding.
func (mr *MockRBACStoreMockRecorder) GetRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleBinding", reflect.TypeOf((*MockRBACStore)(nil).GetRoleBinding), arg0, arg1)
}

// ListRoleBindings mocks base method.
func (m *MockRBACStore) ListRoleBindings(arg0 context.Context) (*core.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleBindings", arg0)
	ret0, _ := ret[0].(*core.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleBindings indicates an expected call of ListRoleBindings.
func (mr *MockRBACStoreMockRecorder) ListRoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleBindings", reflect.TypeOf((*MockRBACStore)(nil).ListRoleBindings), arg0)
}

// ListRoles mocks base method.
func (m *MockRBACStore) ListRoles(arg0 context.Context) (*core.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", arg0)
	ret0, _ := ret[0].(*core.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockRBACStoreMockRecorder) ListRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockRBACStore)(nil).ListRoles), arg0)
}

// MockKeyringStore is a mock of KeyringStore interface.
type MockKeyringStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyringStoreMockRecorder
}

// MockKeyringStoreMockRecorder is the mock recorder for MockKeyringStore.
type MockKeyringStoreMockRecorder struct {
	mock *MockKeyringStore
}

// NewMockKeyringStore creates a new mock instance.
func NewMockKeyringStore(ctrl *gomock.Controller) *MockKeyringStore {
	mock := &MockKeyringStore{ctrl: ctrl}
	mock.recorder = &MockKeyringStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyringStore) EXPECT() *MockKeyringStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKeyringStore) Get(ctx context.Context) (keyring.Keyring, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(keyring.Keyring)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyringStoreMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyringStore)(nil).Get), ctx)
}

// Put mocks base method.
func (m *MockKeyringStore) Put(ctx context.Context, keyring keyring.Keyring) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, keyring)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockKeyringStoreMockRecorder) Put(ctx, keyring interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKeyringStore)(nil).Put), ctx, keyring)
}

// MockKeyValueStore is a mock of KeyValueStore interface.
type MockKeyValueStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyValueStoreMockRecorder
}

// MockKeyValueStoreMockRecorder is the mock recorder for MockKeyValueStore.
type MockKeyValueStoreMockRecorder struct {
	mock *MockKeyValueStore
}

// NewMockKeyValueStore creates a new mock instance.
func NewMockKeyValueStore(ctrl *gomock.Controller) *MockKeyValueStore {
	mock := &MockKeyValueStore{ctrl: ctrl}
	mock.recorder = &MockKeyValueStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyValueStore) EXPECT() *MockKeyValueStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKeyValueStore) Get(ctx context.Context, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyValueStoreMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyValueStore)(nil).Get), ctx, key)
}

// ListKeys mocks base method.
func (m *MockKeyValueStore) ListKeys(ctx context.Context, prefix string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, prefix)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockKeyValueStoreMockRecorder) ListKeys(ctx, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKeyValueStore)(nil).ListKeys), ctx, prefix)
}

// Put mocks base method.
func (m *MockKeyValueStore) Put(ctx context.Context, key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockKeyValueStoreMockRecorder) Put(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKeyValueStore)(nil).Put), ctx, key, value)
}

// MockKeyringStoreBroker is a mock of KeyringStoreBroker interface.
type MockKeyringStoreBroker struct {
	ctrl     *gomock.Controller
	recorder *MockKeyringStoreBrokerMockRecorder
}

// MockKeyringStoreBrokerMockRecorder is the mock recorder for MockKeyringStoreBroker.
type MockKeyringStoreBrokerMockRecorder struct {
	mock *MockKeyringStoreBroker
}

// NewMockKeyringStoreBroker creates a new mock instance.
func NewMockKeyringStoreBroker(ctrl *gomock.Controller) *MockKeyringStoreBroker {
	mock := &MockKeyringStoreBroker{ctrl: ctrl}
	mock.recorder = &MockKeyringStoreBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyringStoreBroker) EXPECT() *MockKeyringStoreBrokerMockRecorder {
	return m.recorder
}

// KeyringStore mocks base method.
func (m *MockKeyringStoreBroker) KeyringStore(ctx context.Context, namespace string, ref *core.Reference) (storage.KeyringStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyringStore", ctx, namespace, ref)
	ret0, _ := ret[0].(storage.KeyringStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyringStore indicates an expected call of KeyringStore.
func (mr *MockKeyringStoreBrokerMockRecorder) KeyringStore(ctx, namespace, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyringStore", reflect.TypeOf((*MockKeyringStoreBroker)(nil).KeyringStore), ctx, namespace, ref)
}

// MockKeyValueStoreBroker is a mock of KeyValueStoreBroker interface.
type MockKeyValueStoreBroker struct {
	ctrl     *gomock.Controller
	recorder *MockKeyValueStoreBrokerMockRecorder
}

// MockKeyValueStoreBrokerMockRecorder is the mock recorder for MockKeyValueStoreBroker.
type MockKeyValueStoreBrokerMockRecorder struct {
	mock *MockKeyValueStoreBroker
}

// NewMockKeyValueStoreBroker creates a new mock instance.
func NewMockKeyValueStoreBroker(ctrl *gomock.Controller) *MockKeyValueStoreBroker {
	mock := &MockKeyValueStoreBroker{ctrl: ctrl}
	mock.recorder = &MockKeyValueStoreBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyValueStoreBroker) EXPECT() *MockKeyValueStoreBrokerMockRecorder {
	return m.recorder
}

// KeyValueStore mocks base method.
func (m *MockKeyValueStoreBroker) KeyValueStore(namespace string) (storage.KeyValueStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyValueStore", namespace)
	ret0, _ := ret[0].(storage.KeyValueStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyValueStore indicates an expected call of KeyValueStore.
func (mr *MockKeyValueStoreBrokerMockRecorder) KeyValueStore(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyValueStore", reflect.TypeOf((*MockKeyValueStoreBroker)(nil).KeyValueStore), namespace)
}

// MockSubjectAccessCapableStore is a mock of SubjectAccessCapableStore interface.
type MockSubjectAccessCapableStore struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectAccessCapableStoreMockRecorder
}

// MockSubjectAccessCapableStoreMockRecorder is the mock recorder for MockSubjectAccessCapableStore.
type MockSubjectAccessCapableStoreMockRecorder struct {
	mock *MockSubjectAccessCapableStore
}

// NewMockSubjectAccessCapableStore creates a new mock instance.
func NewMockSubjectAccessCapableStore(ctrl *gomock.Controller) *MockSubjectAccessCapableStore {
	mock := &MockSubjectAccessCapableStore{ctrl: ctrl}
	mock.recorder = &MockSubjectAccessCapableStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectAccessCapableStore) EXPECT() *MockSubjectAccessCapableStoreMockRecorder {
	return m.recorder
}

// GetRole mocks base method.
func (m *MockSubjectAccessCapableStore) GetRole(ctx context.Context, ref *core.Reference) (*core.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", ctx, ref)
	ret0, _ := ret[0].(*core.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockSubjectAccessCapableStoreMockRecorder) GetRole(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockSubjectAccessCapableStore)(nil).GetRole), ctx, ref)
}

// ListClusters mocks base method.
func (m *MockSubjectAccessCapableStore) ListClusters(ctx context.Context, matchLabels *core.LabelSelector, matchOptions core.MatchOptions) (*core.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", ctx, matchLabels, matchOptions)
	ret0, _ := ret[0].(*core.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockSubjectAccessCapableStoreMockRecorder) ListClusters(ctx, matchLabels, matchOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockSubjectAccessCapableStore)(nil).ListClusters), ctx, matchLabels, matchOptions)
}

// ListRoleBindings mocks base method.
func (m *MockSubjectAccessCapableStore) ListRoleBindings(ctx context.Context) (*core.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleBindings", ctx)
	ret0, _ := ret[0].(*core.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleBindings indicates an expected call of ListRoleBindings.
func (mr *MockSubjectAccessCapableStoreMockRecorder) ListRoleBindings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleBindings", reflect.TypeOf((*MockSubjectAccessCapableStore)(nil).ListRoleBindings), ctx)
}
