package test

//func Test100(t *testing.T) {
//	order := getOrderFromMessage(1)
//	//序列化成二进制数据
//	orderBytes, err := proto.Marshal(&order)
//	if err != nil {
//		println(err)
//	}
//	//假设获取
//	//反序列化二进制数据
//	thisOrder := &message.Order{}
//	proto.Unmarshal(orderBytes, thisOrder)
//	println(*thisOrder.OrderStatus)
//	status := "1"
//	thisOrder.OrderStatus = &status
//	println(*thisOrder.OrderStatus)
//	order2 := &message.Order{
//		UserId:       thisOrder.UserId,
//		OrderType:    thisOrder.OrderType,
//		OrderStatus:  proto.String("1"),
//		OrderBalance: thisOrder.OrderBalance,
//	}
//	//序列化成二进制数据
//	order2Bytes, err2 := proto.Marshal(order2)
//	if err2 != nil {
//		println(err2)
//	}
//	thisOrder2 := &message.Order{}
//	proto.Unmarshal(order2Bytes, thisOrder2)
//	println(*thisOrder2.OrderStatus)
//
//}
//func getOrderFromMessage(orderMessageId int32) message.Order {
//	order := &message.Order{
//		OrderMessageId:      proto.Int32(orderMessageId),
//		UserId:       proto.String("1"),
//		OrderType:    proto.String("pay"),
//		OrderStatus:  proto.String("0"),
//		OrderBalance: proto.Float64(100.0),
//
//	}
//	return *order
//}
//func Test101(t *testing.T) {
//	config :=sarama.NewConfig()
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	//随机的分区类型
//	config.Producer.Partitioner = sarama.NewRandomPartitioner
//	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
//	config.Producer.Return.Successes = true
//	config.Producer.Return.Errors = true
//	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
//	config.Version = sarama.V0_11_0_0
//
//	//使用配置,新建一个异步生产者
//	producer, e := sarama.NewAsyncProducer([]string{"IP:9092","IP:9092","IP:9092"}, config)
//	if e != nil {
//		panic(e)
//	}
//	defer producer.AsyncClose()
//
//	msg := &sarama.ProducerMessage{
//		Topic: "test_topic",
//		Key: sarama.StringEncoder("test"),
//	}
//	order := &message.Order{
//		OrderMessageId:      proto.Int32(1),
//		UserId:       proto.String("1"),
//		OrderType:    proto.String("pay"),
//		OrderStatus:  proto.String("0"),
//		OrderBalance: proto.Float64(100.0),
//	}
//	orderBytes, err := proto.Marshal(order)
//	if err != nil {
//		println(err)
//	}
//	for {
//		//设置发送的真正内容
//		fmt.Scanln(orderBytes)
//		//将字符串转化为字节数组
//		msg.Value = sarama.ByteEncoder(orderBytes)
//		fmt.Println(orderBytes)
//		//使用通道发送
//		producer.Input() <- msg
//		//循环判断哪个通道发送过来数据.
//		select {
//		case suc := <-producer.Successes():
//			fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
//		case fail := <-producer.Errors():
//			fmt.Println("err: ", fail.Err)
//		}
//	}
//}
