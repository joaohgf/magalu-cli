package render

import (
	"context"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	t.Run("successfully", func(t *testing.T) {
		t.Run("JSON", renderJSONSuccessfully)
		t.Run("YAML", renderYAMLSuccessfully)
		t.Run("Table", renderTableSuccessfully)
	})
	t.Run("with error", func(t *testing.T) {
		t.Run("JSON", renderJSONWithError)
		t.Run("YAML", renderYAMLWithError)
		t.Run("Table", renderTableWithError)
	})
}

func renderJSONSuccessfully(t *testing.T) {
	render := NewRender(new(mockView))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputJSON)
	err := render.Render(ctx, struct{}{})
	assert.NoError(t, err)
}

func renderYAMLSuccessfully(t *testing.T) {
	render := NewRender(new(mockView))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputYAML)
	err := render.Render(ctx, struct{}{})
	assert.NoError(t, err)
}

func renderTableSuccessfully(t *testing.T) {
	render := NewRender(new(mockView))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputTable)
	err := render.Render(ctx, struct{}{})
	assert.NoError(t, err)
}

func renderJSONWithError(t *testing.T) {
	render := NewRender(new(mockViewWithError))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputJSON)
	err := render.Render(ctx, struct{}{})
	assert.Error(t, err)
}

func renderYAMLWithError(t *testing.T) {
	render := NewRender(new(mockViewWithError))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputYAML)
	err := render.Render(ctx, struct{}{})
	assert.Error(t, err)
}

func renderTableWithError(t *testing.T) {
	render := NewRender(new(mockViewWithError))
	ctx := context.WithValue(t.Context(), enum.OutputDefaultKey.String(), enum.OutputTable)
	err := render.Render(ctx, struct{}{})
	assert.Error(t, err)
}

/*
Mocks below
*/

type (
	mockView          struct{}
	mockViewWithError struct{}
)

func (m *mockView) JSON(_ any) error {
	return nil
}

func (m *mockView) YAML(_ any) error {
	return nil
}

func (m *mockView) Table(_ any) error {
	return nil
}

func (m *mockViewWithError) JSON(_ any) error {
	return errors.New("mock error")
}

func (m *mockViewWithError) YAML(_ any) error {
	return errors.New("mock error")
}

func (m *mockViewWithError) Table(_ any) error {
	return errors.New("mock error")
}
