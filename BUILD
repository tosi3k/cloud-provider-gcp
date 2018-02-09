package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "gce.go",
        "gce_address_manager.go",
        "gce_addresses.go",
        "gce_addresses_fakes.go",
        "gce_alpha.go",
        "gce_annotations.go",
        "gce_backendservice.go",
        "gce_cert.go",
        "gce_clusterid.go",
        "gce_clusters.go",
        "gce_disks.go",
        "gce_firewall.go",
        "gce_forwardingrule.go",
        "gce_forwardingrule_fakes.go",
        "gce_healthchecks.go",
        "gce_instancegroup.go",
        "gce_instances.go",
        "gce_interfaces.go",
        "gce_loadbalancer.go",
        "gce_loadbalancer_external.go",
        "gce_loadbalancer_internal.go",
        "gce_loadbalancer_naming.go",
        "gce_networkendpointgroup.go",
        "gce_op.go",
        "gce_routes.go",
        "gce_targetpool.go",
        "gce_targetproxy.go",
        "gce_urlmap.go",
        "gce_util.go",
        "gce_zones.go",
        "metrics.go",
        "support.go",
        "token_source.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/cloudprovider/providers/gce",
    deps = [
        "//pkg/api/v1/service:go_default_library",
        "//pkg/cloudprovider:go_default_library",
        "//pkg/cloudprovider/providers/gce/cloud:go_default_library",
        "//pkg/cloudprovider/providers/gce/cloud/filter:go_default_library",
        "//pkg/cloudprovider/providers/gce/cloud/meta:go_default_library",
        "//pkg/controller:go_default_library",
        "//pkg/kubelet/apis:go_default_library",
        "//pkg/master/ports:go_default_library",
        "//pkg/util/net/sets:go_default_library",
        "//pkg/util/version:go_default_library",
        "//pkg/version:go_default_library",
        "//pkg/volume:go_default_library",
        "//pkg/volume/util:go_default_library",
        "//vendor/cloud.google.com/go/compute/metadata:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/golang.org/x/oauth2/google:go_default_library",
        "//vendor/google.golang.org/api/compute/v0.alpha:go_default_library",
        "//vendor/google.golang.org/api/compute/v0.beta:go_default_library",
        "//vendor/google.golang.org/api/compute/v1:go_default_library",
        "//vendor/google.golang.org/api/container/v1:go_default_library",
        "//vendor/google.golang.org/api/googleapi:go_default_library",
        "//vendor/gopkg.in/gcfg.v1:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/fields:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/informers:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/scheme:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/core/v1:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//vendor/k8s.io/client-go/tools/record:go_default_library",
        "//vendor/k8s.io/client-go/util/flowcontrol:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "gce_address_manager_test.go",
        "gce_annotations_test.go",
        "gce_disks_test.go",
        "gce_healthchecks_test.go",
        "gce_loadbalancer_external_test.go",
        "gce_test.go",
        "gce_util_test.go",
        "metrics_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/pkg/cloudprovider/providers/gce",
    deps = [
        "//pkg/cloudprovider:go_default_library",
        "//pkg/cloudprovider/providers/gce/cloud:go_default_library",
        "//pkg/cloudprovider/providers/gce/cloud/mock:go_default_library",
        "//pkg/kubelet/apis:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
        "//vendor/golang.org/x/oauth2/google:go_default_library",
        "//vendor/google.golang.org/api/compute/v0.alpha:go_default_library",
        "//vendor/google.golang.org/api/compute/v0.beta:go_default_library",
        "//vendor/google.golang.org/api/compute/v1:go_default_library",
        "//vendor/google.golang.org/api/googleapi:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//pkg/cloudprovider/providers/gce/cloud:all-srcs",
    ],
    tags = ["automanaged"],
)
