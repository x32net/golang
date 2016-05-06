package server
 
import (
    "git.apache.org/thrift.git/lib/go/thrift"
    "net"
    "time"
)
 
//多服务服务端
type MultiServer struct {
    ListenHost string
    ListenPort string
    TimeOut    int
 
    ServerSocket     *thrift.TServerSocket
    TransportFactory thrift.TTransportFactory
    protocolFactory  *thrift.TBinaryProtocolFactory
    Processor        *thrift.TMultiplexedProcessor
    Server           *thrift.TSimpleServer
}
 
//单个服务处理
type ServerProcessor struct {
    ServaerName string
    Processor   thrift.TProcessor
}
 
func NewMultiServer(host, port string, timeout int) (ret *MultiServer, err error) {
    ret = &MultiServer{ListenHost: host, ListenPort: port, TimeOut: timeout}
 
    ret.ServerSocket, err = thrift.NewTServerSocketTimeout(net.JoinHostPort(ret.ListenHost, ret.ListenPort), time.Duration(ret.TimeOut)*time.Second)
    if err != nil {
        return nil, err
    }
 
    ret.TransportFactory = thrift.NewTTransportFactory()
    ret.TransportFactory = thrift.NewTFramedTransportFactory(ret.TransportFactory)
    ret.protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    ret.Processor = thrift.NewTMultiplexedProcessor()
    ret.Server = thrift.NewTSimpleServer4(ret.Processor, ret.ServerSocket, ret.TransportFactory, ret.protocolFactory)
 
    return ret, nil
}
