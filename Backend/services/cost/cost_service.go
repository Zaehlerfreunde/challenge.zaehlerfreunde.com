package cost

type costService struct {
}

type CostService interface {
}

func NewCostServiceService() CostService {
	return costService{}
}
