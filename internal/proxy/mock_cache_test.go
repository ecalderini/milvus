// Code generated by mockery v2.16.0. DO NOT EDIT.

package proxy

import (
	"context"

	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/stretchr/testify/mock"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"

	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

type MockCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCache) EXPECT() *MockCache_Expecter {
	return &MockCache_Expecter{mock: &_m.Mock}
}

// DeprecateShardCache provides a mock function with given fields: database, collectionName
func (_m *MockCache) DeprecateShardCache(database string, collectionName string) {
	_m.Called(database, collectionName)
}

// MockCache_DeprecateShardCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeprecateShardCache'
type MockCache_DeprecateShardCache_Call struct {
	*mock.Call
}

// DeprecateShardCache is a helper method to define mock.On call
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) DeprecateShardCache(database interface{}, collectionName interface{}) *MockCache_DeprecateShardCache_Call {
	return &MockCache_DeprecateShardCache_Call{Call: _e.mock.On("DeprecateShardCache", database, collectionName)}
}

func (_c *MockCache_DeprecateShardCache_Call) Run(run func(database string, collectionName string)) *MockCache_DeprecateShardCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockCache_DeprecateShardCache_Call) Return() *MockCache_DeprecateShardCache_Call {
	_c.Call.Return()
	return _c
}

// GetCollectionID provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetCollectionID(ctx context.Context, database string, collectionName string) (int64, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionID'
type MockCache_GetCollectionID_Call struct {
	*mock.Call
}

// GetCollectionID is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetCollectionID(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetCollectionID_Call {
	return &MockCache_GetCollectionID_Call{Call: _e.mock.On("GetCollectionID", ctx, database, collectionName)}
}

func (_c *MockCache_GetCollectionID_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetCollectionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetCollectionID_Call) Return(_a0 int64, _a1 error) *MockCache_GetCollectionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCollectionInfo provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetCollectionInfo(ctx context.Context, database string, collectionName string) (*collectionInfo, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 *collectionInfo
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *collectionInfo); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collectionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionInfo'
type MockCache_GetCollectionInfo_Call struct {
	*mock.Call
}

// GetCollectionInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetCollectionInfo(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetCollectionInfo_Call {
	return &MockCache_GetCollectionInfo_Call{Call: _e.mock.On("GetCollectionInfo", ctx, database, collectionName)}
}

func (_c *MockCache_GetCollectionInfo_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetCollectionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetCollectionInfo_Call) Return(_a0 *collectionInfo, _a1 error) *MockCache_GetCollectionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCollectionSchema provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetCollectionSchema(ctx context.Context, database string, collectionName string) (*schemapb.CollectionSchema, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 *schemapb.CollectionSchema
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *schemapb.CollectionSchema); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schemapb.CollectionSchema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionSchema'
type MockCache_GetCollectionSchema_Call struct {
	*mock.Call
}

// GetCollectionSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetCollectionSchema(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetCollectionSchema_Call {
	return &MockCache_GetCollectionSchema_Call{Call: _e.mock.On("GetCollectionSchema", ctx, database, collectionName)}
}

func (_c *MockCache_GetCollectionSchema_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetCollectionSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetCollectionSchema_Call) Return(_a0 *schemapb.CollectionSchema, _a1 error) *MockCache_GetCollectionSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCredentialInfo provides a mock function with given fields: ctx, username
func (_m *MockCache) GetCredentialInfo(ctx context.Context, username string) (*internalpb.CredentialInfo, error) {
	ret := _m.Called(ctx, username)

	var r0 *internalpb.CredentialInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) *internalpb.CredentialInfo); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.CredentialInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCredentialInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCredentialInfo'
type MockCache_GetCredentialInfo_Call struct {
	*mock.Call
}

// GetCredentialInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockCache_Expecter) GetCredentialInfo(ctx interface{}, username interface{}) *MockCache_GetCredentialInfo_Call {
	return &MockCache_GetCredentialInfo_Call{Call: _e.mock.On("GetCredentialInfo", ctx, username)}
}

