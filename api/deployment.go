package api

import "time"

type Deployment struct {
	Code int `json:"code"`
	Date string
	Data struct {
		Chart []struct {
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
			Items []struct {
				Name       string  `json:"name"`
				Cost       float64 `json:"cost"`
				Efficiency float64 `json:"efficiency"`
			} `json:"items"`
		} `json:"chart"`
		Totals struct {
			Name                  string  `json:"name"`
			CPUCost               float64 `json:"cpuCost"`
			GpuCost               int     `json:"gpuCost"`
			RAMCost               float64 `json:"ramCost"`
			PvCost                float64 `json:"pvCost"`
			NetworkCost           int     `json:"networkCost"`
			LoadBalancerCost      float64 `json:"loadBalancerCost"`
			SharedCost            float64 `json:"sharedCost"`
			ExternalCost          int     `json:"externalCost"`
			AverageCPUUtilization float64 `json:"averageCpuUtilization"`
			AverageRAMUtilization float64 `json:"averageRamUtilization"`
			Efficiency            float64 `json:"efficiency"`
			TotalCost             float64 `json:"totalCost"`
		} `json:"totals"`
		Items struct {
			Page    int `json:"page"`
			PerPage int `json:"perPage"`
			Items   []struct {
				Name                  string  `json:"name"`
				CPUCost               float64 `json:"cpuCost"`
				GpuCost               int     `json:"gpuCost"`
				RAMCost               float64 `json:"ramCost"`
				PvCost                float64 `json:"pvCost"`
				NetworkCost           int     `json:"networkCost"`
				LoadBalancerCost      float64 `json:"loadBalancerCost"`
				SharedCost            float64 `json:"sharedCost"`
				ExternalCost          int     `json:"externalCost"`
				AverageCPUUtilization float64 `json:"averageCpuUtilization"`
				AverageRAMUtilization float64 `json:"averageRamUtilization"`
				Efficiency            float64 `json:"efficiency"`
				TotalCost             float64 `json:"totalCost"`
			} `json:"items"`
		} `json:"items"`
	} `json:"data"`
}
