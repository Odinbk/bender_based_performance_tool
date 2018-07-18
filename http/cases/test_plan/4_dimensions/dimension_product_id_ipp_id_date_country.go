package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/http/cases"
)

type ProductIdDateIppIdCountryDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionOneDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
				"dimensions": ["product_id", "date", "ipp_id", "country"],
				"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionOneDayAllCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01"
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionAllDayOneCountry() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"country": ["UK"]
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue", "units"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"]
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s,
						"start_date": "2017-03-01",
						"end_date": "2017-03-01",
						"country": ["UK"]
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}

func (p ProductIdDateIppIdCountryDimensionTestGroup) TestProductIdDateIppIdCountryDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	query := `{
				"user_id": %s,
				"query": {
					"filters": {
						"product_ids": %s
					},
					"dimensions": ["product_id", "date", "ipp_id", "country"],
					"metrics": ["revenue"],
					"order_by": ["revenue", "asc"],
					"currency_code": "CNY"
				}
			}`
	return C.GenerateTestCase(query, p.Data)
}
