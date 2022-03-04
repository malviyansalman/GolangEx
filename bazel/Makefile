BASE_DIR := $(shell pwd)
PROJECT_NAME := Contracts

.SILENT: start

.PHONY: all

all: clean package

clean:
	mvn clean

compile:
	mvn compile

package:
	mvn package

install:
	mvn clean install
install_common_local:
	bazel build //Common:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCommon
	cp -f -R ./bazel-bin/Common/libcommon_enums_proto-speed.jar Common/Gen/JavaCommon
	mvn install:install-file -Dfile=Common/Gen/JavaCommon/libcommon_enums_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=common_enums_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_cms_local:
	bazel build //ContentManagementService:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCMS
	cp -f -R ./bazel-bin/ContentManagementService/libcontent_management_service_proto-speed.jar ContentManagementService/Gen/JavaContentManagementService
	mvn install:install-file -Dfile=ContentManagementService/Gen/JavaContentManagementService/libcontent_management_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=content_management_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

force_clean:
	mvn dependency:purge-local-repository clean install

build_GreeterService:
	bazel build //TestService:all  --sandbox_debug --verbose_failures
		./copy_go_libs_and_sources.sh

