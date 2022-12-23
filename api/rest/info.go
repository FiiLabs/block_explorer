package rest

import (
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoController struct {
}

//	@Summary		查询区块总数
//	@Description	查询区块总数
// @Tags info
//	@Success		200	{object}	int
//	@Router			/info/blocks [get]
func (infoctl *InfoController) QueryLatestHeight(c *gin.Context) {
	//client := pool.GetClient()
	//defer func() {
	//	client.Release()
	//}()
	//status, e := client.Status(context.Background())
	block, e := new(models.Block).GetMaxBlockHeight()
	if e != nil {
		err :=errors.Wrap(e)
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	//res:=status.SyncInfo.LatestBlockHeight
	res:=block.Height
	c.JSON(http.StatusOK, response.Success(res))
}
//	@Summary		查询交易总数
//	@Description	查询交易总数
// @Tags info
//	@Success		200	{object}	int
//	@Router			/info/txs [get]
func (infoctl *InfoController) QueryLatestTx(c *gin.Context) {

	res, err := txService.GetTxsCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
//	@Summary		查询NFT类别总数
//	@Description	查询NFT类别总数
// @Tags info
//	@Success		200	{object}	int
//	@Router			/info/nftclses [get]
func (infoctl *InfoController) QueryNFTClsCount(c *gin.Context) {

	res, err := nftClsService.GetNFTClsCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
//	@Summary		查询NFT总数
//	@Description	查询NFT总数
// @Tags info
//	@Success		200	{object}	int
//	@Router			/info/nfts [get]
func (infoctl *InfoController) QueryNFTCount(c *gin.Context) {

	res, err := nftService.GetNFTCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
