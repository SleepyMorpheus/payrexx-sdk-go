package payrexxsdk

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClientWithValidParams(t *testing.T) {
	client, err := NewClient("test_instance", "test_secret", APIBaseDefault)
	assert.NotNil(t, client)
	assert.Nil(t, err)

	assert.Equal(t, "test_instance", client.InstanceName)
	assert.Equal(t, "test_secret", client.Secret)
	assert.Equal(t, APIBaseDefault, client.ApiUrl)
}

func TestNewClientWithInvalidParams(t *testing.T) {
	client, err := NewClient("", "test_secret", APIBaseDefault)
	assert.Nil(t, client)
	assert.NotNil(t, err)

	client, err = NewClient("test_instance", "", APIBaseDefault)
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestClient_CheckSignatureWithInvalidParams(t *testing.T) {
	client, _ := NewClient("instance_name", "test_secret", APIBaseDefault)
	err := client.CheckSignature(context.Background())
	assert.NotNil(t, err)

	var APIError *APIError
	if !errors.As(err, &APIError) {
		t.Errorf("Expected error of type *APIError, got %T", err)
	}
}
