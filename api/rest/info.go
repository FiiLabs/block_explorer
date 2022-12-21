package rest

import (
	"context"
	"github.com/FiiLabs/block_explorer/api/response"
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/libs/pool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoController struct {
}

func (infoctl *InfoController) QueryLatestHeight(c *gin.Context) {
	client := pool.GetClient()
	defer func() {
		client.Release()
	}()
	status, e := client.Status(context.Background())
	if e != nil {
		err :=errors.Wrap(e)
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	res:=status.SyncInfo.LatestBlockHeight
	c.JSON(http.StatusOK, response.Success(res))
}

func (infoctl *InfoController) QueryLatestTx(c *gin.Context) {

	res, err := txService.GetTxsCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (infoctl *InfoController) QueryNFTClsCount(c *gin.Context) {

	res, err := nftClsService.GetNFTClsCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (infoctl *InfoController) QueryNFTCount(c *gin.Context) {

	res, err := nftService.GetNFTCount()
	if err != nil {
		c.JSON(response.HttpCode(err), response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
