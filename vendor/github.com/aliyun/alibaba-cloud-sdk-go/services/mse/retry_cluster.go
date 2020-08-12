package mse

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

// RetryCluster invokes the mse.RetryCluster API synchronously
// api document: https://help.aliyun.com/api/mse/retrycluster.html
func (client *Client) RetryCluster(request *RetryClusterRequest) (response *RetryClusterResponse, err error) {
	response = CreateRetryClusterResponse()
	err = client.DoAction(request, response)
	return
}

// RetryClusterWithChan invokes the mse.RetryCluster API asynchronously
// api document: https://help.aliyun.com/api/mse/retrycluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RetryClusterWithChan(request *RetryClusterRequest) (<-chan *RetryClusterResponse, <-chan error) {
	responseChan := make(chan *RetryClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetryCluster(request)
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

// RetryClusterWithCallback invokes the mse.RetryCluster API asynchronously
// api document: https://help.aliyun.com/api/mse/retrycluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RetryClusterWithCallback(request *RetryClusterRequest, callback func(response *RetryClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetryClusterResponse
		var err error
		defer close(result)
		response, err = client.RetryCluster(request)
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

// RetryClusterRequest is the request struct for api RetryCluster
type RetryClusterRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	RequestPars string `position:"Query" name:"RequestPars"`
}

// RetryClusterResponse is the response struct for api RetryCluster
type RetryClusterResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRetryClusterRequest creates a request to invoke RetryCluster API
func CreateRetryClusterRequest() (request *RetryClusterRequest) {
	request = &RetryClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "RetryCluster", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetryClusterResponse creates a response to parse from RetryCluster response
func CreateRetryClusterResponse() (response *RetryClusterResponse) {
	response = &RetryClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
