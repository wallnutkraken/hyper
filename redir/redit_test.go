package redir

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUniqueID(t *testing.T) {
	r, _ := New("http://google.com/")
	assert.NotEmpty(t, r.ID())
}

func TestValidURL(t *testing.T) {
	_, err := New("http://google.com/")
	assert.NoError(t, err)
}

func TestInvalidURL(t *testing.T) {
	_, err := New("jsdhfiusdhfuisdhfdsuihfdsuk")
	assert.Error(t, err)
}
