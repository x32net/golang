package server
 
import (
    "errors"
    "pmw/utils"
    "time"
)
 
//服务接口群
var ServerArray = []*ServerProcessor{AprioriServerProcessor}
 
type ServerCluster struct {
    Servers *MultiServer
}
 
func NewServerCluster(host, port string, timeout int) (ret *ServerCluster, err error) {
    ret = &ServerCluster{}
    ret.Servers, err = NewMultiServer(host, port, timeout)
    if err != nil {
        return nil, err
    }
    for _, val := range ServerArray {
        ret.Servers.Processor.RegisterProcessor(val.ServaerName, val.Processor)
    }
    return ret, err
}
 
func (self *ServerCluster) StartServer() error {
    if self.Servers.Server == nil {
        return errors.New("Server do not build")
    }
 
    return self.Servers.Server.Serve()
}
 
func (self *ServerCluster) StopServer() error {
    if self.Servers.Server == nil {
        return errors.New("Server do not build")
    }
 
    return self.Servers.Server.Stop()
}
 
func AutoRestartServerCluster(host, port string, timeout, RSTime, timeinterval int) (err error) {
    for i := 0; i < RSTime; i++ {
        server, err := NewServerCluster(host, port, timeout)
        if err != nil { //出错原因是网络原因，所以设置了等待重连时间以免循环重连而导致重连次数无作用
            utils.Log.Error("[pmw.pipeline.thrift.NewServerCluster]: %v", err)
            time.Sleep(5e9)
            continue
        }
 
        err = server.StartServer() //任意一个服务死掉都会引起重启,重启要尽量快，以免影响反应
        if err != nil {
            utils.Log.Error("[pmw.pipeline.thrift.Servererrors]: %v", err)
        }
    }
    return
}