func (_c *MockCache_GetCredentialInfo_Call) Run(run func(ctx context.Context, username string)) *MockCache_GetCredentialInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_GetCredentialInfo_Call) Return(_a0 *internalpb.CredentialInfo, _a1 error) *MockCache_GetCredentialInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDatabaseAndCollectionName provides a mock function with given fields: ctx, collectionID
func (_m *MockCache) GetDatabaseAndCollectionName(ctx context.Context, collectionID int64) (string, string, error) {
	ret := _m.Called(ctx, collectionID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, collectionID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, int64) string); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64) error); ok {
		r2 = rf(ctx, collectionID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCache_GetDatabaseAndCollectionName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseAndCollectionName'
type MockCache_GetDatabaseAndCollectionName_Call struct {
	*mock.Call
}

// GetDatabaseAndCollectionName is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockCache_Expecter) GetDatabaseAndCollectionName(ctx interface{}, collectionID interface{}) *MockCache_GetDatabaseAndCollectionName_Call {
	return &MockCache_GetDatabaseAndCollectionName_Call{Call: _e.mock.On("GetDatabaseAndCollectionName", ctx, collectionID)}
}

func (_c *MockCache_GetDatabaseAndCollectionName_Call) Run(run func(ctx context.Context, collectionID int64)) *MockCache_GetDatabaseAndCollectionName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockCache_GetDatabaseAndCollectionName_Call) Return(_a0 string, _a1 string, _a2 error) *MockCache_GetDatabaseAndCollectionName_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// GetPartitionID provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) GetPartitionID(ctx context.Context, database string, collectionName string, partitionName string) (int64, error) {
	ret := _m.Called(ctx, database, collectionName, partitionName)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) int64); ok {
		r0 = rf(ctx, database, collectionName, partitionName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, database, collectionName, partitionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionID'
type MockCache_GetPartitionID_Call struct {
	*mock.Call
}

// GetPartitionID is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) GetPartitionID(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_GetPartitionID_Call {
	return &MockCache_GetPartitionID_Call{Call: _e.mock.On("GetPartitionID", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_GetPartitionID_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_GetPartitionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitionID_Call) Return(_a0 int64, _a1 error) *MockCache_GetPartitionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPartitionInfo provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) GetPartitionInfo(ctx context.Context, database string, collectionName string, partitionName string) (*partitionInfo, error) {
	ret := _m.Called(ctx, database, collectionName, partitionName)

	var r0 *partitionInfo
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *partitionInfo); ok {
		r0 = rf(ctx, database, collectionName, partitionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*partitionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, database, collectionName, partitionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionInfo'
type MockCache_GetPartitionInfo_Call struct {
	*mock.Call
}

// GetPartitionInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) GetPartitionInfo(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_GetPartitionInfo_Call {
	return &MockCache_GetPartitionInfo_Call{Call: _e.mock.On("GetPartitionInfo", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_GetPartitionInfo_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_GetPartitionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitionInfo_Call) Return(_a0 *partitionInfo, _a1 error) *MockCache_GetPartitionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPartitions provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetPartitions(ctx context.Context, database string, collectionName string) (map[string]int64, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 map[string]int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]int64); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type MockCache_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetPartitions(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetPartitions_Call {
	return &MockCache_GetPartitions_Call{Call: _e.mock.On("GetPartitions", ctx, database, collectionName)}
}

func (_c *MockCache_GetPartitions_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitions_Call) Return(_a0 map[string]int64, _a1 error) *MockCache_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPrivilegeInfo provides a mock function with given fields: ctx
func (_m *MockCache) GetPrivilegeInfo(ctx context.Context) []string {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_GetPrivilegeInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrivilegeInfo'
type MockCache_GetPrivilegeInfo_Call struct {
	*mock.Call
}

// GetPrivilegeInfo is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCache_Expecter) GetPrivilegeInfo(ctx interface{}) *MockCache_GetPrivilegeInfo_Call {
	return &MockCache_GetPrivilegeInfo_Call{Call: _e.mock.On("GetPrivilegeInfo", ctx)}
}

func (_c *MockCache_GetPrivilegeInfo_Call) Run(run func(ctx context.Context)) *MockCache_GetPrivilegeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCache_GetPrivilegeInfo_Call) Return(_a0 []string) *MockCache_GetPrivilegeInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetShards provides a mock function with given fields: ctx, withCache, database, collectionName
func (_m *MockCache) GetShards(ctx context.Context, withCache bool, database string, collectionName string) (map[string][]nodeInfo, error) {
	ret := _m.Called(ctx, withCache, database, collectionName)

	var r0 map[string][]nodeInfo
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, string) map[string][]nodeInfo); ok {
		r0 = rf(ctx, withCache, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]nodeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool, string, string) error); ok {
		r1 = rf(ctx, withCache, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetShards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShards'
type MockCache_GetShards_Call struct {
	*mock.Call
}

// GetShards is a helper method to define mock.On call
//   - ctx context.Context
//   - withCache bool
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetShards(ctx interface{}, withCache interface{}, database interface{}, collectionName interface{}) *MockCache_GetShards_Call {
	return &MockCache_GetShards_Call{Call: _e.mock.On("GetShards", ctx, withCache, database, collectionName)}
}

func (_c *MockCache_GetShards_Call) Run(run func(ctx context.Context, withCache bool, database string, collectionName string)) *MockCache_GetShards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_GetShards_Call) Return(_a0 map[string][]nodeInfo, _a1 error) *MockCache_GetShards_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUserRole provides a mock function with given fields: username
func (_m *MockCache) GetUserRole(username string) []string {
	ret := _m.Called(username)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_GetUserRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserRole'
type MockCache_GetUserRole_Call struct {
	*mock.Call
}

// GetUserRole is a helper method to define mock.On call
//   - username string
func (_e *MockCache_Expecter) GetUserRole(username interface{}) *MockCache_GetUserRole_Call {
	return &MockCache_GetUserRole_Call{Call: _e.mock.On("GetUserRole", username)}
}

func (_c *MockCache_GetUserRole_Call) Run(run func(username string)) *MockCache_GetUserRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCache_GetUserRole_Call) Return(_a0 []string) *MockCache_GetUserRole_Call {
	_c.Call.Return(_a0)
	return _c
}

// InitPolicyInfo provides a mock function with given fields: info, userRoles
func (_m *MockCache) InitPolicyInfo(info []string, userRoles []string) {
	_m.Called(info, userRoles)
}

// MockCache_InitPolicyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitPolicyInfo'
type MockCache_InitPolicyInfo_Call struct {
	*mock.Call
}

// InitPolicyInfo is a helper method to define mock.On call
//   - info []string
//   - userRoles []string
func (_e *MockCache_Expecter) InitPolicyInfo(info interface{}, userRoles interface{}) *MockCache_InitPolicyInfo_Call {
	return &MockCache_InitPolicyInfo_Call{Call: _e.mock.On("InitPolicyInfo", info, userRoles)}
}

func (_c *MockCache_InitPolicyInfo_Call) Run(run func(info []string, userRoles []string)) *MockCache_InitPolicyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].([]string))
	})
	return _c
}

