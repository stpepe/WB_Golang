package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/stpepe/nats-task"
	"strconv"
)

func (h *Handler) rec_cache (c *gin.Context){
	n_data := h.cache.GetAll()
	c.JSON(http.StatusOK, map[string]interface{}{
		"response": n_data, 
	})
}

func (h *Handler) show (c *gin.Context){
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"response": "Invalid id", 
		})
	}
	n_data, stat := h.cache.Get(input)
	if stat{
		c.HTML(http.StatusOK, "order.html", gin.H{
			"OrderUid": n_data.OrderUid,
			"TrackNumber": n_data.TrackNumber,
			"Entry": n_data.Entry  ,
			"Delivery": n_data.Delivery,
			"Payment": n_data.Payment,
			"Items": n_data.Items,
			"Locale": n_data.Locale,
			"InternalSignature": n_data.InternalSignature,
			"CustomerID": n_data.CustomerID,
			"DeliveryService": n_data.DeliveryService,
			"ShardKey": n_data.Shardkey,
			"SmID": n_data.SmID,
			"DateCreated": n_data.DateCreated,
			"OofShard": n_data.OofShard,
		})
		return

		// c.JSON(http.StatusOK, map[string]interface{}{
		// 	"response": n_data, 
		// })
		// return
	}
	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"response": "error", 
	})
	return
}

func (h *Handler) send (c *gin.Context){
	var input testapp.Order

	if err := c.BindJSON(&input); err != nil {
		return
	}

	err := h.publisher.PublishOrder(input)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"err": err, 
		})
		return
	}

	id := <-h.channel
	if id != 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "complete", 
			"id": id,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"order": "was not created",
		"error": "Invalid data", 
	})
}