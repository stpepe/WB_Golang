package service

import (
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"log"
	"github.com/stpepe/nats-task"
	"time"
)

const (
	NATS_SUBJECT_ORDER = "order"
)

type NATS struct {
	sc stan.Conn
}

type SubNATS struct{
	NATS
	order_service *Service
}

type PubNATS struct{
	NATS
}

func CreateSubNATS(service *Service) *SubNATS {
	n := SubNATS{
		NATS: NATS{},
		order_service: service,
	}
	return &n
}

func CreatePubNATS() *PubNATS {
	n := PubNATS{
		NATS: NATS{},
	}
	return &n
}

func (v *PubNATS) PublishOrder(order testapp.Order) error {
	jsonObj, err := json.Marshal(order)
	if err != nil {
		return err
	}
	msg:=[]byte(string(jsonObj))
	err = v.NATS.sc.Publish(NATS_SUBJECT_ORDER, msg)

	return err
}

func (v *SubNATS) Subscribe(cache *Cache, channel chan int) error {
	_, err := v.NATS.sc.Subscribe(NATS_SUBJECT_ORDER, func(m *stan.Msg) {
		var order testapp.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println(err)
			return
		}

		err = m.Ack()
		if err != nil {
			log.Println(err)
			return
		}

		id, err := v.order_service.Send.CreateOrder(order)
		channel <- id
		if err != nil {
			log.Println(err)
			return
		}

		cache.Set(id, order)

	}, stan.SetManualAckMode(), stan.AckWait(time.Second*50))

	return err
}

func (n *NATS) Connect(ClusterID string, ClientID string, NatsURL string) error {
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(NatsURL))
	if err != nil {
		return err
	}
	n.sc = sc
	return nil
}

func (nats *NATS) Close() {
	if nats.sc != nil {
		nats.sc.Close()
	}
}