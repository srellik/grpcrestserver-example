package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "server",
    srcs = ["main.go"],
    pure = "on",
    deps = [
        "@srellik_grpcrestserver//:go_default_library",
        "//students_app/proto:students_service_go_proto",
        "//students_app/service_impl:studentsservice",
         "@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection:go_default_library",
        "@org_golang_google_grpc//credentials:go_default_library",
        "@org_golang_x_net//context:go_default_library",
        ],
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    "go_image",
    GO_DEFAULT_BASE = "DEFAULT_BASE",
)

go_image(
    name = "server_go_image",
    binary = ":server"
)
