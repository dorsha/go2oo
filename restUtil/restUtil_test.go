package restUtil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHTTPClient(t *testing.T) {

	client := CreateHTTPClient()

	assert.NotNil(t, client, "HTTP client cannot be nil")
	assert.NotNil(t, client, "HTTP Transport cannot be nil")
}
