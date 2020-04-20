package schema

import (
	tfjson "github.com/hashicorp/terraform-json"
)

func MockStorage(ps *tfjson.ProviderSchemas) *Storage {
	s := NewStorage()
	if ps == nil {
		ps = &tfjson.ProviderSchemas{}
	}
	s.ps = ps
	s.sync = true
	s.w = &MockWatcher{}
	return s
}

type MockReader struct {
	ProviderSchemas *tfjson.ProviderSchemas

	ProviderSchemaErr   error
	ResourceSchemaErr   error
	DataSourceSchemaErr error
}

func (r *MockReader) storage() *Storage {
	return MockStorage(r.ProviderSchemas)
}

func (r *MockReader) ProviderConfigSchema(name string) (*tfjson.Schema, error) {
	if r.ProviderSchemaErr != nil {
		return nil, r.ProviderSchemaErr
	}
	return r.storage().ProviderConfigSchema(name)
}

func (r *MockReader) ResourceSchema(rType string) (*tfjson.Schema, error) {
	if r.ResourceSchemaErr != nil {
		return nil, r.ResourceSchemaErr
	}
	return r.storage().ResourceSchema(rType)
}

func (r *MockReader) DataSourceSchema(dsType string) (*tfjson.Schema, error) {
	if r.DataSourceSchemaErr != nil {
		return nil, r.DataSourceSchemaErr
	}
	return r.storage().DataSourceSchema(dsType)
}
