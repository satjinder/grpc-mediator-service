package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

//https://stackoverflow.com/questions/65242456/convert-protobuf-serialized-messages-to-json-without-precompiling-go-code

func main() {
	_, err := createProtoRegistry("statsservice", "stats.proto")
	if err != nil {
		panic(err)
	}
}

func createProtoRegistry(srcDir string, filename string) (*protoregistry.Files, error) {
	// Create descriptors using the protoc binary.
	// Imported dependencies are included so that the descriptors are self-contained.
	tmpFile := "schemaregister/" + filename + "-tmp.pb"

	//fmt.Println(tmpFile)
	cmd := exec.Command("protoc",
		"--include_imports",
		"--descriptor_set_out="+tmpFile,
		"-I"+srcDir,
		path.Join(srcDir, filename))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// defer os.Remove(tmpFile)

	marshalledDescriptorSet, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	//fmt.Println(descriptorSet)
	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}
	return files, nil
}
