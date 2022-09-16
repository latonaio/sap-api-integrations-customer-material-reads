package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-customer-material-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetCustomerMaterial(salesOrganization, distributionChannel, customer, material string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CustomerMaterial":
			func() {
				c.CustomerMaterial(salesOrganization, distributionChannel, customer, material)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CustomerMaterial(salesOrganization, distributionChannel, customer, material string) {
	data, err := c.callCustomerMaterialSrvAPIRequirementCustomerMaterial("A_CustomerMaterial", salesOrganization, distributionChannel, customer, material)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callCustomerMaterialSrvAPIRequirementCustomerMaterial(api, salesOrganization, distributionChannel, customer, material string) ([]sap_api_output_formatter.CustomerMaterial, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMER_MATERIAL_SRV", api}, "/")
	param := c.getQueryWithCustomerMaterial(map[string]string{}, salesOrganization, distributionChannel, customer, material)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerMaterial(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithCustomerMaterial(params map[string]string, salesOrganization, distributionChannel, customer, material string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SalesOrganization eq '%s' and DistributionChannel eq '%s' and Customer eq '%s' and Material eq '%s'", salesOrganization, distributionChannel, customer, material)
	return params
}
