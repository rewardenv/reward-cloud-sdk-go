/*
Reward Cloud

Testing ComponentVersionEnvVarExampleApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rewardcloud

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rewardcloud_ComponentVersionEnvVarExampleApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentVersionEnvVarExampleApiService ApiComponentVersionEnvVarExamplesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}