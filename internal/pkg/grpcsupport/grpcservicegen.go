package grpcsupport

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const (
	serviceName = "\nservice "
	grpcGenPath = "github.com/TIBCOSoftware/mashling/gen/grpc"
)

var (
	protoPath     string
	protoFileName string
	protoImpPath  string
	goPath        string
	cmdExePath    string
	appPath       string
)

//MethodInfoTree holds method information
type MethodInfoTree struct {
	MethodName    string
	MethodReqName string
	MethodResName string
	serviceName   string
}

//ProtoData holds proto file data
type ProtoData struct {
	Timestamp      time.Time
	MethodInfo     []MethodInfoTree
	ProtoImpPath   string
	RegServiceName string
	ProtoName      string
}

func assignValues() {
	appPath = "github.com/TIBCOSoftware/mashling"
	goPath = os.Getenv("GOPATH")
}

//GenerateSupportFiles creates auto genearted code
func GenerateSupportFiles(path string) error {

	path, _ = filepath.Abs(path)
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal("file path provided is invalid")
		return err
	}

	protoPath = path[:strings.LastIndex(path, string(filepath.Separator))]
	protoFileName = path[strings.LastIndex(path, string(filepath.Separator))+1:]

	log.Println("values derived from path protoPath:[", protoPath, "] protoFileName:[", protoFileName, "]")

	assignValues()

	protoImpPath = grpcGenPath + "/" + strings.Split(protoFileName, ".")[0]

	err = generatePbFiles()
	if err != nil {
		return err
	}

	pdArr, err := generateProtoData(path)
	if err != nil {
		return err
	}

	err = generateServiceImplFile(pdArr, "server")
	if err != nil {
		return err
	}

	err = generateServiceImplFile(pdArr, "client")
	if err != nil {
		return err
	}

	//buildGateway()
	return nil
}

//server template to create trigger support files
var registryServerTemplate = template.Must(template.New("").Parse(`// This file registers with grpc service. This file was auto-generated by mashling at
// {{ .Timestamp }}
package server

import (
	"encoding/json"
	"errors"
	"log"
	"fmt"
	"strings"
	servInfo "github.com/TIBCOSoftware/mashling/ext/flogo/trigger/grpc"
  	pb "{{.ProtoImpPath}}"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
{{$serviceName := .RegServiceName}}
{{$protoName := .ProtoName}}
type serviceImpl{{$protoName}}{{$serviceName}} struct {
	trigger *servInfo.GRPCTrigger
	serviceInfo *servInfo.ServiceInfo
}

var serviceInfo{{$protoName}}{{$serviceName}} = &servInfo.ServiceInfo{
	ProtoName: "{{$protoName}}",
	ServiceName: "{{$serviceName}}",
}

func init() {
	servInfo.ServiceRegistery.RegisterServerService(&serviceImpl{{$protoName}}{{$serviceName}}{serviceInfo: serviceInfo{{$protoName}}{{$serviceName}}})
}

//RunRegisterServerService registers server method implimentaion with grpc
func (s *serviceImpl{{$protoName}}{{$serviceName}}) RunRegisterServerService(serv *grpc.Server, trigger *servInfo.GRPCTrigger) {
	service := &serviceImpl{{$protoName}}{{$serviceName}}{
		trigger: trigger,
		serviceInfo: serviceInfo{{$protoName}}{{$serviceName}},
	}
	pb.Register{{$serviceName}}Server(serv, service)
}


{{- range .MethodInfo }}

func (s *serviceImpl{{$protoName}}{{$serviceName}}) {{.MethodName}}(ctx context.Context, req *pb.{{.MethodReqName}}) (res *pb.{{.MethodResName}},err error) {

	methodName := "{{.MethodName}}"

	grpcData := make(map[string]interface{})
	grpcData["methodname"] = methodName
	grpcData["contextdata"] = ctx
	grpcData["reqdata"] = req

	_, replyData, err := s.trigger.CallHandler(grpcData)

	if err != nil {
		log.Println("error: ", err)
	}

	typeHandRes := fmt.Sprintf("%T", replyData)
	if strings.Compare(typeHandRes, "*status.statusError") == 0 {
		return res, replyData.(error)
	}
	typeMethodRes := fmt.Sprintf("%T", res)
	if strings.Compare(typeHandRes, typeMethodRes) == 0 {
		res = replyData.(*pb.{{.MethodResName}})
	} else {
		var errValue = replyData.(map[string]interface{})["error"]
		if errValue != nil && len(errValue.(string)) != 0 {
			return res, errors.New(errValue.(string))
		} else {
			rDBytes, err := json.Marshal(replyData)
			if err != nil {
				log.Println("error: ", err)
			}

			err = json.Unmarshal(rDBytes, &res)
			if err != nil {
				log.Println("error: ", err)
			}
		}
	}
	log.Println("response: ", res)

	//User implimentation area

	return res, err
}

{{- end }}

func (s *serviceImpl{{$protoName}}{{$serviceName}}) ServiceInfo() *servInfo.ServiceInfo {
	return s.serviceInfo
}

`))

