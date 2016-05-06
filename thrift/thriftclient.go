package client
 
import (
    "git.apache.org/thrift.git/lib/go/thrift"
    "net"
    "time"
)
 
var (
    transportFactory thrift.TTransportFactory       = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory  *thrift.TBinaryProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
)
 
//链接提供多种服务的客户端个体，需要指定其服务名称来选择对应的服务
type ThriftClient struct {
    host       string //服务器ip
    port       string //服务器端口
    timeOut    int    //等待服务反应的超时时间
    ServerName string //服务名称
 
    socket    *thrift.TSocket          //与服务器的链接
    transport *thrift.TFramedTransport //thrift运输框架
    protocol  thrift.TProtocol         //thrift远程代理
    Client    interface{}              //具体用来操作的客户端结构，结构由客户端和服务器端所共同用的thrift文件中服务结构决定
}
 
//初始化客户端结构，对所有面对提供多服务的客户端结构都合适
func NewthriftClient(host, port, name string, timeout int) (ret *ThriftClient, err error) {
    ret = &ThriftClient{host: host, port: port, timeOut: timeout, ServerName: name}
    ret.socket, err = thrift.NewTSocketTimeout(net.JoinHostPort(host, port), time.Duration(ret.timeOut)*time.Second)
    if err != nil {
        return nil, err
    }
    ret.transport = thrift.NewTFramedTransport(ret.socket)
    ret.protocol = thrift.NewTBinaryProtocolTransport(ret.transport)
    ret.protocol = thrift.NewTMultiplexedProtocol(ret.protocol, ret.ServerName)
    return ret, nil
}
