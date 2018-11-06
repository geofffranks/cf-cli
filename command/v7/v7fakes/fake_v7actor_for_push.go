// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeV7ActorForPush struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(string, string, bool, v7action.RouteActor) (v7action.ApplicationSummary, v7action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 v7action.RouteActor
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v7action.ApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.ApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	PollStartStub        func(string, chan<- v7action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		arg1 string
		arg2 chan<- v7action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	RestartApplicationStub        func(string) (v7action.Warnings, error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		arg1 string
	}
	restartApplicationReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV7ActorForPush) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeV7ActorForPush) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV7ActorForPush) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeV7ActorForPush) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV7ActorForPush) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpace(arg1 string, arg2 string, arg3 bool, arg4 v7action.RouteActor) (v7action.ApplicationSummary, v7action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 v7action.RouteActor
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{arg1, arg2, arg3, arg4})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationSummaryByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpaceCalls(stub func(string, string, bool, v7action.RouteActor) (v7action.ApplicationSummary, v7action.Warnings, error)) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = stub
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string, bool, v7action.RouteActor) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationSummaryByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpaceReturns(result1 v7action.ApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v7action.ApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v7action.ApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.ApplicationSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.ApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v7action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v7action.LogMessage
			result2 <-chan error
			result3 v7action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV7ActorForPush) PollStart(arg1 string, arg2 chan<- v7action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		arg1 string
		arg2 chan<- v7action.Warnings
	}{arg1, arg2})
	fake.recordInvocation("PollStart", []interface{}{arg1, arg2})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pollStartReturns
	return fakeReturns.result1
}

func (fake *FakeV7ActorForPush) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeV7ActorForPush) PollStartCalls(stub func(string, chan<- v7action.Warnings) error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = stub
}

func (fake *FakeV7ActorForPush) PollStartArgsForCall(i int) (string, chan<- v7action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	argsForCall := fake.pollStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7ActorForPush) PollStartReturns(result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeV7ActorForPush) PollStartReturnsOnCall(i int, result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeV7ActorForPush) RestartApplication(arg1 string) (v7action.Warnings, error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RestartApplication", []interface{}{arg1})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.restartApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7ActorForPush) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakeV7ActorForPush) RestartApplicationCalls(stub func(string) (v7action.Warnings, error)) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = stub
}

func (fake *FakeV7ActorForPush) RestartApplicationArgsForCall(i int) string {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	argsForCall := fake.restartApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV7ActorForPush) RestartApplicationReturns(result1 v7action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) RestartApplicationReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV7ActorForPush) recordInvocation(key string, args []interface{}) {
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

var _ v7.V7ActorForPush = new(FakeV7ActorForPush)
