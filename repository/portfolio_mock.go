package repository

import (
	"errors"
)

type (
	portfolioRepositoryMock struct {
		portfolios Portfolios
	}
)

func NewPortfolioRepositoryMock() portfolioRepositoryMock {
	portfolios := Portfolios{
		Portfolio{ID: 1, UserId: 1, Symbol: "ETH", Available: 0.5},
		Portfolio{ID: 2, UserId: 1, Symbol: "BNB", Available: 0.05},
		Portfolio{ID: 3, UserId: 1, Symbol: "BTC", Available: 0.005},
	}
	return portfolioRepositoryMock{portfolios}
}

func (r portfolioRepositoryMock) GetAll(userid int) (resp Portfolios, err error) {
	resp = Portfolios{}
	for _, port := range r.portfolios {
		if port.UserId == userid {
			resp = append(resp, port)
		}
	}
	return
}

func (r portfolioRepositoryMock) GetBySymbol(userid int, symbol string) (resp *Portfolio, err error) {
	for _, port := range r.portfolios {
		if port.Symbol == symbol && port.UserId == userid {
			return &port, nil
		}
	}
	return nil, errors.New("asset not found")
}
