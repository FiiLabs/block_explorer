package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NFTController struct {
}
// @Summary 查询NFT
// @Description 查询NFT，可以分页查询
// @Param page query		int	false	"page"
// @Param size query		int	false	"size"
// @Tags NFT
// @Success 200 {object} vo.NFTsResp
// @Router /nft/nfts [get]
func (nctl *NFTController) QueryNFT(c *gin.Context) {
	q_page := c.DefaultQuery("page", "0")
	q_size := c.DefaultQuery("size", "10")

	res, e := nftService.QueryNFT(q_page, q_size)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
