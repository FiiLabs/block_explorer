package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TxController struct {
}
// @Summary 通过hash查询交易
// @Description 通过hash查询交易
// @Param tx_hash path string true "hash"
// @Tags transaction
// @Router /tx/hash/{tx_hash} [get]
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

// @Summary 查询区块的所有交易
// @Description 查询区块的所有交易
// @Param block_height path int true "height"
// @Tags transaction
// @Router /tx/block_txs/{block_height} [get]
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
// @Summary 查询最近10笔交易
// @Description 查询最近10笔交易，用在首页展示
// @Tags transaction
// @Router /tx/ltxs [get]
func (txctl *TxController) QueryLTxs(c *gin.Context) {

	res, e := txService.GetLTxs()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}

// @Summary 查询最近交易
// @Description 查询最近交易，最多可以查询50笔
// @Param num query		int	false	"number"
// @Tags transaction
// @Router /tx/txs [get]
func (txctl *TxController) QueryTxs(c *gin.Context) {
	q_num := c.DefaultQuery("num", "10")

	res, e := txService.GetTxs(q_num)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}