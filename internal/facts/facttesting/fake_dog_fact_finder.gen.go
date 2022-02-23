// Code generated by counterfeiter. DO NOT EDIT.
package facttesting

import (
	"context"
	"sync"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/service"
)

type FakeDogFactFinder struct {
	FindDogFactStub        func(context.Context) (facts.FactTopic, error)
	findDogFactMutex       sync.RWMutex
	findDogFactArgsForCall []struct {
		arg1 context.Context
	}
	findDogFactReturns struct {
		result1 facts.FactTopic
		result2 error
	}
	findDogFactReturnsOnCall map[int]struct {
		result1 facts.FactTopic
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDogFactFinder) FindDogFact(arg1 context.Context) (facts.FactTopic, error) {
	fake.findDogFactMutex.Lock()
	ret, specificReturn := fake.findDogFactReturnsOnCall[len(fake.findDogFactArgsForCall)]
	fake.findDogFactArgsForCall = append(fake.findDogFactArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FindDogFactStub
	fakeReturns := fake.findDogFactReturns
	fake.recordInvocation("FindDogFact", []interface{}{arg1})
	fake.findDogFactMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDogFactFinder) FindDogFactCallCount() int {
	fake.findDogFactMutex.RLock()
	defer fake.findDogFactMutex.RUnlock()
	return len(fake.findDogFactArgsForCall)
}

func (fake *FakeDogFactFinder) FindDogFactCalls(stub func(context.Context) (facts.FactTopic, error)) {
	fake.findDogFactMutex.Lock()
	defer fake.findDogFactMutex.Unlock()
	fake.FindDogFactStub = stub
}

func (fake *FakeDogFactFinder) FindDogFactArgsForCall(i int) context.Context {
	fake.findDogFactMutex.RLock()
	defer fake.findDogFactMutex.RUnlock()
	argsForCall := fake.findDogFactArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDogFactFinder) FindDogFactReturns(result1 facts.FactTopic, result2 error) {
	fake.findDogFactMutex.Lock()
	defer fake.findDogFactMutex.Unlock()
	fake.FindDogFactStub = nil
	fake.findDogFactReturns = struct {
		result1 facts.FactTopic
		result2 error
	}{result1, result2}
}

func (fake *FakeDogFactFinder) FindDogFactReturnsOnCall(i int, result1 facts.FactTopic, result2 error) {
	fake.findDogFactMutex.Lock()
	defer fake.findDogFactMutex.Unlock()
	fake.FindDogFactStub = nil
	if fake.findDogFactReturnsOnCall == nil {
		fake.findDogFactReturnsOnCall = make(map[int]struct {
			result1 facts.FactTopic
			result2 error
		})
	}
	fake.findDogFactReturnsOnCall[i] = struct {
		result1 facts.FactTopic
		result2 error
	}{result1, result2}
}

func (fake *FakeDogFactFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findDogFactMutex.RLock()
	defer fake.findDogFactMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDogFactFinder) recordInvocation(key string, args []interface{}) {
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

var _ service.DogFactFinder = new(FakeDogFactFinder)