func (_c *MockCache_InitPolicyInfo_Call) Return() *MockCache_InitPolicyInfo_Call {
	_c.Call.Return()
	return _c
}

// RefreshPolicyInfo provides a mock function with given fields: op
func (_m *MockCache) RefreshPolicyInfo(op typeutil.CacheOp) error {
	ret := _m.Called(op)

	var r0 error
	if rf, ok := ret.Get(0).(func(typeutil.CacheOp) error); ok {
		r0 = rf(op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCache_RefreshPolicyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshPolicyInfo'
type MockCache_RefreshPolicyInfo_Call struct {
	*mock.Call
}

// RefreshPolicyInfo is a helper method to define mock.On call
//   - op typeutil.CacheOp
func (_e *MockCache_Expecter) RefreshPolicyInfo(op interface{}) *MockCache_RefreshPolicyInfo_Call {
	return &MockCache_RefreshPolicyInfo_Call{Call: _e.mock.On("RefreshPolicyInfo", op)}
}

func (_c *MockCache_RefreshPolicyInfo_Call) Run(run func(op typeutil.CacheOp)) *MockCache_RefreshPolicyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(typeutil.CacheOp))
	})
	return _c
}

func (_c *MockCache_RefreshPolicyInfo_Call) Return(_a0 error) *MockCache_RefreshPolicyInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveCollection provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) RemoveCollection(ctx context.Context, database string, collectionName string) {
	_m.Called(ctx, database, collectionName)
}

// MockCache_RemoveCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCollection'
type MockCache_RemoveCollection_Call struct {
	*mock.Call
}

// RemoveCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) RemoveCollection(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_RemoveCollection_Call {
	return &MockCache_RemoveCollection_Call{Call: _e.mock.On("RemoveCollection", ctx, database, collectionName)}
}

func (_c *MockCache_RemoveCollection_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_RemoveCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_RemoveCollection_Call) Return() *MockCache_RemoveCollection_Call {
	_c.Call.Return()
	return _c
}

// RemoveCollectionsByID provides a mock function with given fields: ctx, collectionID
func (_m *MockCache) RemoveCollectionsByID(ctx context.Context, collectionID int64) []string {
	ret := _m.Called(ctx, collectionID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, int64) []string); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_RemoveCollectionsByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCollectionsByID'
type MockCache_RemoveCollectionsByID_Call struct {
	*mock.Call
}

// RemoveCollectionsByID is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockCache_Expecter) RemoveCollectionsByID(ctx interface{}, collectionID interface{}) *MockCache_RemoveCollectionsByID_Call {
	return &MockCache_RemoveCollectionsByID_Call{Call: _e.mock.On("RemoveCollectionsByID", ctx, collectionID)}
}

