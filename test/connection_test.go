package test

import (
	"github.com/adibhauzan/master_data/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDB(t *testing.T) {
	db := app.NewDB()
	assert.NotNil(t, db)
}
