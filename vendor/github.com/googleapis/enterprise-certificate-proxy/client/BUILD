load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "client",
    srcs = ["client.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/googleapis/enterprise-certificate-proxy/client",
    importpath = "github.com/googleapis/enterprise-certificate-proxy/client",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/googleapis/enterprise-certificate-proxy/client/util"],
)