func (_c *MockCache_RemoveCollectionsByID_Call) Run(run func(ctx context.Context, collectionID int64)) *MockCache_RemoveCollectionsByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockCache_RemoveCollectionsByID_Call) Return(_a0 []string) *MockCache_RemoveCollectionsByID_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveCredential provides a mock function with given fields: username
func (_m *MockCache) RemoveCredential(username string) {
	_m.Called(username)
}

// MockCache_RemoveCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCredential'
type MockCache_RemoveCredential_Call struct {
	*mock.Call
}

// RemoveCredential is a helper method to define mock.On call
//   - username string
func (_e *MockCache_Expecter) RemoveCredential(username interface{}) *MockCache_RemoveCredential_Call {
	return &MockCache_RemoveCredential_Call{Call: _e.mock.On("RemoveCredential", username)}
}

func (_c *MockCache_RemoveCredential_Call) Run(run func(username string)) *MockCache_RemoveCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCache_RemoveCredential_Call) Return() *MockCache_RemoveCredential_Call {
	_c.Call.Return()
	return _c
}

// RemoveDatabase provides a mock function with given fields: ctx, database
func (_m *MockCache) RemoveDatabase(ctx context.Context, database string) {
	_m.Called(ctx, database)
}

// MockCache_RemoveDatabase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveDatabase'
type MockCache_RemoveDatabase_Call struct {
	*mock.Call
}

// RemoveDatabase is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
func (_e *MockCache_Expecter) RemoveDatabase(ctx interface{}, database interface{}) *MockCache_RemoveDatabase_Call {
	return &MockCache_RemoveDatabase_Call{Call: _e.mock.On("RemoveDatabase", ctx, database)}
}

func (_c *MockCache_RemoveDatabase_Call) Run(run func(ctx context.Context, database string)) *MockCache_RemoveDatabase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_RemoveDatabase_Call) Return() *MockCache_RemoveDatabase_Call {
	_c.Call.Return()
	return _c
}

// RemovePartition provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) RemovePartition(ctx context.Context, database string, collectionName string, partitionName string) {
	_m.Called(ctx, database, collectionName, partitionName)
}

// MockCache_RemovePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePartition'
type MockCache_RemovePartition_Call struct {
	*mock.Call
}

// RemovePartition is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) RemovePartition(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_RemovePartition_Call {
	return &MockCache_RemovePartition_Call{Call: _e.mock.On("RemovePartition", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_RemovePartition_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_RemovePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_RemovePartition_Call) Return() *MockCache_RemovePartition_Call {
	_c.Call.Return()
	return _c
}

// UpdateCredential provides a mock function with given fields: credInfo
func (_m *MockCache) UpdateCredential(credInfo *internalpb.CredentialInfo) {
	_m.Called(credInfo)
}

// MockCache_UpdateCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCredential'
type MockCache_UpdateCredential_Call struct {
	*mock.Call
}

// UpdateCredential is a helper method to define mock.On call
//   - credInfo *internalpb.CredentialInfo
func (_e *MockCache_Expecter) UpdateCredential(credInfo interface{}) *MockCache_UpdateCredential_Call {
	return &MockCache_UpdateCredential_Call{Call: _e.mock.On("UpdateCredential", credInfo)}
}

func (_c *MockCache_UpdateCredential_Call) Run(run func(credInfo *internalpb.CredentialInfo)) *MockCache_UpdateCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*internalpb.CredentialInfo))
	})
	return _c
}

func (_c *MockCache_UpdateCredential_Call) Return() *MockCache_UpdateCredential_Call {
	_c.Call.Return()
	return _c
}

// expireShardLeaderCache provides a mock function with given fields: ctx
func (_m *MockCache) expireShardLeaderCache(ctx context.Context) {
	_m.Called(ctx)
}

// MockCache_expireShardLeaderCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'expireShardLeaderCache'
type MockCache_expireShardLeaderCache_Call struct {
	*mock.Call
}

// expireShardLeaderCache is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCache_Expecter) expireShardLeaderCache(ctx interface{}) *MockCache_expireShardLeaderCache_Call {
	return &MockCache_expireShardLeaderCache_Call{Call: _e.mock.On("expireShardLeaderCache", ctx)}
}

func (_c *MockCache_expireShardLeaderCache_Call) Run(run func(ctx context.Context)) *MockCache_expireShardLeaderCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCache_expireShardLeaderCache_Call) Return() *MockCache_expireShardLeaderCache_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCache(t mockConstructorTestingTNewMockCache) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
