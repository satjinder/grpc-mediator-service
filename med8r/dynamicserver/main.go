// package main

// import (
// 	"context"
// 	"log"
// 	"strings"

// 	"google.golang.org/grpc"
// )

// // func startService(sr func(*grpc.Server)) error {
// //     lis, err := net.Listen("tcp", port)
// //     if err != nil {
// //         return err
// //     }
// //     s := grpc.NewServer()
// //     sr(s)
// //     reflection.Register(s)
// //     return s.Serve(lis)
// // }

// // func main() {
// //     err := startService(func(grpcServer *grpc.Server) {
// //             pb.RegisterCollectionServer(grpcServer, &server.Server{})
// //         }, &server.Server{})

// //     if err != nil {
// //         log.Fatalf("failed to start Service: %v", err)
// //     }
// // }

// func main() {
// 	grpc.LoadSpec("abc")
// }

// // https://stackoverflow.com/questions/65561125/grpc-go-single-generic-service-handler
// //Parse protofile, create grpc.ServiceDesc, register
// func (s *GRPCService) LoadSpec(protoFileName string) {
// 	p := protoparse.Parser{}
// 	fdlist, _ := p.ParseFiles(protoFileName)
// 	for _, fd := range fdlist {
// 		for _, rsd := range fd.GetServices() {
// 			s.sdMap[rsd.GetName()] = rsd
// 			gsd := grpc.ServiceDesc{ServiceName: rsd.GetName(), HandlerType: (*interface{})(nil)}
// 			for _, m := range rsd.GetMethods() {
// 				gsd.Methods = append(gsd.Methods, grpc.MethodDesc{MethodName: m.GetName(), Handler: s.Handler})
// 			}
// 			s.grpcServer.RegisterService(&gsd, s)
// 		}
// 	}
// }

// func (s *GRPCService) Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	stream := grpc.ServerTransportStreamFromContext(ctx)
// 	arr := strings.Split(stream.Method(), "/")
// 	serviceName := arr[1]
// 	methodName := arr[2]
// 	service := s.sdMap[serviceName]
// 	method := service.FindMethodByName(methodName)
// 	input := dynamic.NewMessage(method.GetInputType())

// 	dec(input)
// 	jsonInput, err := input.MarshalJSON()
// 	log.Printf("Input:%s Err:%v \n", jsonInput, err)
// 	//jsonOutput:=invokeServiceViaReflectionOrHttp(jsonInput)
// 	jsonOutput := `{"message":"response"}`

// 	output := dynamic.NewMessage(method.GetOutputType())
// 	output.UnmarshalJSON([]byte(jsonOutput))
// 	return output, nil
// }
