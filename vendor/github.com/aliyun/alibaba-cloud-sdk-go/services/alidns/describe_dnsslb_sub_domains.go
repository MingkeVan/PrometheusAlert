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

// DescribeDNSSLBSubDomains invokes the alidns.DescribeDNSSLBSubDomains API synchronously
// api document: https://help.aliyun.com/api/alidns/describednsslbsubdomains.html
func (client *Client) DescribeDNSSLBSubDomains(request *DescribeDNSSLBSubDomainsRequest) (response *DescribeDNSSLBSubDomainsResponse, err error) {
	response = CreateDescribeDNSSLBSubDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDNSSLBSubDomainsWithChan invokes the alidns.DescribeDNSSLBSubDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/describednsslbsubdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDNSSLBSubDomainsWithChan(request *DescribeDNSSLBSubDomainsRequest) (<-chan *DescribeDNSSLBSubDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDNSSLBSubDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDNSSLBSubDomains(request)
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

// DescribeDNSSLBSubDomainsWithCallback invokes the alidns.DescribeDNSSLBSubDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/describednsslbsubdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDNSSLBSubDomainsWithCallback(request *DescribeDNSSLBSubDomainsRequest, callback func(response *DescribeDNSSLBSubDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDNSSLBSubDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDNSSLBSubDomains(request)
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

// DescribeDNSSLBSubDomainsRequest is the request struct for api DescribeDNSSLBSubDomains
type DescribeDNSSLBSubDomainsRequest struct {
	*requests.RpcRequest
	DomainName   string           `position:"Query" name:"DomainName"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeDNSSLBSubDomainsResponse is the response struct for api DescribeDNSSLBSubDomains
type DescribeDNSSLBSubDomainsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	TotalCount    int64         `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int64         `json:"PageNumber" xml:"PageNumber"`
	PageSize      int64         `json:"PageSize" xml:"PageSize"`
	SlbSubDomains SlbSubDomains `json:"SlbSubDomains" xml:"SlbSubDomains"`
}

// CreateDescribeDNSSLBSubDomainsRequest creates a request to invoke DescribeDNSSLBSubDomains API
func CreateDescribeDNSSLBSubDomainsRequest() (request *DescribeDNSSLBSubDomainsRequest) {
	request = &DescribeDNSSLBSubDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDNSSLBSubDomains", "alidns", "openAPI")
	return
}

// CreateDescribeDNSSLBSubDomainsResponse creates a response to parse from DescribeDNSSLBSubDomains response
func CreateDescribeDNSSLBSubDomainsResponse() (response *DescribeDNSSLBSubDomainsResponse) {
	response = &DescribeDNSSLBSubDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}