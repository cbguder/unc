// This file was generated by counterfeiter
package converterfakes

import (
	"sync"

	"github.com/cbguder/v2e/converter"
	"github.com/cbguder/v2e/models"
)

type FakeImporter struct {
	ImportStub        func(string) ([]models.Note, error)
	importMutex       sync.RWMutex
	importArgsForCall []struct {
		arg1 string
	}
	importReturns struct {
		result1 []models.Note
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImporter) Import(arg1 string) ([]models.Note, error) {
	fake.importMutex.Lock()
	fake.importArgsForCall = append(fake.importArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Import", []interface{}{arg1})
	fake.importMutex.Unlock()
	if fake.ImportStub != nil {
		return fake.ImportStub(arg1)
	} else {
		return fake.importReturns.result1, fake.importReturns.result2
	}
}

func (fake *FakeImporter) ImportCallCount() int {
	fake.importMutex.RLock()
	defer fake.importMutex.RUnlock()
	return len(fake.importArgsForCall)
}

func (fake *FakeImporter) ImportArgsForCall(i int) string {
	fake.importMutex.RLock()
	defer fake.importMutex.RUnlock()
	return fake.importArgsForCall[i].arg1
}

func (fake *FakeImporter) ImportReturns(result1 []models.Note, result2 error) {
	fake.ImportStub = nil
	fake.importReturns = struct {
		result1 []models.Note
		result2 error
	}{result1, result2}
}

func (fake *FakeImporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.importMutex.RLock()
	defer fake.importMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImporter) recordInvocation(key string, args []interface{}) {
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

var _ converter.Importer = new(FakeImporter)