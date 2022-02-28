package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/kinmaBackend/util"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T, account Account) Product{
	
	tagList := util.RandomTagList()
	
	productArg := CreateProductParams{
		AccountID	 : account.ID,
		Title      : util.RandomString(5),
		Content    : util.RandomString(20) + "_Content",
		ProductTag : tagList,
	}

	product, err := testQueries.CreateProduct(context.Background(),productArg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, product.AccountID, productArg.AccountID)
	require.Equal(t, product.Title, productArg.Title)
	require.Equal(t, product.ProductTag, productArg.ProductTag)
	var array []string
	require.IsType(t, array, product.ProductTag)

	require.NotZero(t, product.ID)
	require.Equal(t, product.AccountID, account.ID)

	return product
}

func TestCreateProduct(t *testing.T){
	account := createRandomAccount(t)
	createRandomProduct(t, account)
}

func TestDeleteProduct(t *testing.T){
	account1 := createRandomAccount(t)
	product1 := createRandomProduct(t, account1)
	err := testQueries.DeleteAccountProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestGetProduct(t *testing.T) {
	account1 := createRandomAccount(t)
	product1 := createRandomProduct(t, account1)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestListProduct(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		
		tagList := util.RandomTagList()

		productArg := CreateProductParams{
			AccountID	 : account.ID,
			Title      : util.RandomString(5),
			Content    : util.RandomString(20) + "_Content",
			ProductTag : tagList,
		}
		testQueries.CreateProduct(context.Background(),productArg)
	}

	arg := ListMyProductParams{
		AccountID : account.ID,
		Limit		  : 5,
		Offset		: 5,
	}

	products, err := testQueries.ListMyProduct(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestUpdateProduct(t *testing.T) {
	account := createRandomAccount(t)
	product1 := createRandomProduct(t, account)

	tagList := util.RandomTagList()

	arg := UpdateProductDetailParams{
		ID					:		product1.ID,
		Title 			:	 	util.RandomString(5),
		Content 		:		util.RandomString(20) + "_Content",
		ProductTag	: 	tagList,
	}

	updateProduct1, err := testQueries.UpdateProductDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateProduct1)
	require.Equal(t, updateProduct1.ID, arg.ID)
	require.Equal(t, updateProduct1.Title, arg.Title)
	require.Equal(t, updateProduct1.Content, arg.Content)
	require.Equal(t, updateProduct1.ProductTag, arg.ProductTag)
	require.NotEmpty(t, updateProduct1.LastUpdate)
}