package __dimension

import (
	E "git.appannie.org/appannie/performance/erik"
	C "git.appannie.org/appannie/performance/thrift/cases"
	"time"
	"fmt"
)

type CountryDimensionTestGroup struct {
	Data map[string]interface{}
}

func (p CountryDimensionTestGroup) TestCountryDimensionOneDayOneCountry() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "date", "start": "2017-03-01", "end": "2017-03-01"},
			{"name": "country", "values": []string{"UK"}},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionOneDayAllCountry() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "date", "start": "2017-03-01", "end": "2017-03-01"},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionAllDayOneCountry() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "country", "values": []string{"UK"}},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionOneDayOneCountryTwoMetrics() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "date", "start": "2017-03-01", "end": "2017-03-01"},
			{"name": "country", "values": []string{"UK"}},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
			{"name": "units"},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionOneDayOneCountryOrderByRevenue() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "date", "start": "2017-03-01", "end": "2017-03-01"},
			{"name": "country", "values": []string{"UK"}},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
		"sorts": []map[string]interface{}{
			{"name": "revenue", "desc": false},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionOneDayOneCountryOrderByRevenueWithCurrency() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
			{"name": "date", "start": "2017-03-01", "end": "2017-03-01"},
			{"name": "country", "values": []string{"UK"}},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
		"sorts": []map[string]interface{}{
			{"name": "revenue", "desc": false},
		},
		"meta": map[string]interface{} {
			"code": "CNY",
			"currencies": map[string]float64 {
				"2017-03-01": 6.12345,
			},
		},
	}
	return C.GenerateTestCase(query, p.Data)
}

func (p CountryDimensionTestGroup) TestCountryDimensionAllDayAllCountryOrderByRevenueWithCurrency() E.ITestCase {
	userProductIds := p.Data["user_product_ids"].([]int64)
	currencies := map[string]float64{}
	now := time.Now()
	for i := 0; i < 365; i++ {
		year, month, day := now.AddDate(0, 0, -1).Date()
		dateStr := fmt.Sprintf("%d-%02d-%02d", year, month, day)
		currencies[dateStr] = 6.12345
	}
	query := map[string]interface{}{
		"dimensions": []map[string]string{
			{"name": "country"},
		},
		"filters": []map[string]interface{}{
			{"name": "user_product_id", "values": userProductIds},
		},
		"metrics": []map[string]string{
			{"name": "revenue"},
		},
		"sorts": []map[string]interface{}{
			{"name": "revenue", "desc": false},
		},
		"meta": map[string]interface{} {
			"code": "CNY",
			"currencies": currencies,
		},
	}
	return C.GenerateTestCase(query, p.Data)
}
