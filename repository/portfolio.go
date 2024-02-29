package repository

type (
	// Assert struct {
	// 	Name      string  `db:"name"`
	// 	Symbol    string  `db:"symbol"`
	// 	Available float64 `db:"available"`
	// }

	Portfolio struct {
		ID        int     `json:"id" db:"id"`
		UserId    int     `json:"userid" db:"userid"`
		Symbol    string  `json:"symbol" db:"symbol"`
		Available float64 `json:"available" db:"available"`
		Reserved  float64 `json:"reserved" db:"reserved"`
		// Last      float64 `json:"last" db:"last"`
		// Thb       float64 `json:"thb" db:"thb"`
	}

	//
	Portfolios []Portfolio

	// PortfolioRepository is Port
	PortfolioRepository interface {
		GetAll(int) (Portfolios, error)
		GetBySymbol(int, string) (*Portfolio, error)
	}
)
