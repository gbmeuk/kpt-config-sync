module kpt.dev/configsync

go 1.21

require (
	cloud.google.com/go/compute/metadata v0.2.3
	cloud.google.com/go/monitoring v1.13.0
	cloud.google.com/go/trace v1.9.0
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	github.com/GoogleContainerTools/kpt v1.0.0-beta.46
	github.com/GoogleContainerTools/kpt-functions-catalog/functions/go/set-namespace v0.4.1-0.20220713210718-d955e7d3a800
	github.com/GoogleContainerTools/kpt-functions-sdk/go/fn v0.0.0-20220706221933-7181f451a663
	github.com/Masterminds/semver v1.5.0
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/davecgh/go-spew v1.1.1
	github.com/ettle/strcase v0.1.1
	github.com/evanphx/json-patch v4.12.0+incompatible
	github.com/go-logr/logr v1.2.3
	github.com/golang/protobuf v1.5.3
	github.com/google/gnostic v0.6.9
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.14.0
	github.com/google/uuid v1.3.0
	github.com/jstemmer/go-junit-report/v2 v2.0.0
	github.com/kylelemons/godebug v1.1.0
	github.com/open-policy-agent/cert-controller v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/common v0.39.0
	github.com/spf13/cobra v1.6.1
	github.com/spyzhov/ajson v0.7.2
	github.com/stretchr/testify v1.8.1
	github.com/wk8/go-ordered-map v1.0.0
	go.opencensus.io v0.24.0
	go.uber.org/multierr v1.6.0
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
	golang.org/x/mod v0.11.0
	golang.org/x/net v0.17.0
	golang.org/x/oauth2 v0.7.0
	google.golang.org/api v0.114.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.26.9
	k8s.io/apiextensions-apiserver v0.26.9
	k8s.io/apimachinery v0.26.9
	k8s.io/cli-runtime v0.26.9
	k8s.io/client-go v0.26.9
	k8s.io/cluster-registry v0.0.6
	k8s.io/klog/v2 v2.90.1
	k8s.io/kube-aggregator v0.26.7
	k8s.io/kube-openapi v0.0.0-20230109183929-3758b55a6596
	k8s.io/kubectl v0.26.9
	k8s.io/kubernetes v1.26.9
	k8s.io/utils v0.0.0-20230115233650-391b47cb4029
	sigs.k8s.io/cli-utils v0.35.0
	sigs.k8s.io/controller-runtime v0.14.1
	sigs.k8s.io/kind v0.20.0
	sigs.k8s.io/kustomize/api v0.12.1
	sigs.k8s.io/kustomize/kyaml v0.13.9
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3
	sigs.k8s.io/yaml v1.3.0
)

require (
	cloud.google.com/go/compute v1.19.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/docker/cli v23.0.6+incompatible // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fvbommel/sortorder v1.0.1 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/safetext v0.0.0-20220905092116-b49f7bc46da2 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/gomega v1.24.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vbatts/tar-split v0.11.2 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	go.starlark.net v0.0.0-20210901212718-87f333178d59 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.56.3 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiserver v0.26.9 // indirect
	k8s.io/component-base v0.26.9 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
)
