load("@rules_cc//cc:defs.bzl", "cc_proto_library")
load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")


load("@rules_java//java:defs.bzl", "java_proto_library")
load("@io_grpc_grpc_java//:java_grpc_library.bzl", "java_grpc_library")

#pb file
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

#proto compiler
load("@rules_proto//proto:defs.bzl", "proto_library")


package(default_visibility = ["//visibility:public"])

# Proto definitions
proto_library(
    name = "Greeter_proto",
    srcs = ["Greeter.proto"],
)




go_proto_library(
    name = "Greeter_go_proto",

    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "Go/bazel/GreeterService/Gen/GoGreeterService",

    protos = [":Greeter_proto",],
)
