// Code generated by counterfeiter. DO NOT EDIT.
package filterfakes

import (
	"context"
	"sync"

	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/awsapi"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils/filter"
)

type FakeNodegroupFilter struct {
	LogInfoStub        func(*v1alpha5.ClusterConfig)
	logInfoMutex       sync.RWMutex
	logInfoArgsForCall []struct {
		arg1 *v1alpha5.ClusterConfig
	}
	MatchStub        func(string) bool
	matchMutex       sync.RWMutex
	matchArgsForCall []struct {
		arg1 string
	}
	matchReturns struct {
		result1 bool
	}
	matchReturnsOnCall map[int]struct {
		result1 bool
	}
	SetOnlyLocalStub        func(context.Context, awsapi.EKS, filter.StackLister, *v1alpha5.ClusterConfig) error
	setOnlyLocalMutex       sync.RWMutex
	setOnlyLocalArgsForCall []struct {
		arg1 context.Context
		arg2 awsapi.EKS
		arg3 filter.StackLister
		arg4 *v1alpha5.ClusterConfig
	}
	setOnlyLocalReturns struct {
		result1 error
	}
	setOnlyLocalReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNodegroupFilter) LogInfo(arg1 *v1alpha5.ClusterConfig) {
	fake.logInfoMutex.Lock()
	fake.logInfoArgsForCall = append(fake.logInfoArgsForCall, struct {
		arg1 *v1alpha5.ClusterConfig
	}{arg1})
	stub := fake.LogInfoStub
	fake.recordInvocation("LogInfo", []interface{}{arg1})
	fake.logInfoMutex.Unlock()
	if stub != nil {
		fake.LogInfoStub(arg1)
	}
}

func (fake *FakeNodegroupFilter) LogInfoCallCount() int {
	fake.logInfoMutex.RLock()
	defer fake.logInfoMutex.RUnlock()
	return len(fake.logInfoArgsForCall)
}

func (fake *FakeNodegroupFilter) LogInfoCalls(stub func(*v1alpha5.ClusterConfig)) {
	fake.logInfoMutex.Lock()
	defer fake.logInfoMutex.Unlock()
	fake.LogInfoStub = stub
}

func (fake *FakeNodegroupFilter) LogInfoArgsForCall(i int) *v1alpha5.ClusterConfig {
	fake.logInfoMutex.RLock()
	defer fake.logInfoMutex.RUnlock()
	argsForCall := fake.logInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodegroupFilter) Match(arg1 string) bool {
	fake.matchMutex.Lock()
	ret, specificReturn := fake.matchReturnsOnCall[len(fake.matchArgsForCall)]
	fake.matchArgsForCall = append(fake.matchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MatchStub
	fakeReturns := fake.matchReturns
	fake.recordInvocation("Match", []interface{}{arg1})
	fake.matchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodegroupFilter) MatchCallCount() int {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	return len(fake.matchArgsForCall)
}

func (fake *FakeNodegroupFilter) MatchCalls(stub func(string) bool) {
	fake.matchMutex.Lock()
	defer fake.matchMutex.Unlock()
	fake.MatchStub = stub
}

func (fake *FakeNodegroupFilter) MatchArgsForCall(i int) string {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	argsForCall := fake.matchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodegroupFilter) MatchReturns(result1 bool) {
	fake.matchMutex.Lock()
	defer fake.matchMutex.Unlock()
	fake.MatchStub = nil
	fake.matchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeNodegroupFilter) MatchReturnsOnCall(i int, result1 bool) {
	fake.matchMutex.Lock()
	defer fake.matchMutex.Unlock()
	fake.MatchStub = nil
	if fake.matchReturnsOnCall == nil {
		fake.matchReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.matchReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeNodegroupFilter) SetOnlyLocal(arg1 context.Context, arg2 awsapi.EKS, arg3 filter.StackLister, arg4 *v1alpha5.ClusterConfig) error {
	fake.setOnlyLocalMutex.Lock()
	ret, specificReturn := fake.setOnlyLocalReturnsOnCall[len(fake.setOnlyLocalArgsForCall)]
	fake.setOnlyLocalArgsForCall = append(fake.setOnlyLocalArgsForCall, struct {
		arg1 context.Context
		arg2 awsapi.EKS
		arg3 filter.StackLister
		arg4 *v1alpha5.ClusterConfig
	}{arg1, arg2, arg3, arg4})
	stub := fake.SetOnlyLocalStub
	fakeReturns := fake.setOnlyLocalReturns
	fake.recordInvocation("SetOnlyLocal", []interface{}{arg1, arg2, arg3, arg4})
	fake.setOnlyLocalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodegroupFilter) SetOnlyLocalCallCount() int {
	fake.setOnlyLocalMutex.RLock()
	defer fake.setOnlyLocalMutex.RUnlock()
	return len(fake.setOnlyLocalArgsForCall)
}

func (fake *FakeNodegroupFilter) SetOnlyLocalCalls(stub func(context.Context, awsapi.EKS, filter.StackLister, *v1alpha5.ClusterConfig) error) {
	fake.setOnlyLocalMutex.Lock()
	defer fake.setOnlyLocalMutex.Unlock()
	fake.SetOnlyLocalStub = stub
}

func (fake *FakeNodegroupFilter) SetOnlyLocalArgsForCall(i int) (context.Context, awsapi.EKS, filter.StackLister, *v1alpha5.ClusterConfig) {
	fake.setOnlyLocalMutex.RLock()
	defer fake.setOnlyLocalMutex.RUnlock()
	argsForCall := fake.setOnlyLocalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNodegroupFilter) SetOnlyLocalReturns(result1 error) {
	fake.setOnlyLocalMutex.Lock()
	defer fake.setOnlyLocalMutex.Unlock()
	fake.SetOnlyLocalStub = nil
	fake.setOnlyLocalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodegroupFilter) SetOnlyLocalReturnsOnCall(i int, result1 error) {
	fake.setOnlyLocalMutex.Lock()
	defer fake.setOnlyLocalMutex.Unlock()
	fake.SetOnlyLocalStub = nil
	if fake.setOnlyLocalReturnsOnCall == nil {
		fake.setOnlyLocalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setOnlyLocalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodegroupFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.logInfoMutex.RLock()
	defer fake.logInfoMutex.RUnlock()
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	fake.setOnlyLocalMutex.RLock()
	defer fake.setOnlyLocalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNodegroupFilter) recordInvocation(key string, args []interface{}) {
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

var _ filter.NodegroupFilter = new(FakeNodegroupFilter)
