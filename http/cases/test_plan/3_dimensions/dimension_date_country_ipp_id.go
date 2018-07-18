package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/http/cases"
)

type DateCountryIppIdDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionOneDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
				"dimensions": ["date", "country", "ipp_id"],
				"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionOneDayAllCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01"
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionAllDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"country": ["UK"]
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue", "units"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p DateCountryIppIdDimensionTestGroup) TestDateCountryIppIdDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s
					},
					"dimensions": ["date", "country", "ipp_id"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}
