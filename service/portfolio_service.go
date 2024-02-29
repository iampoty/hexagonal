package service

import "hexagonal/repository"

type (
	portfolioService struct {
		portfolioRepo repository.PortfolioRepository
	}
)

func NewPortfolioService(portfolioRepo repository.PortfolioRepository) portfolioService {

	return portfolioService{portfolioRepo: portfolioRepo}
}

func (s portfolioService) GetPortfolios(userid int) (resp PortfoliosResponse, err error) {
	var ports repository.Portfolios

	if ports, err = s.portfolioRepo.GetAll(userid); err != nil {
		return
	}

	for _, port := range ports {
		portResp := PortfolioResponse{
			Symbol:    port.Symbol,
			Available: port.Available,
			Reserved:  port.Reserved,
			Last:      0.0,
			Thb:       0.0,
		}
		resp = append(resp, portResp)
	}

	return
}

func (s portfolioService) GetPortfolio(userid int, symbol string) (resp *PortfolioResponse, err error) {

	var port *repository.Portfolio
	if port, err = s.portfolioRepo.GetBySymbol(userid, symbol); err != nil {
		return
	}

	resp = &PortfolioResponse{
		Symbol:    port.Symbol,
		Available: port.Available,
		Reserved:  port.Reserved,
		Last:      0.0,
		Thb:       0.0,
	}

	return
}
