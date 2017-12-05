package main

import (
	"flag"
	"os"
	"path/filepath"

	pb "github.com/kwins/iceberg/demo/s1/pb"
	"github.com/kwins/iceberg/frame/config"
)

var (
	cfgFile  = flag.String("config-path", "icetest_conf.json", "config file")
	logLevel = flag.String("level", "debug", "log level")
	logPath  = flag.String("logPath", "", "log path")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 设置进程的当前目录为程序所在的路径
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	var baseCfg config.BaseCfg

	var etcdCfg config.EtcdCfg
	etcdCfg.EndPoints = []string{"http://127.0.0.1:2379"}
	etcdCfg.User = "iceberg"
	etcdCfg.Psw = "123456"
	etcdCfg.Timeout = 3

	var zipkinCfg config.ZipkinCfg
	// zipkinCfg.EndPoints = "http://localhost:9411/api/v1/spans"
	baseCfg.Etcd = etcdCfg
	baseCfg.Zipkin = zipkinCfg
	// 直接注册就行了
	pb.RegisterHelloServer(new(Hello), &baseCfg)
}
