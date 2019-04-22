// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type UploadStemcellService struct {
	GetDiagnosticReportStub        func() (api.DiagnosticReport, error)
	getDiagnosticReportMutex       sync.RWMutex
	getDiagnosticReportArgsForCall []struct {
	}
	getDiagnosticReportReturns struct {
		result1 api.DiagnosticReport
		result2 error
	}
	getDiagnosticReportReturnsOnCall map[int]struct {
		result1 api.DiagnosticReport
		result2 error
	}
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
	UploadStemcellStub        func(api.StemcellUploadInput) (api.StemcellUploadOutput, error)
	uploadStemcellMutex       sync.RWMutex
	uploadStemcellArgsForCall []struct {
		arg1 api.StemcellUploadInput
	}
	uploadStemcellReturns struct {
		result1 api.StemcellUploadOutput
		result2 error
	}
	uploadStemcellReturnsOnCall map[int]struct {
		result1 api.StemcellUploadOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UploadStemcellService) GetDiagnosticReport() (api.DiagnosticReport, error) {
	fake.getDiagnosticReportMutex.Lock()
	ret, specificReturn := fake.getDiagnosticReportReturnsOnCall[len(fake.getDiagnosticReportArgsForCall)]
	fake.getDiagnosticReportArgsForCall = append(fake.getDiagnosticReportArgsForCall, struct {
	}{})
	fake.recordInvocation("GetDiagnosticReport", []interface{}{})
	fake.getDiagnosticReportMutex.Unlock()
	if fake.GetDiagnosticReportStub != nil {
		return fake.GetDiagnosticReportStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDiagnosticReportReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadStemcellService) GetDiagnosticReportCallCount() int {
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	return len(fake.getDiagnosticReportArgsForCall)
}

func (fake *UploadStemcellService) GetDiagnosticReportCalls(stub func() (api.DiagnosticReport, error)) {
	fake.getDiagnosticReportMutex.Lock()
	defer fake.getDiagnosticReportMutex.Unlock()
	fake.GetDiagnosticReportStub = stub
}

func (fake *UploadStemcellService) GetDiagnosticReportReturns(result1 api.DiagnosticReport, result2 error) {
	fake.getDiagnosticReportMutex.Lock()
	defer fake.getDiagnosticReportMutex.Unlock()
	fake.GetDiagnosticReportStub = nil
	fake.getDiagnosticReportReturns = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *UploadStemcellService) GetDiagnosticReportReturnsOnCall(i int, result1 api.DiagnosticReport, result2 error) {
	fake.getDiagnosticReportMutex.Lock()
	defer fake.getDiagnosticReportMutex.Unlock()
	fake.GetDiagnosticReportStub = nil
	if fake.getDiagnosticReportReturnsOnCall == nil {
		fake.getDiagnosticReportReturnsOnCall = make(map[int]struct {
			result1 api.DiagnosticReport
			result2 error
		})
	}
	fake.getDiagnosticReportReturnsOnCall[i] = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *UploadStemcellService) Info() (api.Info, error) {
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

func (fake *UploadStemcellService) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *UploadStemcellService) InfoCalls(stub func() (api.Info, error)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *UploadStemcellService) InfoReturns(result1 api.Info, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 api.Info
		result2 error
	}{result1, result2}
}

func (fake *UploadStemcellService) InfoReturnsOnCall(i int, result1 api.Info, result2 error) {
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

func (fake *UploadStemcellService) UploadStemcell(arg1 api.StemcellUploadInput) (api.StemcellUploadOutput, error) {
	fake.uploadStemcellMutex.Lock()
	ret, specificReturn := fake.uploadStemcellReturnsOnCall[len(fake.uploadStemcellArgsForCall)]
	fake.uploadStemcellArgsForCall = append(fake.uploadStemcellArgsForCall, struct {
		arg1 api.StemcellUploadInput
	}{arg1})
	fake.recordInvocation("UploadStemcell", []interface{}{arg1})
	fake.uploadStemcellMutex.Unlock()
	if fake.UploadStemcellStub != nil {
		return fake.UploadStemcellStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.uploadStemcellReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadStemcellService) UploadStemcellCallCount() int {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	return len(fake.uploadStemcellArgsForCall)
}

func (fake *UploadStemcellService) UploadStemcellCalls(stub func(api.StemcellUploadInput) (api.StemcellUploadOutput, error)) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = stub
}

func (fake *UploadStemcellService) UploadStemcellArgsForCall(i int) api.StemcellUploadInput {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	argsForCall := fake.uploadStemcellArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UploadStemcellService) UploadStemcellReturns(result1 api.StemcellUploadOutput, result2 error) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = nil
	fake.uploadStemcellReturns = struct {
		result1 api.StemcellUploadOutput
		result2 error
	}{result1, result2}
}

func (fake *UploadStemcellService) UploadStemcellReturnsOnCall(i int, result1 api.StemcellUploadOutput, result2 error) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = nil
	if fake.uploadStemcellReturnsOnCall == nil {
		fake.uploadStemcellReturnsOnCall = make(map[int]struct {
			result1 api.StemcellUploadOutput
			result2 error
		})
	}
	fake.uploadStemcellReturnsOnCall[i] = struct {
		result1 api.StemcellUploadOutput
		result2 error
	}{result1, result2}
}

func (fake *UploadStemcellService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *UploadStemcellService) recordInvocation(key string, args []interface{}) {
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
