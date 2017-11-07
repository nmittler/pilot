load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "discovery.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//model:go_default_library",
        "//proxy:go_default_library",
        "@com_github_emicklei_go_restful//:go_default_library",
        "@com_github_golang_glog//:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_golang_protobuf//ptypes/duration:go_default_library",
        "@com_github_hashicorp_go_multierror//:go_default_library",
        "@com_github_howeyc_fsnotify//:go_default_library",
        "@io_istio_api_mixer_client//:mixer/v1/config/client",
        "@io_istio_api_mixer//:mixer/v1",
        "@io_istio_api//:go_default_library",
    ],
)

