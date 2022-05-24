// Code generated by go-mockgen 1.2.0; DO NOT EDIT.

package featureflag

import (
	"context"
	"sync"
)

// MockStore is a mock implementation of the Store interface (from the
// package github.com/sourcegraph/sourcegraph/internal/featureflag) used for
// unit testing.
type MockStore struct {
	// GetAnonymousUserFlagFunc is an instance of a mock function object
	// controlling the behavior of the method GetAnonymousUserFlag.
	GetAnonymousUserFlagFunc *StoreGetAnonymousUserFlagFunc
	// GetFeatureFlagFunc is an instance of a mock function object
	// controlling the behavior of the method GetFeatureFlag.
	GetFeatureFlagFunc *StoreGetFeatureFlagFunc
	// GetFeatureFlagsFunc is an instance of a mock function object
	// controlling the behavior of the method GetFeatureFlags.
	GetFeatureFlagsFunc *StoreGetFeatureFlagsFunc
	// GetGlobalFeatureFlagFunc is an instance of a mock function object
	// controlling the behavior of the method GetGlobalFeatureFlag.
	GetGlobalFeatureFlagFunc *StoreGetGlobalFeatureFlagFunc
	// GetOrgOverrideForUserFunc is an instance of a mock function object
	// controlling the behavior of the method GetOrgOverrideForUser.
	GetOrgOverrideForUserFunc *StoreGetOrgOverrideForUserFunc
	// GetUserFlagFunc is an instance of a mock function object controlling
	// the behavior of the method GetUserFlag.
	GetUserFlagFunc *StoreGetUserFlagFunc
	// GetUserOverrideFunc is an instance of a mock function object
	// controlling the behavior of the method GetUserOverride.
	GetUserOverrideFunc *StoreGetUserOverrideFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		GetAnonymousUserFlagFunc: &StoreGetAnonymousUserFlagFunc{
			defaultHook: func(context.Context, string, string) (r0 *bool, r1 error) {
				return
			},
		},
		GetFeatureFlagFunc: &StoreGetFeatureFlagFunc{
			defaultHook: func(context.Context, string) (r0 *FeatureFlag, r1 error) {
				return
			},
		},
		GetFeatureFlagsFunc: &StoreGetFeatureFlagsFunc{
			defaultHook: func(context.Context) (r0 []*FeatureFlag, r1 error) {
				return
			},
		},
		GetGlobalFeatureFlagFunc: &StoreGetGlobalFeatureFlagFunc{
			defaultHook: func(context.Context, string) (r0 *bool, r1 error) {
				return
			},
		},
		GetOrgOverrideForUserFunc: &StoreGetOrgOverrideForUserFunc{
			defaultHook: func(context.Context, int32, string) (r0 *Override, r1 error) {
				return
			},
		},
		GetUserFlagFunc: &StoreGetUserFlagFunc{
			defaultHook: func(context.Context, int32, string) (r0 *bool, r1 error) {
				return
			},
		},
		GetUserOverrideFunc: &StoreGetUserOverrideFunc{
			defaultHook: func(context.Context, int32, string) (r0 *Override, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockStore creates a new mock of the Store interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockStore() *MockStore {
	return &MockStore{
		GetAnonymousUserFlagFunc: &StoreGetAnonymousUserFlagFunc{
			defaultHook: func(context.Context, string, string) (*bool, error) {
				panic("unexpected invocation of MockStore.GetAnonymousUserFlag")
			},
		},
		GetFeatureFlagFunc: &StoreGetFeatureFlagFunc{
			defaultHook: func(context.Context, string) (*FeatureFlag, error) {
				panic("unexpected invocation of MockStore.GetFeatureFlag")
			},
		},
		GetFeatureFlagsFunc: &StoreGetFeatureFlagsFunc{
			defaultHook: func(context.Context) ([]*FeatureFlag, error) {
				panic("unexpected invocation of MockStore.GetFeatureFlags")
			},
		},
		GetGlobalFeatureFlagFunc: &StoreGetGlobalFeatureFlagFunc{
			defaultHook: func(context.Context, string) (*bool, error) {
				panic("unexpected invocation of MockStore.GetGlobalFeatureFlag")
			},
		},
		GetOrgOverrideForUserFunc: &StoreGetOrgOverrideForUserFunc{
			defaultHook: func(context.Context, int32, string) (*Override, error) {
				panic("unexpected invocation of MockStore.GetOrgOverrideForUser")
			},
		},
		GetUserFlagFunc: &StoreGetUserFlagFunc{
			defaultHook: func(context.Context, int32, string) (*bool, error) {
				panic("unexpected invocation of MockStore.GetUserFlag")
			},
		},
		GetUserOverrideFunc: &StoreGetUserOverrideFunc{
			defaultHook: func(context.Context, int32, string) (*Override, error) {
				panic("unexpected invocation of MockStore.GetUserOverride")
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i Store) *MockStore {
	return &MockStore{
		GetAnonymousUserFlagFunc: &StoreGetAnonymousUserFlagFunc{
			defaultHook: i.GetAnonymousUserFlag,
		},
		GetFeatureFlagFunc: &StoreGetFeatureFlagFunc{
			defaultHook: i.GetFeatureFlag,
		},
		GetFeatureFlagsFunc: &StoreGetFeatureFlagsFunc{
			defaultHook: i.GetFeatureFlags,
		},
		GetGlobalFeatureFlagFunc: &StoreGetGlobalFeatureFlagFunc{
			defaultHook: i.GetGlobalFeatureFlag,
		},
		GetOrgOverrideForUserFunc: &StoreGetOrgOverrideForUserFunc{
			defaultHook: i.GetOrgOverrideForUser,
		},
		GetUserFlagFunc: &StoreGetUserFlagFunc{
			defaultHook: i.GetUserFlag,
		},
		GetUserOverrideFunc: &StoreGetUserOverrideFunc{
			defaultHook: i.GetUserOverride,
		},
	}
}

// StoreGetAnonymousUserFlagFunc describes the behavior when the
// GetAnonymousUserFlag method of the parent MockStore instance is invoked.
type StoreGetAnonymousUserFlagFunc struct {
	defaultHook func(context.Context, string, string) (*bool, error)
	hooks       []func(context.Context, string, string) (*bool, error)
	history     []StoreGetAnonymousUserFlagFuncCall
	mutex       sync.Mutex
}

// GetAnonymousUserFlag delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) GetAnonymousUserFlag(v0 context.Context, v1 string, v2 string) (*bool, error) {
	r0, r1 := m.GetAnonymousUserFlagFunc.nextHook()(v0, v1, v2)
	m.GetAnonymousUserFlagFunc.appendCall(StoreGetAnonymousUserFlagFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetAnonymousUserFlag
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreGetAnonymousUserFlagFunc) SetDefaultHook(hook func(context.Context, string, string) (*bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetAnonymousUserFlag method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreGetAnonymousUserFlagFunc) PushHook(hook func(context.Context, string, string) (*bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetAnonymousUserFlagFunc) SetDefaultReturn(r0 *bool, r1 error) {
	f.SetDefaultHook(func(context.Context, string, string) (*bool, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetAnonymousUserFlagFunc) PushReturn(r0 *bool, r1 error) {
	f.PushHook(func(context.Context, string, string) (*bool, error) {
		return r0, r1
	})
}

func (f *StoreGetAnonymousUserFlagFunc) nextHook() func(context.Context, string, string) (*bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetAnonymousUserFlagFunc) appendCall(r0 StoreGetAnonymousUserFlagFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetAnonymousUserFlagFuncCall objects
// describing the invocations of this function.
func (f *StoreGetAnonymousUserFlagFunc) History() []StoreGetAnonymousUserFlagFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetAnonymousUserFlagFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetAnonymousUserFlagFuncCall is an object that describes an
// invocation of method GetAnonymousUserFlag on an instance of MockStore.
type StoreGetAnonymousUserFlagFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetAnonymousUserFlagFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetAnonymousUserFlagFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetFeatureFlagFunc describes the behavior when the GetFeatureFlag
// method of the parent MockStore instance is invoked.
type StoreGetFeatureFlagFunc struct {
	defaultHook func(context.Context, string) (*FeatureFlag, error)
	hooks       []func(context.Context, string) (*FeatureFlag, error)
	history     []StoreGetFeatureFlagFuncCall
	mutex       sync.Mutex
}

// GetFeatureFlag delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) GetFeatureFlag(v0 context.Context, v1 string) (*FeatureFlag, error) {
	r0, r1 := m.GetFeatureFlagFunc.nextHook()(v0, v1)
	m.GetFeatureFlagFunc.appendCall(StoreGetFeatureFlagFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetFeatureFlag
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreGetFeatureFlagFunc) SetDefaultHook(hook func(context.Context, string) (*FeatureFlag, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetFeatureFlag method of the parent MockStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreGetFeatureFlagFunc) PushHook(hook func(context.Context, string) (*FeatureFlag, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetFeatureFlagFunc) SetDefaultReturn(r0 *FeatureFlag, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*FeatureFlag, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetFeatureFlagFunc) PushReturn(r0 *FeatureFlag, r1 error) {
	f.PushHook(func(context.Context, string) (*FeatureFlag, error) {
		return r0, r1
	})
}

func (f *StoreGetFeatureFlagFunc) nextHook() func(context.Context, string) (*FeatureFlag, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetFeatureFlagFunc) appendCall(r0 StoreGetFeatureFlagFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetFeatureFlagFuncCall objects
// describing the invocations of this function.
func (f *StoreGetFeatureFlagFunc) History() []StoreGetFeatureFlagFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetFeatureFlagFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetFeatureFlagFuncCall is an object that describes an invocation of
// method GetFeatureFlag on an instance of MockStore.
type StoreGetFeatureFlagFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *FeatureFlag
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetFeatureFlagFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetFeatureFlagFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetFeatureFlagsFunc describes the behavior when the GetFeatureFlags
// method of the parent MockStore instance is invoked.
type StoreGetFeatureFlagsFunc struct {
	defaultHook func(context.Context) ([]*FeatureFlag, error)
	hooks       []func(context.Context) ([]*FeatureFlag, error)
	history     []StoreGetFeatureFlagsFuncCall
	mutex       sync.Mutex
}

// GetFeatureFlags delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) GetFeatureFlags(v0 context.Context) ([]*FeatureFlag, error) {
	r0, r1 := m.GetFeatureFlagsFunc.nextHook()(v0)
	m.GetFeatureFlagsFunc.appendCall(StoreGetFeatureFlagsFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetFeatureFlags
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreGetFeatureFlagsFunc) SetDefaultHook(hook func(context.Context) ([]*FeatureFlag, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetFeatureFlags method of the parent MockStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreGetFeatureFlagsFunc) PushHook(hook func(context.Context) ([]*FeatureFlag, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetFeatureFlagsFunc) SetDefaultReturn(r0 []*FeatureFlag, r1 error) {
	f.SetDefaultHook(func(context.Context) ([]*FeatureFlag, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetFeatureFlagsFunc) PushReturn(r0 []*FeatureFlag, r1 error) {
	f.PushHook(func(context.Context) ([]*FeatureFlag, error) {
		return r0, r1
	})
}

func (f *StoreGetFeatureFlagsFunc) nextHook() func(context.Context) ([]*FeatureFlag, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetFeatureFlagsFunc) appendCall(r0 StoreGetFeatureFlagsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetFeatureFlagsFuncCall objects
// describing the invocations of this function.
func (f *StoreGetFeatureFlagsFunc) History() []StoreGetFeatureFlagsFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetFeatureFlagsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetFeatureFlagsFuncCall is an object that describes an invocation of
// method GetFeatureFlags on an instance of MockStore.
type StoreGetFeatureFlagsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*FeatureFlag
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetFeatureFlagsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetFeatureFlagsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetGlobalFeatureFlagFunc describes the behavior when the
// GetGlobalFeatureFlag method of the parent MockStore instance is invoked.
type StoreGetGlobalFeatureFlagFunc struct {
	defaultHook func(context.Context, string) (*bool, error)
	hooks       []func(context.Context, string) (*bool, error)
	history     []StoreGetGlobalFeatureFlagFuncCall
	mutex       sync.Mutex
}

// GetGlobalFeatureFlag delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) GetGlobalFeatureFlag(v0 context.Context, v1 string) (*bool, error) {
	r0, r1 := m.GetGlobalFeatureFlagFunc.nextHook()(v0, v1)
	m.GetGlobalFeatureFlagFunc.appendCall(StoreGetGlobalFeatureFlagFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetGlobalFeatureFlag
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreGetGlobalFeatureFlagFunc) SetDefaultHook(hook func(context.Context, string) (*bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetGlobalFeatureFlag method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreGetGlobalFeatureFlagFunc) PushHook(hook func(context.Context, string) (*bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetGlobalFeatureFlagFunc) SetDefaultReturn(r0 *bool, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*bool, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetGlobalFeatureFlagFunc) PushReturn(r0 *bool, r1 error) {
	f.PushHook(func(context.Context, string) (*bool, error) {
		return r0, r1
	})
}

func (f *StoreGetGlobalFeatureFlagFunc) nextHook() func(context.Context, string) (*bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetGlobalFeatureFlagFunc) appendCall(r0 StoreGetGlobalFeatureFlagFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetGlobalFeatureFlagFuncCall objects
// describing the invocations of this function.
func (f *StoreGetGlobalFeatureFlagFunc) History() []StoreGetGlobalFeatureFlagFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetGlobalFeatureFlagFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetGlobalFeatureFlagFuncCall is an object that describes an
// invocation of method GetGlobalFeatureFlag on an instance of MockStore.
type StoreGetGlobalFeatureFlagFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetGlobalFeatureFlagFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetGlobalFeatureFlagFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetOrgOverrideForUserFunc describes the behavior when the
// GetOrgOverrideForUser method of the parent MockStore instance is invoked.
type StoreGetOrgOverrideForUserFunc struct {
	defaultHook func(context.Context, int32, string) (*Override, error)
	hooks       []func(context.Context, int32, string) (*Override, error)
	history     []StoreGetOrgOverrideForUserFuncCall
	mutex       sync.Mutex
}

// GetOrgOverrideForUser delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStore) GetOrgOverrideForUser(v0 context.Context, v1 int32, v2 string) (*Override, error) {
	r0, r1 := m.GetOrgOverrideForUserFunc.nextHook()(v0, v1, v2)
	m.GetOrgOverrideForUserFunc.appendCall(StoreGetOrgOverrideForUserFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// GetOrgOverrideForUser method of the parent MockStore instance is invoked
// and the hook queue is empty.
func (f *StoreGetOrgOverrideForUserFunc) SetDefaultHook(hook func(context.Context, int32, string) (*Override, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetOrgOverrideForUser method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreGetOrgOverrideForUserFunc) PushHook(hook func(context.Context, int32, string) (*Override, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetOrgOverrideForUserFunc) SetDefaultReturn(r0 *Override, r1 error) {
	f.SetDefaultHook(func(context.Context, int32, string) (*Override, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetOrgOverrideForUserFunc) PushReturn(r0 *Override, r1 error) {
	f.PushHook(func(context.Context, int32, string) (*Override, error) {
		return r0, r1
	})
}

func (f *StoreGetOrgOverrideForUserFunc) nextHook() func(context.Context, int32, string) (*Override, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetOrgOverrideForUserFunc) appendCall(r0 StoreGetOrgOverrideForUserFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetOrgOverrideForUserFuncCall objects
// describing the invocations of this function.
func (f *StoreGetOrgOverrideForUserFunc) History() []StoreGetOrgOverrideForUserFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetOrgOverrideForUserFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetOrgOverrideForUserFuncCall is an object that describes an
// invocation of method GetOrgOverrideForUser on an instance of MockStore.
type StoreGetOrgOverrideForUserFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int32
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *Override
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetOrgOverrideForUserFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetOrgOverrideForUserFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetUserFlagFunc describes the behavior when the GetUserFlag method
// of the parent MockStore instance is invoked.
type StoreGetUserFlagFunc struct {
	defaultHook func(context.Context, int32, string) (*bool, error)
	hooks       []func(context.Context, int32, string) (*bool, error)
	history     []StoreGetUserFlagFuncCall
	mutex       sync.Mutex
}

// GetUserFlag delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockStore) GetUserFlag(v0 context.Context, v1 int32, v2 string) (*bool, error) {
	r0, r1 := m.GetUserFlagFunc.nextHook()(v0, v1, v2)
	m.GetUserFlagFunc.appendCall(StoreGetUserFlagFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetUserFlag method
// of the parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreGetUserFlagFunc) SetDefaultHook(hook func(context.Context, int32, string) (*bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetUserFlag method of the parent MockStore instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreGetUserFlagFunc) PushHook(hook func(context.Context, int32, string) (*bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetUserFlagFunc) SetDefaultReturn(r0 *bool, r1 error) {
	f.SetDefaultHook(func(context.Context, int32, string) (*bool, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetUserFlagFunc) PushReturn(r0 *bool, r1 error) {
	f.PushHook(func(context.Context, int32, string) (*bool, error) {
		return r0, r1
	})
}

func (f *StoreGetUserFlagFunc) nextHook() func(context.Context, int32, string) (*bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetUserFlagFunc) appendCall(r0 StoreGetUserFlagFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetUserFlagFuncCall objects describing
// the invocations of this function.
func (f *StoreGetUserFlagFunc) History() []StoreGetUserFlagFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetUserFlagFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetUserFlagFuncCall is an object that describes an invocation of
// method GetUserFlag on an instance of MockStore.
type StoreGetUserFlagFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int32
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetUserFlagFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetUserFlagFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreGetUserOverrideFunc describes the behavior when the GetUserOverride
// method of the parent MockStore instance is invoked.
type StoreGetUserOverrideFunc struct {
	defaultHook func(context.Context, int32, string) (*Override, error)
	hooks       []func(context.Context, int32, string) (*Override, error)
	history     []StoreGetUserOverrideFuncCall
	mutex       sync.Mutex
}

// GetUserOverride delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) GetUserOverride(v0 context.Context, v1 int32, v2 string) (*Override, error) {
	r0, r1 := m.GetUserOverrideFunc.nextHook()(v0, v1, v2)
	m.GetUserOverrideFunc.appendCall(StoreGetUserOverrideFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetUserOverride
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreGetUserOverrideFunc) SetDefaultHook(hook func(context.Context, int32, string) (*Override, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetUserOverride method of the parent MockStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreGetUserOverrideFunc) PushHook(hook func(context.Context, int32, string) (*Override, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetUserOverrideFunc) SetDefaultReturn(r0 *Override, r1 error) {
	f.SetDefaultHook(func(context.Context, int32, string) (*Override, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetUserOverrideFunc) PushReturn(r0 *Override, r1 error) {
	f.PushHook(func(context.Context, int32, string) (*Override, error) {
		return r0, r1
	})
}

func (f *StoreGetUserOverrideFunc) nextHook() func(context.Context, int32, string) (*Override, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetUserOverrideFunc) appendCall(r0 StoreGetUserOverrideFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetUserOverrideFuncCall objects
// describing the invocations of this function.
func (f *StoreGetUserOverrideFunc) History() []StoreGetUserOverrideFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetUserOverrideFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetUserOverrideFuncCall is an object that describes an invocation of
// method GetUserOverride on an instance of MockStore.
type StoreGetUserOverrideFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int32
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *Override
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetUserOverrideFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetUserOverrideFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
