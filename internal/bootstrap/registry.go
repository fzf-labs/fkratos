package bootstrap

import (
	"fkratos/internal/bootstrap/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
	"path/filepath"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// nacos
	nacosKratos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	// kubernetes
	k8sRegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	k8s "k8s.io/client-go/kubernetes"
	k8sRest "k8s.io/client-go/rest"
	k8sTools "k8s.io/client-go/tools/clientcmd"
	k8sUtil "k8s.io/client-go/util/homedir"

	// polaris
	polarisKratos "github.com/go-kratos/kratos/contrib/registry/polaris/v2"
)

type RegistryType string

const (
	RegistryTypeConsul   RegistryType = "consul"
	LoggerTypeEtcd       RegistryType = "etcd"
	LoggerTypeNacos      RegistryType = "nacos"
	LoggerTypeKubernetes RegistryType = "kubernetes"
	LoggerTypePolaris    RegistryType = "polaris"
)

// NewRegistryAndDiscovery 创建一个注册和发现客户端
func NewRegistryAndDiscovery(cfg *conf.Registry) (registry.Registrar, registry.Discovery) {
	if cfg == nil {
		return nil, nil
	}
	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		res := NewConsulRegistry(cfg)
		return res, res
	case LoggerTypeEtcd:
		res := NewEtcdRegistry(cfg)
		return res, res
	case LoggerTypeNacos:
		res := NewNacosRegistry(cfg)
		return res, res
	case LoggerTypeKubernetes:
		res := NewKubernetesRegistry(cfg)
		return res, res
	case LoggerTypePolaris:
		res := NewPolarisRegistry(cfg)
		return res, res
	}
	fmt.Println("Bootstrap NewRegistry Success")
	return nil, nil
}

// NewConsulRegistry 创建一个注册发现客户端 - Consul
func NewConsulRegistry(c *conf.Registry) *consulKratos.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := consulKratos.New(cli, consulKratos.WithHealthCheck(c.Consul.HealthCheck))

	return reg
}

// NewEtcdRegistry 创建一个注册发现客户端 - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdKratos.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		log.Fatal(err)
	}

	reg := etcdKratos.New(cli)

	return reg
}

// NewNacosRegistry 创建一个注册发现客户端 - Nacos
func NewNacosRegistry(c *conf.Registry) *nacosKratos.Registry {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(c.Nacos.Address, c.Nacos.Port),
	}

	cliConf := nacosConstant.ClientConfig{
		NamespaceId:          c.Nacos.NamespaceId,
		TimeoutMs:            uint64(c.Nacos.Timeout.AsDuration().Milliseconds()), // http请求超时时间，单位毫秒
		BeatInterval:         c.Nacos.BeatInterval.AsDuration().Milliseconds(),    // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      int(c.Nacos.UpdateThreadNum),                        // 更新服务的线程数
		LogLevel:             c.Nacos.LogLevel,
		CacheDir:             c.Nacos.CacheDir,             // 缓存目录
		LogDir:               c.Nacos.LogDir,               // 日志目录
		NotLoadCacheAtStart:  c.Nacos.NotLoadCacheAtStart,  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: c.Nacos.UpdateCacheWhenEmpty, // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	cli, err := nacosClients.NewNamingClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	reg := nacosKratos.New(cli)

	return reg
}

// NewKubernetesRegistry 创建一个注册发现客户端 - Kubernetes
func NewKubernetesRegistry(_ *conf.Registry) *k8sRegistry.Registry {
	restConfig, err := k8sRest.InClusterConfig()
	if err != nil {
		home := k8sUtil.HomeDir()
		kubeConfig := filepath.Join(home, ".kube", "config")
		restConfig, err = k8sTools.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	clientSet, err := k8s.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	reg := k8sRegistry.NewRegistry(clientSet)

	return reg
}

// NewPolarisRegistry 创建一个注册发现客户端 - Polaris
func NewPolarisRegistry(c *conf.Registry) *polarisKratos.Registry {
	address := fmt.Sprintf("%s:%d", c.Polaris.Address, c.Polaris.Port)
	configuration := config.NewDefaultConfiguration([]string{address})
	consumer, err := api.NewConsumerAPIByConfig(configuration)
	if err != nil {
		panic(err)
	}
	provider := api.NewProviderAPIByContext(consumer.SDKContext())
	reg := polarisKratos.NewRegistry(
		provider,
		consumer,
		polarisKratos.WithNamespace(c.Polaris.Namespace),
		polarisKratos.WithServiceToken(c.Polaris.Token),
		polarisKratos.WithTTL(10),
	)
	return reg
}
