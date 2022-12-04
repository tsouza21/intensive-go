package usecase

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tsouza21/intensive-go/internal/order/entity"
	"github.com/tsouza21/intensive-go/internal/order/infra/database"

	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.NoError(err)

	suite.Db = db
	suite.OrderRepository = database.NewOrderRepository(db)

}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestGivenAnOrderInput_WhenExecute_ThenCalculateFinalPriceAndSave() {
	order, err := entity.NewOrder("1", 10, 2)
	suite.NoError(err)
	order.CalculateFinalPrice()

	calculateFinalPriceInput := OrderInputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)

	calculateFinalPriceOutput, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, calculateFinalPriceOutput.ID)
	suite.Equal(order.Price, calculateFinalPriceOutput.Price)
	suite.Equal(order.Tax, calculateFinalPriceOutput.Tax)
	suite.Equal(order.FinalPrice, calculateFinalPriceOutput.FinalPrice)
}