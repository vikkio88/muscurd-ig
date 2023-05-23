package context_test

import (
	"muscurdig/context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialContextState(t *testing.T) {
	db := &MockDb{}
	ctx := context.NewAppContext(context.Setup, &MockDb{}, WindowMock{})
	assert.Equal(t, context.Setup, ctx.CurrentRoute())

	assert.IsType(t, WindowMock{}, ctx.GetWindow())
	assert.Equal(t, db, ctx.Db)
}

func TestRouting(t *testing.T) {
	ctx := context.NewAppContext(context.Setup, &MockDb{}, WindowMock{})
	assert.Equal(t, context.Setup, ctx.CurrentRoute())
	ctx.NavigateTo(context.About)
	assert.Equal(t, context.About, ctx.CurrentRoute())

	ctx.NavigateToWithParam(context.Details, "someId")
	assert.Equal(t, context.Details, ctx.CurrentRoute())
	assert.Equal(t, "someId", ctx.RouteParam.(string))

	ctx.NavigateTo(context.List)
	assert.Equal(t, context.List, ctx.CurrentRoute())
	assert.Nil(t, ctx.RouteParam)
}
