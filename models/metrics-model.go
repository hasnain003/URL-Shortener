package models

type MetricsResponse struct {
	Score  int
	Member string
}

func NewMetricResponse(score int, member string) MetricsResponse {
	return MetricsResponse{
		Score:  score,
		Member: member,
	}
}

type TopMetrics struct {
	Responses []MetricsResponse
}
