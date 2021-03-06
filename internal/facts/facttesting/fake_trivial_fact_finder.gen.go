// Code generated by counterfeiter. DO NOT EDIT.
package facttesting

import (
	"context"
	"sync"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/service"
)

type FakeTrivialFactFinder struct {
	FindTrivialFactStub        func(context.Context) (facts.Topic, error)
	findTrivialFactMutex       sync.RWMutex
	findTrivialFactArgsForCall []struct {
		arg1 context.Context
	}
	findTrivialFactReturns struct {
		result1 facts.Topic
		result2 error
	}
	findTrivialFactReturnsOnCall map[int]struct {
		result1 facts.Topic
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTrivialFactFinder) FindTrivialFact(arg1 context.Context) (facts.Topic, error) {
	fake.findTrivialFactMutex.Lock()
	ret, specificReturn := fake.findTrivialFactReturnsOnCall[len(fake.findTrivialFactArgsForCall)]
	fake.findTrivialFactArgsForCall = append(fake.findTrivialFactArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FindTrivialFactStub
	fakeReturns := fake.findTrivialFactReturns
	fake.recordInvocation("FindTrivialFact", []interface{}{arg1})
	fake.findTrivialFactMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTrivialFactFinder) FindTrivialFactCallCount() int {
	fake.findTrivialFactMutex.RLock()
	defer fake.findTrivialFactMutex.RUnlock()
	return len(fake.findTrivialFactArgsForCall)
}

func (fake *FakeTrivialFactFinder) FindTrivialFactCalls(stub func(context.Context) (facts.Topic, error)) {
	fake.findTrivialFactMutex.Lock()
	defer fake.findTrivialFactMutex.Unlock()
	fake.FindTrivialFactStub = stub
}

func (fake *FakeTrivialFactFinder) FindTrivialFactArgsForCall(i int) context.Context {
	fake.findTrivialFactMutex.RLock()
	defer fake.findTrivialFactMutex.RUnlock()
	argsForCall := fake.findTrivialFactArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTrivialFactFinder) FindTrivialFactReturns(result1 facts.Topic, result2 error) {
	fake.findTrivialFactMutex.Lock()
	defer fake.findTrivialFactMutex.Unlock()
	fake.FindTrivialFactStub = nil
	fake.findTrivialFactReturns = struct {
		result1 facts.Topic
		result2 error
	}{result1, result2}
}

func (fake *FakeTrivialFactFinder) FindTrivialFactReturnsOnCall(i int, result1 facts.Topic, result2 error) {
	fake.findTrivialFactMutex.Lock()
	defer fake.findTrivialFactMutex.Unlock()
	fake.FindTrivialFactStub = nil
	if fake.findTrivialFactReturnsOnCall == nil {
		fake.findTrivialFactReturnsOnCall = make(map[int]struct {
			result1 facts.Topic
			result2 error
		})
	}
	fake.findTrivialFactReturnsOnCall[i] = struct {
		result1 facts.Topic
		result2 error
	}{result1, result2}
}

func (fake *FakeTrivialFactFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findTrivialFactMutex.RLock()
	defer fake.findTrivialFactMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTrivialFactFinder) recordInvocation(key string, args []interface{}) {
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

var _ service.TrivialFactFinder = new(FakeTrivialFactFinder)
