package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"examples/gin-sample/adapter/gateway"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	db := Connect()

	var u gateway.User
	db.First(&u)

	assert.Equal(u.Username, "kuzu", "test")
	defer CloseConn()
}

//
// func TestCloseConn(t *testing.T) {
// 	assert := assert.New(t)

// }
