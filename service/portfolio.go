package service

type (
	PortfolioResponse struct {
		Symbol    string  `json:"symbol"`
		Available float64 `json:"available" db:"available"`
		Reserved  float64 `json:"reserved" db:"reserved"`
		Last      float64 `json:"last" db:"last"`
		Thb       float64 `json:"thb" db:"thb"`
	}

	PortfoliosResponse []PortfolioResponse

	PortfolioService interface {
		GetPortfolios(int) (PortfoliosResponse, error)
		GetPortfolio(int, string) (*PortfolioResponse, error)
	}
)
