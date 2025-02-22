// Code generated by counterfeiter. DO NOT EDIT.
package apkfakes

import (
	"sync"

	"chainguard.dev/apko/pkg/build/types"
	"chainguard.dev/apko/pkg/exec"
	"chainguard.dev/apko/pkg/options"
)

type FakeApkImplementation struct {
	FixateWorldStub        func(*options.Options, *exec.Executor) error
	fixateWorldMutex       sync.RWMutex
	fixateWorldArgsForCall []struct {
		arg1 *options.Options
		arg2 *exec.Executor
	}
	fixateWorldReturns struct {
		result1 error
	}
	fixateWorldReturnsOnCall map[int]struct {
		result1 error
	}
	InitDBStub        func(*options.Options, exec.Executor) error
	initDBMutex       sync.RWMutex
	initDBArgsForCall []struct {
		arg1 *options.Options
		arg2 exec.Executor
	}
	initDBReturns struct {
		result1 error
	}
	initDBReturnsOnCall map[int]struct {
		result1 error
	}
	InitKeyringStub        func(*options.Options, *types.ImageConfiguration) error
	initKeyringMutex       sync.RWMutex
	initKeyringArgsForCall []struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}
	initKeyringReturns struct {
		result1 error
	}
	initKeyringReturnsOnCall map[int]struct {
		result1 error
	}
	InitRepositoriesStub        func(*options.Options, *types.ImageConfiguration) error
	initRepositoriesMutex       sync.RWMutex
	initRepositoriesArgsForCall []struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}
	initRepositoriesReturns struct {
		result1 error
	}
	initRepositoriesReturnsOnCall map[int]struct {
		result1 error
	}
	InitWorldStub        func(*options.Options, *types.ImageConfiguration) error
	initWorldMutex       sync.RWMutex
	initWorldArgsForCall []struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}
	initWorldReturns struct {
		result1 error
	}
	initWorldReturnsOnCall map[int]struct {
		result1 error
	}
	LoadSystemKeyringStub        func(*options.Options, ...string) ([]string, error)
	loadSystemKeyringMutex       sync.RWMutex
	loadSystemKeyringArgsForCall []struct {
		arg1 *options.Options
		arg2 []string
	}
	loadSystemKeyringReturns struct {
		result1 []string
		result2 error
	}
	loadSystemKeyringReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	NormalizeScriptsTarStub        func(*options.Options) error
	normalizeScriptsTarMutex       sync.RWMutex
	normalizeScriptsTarArgsForCall []struct {
		arg1 *options.Options
	}
	normalizeScriptsTarReturns struct {
		result1 error
	}
	normalizeScriptsTarReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApkImplementation) FixateWorld(arg1 *options.Options, arg2 *exec.Executor) error {
	fake.fixateWorldMutex.Lock()
	ret, specificReturn := fake.fixateWorldReturnsOnCall[len(fake.fixateWorldArgsForCall)]
	fake.fixateWorldArgsForCall = append(fake.fixateWorldArgsForCall, struct {
		arg1 *options.Options
		arg2 *exec.Executor
	}{arg1, arg2})
	stub := fake.FixateWorldStub
	fakeReturns := fake.fixateWorldReturns
	fake.recordInvocation("FixateWorld", []interface{}{arg1, arg2})
	fake.fixateWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) FixateWorldCallCount() int {
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	return len(fake.fixateWorldArgsForCall)
}

func (fake *FakeApkImplementation) FixateWorldCalls(stub func(*options.Options, *exec.Executor) error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = stub
}

