package sap_api_output_formatter

type CustomerMaterialReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	Material      string `json:"material_code"`
	Deleted       bool   `json:"deleted"`
}

type CustomerMaterial struct {
	SalesOrganization              string `json:"SalesOrganization"`
	DistributionChannel            string `json:"DistributionChannel"`
	Customer                       string `json:"Customer"`
	Material                       string `json:"Material"`
	MaterialByCustomer             string `json:"MaterialByCustomer"`
	MaterialDescriptionByCustomer  string `json:"MaterialDescriptionByCustomer"`
	Plant                          string `json:"Plant"`
	DeliveryPriority               string `json:"DeliveryPriority"`
	MinDeliveryQtyInBaseUnit       string `json:"MinDeliveryQtyInBaseUnit"`
	BaseUnit                       string `json:"BaseUnit"`
	PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
	MaxNmbrOfPartialDelivery       string `json:"MaxNmbrOfPartialDelivery"`
	UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
	OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
	UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
	CustomerMaterialItemUsage      string `json:"CustomerMaterialItemUsage"`
	SalesUnit                      string `json:"SalesUnit"`
	SalesQtyToBaseQtyDnmntr        string `json:"SalesQtyToBaseQtyDnmntr"`
	SalesQtyToBaseQtyNmrtr         string `json:"SalesQtyToBaseQtyNmrtr"`
}
