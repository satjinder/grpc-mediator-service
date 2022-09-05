package localschemaregistry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type LocalSchemaRegistry struct {
	basePath string
	regFiles map[string]*protoregistry.Files
}

func New(basePath string) *LocalSchemaRegistry {
	return &LocalSchemaRegistry{
		basePath: basePath,
		regFiles: make(map[string]*protoregistry.Files)}
}

func createProtoRegistry(path string) (*protoregistry.Files, error) {
	marshalledDescriptorSet, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (ls *LocalSchemaRegistry) Get(registryName string, protoFile string) (protoreflect.FileDescriptor, error) {
	if _, ok := ls.regFiles[registryName]; !ok {
		path := filepath.Join(ls.basePath, registryName)
		reg, err := createProtoRegistry(path)
		if err != nil {
			return nil, err
		}
		ls.regFiles[registryName] = reg
	}

	files := ls.regFiles[registryName]
	return files.FindFileByPath(protoFile)

}
