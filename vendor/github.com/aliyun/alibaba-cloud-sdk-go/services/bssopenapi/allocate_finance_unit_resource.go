package bssopenapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AllocateFinanceUnitResource invokes the bssopenapi.AllocateFinanceUnitResource API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatefinanceunitresource.html
func (client *Client) AllocateFinanceUnitResource(request *AllocateFinanceUnitResourceRequest) (response *AllocateFinanceUnitResourceResponse, err error) {
	response = CreateAllocateFinanceUnitResourceResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateFinanceUnitResourceWithChan invokes the bssopenapi.AllocateFinanceUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatefinanceunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateFinanceUnitResourceWithChan(request *AllocateFinanceUnitResourceRequest) (<-chan *AllocateFinanceUnitResourceResponse, <-chan error) {
	responseChan := make(chan *AllocateFinanceUnitResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateFinanceUnitResource(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AllocateFinanceUnitResourceWithCallback invokes the bssopenapi.AllocateFinanceUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatefinanceunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateFinanceUnitResourceWithCallback(request *AllocateFinanceUnitResourceRequest, callback func(response *AllocateFinanceUnitResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateFinanceUnitResourceResponse
		var err error
		defer close(result)
		response, err = client.AllocateFinanceUnitResource(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AllocateFinanceUnitResourceRequest is the request struct for api AllocateFinanceUnitResource
type AllocateFinanceUnitResourceRequest struct {
	*requests.RpcRequest
	ResourceInstanceList *[]AllocateFinanceUnitResourceResourceInstanceList `position:"Query" name:"ResourceInstanceList"  type:"Repeated"`
	FromUnitId           requests.Integer                                   `position:"Query" name:"FromUnitId"`
	ToUnitId             requests.Integer                                   `position:"Query" name:"ToUnitId"`
	FromUnitUserId       requests.Integer                                   `position:"Query" name:"FromUnitUserId"`
	ToUnitUserId         requests.Integer                                   `position:"Query" name:"ToUnitUserId"`
}

// AllocateFinanceUnitResourceResourceInstanceList is a repeated param struct in AllocateFinanceUnitResourceRequest
type AllocateFinanceUnitResourceResourceInstanceList struct {
	ResourceId     string `name:"ResourceId"`
	CommodityCode  string `name:"CommodityCode"`
	ResourceUserId string `name:"ResourceUserId"`
}

// AllocateFinanceUnitResourceResponse is the response struct for api AllocateFinanceUnitResource
type AllocateFinanceUnitResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAllocateFinanceUnitResourceRequest creates a request to invoke AllocateFinanceUnitResource API
func CreateAllocateFinanceUnitResourceRequest() (request *AllocateFinanceUnitResourceRequest) {
	request = &AllocateFinanceUnitResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "AllocateFinanceUnitResource", "bssopenapi", "openAPI")
	return
}

// CreateAllocateFinanceUnitResourceResponse creates a response to parse from AllocateFinanceUnitResource response
func CreateAllocateFinanceUnitResourceResponse() (response *AllocateFinanceUnitResourceResponse) {
	response = &AllocateFinanceUnitResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}