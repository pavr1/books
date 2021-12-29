package inventory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InventoryTestSuite struct {
	suite.Suite
	Inventory Inventory
}

func TestInventoryTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryTestSuite))
}

func (i *InventoryTestSuite) SetupSuite() {
	i.Inventory = NewInventory()
}

func (i *InventoryTestSuite) AfterTest(suiteName, testName string) {
	i.Inventory = NewInventory()
}

func (i InventoryTestSuite) Test_AddBook_Success() {
	//arrange

	//act
	err := i.Inventory.AddBook("Lord of the rings", 10)

	//assert
	i.Assert().Nil(err)
}

func (i InventoryTestSuite) Test_ListAllBooks_Success() {
	//arrange

	//act
	err := i.Inventory.AddBook("Lord of the rings", 10)

	//assert
	i.Assert().Nil(err)

	listedBooks := i.Inventory.ListAllBooks()
	i.Assert().Equal(1, len(listedBooks))
	i.Assert().Equal(10, listedBooks[0].Stock)
}

func (i InventoryTestSuite) Test_GetBookByTitle_Success() {
	//arrange

	//act
	err := i.Inventory.AddBook("Lord of the rings", 10)

	//assert
	i.Assert().Nil(err)

	listedBook := i.Inventory.GetBookByTitle("Lord of the rings")
	i.Assert().NotNil(listedBook)
}

func (i InventoryTestSuite) Test_GetBookByTitle_Failed() {
	//arrange

	//act
	err := i.Inventory.AddBook("Lord of the rings", 10)

	//assert
	i.Assert().Nil(err)

	listedBook := i.Inventory.GetBookByTitle("Lord of the rings 2")
	i.Assert().Nil(listedBook)
}

func (i InventoryTestSuite) Test_EditBook_EditTile_Success() {
	//arrange

	//act
	err := i.Inventory.AddBook("Lord of the rings", 10)
	i.Assert().Nil(err)

	book := i.Inventory.GetBookByTitle("Lord of the rings")
	i.Assert().NotNil(book)

	err = i.Inventory.EditBook("Lord of the rings", "Lord of the rings 2", 0)
	i.Assert().Nil(err)

	book = i.Inventory.GetBookByTitle("Lord of the rings")
	i.Assert().Nil(book)

	book = i.Inventory.GetBookByTitle("Lord of the rings 2")
	i.Assert().NotNil(book)
}
