// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type PreDeployCheckService struct {
	InfoStub        func() (api.Info, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
	}
	infoReturns struct {
		result1 api.Info
		result2 error
	}
	infoReturnsOnCall map[int]struct {
		result1 api.Info
		result2 error
	}
	ListAllPendingProductChangesStub        func() ([]api.PendingProductChangesOutput, error)
	listAllPendingProductChangesMutex       sync.RWMutex
	listAllPendingProductChangesArgsForCall []struct {
	}
	listAllPendingProductChangesReturns struct {
		result1 []api.PendingProductChangesOutput
		result2 error
	}
	listAllPendingProductChangesReturnsOnCall map[int]struct {
		result1 []api.PendingProductChangesOutput
		result2 error
	}
	ListPendingDirectorChangesStub        func() (api.PendingDirectorChangesOutput, error)
	listPendingDirectorChangesMutex       sync.RWMutex
	listPendingDirectorChangesArgsForCall []struct {
	}
	listPendingDirectorChangesReturns struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}
	listPendingDirectorChangesReturnsOnCall map[int]struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PreDeployCheckService) Info() (api.Info, error) {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
	}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.infoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PreDeployCheckService) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *PreDeployCheckService) InfoCalls(stub func() (api.Info, error)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *PreDeployCheckService) InfoReturns(result1 api.Info, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 api.Info
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) InfoReturnsOnCall(i int, result1 api.Info, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 api.Info
			result2 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 api.Info
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) ListAllPendingProductChanges() ([]api.PendingProductChangesOutput, error) {
	fake.listAllPendingProductChangesMutex.Lock()
	ret, specificReturn := fake.listAllPendingProductChangesReturnsOnCall[len(fake.listAllPendingProductChangesArgsForCall)]
	fake.listAllPendingProductChangesArgsForCall = append(fake.listAllPendingProductChangesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListAllPendingProductChanges", []interface{}{})
	fake.listAllPendingProductChangesMutex.Unlock()
	if fake.ListAllPendingProductChangesStub != nil {
		return fake.ListAllPendingProductChangesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listAllPendingProductChangesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PreDeployCheckService) ListAllPendingProductChangesCallCount() int {
	fake.listAllPendingProductChangesMutex.RLock()
	defer fake.listAllPendingProductChangesMutex.RUnlock()
	return len(fake.listAllPendingProductChangesArgsForCall)
}

func (fake *PreDeployCheckService) ListAllPendingProductChangesCalls(stub func() ([]api.PendingProductChangesOutput, error)) {
	fake.listAllPendingProductChangesMutex.Lock()
	defer fake.listAllPendingProductChangesMutex.Unlock()
	fake.ListAllPendingProductChangesStub = stub
}

func (fake *PreDeployCheckService) ListAllPendingProductChangesReturns(result1 []api.PendingProductChangesOutput, result2 error) {
	fake.listAllPendingProductChangesMutex.Lock()
	defer fake.listAllPendingProductChangesMutex.Unlock()
	fake.ListAllPendingProductChangesStub = nil
	fake.listAllPendingProductChangesReturns = struct {
		result1 []api.PendingProductChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) ListAllPendingProductChangesReturnsOnCall(i int, result1 []api.PendingProductChangesOutput, result2 error) {
	fake.listAllPendingProductChangesMutex.Lock()
	defer fake.listAllPendingProductChangesMutex.Unlock()
	fake.ListAllPendingProductChangesStub = nil
	if fake.listAllPendingProductChangesReturnsOnCall == nil {
		fake.listAllPendingProductChangesReturnsOnCall = make(map[int]struct {
			result1 []api.PendingProductChangesOutput
			result2 error
		})
	}
	fake.listAllPendingProductChangesReturnsOnCall[i] = struct {
		result1 []api.PendingProductChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) ListPendingDirectorChanges() (api.PendingDirectorChangesOutput, error) {
	fake.listPendingDirectorChangesMutex.Lock()
	ret, specificReturn := fake.listPendingDirectorChangesReturnsOnCall[len(fake.listPendingDirectorChangesArgsForCall)]
	fake.listPendingDirectorChangesArgsForCall = append(fake.listPendingDirectorChangesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListPendingDirectorChanges", []interface{}{})
	fake.listPendingDirectorChangesMutex.Unlock()
	if fake.ListPendingDirectorChangesStub != nil {
		return fake.ListPendingDirectorChangesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listPendingDirectorChangesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesCallCount() int {
	fake.listPendingDirectorChangesMutex.RLock()
	defer fake.listPendingDirectorChangesMutex.RUnlock()
	return len(fake.listPendingDirectorChangesArgsForCall)
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesCalls(stub func() (api.PendingDirectorChangesOutput, error)) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = stub
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesReturns(result1 api.PendingDirectorChangesOutput, result2 error) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = nil
	fake.listPendingDirectorChangesReturns = struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesReturnsOnCall(i int, result1 api.PendingDirectorChangesOutput, result2 error) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = nil
	if fake.listPendingDirectorChangesReturnsOnCall == nil {
		fake.listPendingDirectorChangesReturnsOnCall = make(map[int]struct {
			result1 api.PendingDirectorChangesOutput
			result2 error
		})
	}
	fake.listPendingDirectorChangesReturnsOnCall[i] = struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.listAllPendingProductChangesMutex.RLock()
	defer fake.listAllPendingProductChangesMutex.RUnlock()
	fake.listPendingDirectorChangesMutex.RLock()
	defer fake.listPendingDirectorChangesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PreDeployCheckService) recordInvocation(key string, args []interface{}) {
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