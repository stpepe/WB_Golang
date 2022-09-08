package main

import (
    "log"
    // "fmt"
    // "context"
	// "net"
    "github.com/stpepe/nats-task"
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
    "github.com/stpepe/nats-task/pkg/repository"
    "github.com/stpepe/nats-task/pkg/handler"
    "github.com/stpepe/nats-task/pkg/service"
    nats "github.com/nats-io/nats.go"
    // stan "github.com/nats-io/stan.go"
)

func main() {
    if err := InitConfig(); err != nil {
        log.Fatalf("error initializing configs: %s", err.Error())
    } 
        
    db, err := repository.NewPostgresDB(repository.Config{
        Host: viper.GetString("db.host"),
        Port: viper.GetString("db.port"),
        Username: viper.GetString("db.username"),
        Password: viper.GetString("db.password"),
        DBName: viper.GetString("db.dbname"),
        SSLMode: viper.GetString("db.sslmode"),
    })
    if err != nil {
        log.Fatalf("error initializing database: %s", err.Error())
    } 

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

    channel := make(chan int)
    defer close(channel)

    repos := repository.NewRepository(db)
    services := service.NewService(repos)
    cache := service.NewCache()
    cache.RecoverCache(repos)

    publisher := service.CreatePubNATS()
    subscriber := service.CreateSubNATS(services)

    handlers := handler.NewHandler(services, publisher, cache, channel)

    publisher.NATS.Connect(
        viper.GetString("nats.cluster"),
        viper.GetString("nats.sender"),
        viper.GetString("nats.client_port"),
    )
    defer publisher.NATS.Close()

    subscriber.NATS.Connect(
        viper.GetString("nats.cluster"),
        viper.GetString("nats.receiver"),
        viper.GetString("nats.client_port"),
    )
    defer subscriber.NATS.Close()

    subscriber.Subscribe(cache, channel)

    srv := new(testapp.Server)
    if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
        log.Fatalf(err.Error())
    }
}

func InitConfig() error{
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}