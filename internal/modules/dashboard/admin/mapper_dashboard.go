package admin

func ToRevenueSummaryResponse(summary *RevenueSummary) RevenueSummaryResponse {
	if summary == nil {
		return RevenueSummaryResponse{}
	}

	return RevenueSummaryResponse{
		TotalRevenue:    summary.TotalRevenue,
		RevenuePerWeek:  toRevenuePerWeekResponses(summary.RevenuePerWeek),
		RevenuePerMonth: toRevenuePerMonthResponses(summary.RevenuePerMonth),
		RevenuePerYear:  toRevenuePerYearResponses(summary.RevenuePerYear),
	}
}

func toRevenuePerWeekResponses(weeks []RevenuePerWeek) []RevenuePerWeek {
	if weeks == nil {
		return []RevenuePerWeek{}
	}

	responses := make([]RevenuePerWeek, len(weeks))
	for i, w := range weeks {
		responses[i] = toRevenuePerWeekResponse(w)
	}
	return responses
}

func toRevenuePerWeekResponse(week RevenuePerWeek) RevenuePerWeek {
	return RevenuePerWeek{
		Year:    week.Year,
		Week:    week.Week,
		Revenue: week.Revenue,
	}
}

func toRevenuePerMonthResponses(months []RevenuePerMonth) []RevenuePerMonth {
	if months == nil {
		return []RevenuePerMonth{}
	}

	responses := make([]RevenuePerMonth, len(months))
	for i, m := range months {
		responses[i] = toRevenuePerMonthResponse(m)
	}
	return responses
}

func toRevenuePerMonthResponse(month RevenuePerMonth) RevenuePerMonth {
	return RevenuePerMonth{
		Year:      month.Year,
		Month:     month.Month,
		MonthName: month.MonthName,
		Revenue:   month.Revenue,
	}
}
func toRevenuePerYearResponses(years []RevenuePerYear) []RevenuePerYear {
    if years == nil {
        return []RevenuePerYear{}
    }

    responses := make([]RevenuePerYear, len(years))
    for i, y := range years {
        responses[i] = toRevenuePerYearResponse(y)
    }
    return responses
}

func toRevenuePerYearResponse(year RevenuePerYear) RevenuePerYear {
    return RevenuePerYear{
        Year:    year.Year,
        Revenue: year.Revenue,
    }
}