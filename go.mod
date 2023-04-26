module fkratos

go 1.20

require (
	github.com/dtm-labs/rockscache v0.1.0
	github.com/fzf-labs/fpkg v1.1.7
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/config/kubernetes/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/config/polaris/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/log/aliyun/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/contrib/registry/polaris/v2 v2.0.0-20230410063538-3958f9d5c06e
	github.com/go-kratos/kratos/v2 v2.6.1
	github.com/go-redis/redis/v8 v8.11.6-0.20220405070650-99c79f7041fc
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.20.0
	github.com/hibiken/asynq v0.23.0
	github.com/nacos-group/nacos-sdk-go v1.1.4
	github.com/polarismesh/polaris-go v1.3.0
	github.com/prometheus/client_golang v1.14.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.15.0
	go.etcd.io/etcd/client/v3 v3.5.6
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/exporters/zipkin v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/zap v1.23.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/gen v0.3.21
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
	gorm.io/plugin/dbresolver v1.4.1
	k8s.io/client-go v0.27.0
)

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1716 // indirect
	github.com/aliyun/aliyun-log-go-sdk v0.1.43 // indirect
	github.com/apolloconfig/agollo/v4 v4.3.0 // indirect
	github.com/armon/go-metrics v0.4.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-kratos/aegis v0.1.4 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/lithammer/shortuuid v3.0.0+incompatible // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.41.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.6 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.6 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.1.1 // indirect
	gorm.io/driver/mysql v1.4.7 // indirect
	gorm.io/driver/postgres v1.5.0 // indirect
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/opentelemetry v0.1.1 // indirect
	k8s.io/api v0.27.0 // indirect
	k8s.io/apimachinery v0.27.0 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230308215209-15aac26d736a // indirect
	k8s.io/utils v0.0.0-20230209194617-a36077c30491 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
