package attendance

// ============================================
// MAPPER
// ============================================

func ToAttendanceDetailResponse(a AttendanceDetail) AttendanceDetailResponse {
	return AttendanceDetailResponse{
		ID:             a.ID,
		OrderID:        a.OrderID,
		EventID:        a.EventID,
		UserID:         a.UserID,
		TicketTypeID:   a.TicketTypeID,
		CheckInTime:    a.CheckInTime,
		CheckInMethod:  string(a.CheckInMethod),
		Status:         string(a.Status),
		CheckedBy:      a.CheckedBy,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
		UserName:       a.UserName,
		UserEmail:      a.UserEmail,
		EventTitle:     a.EventTitle,
		TicketTypeName: a.TicketTypeName,
		CheckedByName:  a.CheckedByName,
	}
}

func ToAttendanceByEventResponses(rows []AttendanceByEvent) []AttendanceByEventResponse {
	result := make([]AttendanceByEventResponse, 0, len(rows))
	for _, a := range rows {
		result = append(result, AttendanceByEventResponse{
			ID:            a.ID,
			OrderID:       a.OrderID,
			UserID:        a.UserID,
			UserName:      a.UserName,
			UserEmail:     a.UserEmail,
			CheckInTime:   a.CheckInTime,
			CheckInMethod: string(a.CheckInMethod),
			Status:        string(a.Status),
			CheckedBy:     a.CheckedBy,
			CheckedByName: a.CheckedByName,
		})
	}
	return result
}

func ToAttendanceByUserResponses(rows []AttendanceByUser) []AttendanceByUserResponse {
	result := make([]AttendanceByUserResponse, 0, len(rows))
	for _, a := range rows {
		result = append(result, AttendanceByUserResponse{
			ID:            a.ID,
			OrderID:       a.OrderID,
			EventID:       a.EventID,
			UserID:        a.UserID,
			TicketTypeID:  a.TicketTypeID,
			CheckInTime:   a.CheckInTime,
			CheckInMethod: string(a.CheckInMethod),
			Status:        string(a.Status),
			CheckedBy:     a.CheckedBy,
			CreatedAt:     a.CreatedAt,
			EventTitle:    a.EventTitle,
			StartTime:     a.StartTime,
			EndTime:       a.EndTime,
		})
	}
	return result
}

func ToEventAttendanceRateResponses(rows []EventAttendanceRate) []EventAttendanceRateResponse {
	result := make([]EventAttendanceRateResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, EventAttendanceRateResponse{
			EventID:        r.EventID,
			EventTitle:     r.EventTitle,
			TotalOrders:    r.TotalOrders,
			Attended:       r.Attended,
			AttendanceRate: r.AttendanceRate,
		})
	}
	return result
}

func ToCheckInTimelineHourResponses(rows []CheckInTimelineHour) []CheckInTimelineHourResponse {
	result := make([]CheckInTimelineHourResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, CheckInTimelineHourResponse{
			Hour:  r.Hour,
			Count: r.Count,
		})
	}
	return result
}

func ToCheckInTimelineDateResponses(rows []CheckInTimelineDate) []CheckInTimelineDateResponse {
	result := make([]CheckInTimelineDateResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, CheckInTimelineDateResponse{
			Date:  r.Date,
			Count: r.Count,
		})
	}
	return result
}
