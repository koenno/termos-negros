package portal

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnProperRequest(t *testing.T) {
	// given
	sut := NewRequestFactory()

	// when
	req, err := sut.NewRequest(context.Background())

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "portal.czarnetermosy.pl", req.URL.Host)
	assert.Equal(t, "https", req.URL.Scheme)
	assert.Equal(t, "/menu", req.URL.Path)
}
