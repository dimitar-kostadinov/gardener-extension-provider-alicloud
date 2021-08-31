package alidns

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

// ChangeDomainOfDnsProduct invokes the alidns.ChangeDomainOfDnsProduct API synchronously
// api document: https://help.aliyun.com/api/alidns/changedomainofdnsproduct.html
func (client *Client) ChangeDomainOfDnsProduct(request *ChangeDomainOfDnsProductRequest) (response *ChangeDomainOfDnsProductResponse, err error) {
	response = CreateChangeDomainOfDnsProductResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeDomainOfDnsProductWithChan invokes the alidns.ChangeDomainOfDnsProduct API asynchronously
// api document: https://help.aliyun.com/api/alidns/changedomainofdnsproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeDomainOfDnsProductWithChan(request *ChangeDomainOfDnsProductRequest) (<-chan *ChangeDomainOfDnsProductResponse, <-chan error) {
	responseChan := make(chan *ChangeDomainOfDnsProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeDomainOfDnsProduct(request)
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

// ChangeDomainOfDnsProductWithCallback invokes the alidns.ChangeDomainOfDnsProduct API asynchronously
// api document: https://help.aliyun.com/api/alidns/changedomainofdnsproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeDomainOfDnsProductWithCallback(request *ChangeDomainOfDnsProductRequest, callback func(response *ChangeDomainOfDnsProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeDomainOfDnsProductResponse
		var err error
		defer close(result)
		response, err = client.ChangeDomainOfDnsProduct(request)
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

// ChangeDomainOfDnsProductRequest is the request struct for api ChangeDomainOfDnsProduct
type ChangeDomainOfDnsProductRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	NewDomain    string           `position:"Query" name:"NewDomain"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Force        requests.Boolean `position:"Query" name:"Force"`
	Lang         string           `position:"Query" name:"Lang"`
}

// ChangeDomainOfDnsProductResponse is the response struct for api ChangeDomainOfDnsProduct
type ChangeDomainOfDnsProductResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	OriginalDomain string `json:"OriginalDomain" xml:"OriginalDomain"`
}

// CreateChangeDomainOfDnsProductRequest creates a request to invoke ChangeDomainOfDnsProduct API
func CreateChangeDomainOfDnsProductRequest() (request *ChangeDomainOfDnsProductRequest) {
	request = &ChangeDomainOfDnsProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "ChangeDomainOfDnsProduct", "alidns", "openAPI")
	return
}

// CreateChangeDomainOfDnsProductResponse creates a response to parse from ChangeDomainOfDnsProduct response
func CreateChangeDomainOfDnsProductResponse() (response *ChangeDomainOfDnsProductResponse) {
	response = &ChangeDomainOfDnsProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}