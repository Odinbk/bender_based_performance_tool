package main

import (
	E "git.appannie.org/appannie/performance/erik"
	OD "git.appannie.org/appannie/performance/thrift/cases/test_plan/1_dimension"
	TWD "git.appannie.org/appannie/performance/thrift/cases/test_plan/2_dimensions"
	TRD "git.appannie.org/appannie/performance/thrift/cases/test_plan/3_dimensions"
	FD "git.appannie.org/appannie/performance/thrift/cases/test_plan/4_dimensions"
	"time"
	"flag"
	"fmt"
)

func generateTestGroupMapping(c *Config) map[string]interface{} {
	data := map[string]interface{}{
		"user_product_ids": c.UserProductIds,
		"interval":         c.Interval,
		"repeat":           c.Repeat,
		"address":          c.ServerAddress,
		"auth":             c.Auth,
		"timeout":          time.Duration(c.Timeout) * time.Second,
	}

	var testPlanMap = map[string]interface{}{
		"product_id":                     OD.ProductDimensionTestGroup{Data: data},
		"country":                        OD.CountryDimensionTestGroup{Data: data},
		"date":                           OD.DateDimensionTestGroup{Data: data},
		"asset":                          OD.AssetDimensionTestGroup{Data: data},
		"ipp_id":                         OD.IppIdDimensionTestGroup{Data: data},
		"ipp_transaction_type":           OD.IppTransactionTypeDimensionTestGroup{Data: data},
		"asset_country":                  TWD.AssetCountryDimensionTestGroup{Data: data},
		"asset_date":                     TWD.AssetDateDimensionTestGroup{Data: data},
		"asset_product_id":               TWD.AssetProductIdDimensionTestGroup{Data: data},
		"country_ipp_transaction_type":   TWD.CountryIppTransactionTypeDimensionTestGroup{Data: data},
		"date_ipp_transaction_type":      TWD.DateIppTransactionTypeDimensionTestGroup{Data: data},
		"ipp_id_country":                 TWD.IppIdCountryDimensionTestGroup{Data: data},
		"ipp_id_ipp_transaction_type":    TWD.IppIdIppTransactionTypeDimensionTestGroup{Data: data},
		"product_id_country":             TWD.ProductIdCountryDimensionTestGroup{Data: data},
		"product_id_date":                TWD.ProductIdDateDimensionTestGroup{Data: data},
		"date_country":                   TWD.DateCountryDimensionTestGroup{Data: data},
		"asset_date_country":             TRD.AssetDateCountryDimensionTestGroup{Data: data},
		"date_country_ipp_id":            TRD.DateCountryIppIdDimensionTestGroup{Data: data},
		"product_id_date_country":        TRD.ProductIdDateCountryDimensionTestGroup{Data: data},
		"product_id_date_ipp_id":         TRD.ProductIdDateIppIdDimensionTestGroup{Data: data},
		"product_id_ipp_id_date_country": FD.ProductIdDateIppIdCountryDimensionTestGroup{Data: data},
	}

	return testPlanMap
}

type Config struct {
	UserProductIds []int64
	Interval       float64
	Repeat         int
	ServerAddress  string
	Auth           string
	Timeout        int
}

func main() {
	var configFile = flag.String("config", "test.conf", "config file")
	var testPlan = flag.String("tp", "", `
		One dimension:
			product_id
			country
			date
			asset
			ipp_id
			ipp_transaction_type
		Two dimensions:
			asset_country
			asset_date
			asset_product_id
			country_ipp_transaction_type
			date_ipp_transaction_type
			ipp_id_country
			ipp_id_ipp_transaction_type
			product_id_country
			product_id_date
			date_country
		Three dimensions:
			asset_date_country
			date_country_ipp_id
			product_id_date_country
			product_id_date_ipp_id
		Four dimensions:
			product_id_ipp_id_date_country
	`)
	flag.Parse()

	c := new(Config)
	err := E.LoadConfig(*configFile, c)
	if err != nil {
		panic(err)
	}
	testPlanMap := generateTestGroupMapping(c)

	if *testPlan == "" {
		tps := []E.TestPlan{}
		for _, plan := range testPlanMap {
			testCases := E.LoadCases(plan)
			tp := E.TestPlan{
				TestCases: testCases,
			}
			tps = append(tps, tp)
		}
		tr := E.NewTestRun(tps)
		tr.Run()
	} else {
		if plan, ok := testPlanMap[*testPlan]; ok {
			testCases := E.LoadCases(plan)
			tp := E.TestPlan{
				TestCases: testCases,
			}
			tr := E.NewTestRun([]E.TestPlan{tp})
			tr.Run()
		} else {
			fmt.Printf("unsupport dimension or dimension combination: %s\n", testPlan)
		}
	}
}
