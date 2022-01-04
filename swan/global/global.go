//go:binary-only-package

package global

import (
	"swan/librarys/database"
	"github.com/go-ini/ini"
	"sync"
	_ "os"
	_ "os/exec"
	_  "reflect"
	_  "sync"
	_ "swan/librarys/log"
)

type RPCError struct {
	Nil bool
	Err string
}

const (
	PLUGIN_TYPE_ONLINEMSG           = iota 			//联机业务
	PLUGIN_TYPE_ONLINEMSG_ASYSNRULE        			//联机业务异步规则
	PLUGIN_TYPE_ONLINEMSG_CHECKRULE        			//联机业务检查规则
	PLUGIN_TYPE_ONLINEMSG_SENDFAILURE	   			//消息重发失败插件
	PLUGIN_TYPE_ONLINEMSG_OVERTIMERULE	   			//消息超时应答插件
	PLUGIN_TYPE_ONLINEMSG_CLUSTEROVERTIMERULE	   	//消息超时应答插件
	PLUGIN_TYPE_ONLINEMSG_LINKNOTIFY				//通信连接通知插件
	PLUGIN_TYPE_HANDLEFILEUNPACK           			//文件解包处理插件
	PLUGIN_TYPE_FINALFILEUNPACK			   			//文件解包处理结束后插件
	PLUGIN_TYPE_HANDLEFILEPACK             			//文件打包
	PLUGIN_TYPE_HANDLEFILEPACKEND          			//文件打包文件末尾处理
	PLUGIN_TYPE_FINALFILEPACK			   			//文件打包处理结束后插件
	PLUGIN_TYPE_WEBSERVER                  			//web服务
	PLUGIN_TYPE_SAFEWEBSERVER                  		//HTTPS web服务
	PLUGIN_TYPE_BATCHTASK                  			//批处理
	PLUGIN_TYPE_TIMER								//定时服务插件
	PLUGIN_TYPE_SERIAL_MSG							//串口报文插件
	PLUGIN_TYPE_SERIAL_TAC							//串口报文校验插件
	PLUGIN_TYPE_SERIAL_ESCAPE						//串口报文转义插件
	PLUGIN_TYPE_ACCOUNT                  			//账户管理插件
	PLUGIN_TYPE_ISCS//综合监控插件
)

var (
	GConfigFileName string				//配置文件名
	GUseDBFlag		bool
	GConfigSource	string				//项目配置来源
	GCfgSourceFile	string				//项目配置来源--文件路径
	GDatabase       *database.Database	//项目配置来源--数据库信息
	GLogPath        string				//日志路径
	GConfig         *ini.File			//配置文件句柄
	GNodeName 		string 				//节点名
	GManagerIP		string
	GManagerPort	int
	GConfigMng		map[string]string
)


func (this *RPCError)Error() string {
	return ""
}

func SetRPCResult(response *map[string]interface{},err error)  {
	return
}

func Register(servicename string,plugintype int,plugin interface{}) {

}

func GetPlugin(servicename string,plugintype int) (map[string]interface{},*sync.RWMutex) {
	return nil,nil
}
func GetPluginObj(servicename string,plugintype int) []interface{} {
	return nil
}
func RegisterDBInit(f func()error) {
}

func GetDbInitPlugins() []func()error {
	return nil
}

func RegisterCfgInit(f func(*ini.File) error)  {
}

func GetCfgInitPlugins() []func(*ini.File) {
	return nil
}

func AddShareVariable(sharetype string,sharevariable map[string]interface{}) error{
	return nil
}

func DelShareVariable(sharetype string)  {
}

func GetShareVariable(sharetype string,key string) (interface{},error) {
	return nil,nil
}

func GetShareVariables(sharetype string) (map[string]interface{},error) {
	return nil,nil
}

func SetShareVariable(sharetype string,key string,value interface{}) error {
	return nil
}

func CallPlugin(servicename string, plugintype int, pluginname string, params []interface{}) ([]interface{}, error) {
	return nil, nil
}

func AddFunction(name string, fun func(...interface{})) {}

func CallFunction(name string, args ...interface{}) {}