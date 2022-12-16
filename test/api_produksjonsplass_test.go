/*
Endringsservice

Testing ProduksjonsplassApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package endringsservice-api

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_endringsservice-api_ProduksjonsplassApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProduksjonsplassApiService PostProduksjonsplass", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ProduksjonsplassApi.PostProduksjonsplass(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
