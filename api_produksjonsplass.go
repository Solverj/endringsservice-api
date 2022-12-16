/*
Endringsservice

Dette er et api som har i ansvar å ta i mot endringer fra diverse kilder som mener noe om produksjonsplasser

API version: 1.0.0
Contact: no-spam@mattilsynet.no
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endringsservice-api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// ProduksjonsplassApiService ProduksjonsplassApi service
type ProduksjonsplassApiService service

type ApiPostProduksjonsplassRequest struct {
	ctx context.Context
	ApiService *ProduksjonsplassApiService
	produksjonsplass *Produksjonsplass
}

func (r ApiPostProduksjonsplassRequest) Produksjonsplass(produksjonsplass Produksjonsplass) ApiPostProduksjonsplassRequest {
	r.produksjonsplass = &produksjonsplass
	return r
}

func (r ApiPostProduksjonsplassRequest) Execute() (*Produksjonsplass, *http.Response, error) {
	return r.ApiService.PostProduksjonsplassExecute(r)
}

/*
PostProduksjonsplass les inn en produksjonsplass



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostProduksjonsplassRequest
*/
func (a *ProduksjonsplassApiService) PostProduksjonsplass(ctx context.Context) ApiPostProduksjonsplassRequest {
	return ApiPostProduksjonsplassRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Produksjonsplass
func (a *ProduksjonsplassApiService) PostProduksjonsplassExecute(r ApiPostProduksjonsplassRequest) (*Produksjonsplass, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Produksjonsplass
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProduksjonsplassApiService.PostProduksjonsplass")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/produksjonsplass"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.produksjonsplass == nil {
		return localVarReturnValue, nil, reportError("produksjonsplass is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.produksjonsplass
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
