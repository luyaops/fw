package main

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"io/ioutil"
	"luyaops/api-gateway/types"
	"luyaops/fw/common/log"
	"luyaops/fw/third_party/google/api"
	"luyaops/fw/third_party/runtime"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panic("ioutil.ReadAll error:", err)
	}

	request := new(plugin_go.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, request); err != nil {
		log.Panic("proto.Unmarshal error:", err)
	}

	var methods []types.MethodWrapper
	var server string
	for _, allProtoBuff := range request.GetProtoFile() {
		for _, generateProtoBuff := range request.FileToGenerate {
			if *allProtoBuff.Name == generateProtoBuff {
				for _, service := range allProtoBuff.Service {
					for _, md := range service.Method {
						method := types.MethodWrapper{}
						options := make(map[string]interface{})

						method.Package = *allProtoBuff.Package
						method.Service = *service.Name
						method.Method = md
						if ext, err := proto.GetExtension(md.Options, google_api.E_Http); err == nil {
							pattern := types.Pattern{}
							rule := ext.(*google_api.HttpRule)
							verb, path := getVerbAndPath(rule)
							pattern.Verb = verb
							pattern.Path = path
							pattern.Body = rule.Body
							method.Pattern = pattern
							//options[google_api.E_Http.Name] = rule
						}
						if aut, err := proto.GetExtension(md.Options, runtime.E_Authentication); err == nil {
							au := aut.(*bool)
							options[runtime.E_Authentication.Name] = au
						}
						method.Options = options
						method.Method.Name = md.Name
						// 去除点
						inputType := strings.TrimLeft(*md.InputType, ".")
						method.Method.InputType = &inputType
						outputType := strings.TrimLeft(*md.OutputType, ".")
						method.Method.OutputType = &outputType
						server = *allProtoBuff.Package
						methods = append(methods, method)
					}
				}
			}
		}
	}
	if len(methods) > 0 {
		jsonOut, err := json.Marshal(methods)
		if err != nil {
			log.Panic("json.Marshal error:", err)
		}
		header := `package service

// Code generated by protoc-gen-parse. DO NOT EDIT.
import (
	"luyaops/fw/common/etcd"
	"luyaops/fw/common/register"
)
`
		body := "\nvar apiJson = " + strconv.Quote(string(jsonOut)) +
			"\nvar store = etcd.NewStore(\"localhost:2379\")" +
			"\nvar server = \"" + server + "\"\n\n"
		function :=
			`func init() {
	register.Register(store, server, apiJson)
}`
		err = ioutil.WriteFile("internal/impl/service/register.go", []byte(header+body+function), 0644)
		if err != nil {
			log.Panic("ioutil.WriteFile error:", err)
		}
	}
}

func getVerbAndPath(opts *google_api.HttpRule) (string, string) {
	var httpMethod, path string
	switch {
	case opts.GetGet() != "":
		httpMethod = "GET"
		path = opts.GetGet()
	case opts.GetPost() != "":
		httpMethod = "POST"
		path = opts.GetPost()
	case opts.GetPut() != "":
		httpMethod = "PUT"
		path = opts.GetPut()
	case opts.GetDelete() != "":
		httpMethod = "DELETE"
		path = opts.GetDelete()
	case opts.GetPatch() != "":
		httpMethod = "PATCH"
		path = opts.GetPatch()
	case opts.GetCustom() != nil:
		custom := opts.GetCustom()
		httpMethod = custom.Kind
		path = custom.Path
	}
	return httpMethod, path
}
