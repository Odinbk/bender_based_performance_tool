package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/http/cases"
)

type IppIdDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionOneDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
				"dimensions": ["ipp_id"],
				"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionOneDayAllCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01"
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionAllDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"country": ["UK"]
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue", "units"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDimensionTestGroup) TestIppIdDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s
					},
					"dimensions": ["ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}
