/*
Reward Cloud

Testing EnvironmentAccessDevToolsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rewardcloud

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rewardcloud_EnvironmentAccessDevToolsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAccessDevToolsApiService ApiEnvironmentAccessDevToolsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
