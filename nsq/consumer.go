package main

import "github.com/nsqio/go-nsq"
type handler struct {}
func consumer(){
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("events", "stats", config)
	if err != nil{
		panic(err)
	}
	h := &handler{}
	consumer.AddHandler(h)
	err = consumer.ConnectToNSQLookupd("localhost:4171")
	if err != nil{
		panic(err)
	}
	consumer.Stop()
}
func(h *handler) HandleMessage(message *nsq.Message) error{
	return nil
}