package api

import "time"

type Department struct {
	Date time.Time
	Code int `json:"code"`
	Data struct {
		Step int64 `json:"step"`
		Sets []struct {
			AllocationTrends struct {
				Shared01 struct {
					Trends struct {
						Costs struct {
							TotalCost struct {
								RelativeChange struct {
									IsInfinite bool    `json:"isInfinite"`
									IsNaN      bool    `json:"isNaN"`
									Value      float64 `json:"value"`
								} `json:"relativeChange"`
							} `json:"totalCost"`
						} `json:"costs"`
					} `json:"trends"`
				} `json:"Shared-01"`
				Unallocated struct {
					Trends struct {
						Costs struct {
							TotalCost struct {
								RelativeChange struct {
									IsInfinite bool    `json:"isInfinite"`
									IsNaN      bool    `json:"isNaN"`
									Value      float64 `json:"value"`
								} `json:"relativeChange"`
							} `json:"totalCost"`
						} `json:"costs"`
					} `json:"trends"`
				} `json:"__unallocated__"`
				Dashuju01 struct {
					Trends struct {
						Costs struct {
							TotalCost struct {
								RelativeChange struct {
									IsInfinite bool    `json:"isInfinite"`
									IsNaN      bool    `json:"isNaN"`
									Value      float64 `json:"value"`
								} `json:"relativeChange"`
							} `json:"totalCost"`
						} `json:"costs"`
					} `json:"trends"`
				} `json:"dashuju-01"`
				Jishuzhongxin01 struct {
					Trends struct {
						Costs struct {
							TotalCost struct {
								RelativeChange struct {
									IsInfinite bool    `json:"isInfinite"`
									IsNaN      bool    `json:"isNaN"`
									Value      float64 `json:"value"`
								} `json:"relativeChange"`
							} `json:"totalCost"`
						} `json:"costs"`
					} `json:"trends"`
				} `json:"jishuzhongxin-01"`
			} `json:"allocationTrends"`
			Window struct {
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"window"`
		} `json:"sets"`
		Window struct {
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
		} `json:"window"`
	} `json:"data"`
}