//client template to create grpc service support file
var registryClientTemplate = template.Must(template.New("").Parse(`// This file registers with grpc service. This file was auto-generated by mashling at
	// {{ .Timestamp }}
	package client
	
	import (
		servInfo "github.com/TIBCOSoftware/mashling/internal/pkg/model/v2/activity/service/grpc"
		pb "{{.ProtoImpPath}}"
		"google.golang.org/grpc"
	)
	{{$serviceName := .RegServiceName}}
	{{$protoName := .ProtoName}}
	type clientService{{$protoName}}{{$serviceName}} struct {
		serviceInfo *servInfo.ServiceInfo
	}
	
	var serviceInfo{{$protoName}}{{$serviceName}} = &servInfo.ServiceInfo{
		ProtoName: "{{$protoName}}",
		ServiceName: "{{$serviceName}}",
	}
	
	func init() {
		servInfo.ClientServiceRegistery.RegisterClientService(&clientService{{$protoName}}{{$serviceName}}{serviceInfo: serviceInfo{{$protoName}}{{$serviceName}}})
	}
	
	//GetRegisteredClientService returns client implimentaion stub with grpc connection
	func (cs *clientService{{$protoName}}{{$serviceName}}) GetRegisteredClientService(gCC *grpc.ClientConn) interface{} {
		return pb.New{{$serviceName}}Client(gCC)
	}
	
	func (cs *clientService{{$protoName}}{{$serviceName}}) ServiceInfo() *servInfo.ServiceInfo {
		return cs.serviceInfo
	}
	
	`))

// Exec executes a command within the build context.
func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	if len(cmdExePath) != 0 {
		cmd.Dir = cmdExePath
		cmdExePath = ""
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, string(output))
	}

	return nil
}

//generatePbFiles generates stub file based on given proto
func generatePbFiles() error {

	fullPath := filepath.Join(goPath, "src", protoImpPath)

	_, err := os.Stat(fullPath)
	if err != nil {
		os.MkdirAll(fullPath, os.ModePerm)
	}

	err = Exec("protoc", "-I", protoPath, protoPath+string(filepath.Separator)+protoFileName, "--go_out=plugins=grpc:"+fullPath)
	if err != nil {
		_, statErr := os.Stat(fullPath)
		if statErr == nil {
			os.RemoveAll(fullPath)
		}
		fmt.Println("error occured", err)
		return err
	}
	log.Println("generated files at: ", fullPath)
	return nil
}

//generateProtoData reads proto and returns proto data present in proto file
func generateProtoData(protoPath string) ([]ProtoData, error) {
	var regServiceName string
	var methodInfoList []MethodInfoTree

	bytes, err := ioutil.ReadFile(protoPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fullString := string(bytes)

	var ProtodataArr []ProtoData

	tempString := fullString
	log.Println("Number of services present in proto ", strings.Count(fullString, serviceName))
	for i := 0; i < strings.Count(fullString, serviceName); i++ {

		//getting service declaration full string
		tempString = tempString[strings.Index(tempString, serviceName):]

		//getting entire service declaration
		temp := tempString[:strings.Index(tempString, "}")+1]

		regServiceName = strings.TrimSpace(temp[strings.Index(temp, serviceName)+len(serviceName) : strings.Index(temp, "{")])
		temp = temp[strings.Index(temp, "rpc")+len("rpc"):]

		//entire rpc methods content
		methodArr := strings.Split(temp, "rpc")

		for _, mthd := range methodArr {
			methodInfo := MethodInfoTree{}
			mthdDtls := strings.Split(mthd, "(")
			methodInfo.MethodName = strings.TrimSpace(mthdDtls[0])
			methodInfo.MethodReqName = strings.TrimSpace(strings.Split(mthdDtls[1], ")")[0])
			methodInfo.MethodResName = strings.TrimSpace(strings.Split(mthdDtls[2], ")")[0])
			methodInfo.serviceName = regServiceName
			methodInfoList = append(methodInfoList, methodInfo)
		}
		protodata := ProtoData{
			MethodInfo:     methodInfoList,
			Timestamp:      time.Now(),
			ProtoImpPath:   protoImpPath,
			RegServiceName: regServiceName,
			ProtoName:      strings.Split(protoFileName, ".")[0],
		}

		ProtodataArr = append(ProtodataArr, protodata)
		methodInfoList = nil

		//getting next service content
		tempString = tempString[strings.Index(tempString, serviceName)+len(serviceName):]
	}

	return ProtodataArr, nil
}

//generateServiceImplFile creates implimentation files supported for grpc trigger and grpc service
func generateServiceImplFile(pdArr []ProtoData, option string) error {
	for _, pd := range pdArr {
		connectorFile := filepath.Join(goPath, "src", grpcGenPath, option, strings.Split(protoFileName, ".")[0]+"."+pd.RegServiceName+".grpcservice.go")
		log.Printf("Generating %s...\n", connectorFile)
		f, err := os.Create(connectorFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return err
		}
		defer f.Close()

		if strings.Compare(option, "server") == 0 {
			err = registryServerTemplate.Execute(f, pd)
		} else {
			err = registryClientTemplate.Execute(f, pd)
		}
		if err != nil {
			return err
		}
	}
	return nil
}