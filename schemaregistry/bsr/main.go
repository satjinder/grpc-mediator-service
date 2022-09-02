package localschemaregistry

import (
	"context"
	"flag"
	"net/http"

	"github.com/bufbuild/connect-go"
	bpb1 "go.buf.build/bufbuild/connect-go/bufbuild/buf/buf/alpha/registry/v1alpha1"
	bapb1 "go.buf.build/bufbuild/connect-go/bufbuild/buf/buf/alpha/registry/v1alpha1/registryv1alpha1connect"

	//bufpb "go.buf.build/grpc/go/bufbuild/buf/buf/alpha/registry/v1alpha1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	addr = flag.String("addr", "api.buf.build:443", "the address to connect to")
)

type BufSchemaRegistry struct {
	regFiles map[string]*protoregistry.Files
}

func New() *BufSchemaRegistry {
	return &BufSchemaRegistry{
		regFiles: make(map[string]*protoregistry.Files)}
}

func (ls *BufSchemaRegistry) Get(registryName string, protoFile string) (protoreflect.FileDescriptor, error) {
	if _, ok := ls.regFiles[registryName]; !ok {
		reg, err := createProtoRegistry()
		if err != nil {
			return nil, err
		}
		ls.regFiles[registryName] = reg
	}

	files := ls.regFiles[registryName]
	return files.FindFileByPath(protoFile)

}

func createProtoRegistry() (*protoregistry.Files, error) {
	client := bapb1.NewImageServiceClient(
		http.DefaultClient,
		"https://api.buf.build",
	)
	res, err := client.GetImage(
		context.Background(),
		connect.NewRequest(&bpb1.GetImageRequest{
			Owner:             "satjinder",
			Repository:        "schemas",
			Reference:         "main",
			ExcludeSourceInfo: true,
		}),
	)
	if err != nil {
		panic(err)
	}

	// treat buf.alpha.image.v1.Image as google.protobuf.FileDescriptorSet
	bin, err := proto.Marshal(res.Msg.GetImage())
	if err != nil {
		panic(err)
	}
	fileDescriptorSet := &descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(bin, fileDescriptorSet)
	if err != nil {
		panic(err)
	}

	return protodesc.NewFiles(fileDescriptorSet)

}
