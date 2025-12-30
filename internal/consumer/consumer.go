package consumer

func AllConsumers() []func() {
	return []func(){
		TestQueue,
		UserQueue,
		// 继续添加其他队列监听函数
	}
}
