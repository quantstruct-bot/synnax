package writer_test

//func openClient(ctx context.Context, id distribution.NodeID, services map[distribution.NodeID]serviceContainer) writer.ClientStream {
//	client, err := services[id].transport.writerClient.Stream(ctx, "localhost:0")
//	Expect(err).ToNot(HaveOccurred())
//	return client
//}
//
//func openRequest(client writer.ClientStream, keys channel.Keys) (writer.Response, error) {
//	Expect(client.Send(writer.Request{Keys: keys})).To(Succeed())
//	Expect(client.CloseSend()).To(Succeed())
//	return client.Receive()
//}
//
//var _ = Describe("Server", func() {
//	var (
//		log      *zap.Logger
//		services map[distribcore.NodeID]serviceContainer
//		builder  *mock.CoreBuilder
//	)
//	BeforeEach(func() {
//		log = zap.NewNop()
//		builder, services = provisionNServices(1, log)
//		ch := channel.Channel{Name: "SG02", Rate: 1 * telem.Hz, DataType: telem.Float64T, NodeID: 1}
//		Expect(services[1].channel.Create(&ch)).To(Succeed())
//	})
//	BeforeEach(func() {
//		routines := gleak.Goroutines()
//		DeferCleanup(func() {
//			Eventually(gleak.Goroutines).WithTimeout(time.Second).ShouldNot(gleak.HaveLeaked(routines))
//		})
//	})
//	AfterEach(func() {
//		Expect(builder.Close()).To(Succeed())
//		Expect(builder.Cleanup()).To(Succeed())
//	})
//	DescribeTable("Open Request", func(keys channel.Keys, expectedResError, expectedTransportError error) {
//		res, err := openRequest(openClient(ctx, 1, services), keys)
//		Expect(res.Err).To(HaveOccurredAs(expectedResError))
//		Expect(err).To(HaveOccurredAs(expectedTransportError))
//	},
//		Entry("Open the writerClient properly when the keys exist", channel.Keys{channel.NewKey(1, 1)}, nil, freighter.EOF),
//		Entry("Return an error when no keys are provided", channel.Keys{}, nil, errors.New("[segment.w] - server expected Keys to be defined")),
//		Entry("Return an error when invalid keys are provided", channel.Keys{channel.NewKey(1, 2)}, nil, query.NotFound),
//	)
//	Describe("Write Request", func() {
//		It("Should immediately abort all operations when the context is cancelled", func() {
//			ctx, cancel := context.WithCancel(context.TODO())
//			client := openClient(ctx, 1, services)
//			Expect(client.Send(writer.Request{Keys: channel.Keys{channel.NewKey(1, 1)}})).To(Succeed())
//			var s core.Segment
//			s.Data = []byte{1, 2, 3}
//			s.Start = telem.TimeStamp(25)
//			Expect(client.Send(writer.Request{Segments: []core.Segment{s}})).To(Succeed())
//			cancel()
//			res, err := client.Receive()
//			Expect(res.Err).To(BeNil())
//			Expect(err).To(HaveOccurredAs(context.Canceled))
//		})
//		Describe("No Cancellation", func() {
//			var client writer.ClientStream
//			BeforeEach(func() {
//				client = openClient(ctx, 1, services)
//				Expect(client.Send(writer.Request{Keys: channel.Keys{channel.NewKey(1, 1)}})).To(Succeed())
//			})
//			It("Should execute a valid write request", func() {
//				var s core.Segment
//				s.ChannelKey = channel.NewKey(1, 1)
//				s.Data = []byte{1, 2, 3}
//				s.Start = telem.TimeStamp(25)
//				Expect(client.Send(writer.Request{Segments: []core.Segment{s}})).To(Succeed())
//				Expect(client.CloseSend()).To(Succeed())
//				res, err := client.Receive()
//				Expect(err).To(HaveOccurredAs(freighter.EOF))
//				Expect(res.Err).ToNot(HaveOccurred())
//			})
//		})
//
//	})
//})
