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
	org      string
	repo     string
	ref      string
}

func New(org string, repo string, ref string) *BufSchemaRegistry {
	return &BufSchemaRegistry{
		org:      org,
		repo:     repo,
		ref:      ref,
		regFiles: make(map[string]*protoregistry.Files)}
}

func (bs *BufSchemaRegistry) Get(registryName string, protoFile string) (protoreflect.FileDescriptor, error) {
	if _, ok := bs.regFiles[registryName]; !ok {
		reg, err := bs.createProtoRegistry()
		if err != nil {
			return nil, err
		}
		bs.regFiles[registryName] = reg
	}

	files := bs.regFiles[registryName]

	return files.FindFileByPath(protoFile)

}

func (bs *BufSchemaRegistry) createProtoRegistry() (*protoregistry.Files, error) {
	client := bapb1.NewImageServiceClient(
		http.DefaultClient,
		"https://api.buf.build",
	)
	res, err := client.GetImage(
		context.Background(),
		connect.NewRequest(&bpb1.GetImageRequest{
			Owner:             bs.org,
			Repository:        bs.repo,
			Reference:         bs.ref,
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
