package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TxController struct {
}

func (txctl *TxController) QueryByHash(c *gin.Context) {
	txhash := c.Param("tx_hash")
	if txhash == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter tx_hash is required"))
		return
	}

	res, e := txService.QueryTransactionByHash(txhash)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
func (txctl *TxController) QueryTxsByBlockHeight(c *gin.Context) {
	blkHeight := c.Param("block_height")
	if blkHeight == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter block_height is required"))
		return
	}

	res, e := txService.QueryTxsByHeight(blkHeight)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}

func (txctl *TxController) QueryLTxs(c *gin.Context) {

	res, e := txService.GetLTxs()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
func (txctl *TxController) QueryTxs(c *gin.Context) {
	q_page := c.DefaultQuery("page", "0")
	q_size := c.DefaultQuery("size", "10")

	res, e := txService.GetTxs(q_page, q_size)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}