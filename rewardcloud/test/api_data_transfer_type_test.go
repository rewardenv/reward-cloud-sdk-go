/*
Reward Cloud

Testing DataTransferTypeApiService

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

func Test_rewardcloud_DataTransferTypeApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataTransferTypeApiService ApiDataTransferTypesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
