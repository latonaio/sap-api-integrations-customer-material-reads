package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-material-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCustomerMaterial(raw []byte, l *logger.Logger) ([]CustomerMaterial, error) {
	pm := &responses.CustomerMaterial{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerMaterial. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	customerMaterial := make([]CustomerMaterial, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerMaterial = append(customerMaterial, CustomerMaterial{
			SalesOrganization:              data.SalesOrganization,
			DistributionChannel:            data.DistributionChannel,
			Customer:                       data.Customer,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialDescriptionByCustomer:  data.MaterialDescriptionByCustomer,
			Plant:                          data.Plant,
			DeliveryPriority:               data.DeliveryPriority,
			MinDeliveryQtyInBaseUnit:       data.MinDeliveryQtyInBaseUnit,
			BaseUnit:                       data.BaseUnit,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			MaxNmbrOfPartialDelivery:       data.MaxNmbrOfPartialDelivery,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			CustomerMaterialItemUsage:      data.CustomerMaterialItemUsage,
			SalesUnit:                      data.SalesUnit,
			SalesQtyToBaseQtyDnmntr:        data.SalesQtyToBaseQtyDnmntr,
			SalesQtyToBaseQtyNmrtr:         data.SalesQtyToBaseQtyNmrtr,
		})
	}

	return customerMaterial, nil
}
