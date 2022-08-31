package application_test

import (
	"testing"

	"github.com/rafaelbertelli/hexagonal-GO/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "Price must be greather than 0 to enable the product", err.Error())
}

