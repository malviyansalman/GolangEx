load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "bf0e5070b4b99240183b29df78155eee335885e53a8af8683964579c214ad301",
    strip_prefix = "protobuf-3.14.0",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.14.0.zip"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d8c45ee70ec39a57e7a05e5027c32b1576cc7f16d9dd37135b0eddde45cf1b10",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6f111c57fd50baf5b8ee9d63024874dd2a014b069426156c55adbf6d3d22cb7b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.0/rules_go-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.0/rules_go-v0.25.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc",
    sum = "h1:raiipEjMOIC/TO2AvyTxP25XFdLxNIBwzDh3FM3XztI=",
    version = "v1.34.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
    version = "v0.0.0-20190311183353-d8887717615a",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version = "v0.3.0",
)

# bazel_skylib --> 1.0.2
# platforms --> 0.0.1
# org_golang_x_tools -->  # master, as of 2020-08-24
# org_golang_x_errors -->  # master, as of 2020-08-24
# rules_cc --> # master as of 2020-12-01
# org_golang_google_protobuf --> # v1.25.0, latest as of 2020-12-01
# com_github_golang_protobuf --> # v1.4.3
# com_github_mwitkow_go_proto_validators -->    # v0.3.2, latest as of 2020-12-01
# com_github_gogo_protobuf -->   # v1.3.1, latest as of 2020-12-01
# org_golang_google_genproto -->  # master, as of 2020-12-01
# go_googleapis -->  # master, as of 2020-08-24
go_rules_dependencies()
go_register_toolchains(version = "1.15.6")

gazelle_dependencies()

# Running protbuf_deps brings following http_archives with version/commit id
# bazel_skylib --> v1.0.2, skipped
# zlib --> v1.2.11
# six --> 1.12.0
# rules_proto --> 40298556293ae502c66579620a7ce867d5f57311
# rules_cc --> 818289e5613731ae410efb54218a4077fb9dbb03, skipped, due to above call
# rules_java --> 981f06c3d2bd10225e85209904090eb7b5fb26bd
# rules_python --> 4b84ad270387a7c439ebdccfd530e2339601ef27
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "12a4a6f8c06b96e38f8576ded76d0b79bce13efd7560ed22134c2f433bc496ad",
    urls = [
        "https://github.com/grpc/grpc/archive/v1.41.1.tar.gz",
    ],
    strip_prefix = "grpc-1.41.1",
)


# ================================================================
# Python GRPC support requires rules_python
# ================================================================

#github_archive(
#    name = "io_bazel_rules_python",
#    commit = "8b5d0683a7d878b28fffe464779c8a53659fc645",
#    org = "bazelbuild",
#    repo = "rules_python",
#    sha256 = "40499c0a9d55f0c5deb245ed24733da805f05aaf6085cb39027ba486faf1d2e1",
#)

#load("@io_bazel_rules_python//python:pip.bzl", "pip_repositories", "pip_import")

#pip_repositories()

#pip_import(
#   name = "pip_grpcio",
#   requirements = "//python:requirements.txt",
#)


# GRPC deps bring many libraries please refer
# https://github.com/grpc/grpc/blob/d910557f22ab445a3618a358d7a6bb17836c7151/bazel/grpc_deps.bzl
# Important dependencies it brings are
# com_google_protobuf which inturn fetches bazel_skylib, rules_proto, rules_cc, rules_java, rules_python
# io_bazel_rules_go, bazel_toolchains, bazel_skylib
# envoy_api, libuv
# Google absl, flags, re2, benchmark ...
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps", "grpc_test_only_deps")
grpc_deps()
grpc_test_only_deps()

# GRPC extra deps, will bring in
# go_rules_dependencies, go_register_toolchains
# apple_rules_dependencies, apple_support_dependencies
load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")
grpc_extra_deps()


# Get Java GRPC dependencies
http_archive(
    name = "io_grpc_grpc_java",
    sha256 = "7978db73f4bb0fdc523c042a0fc347c046b4345823dbeadfec15ceed99d50737",
    urls = ["https://github.com/grpc/grpc-java/archive/v1.41.1.tar.gz",],
    strip_prefix = "grpc-java-1.41.1",
)

http_archive(
    name = "rules_jvm_external",
    sha256 = "f36441aa876c4f6427bfb2d1f2d723b48e9d930b62662bf723ddfb8fc80f0140",
    strip_prefix = "rules_jvm_external-4.1",
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/4.1.zip",
)

load("@rules_jvm_external//:defs.bzl", "maven_install")
load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_ARTIFACTS")
load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS")

maven_install(
    artifacts = [
        "com.google.api.grpc:grpc-google-cloud-pubsub-v1:0.1.24",
        "com.google.api.grpc:proto-google-cloud-pubsub-v1:0.1.24",
    ] + IO_GRPC_GRPC_JAVA_ARTIFACTS,
    generate_compat_repositories = True,
    override_targets = IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS,
    repositories = [
        "https://repo.maven.apache.org/maven2/",
    ],
)

# Compat repositories build, maven to bazel name conversion
load("@maven//:compat.bzl", "compat_repositories")
compat_repositories()

# GRPC Java repositories will `bind` (create alias) for google libs
# Google libs are gson, guava, error prone
# Also tries to bringing com_google_protobuf, com_google_protobuf_javalite, io_grpc_grpc_proto
load("@io_grpc_grpc_java//:repositories.bzl", "grpc_java_repositories")
grpc_java_repositories()

#load("//python:rules.bzl", "py_proto_repositories")
