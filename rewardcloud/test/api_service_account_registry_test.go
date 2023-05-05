/*
Reward Cloud

Testing ServiceAccountRegistryApiService

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

func Test_rewardcloud_ServiceAccountRegistryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAccountRegistryApiService ApiServiceAccountRegistriesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
