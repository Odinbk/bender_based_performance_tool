package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/http/cases"
)

type IppIdDateDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionOneDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
				"dimensions": ["ipp_id", "date"],
				"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionOneDayAllCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01"
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionAllDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue", "units"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p IppIdDateDimensionTestGroup) TestIppIdDateDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s
					},
					"dimensions": ["ipp_id", "date"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}
