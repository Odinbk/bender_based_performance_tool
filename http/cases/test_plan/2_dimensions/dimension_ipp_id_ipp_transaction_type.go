package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/http/cases"
)

type IppIdIppTransactionTypeDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionOneDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
				"dimensions": ["ipp_id", "ipp_transaction_type"],
				"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionOneDayAllCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01"
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionAllDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue", "units"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdIppTransactionTypeDimensionTestGroup) TestIppIdIppTransactionTypeDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s
					},
					"dimensions": ["ipp_id", "ipp_transaction_type"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}