func (fake *FakeApkImplementation) FixateWorldArgsForCall(i int) (*options.Options, *exec.Executor) {
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	argsForCall := fake.fixateWorldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) FixateWorldReturns(result1 error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = nil
	fake.fixateWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) FixateWorldReturnsOnCall(i int, result1 error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = nil
	if fake.fixateWorldReturnsOnCall == nil {
		fake.fixateWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.fixateWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitDB(arg1 *options.Options, arg2 exec.Executor) error {
	fake.initDBMutex.Lock()
	ret, specificReturn := fake.initDBReturnsOnCall[len(fake.initDBArgsForCall)]
	fake.initDBArgsForCall = append(fake.initDBArgsForCall, struct {
		arg1 *options.Options
		arg2 exec.Executor
	}{arg1, arg2})
	stub := fake.InitDBStub
	fakeReturns := fake.initDBReturns
	fake.recordInvocation("InitDB", []interface{}{arg1, arg2})
	fake.initDBMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitDBCallCount() int {
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	return len(fake.initDBArgsForCall)
}

func (fake *FakeApkImplementation) InitDBCalls(stub func(*options.Options, exec.Executor) error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = stub
}

func (fake *FakeApkImplementation) InitDBArgsForCall(i int) (*options.Options, exec.Executor) {
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	argsForCall := fake.initDBArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) InitDBReturns(result1 error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = nil
	fake.initDBReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitDBReturnsOnCall(i int, result1 error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = nil
	if fake.initDBReturnsOnCall == nil {
		fake.initDBReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initDBReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitKeyring(arg1 *options.Options, arg2 *types.ImageConfiguration) error {
	fake.initKeyringMutex.Lock()
	ret, specificReturn := fake.initKeyringReturnsOnCall[len(fake.initKeyringArgsForCall)]
	fake.initKeyringArgsForCall = append(fake.initKeyringArgsForCall, struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.InitKeyringStub
	fakeReturns := fake.initKeyringReturns
	fake.recordInvocation("InitKeyring", []interface{}{arg1, arg2})
	fake.initKeyringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitKeyringCallCount() int {
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	return len(fake.initKeyringArgsForCall)
}

func (fake *FakeApkImplementation) InitKeyringCalls(stub func(*options.Options, *types.ImageConfiguration) error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = stub
}

func (fake *FakeApkImplementation) InitKeyringArgsForCall(i int) (*options.Options, *types.ImageConfiguration) {
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	argsForCall := fake.initKeyringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) InitKeyringReturns(result1 error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = nil
	fake.initKeyringReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitKeyringReturnsOnCall(i int, result1 error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = nil
	if fake.initKeyringReturnsOnCall == nil {
		fake.initKeyringReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initKeyringReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitRepositories(arg1 *options.Options, arg2 *types.ImageConfiguration) error {
	fake.initRepositoriesMutex.Lock()
	ret, specificReturn := fake.initRepositoriesReturnsOnCall[len(fake.initRepositoriesArgsForCall)]
	fake.initRepositoriesArgsForCall = append(fake.initRepositoriesArgsForCall, struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.InitRepositoriesStub
	fakeReturns := fake.initRepositoriesReturns
	fake.recordInvocation("InitRepositories", []interface{}{arg1, arg2})
	fake.initRepositoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitRepositoriesCallCount() int {
	fake.initRepositoriesMutex.RLock()
	defer fake.initRepositoriesMutex.RUnlock()
	return len(fake.initRepositoriesArgsForCall)
}

func (fake *FakeApkImplementation) InitRepositoriesCalls(stub func(*options.Options, *types.ImageConfiguration) error) {
	fake.initRepositoriesMutex.Lock()
	defer fake.initRepositoriesMutex.Unlock()
	fake.InitRepositoriesStub = stub
}

func (fake *FakeApkImplementation) InitRepositoriesArgsForCall(i int) (*options.Options, *types.ImageConfiguration) {
	fake.initRepositoriesMutex.RLock()
	defer fake.initRepositoriesMutex.RUnlock()
	argsForCall := fake.initRepositoriesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) InitRepositoriesReturns(result1 error) {
	fake.initRepositoriesMutex.Lock()
	defer fake.initRepositoriesMutex.Unlock()
	fake.InitRepositoriesStub = nil
	fake.initRepositoriesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitRepositoriesReturnsOnCall(i int, result1 error) {
	fake.initRepositoriesMutex.Lock()
	defer fake.initRepositoriesMutex.Unlock()
	fake.InitRepositoriesStub = nil
	if fake.initRepositoriesReturnsOnCall == nil {
		fake.initRepositoriesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initRepositoriesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitWorld(arg1 *options.Options, arg2 *types.ImageConfiguration) error {
	fake.initWorldMutex.Lock()
	ret, specificReturn := fake.initWorldReturnsOnCall[len(fake.initWorldArgsForCall)]
	fake.initWorldArgsForCall = append(fake.initWorldArgsForCall, struct {
		arg1 *options.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.InitWorldStub
	fakeReturns := fake.initWorldReturns
	fake.recordInvocation("InitWorld", []interface{}{arg1, arg2})
	fake.initWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitWorldCallCount() int {
	fake.initWorldMutex.RLock()
	defer fake.initWorldMutex.RUnlock()
	return len(fake.initWorldArgsForCall)
}

func (fake *FakeApkImplementation) InitWorldCalls(stub func(*options.Options, *types.ImageConfiguration) error) {
	fake.initWorldMutex.Lock()
	defer fake.initWorldMutex.Unlock()
	fake.InitWorldStub = stub
}

func (fake *FakeApkImplementation) InitWorldArgsForCall(i int) (*options.Options, *types.ImageConfiguration) {
	fake.initWorldMutex.RLock()
	defer fake.initWorldMutex.RUnlock()
	argsForCall := fake.initWorldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) InitWorldReturns(result1 error) {
	fake.initWorldMutex.Lock()
	defer fake.initWorldMutex.Unlock()
	fake.InitWorldStub = nil
	fake.initWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitWorldReturnsOnCall(i int, result1 error) {
	fake.initWorldMutex.Lock()
	defer fake.initWorldMutex.Unlock()
	fake.InitWorldStub = nil
	if fake.initWorldReturnsOnCall == nil {
		fake.initWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) LoadSystemKeyring(arg1 *options.Options, arg2 ...string) ([]string, error) {
	fake.loadSystemKeyringMutex.Lock()
	ret, specificReturn := fake.loadSystemKeyringReturnsOnCall[len(fake.loadSystemKeyringArgsForCall)]
	fake.loadSystemKeyringArgsForCall = append(fake.loadSystemKeyringArgsForCall, struct {
		arg1 *options.Options
		arg2 []string
	}{arg1, arg2})
	stub := fake.LoadSystemKeyringStub
	fakeReturns := fake.loadSystemKeyringReturns
	fake.recordInvocation("LoadSystemKeyring", []interface{}{arg1, arg2})
	fake.loadSystemKeyringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApkImplementation) LoadSystemKeyringCallCount() int {
	fake.loadSystemKeyringMutex.RLock()
	defer fake.loadSystemKeyringMutex.RUnlock()
	return len(fake.loadSystemKeyringArgsForCall)
}

func (fake *FakeApkImplementation) LoadSystemKeyringCalls(stub func(*options.Options, ...string) ([]string, error)) {
	fake.loadSystemKeyringMutex.Lock()
	defer fake.loadSystemKeyringMutex.Unlock()
	fake.LoadSystemKeyringStub = stub
}

func (fake *FakeApkImplementation) LoadSystemKeyringArgsForCall(i int) (*options.Options, []string) {
	fake.loadSystemKeyringMutex.RLock()
	defer fake.loadSystemKeyringMutex.RUnlock()
	argsForCall := fake.loadSystemKeyringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) LoadSystemKeyringReturns(result1 []string, result2 error) {
	fake.loadSystemKeyringMutex.Lock()
	defer fake.loadSystemKeyringMutex.Unlock()
	fake.LoadSystemKeyringStub = nil
	fake.loadSystemKeyringReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) LoadSystemKeyringReturnsOnCall(i int, result1 []string, result2 error) {
	fake.loadSystemKeyringMutex.Lock()
	defer fake.loadSystemKeyringMutex.Unlock()
	fake.LoadSystemKeyringStub = nil
	if fake.loadSystemKeyringReturnsOnCall == nil {
		fake.loadSystemKeyringReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.loadSystemKeyringReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) NormalizeScriptsTar(arg1 *options.Options) error {
	fake.normalizeScriptsTarMutex.Lock()
	ret, specificReturn := fake.normalizeScriptsTarReturnsOnCall[len(fake.normalizeScriptsTarArgsForCall)]
	fake.normalizeScriptsTarArgsForCall = append(fake.normalizeScriptsTarArgsForCall, struct {
		arg1 *options.Options
	}{arg1})
	stub := fake.NormalizeScriptsTarStub
	fakeReturns := fake.normalizeScriptsTarReturns
	fake.recordInvocation("NormalizeScriptsTar", []interface{}{arg1})
	fake.normalizeScriptsTarMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) NormalizeScriptsTarCallCount() int {
	fake.normalizeScriptsTarMutex.RLock()
	defer fake.normalizeScriptsTarMutex.RUnlock()
	return len(fake.normalizeScriptsTarArgsForCall)
}

func (fake *FakeApkImplementation) NormalizeScriptsTarCalls(stub func(*options.Options) error) {
	fake.normalizeScriptsTarMutex.Lock()
	defer fake.normalizeScriptsTarMutex.Unlock()
	fake.NormalizeScriptsTarStub = stub
}

func (fake *FakeApkImplementation) NormalizeScriptsTarArgsForCall(i int) *options.Options {
	fake.normalizeScriptsTarMutex.RLock()
	defer fake.normalizeScriptsTarMutex.RUnlock()
	argsForCall := fake.normalizeScriptsTarArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeApkImplementation) NormalizeScriptsTarReturns(result1 error) {
	fake.normalizeScriptsTarMutex.Lock()
	defer fake.normalizeScriptsTarMutex.Unlock()
	fake.NormalizeScriptsTarStub = nil
	fake.normalizeScriptsTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) NormalizeScriptsTarReturnsOnCall(i int, result1 error) {
	fake.normalizeScriptsTarMutex.Lock()
	defer fake.normalizeScriptsTarMutex.Unlock()
	fake.NormalizeScriptsTarStub = nil
	if fake.normalizeScriptsTarReturnsOnCall == nil {
		fake.normalizeScriptsTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.normalizeScriptsTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	fake.initRepositoriesMutex.RLock()
	defer fake.initRepositoriesMutex.RUnlock()
	fake.initWorldMutex.RLock()
	defer fake.initWorldMutex.RUnlock()
	fake.loadSystemKeyringMutex.RLock()
	defer fake.loadSystemKeyringMutex.RUnlock()
	fake.normalizeScriptsTarMutex.RLock()
	defer fake.normalizeScriptsTarMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApkImplementation) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
