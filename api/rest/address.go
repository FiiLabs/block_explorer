package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressController struct {
}
//	@Summary		获取某地址的交易列表
//	@Description	获取某地址的交易列表
// @Tags address
//	@Param			address	path		string	true	"address"
//	@Param			page	query		int	false	"page"
//	@Param			size	query		int	false	"size"
//	@Router			/address/txs/{address} [get]
func (actl *AddressController) QueryTxs(c *gin.Context) {
	address := c.Param("address")
	q_page := c.DefaultQuery("page", "0")
	q_size := c.DefaultQuery("size", "10")

	res, e := txService.GetTxsByAddress(address,q_page, q_size)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
