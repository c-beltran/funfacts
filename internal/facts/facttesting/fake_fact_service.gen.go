// Code generated by counterfeiter. DO NOT EDIT.
package facttesting

import (
	"context"
	"sync"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/http/rest"
)

type FakeFactSvc struct {
	FindStub        func(context.Context, facts.TopicType) (facts.Topic, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 context.Context
		arg2 facts.TopicType
	}
	findReturns struct {
		result1 facts.Topic
		result2 error
	}
	findReturnsOnCall map[int]struct {
		result1 facts.Topic
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactSvc) Find(arg1 context.Context, arg2 facts.TopicType) (facts.Topic, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 context.Context
		arg2 facts.TopicType
	}{arg1, arg2})
	stub := fake.FindStub
	fakeReturns := fake.findReturns
	fake.recordInvocation("Find", []interface{}{arg1, arg2})
	fake.findMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFactSvc) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeFactSvc) FindCalls(stub func(context.Context, facts.TopicType) (facts.Topic, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeFactSvc) FindArgsForCall(i int) (context.Context, facts.TopicType) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFactSvc) FindReturns(result1 facts.Topic, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 facts.Topic
		result2 error
	}{result1, result2}
}

func (fake *FakeFactSvc) FindReturnsOnCall(i int, result1 facts.Topic, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 facts.Topic
			result2 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 facts.Topic
		result2 error
	}{result1, result2}
}

func (fake *FakeFactSvc) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFactSvc) recordInvocation(key string, args []interface{}) {
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

var _ rest.FactSvc = new(FakeFactSvc)