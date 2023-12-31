#!/usr/bin/env bash

set -e

source ./scripts/test_lib.sh

VER=$1
PROJ="etcd"
REPOSITORY="${REPOSITORY:-https://github.com/etcd-io/etcd.git}"

if [ -z "$1" ]; then
	echo "Usage: ${0} VERSION" >> /dev/stderr
	exit 255
fi

set -u

function setup_env {
	local proj=${1}
	local ver=${2}

	if [ ! -d "${proj}" ]; then
	  log_callout "Cloning ${REPOSITORY}..."
	  git clone "${REPOSITORY}"
	fi

	pushd "${proj}" >/dev/null
		git fetch --all
		git checkout "${ver}"
	popd >/dev/null
}


function package {
	local target=${1}
	local srcdir="${2}/bin"

	local ccdir="${srcdir}/${GOOS}_${GOARCH}"
	if [ -d "${ccdir}" ]; then
		srcdir="${ccdir}"
	fi
	local ext=""
	if [ "${GOOS}" == "windows" ]; then
		ext=".exe"
	fi
	for bin in etcd etcdctl; do
		cp "${srcdir}/${bin}" "${target}/${bin}${ext}"
	done

	cp etcd/README.md "${target}"/README.md
	cp etcd/etcdctl/README.md "${target}"/README-etcdctl.md
	cp etcd/etcdctl/READMEv2.md "${target}"/READMEv2-etcdctl.md

	cp -R etcd/Documentation "${target}"/Documentation
}

function main {
	mkdir -p release
	cd release
	setup_env "${PROJ}" "${VER}"

	tarcmd=tar
	if [[ $(go env GOOS) == "darwin" ]]; then
		echo "Please use linux machine for release builds."
		exit 1
	fi

	for os in darwin windows linux; do
		export GOOS=${os}
		TARGET_ARCHS=("amd64")

		if [ ${GOOS} == "linux" ]; then
			TARGET_ARCHS+=("arm64")
			TARGET_ARCHS+=("ppc64le")
		fi

		for TARGET_ARCH in "${TARGET_ARCHS[@]}"; do
			export GOARCH=${TARGET_ARCH}

			pushd etcd >/dev/null
			GO_LDFLAGS="-s" ./build
			popd >/dev/null

			TARGET="etcd-${VER}-${GOOS}-${GOARCH}"
			mkdir "${TARGET}"
			package "${TARGET}" "${PROJ}"

			if [ ${GOOS} == "linux" ]; then
				${tarcmd} cfz "${TARGET}.tar.gz" "${TARGET}"
				echo "Wrote release/${TARGET}.tar.gz"
			else
				zip -qr "${TARGET}.zip" "${TARGET}"
				echo "Wrote release/${TARGET}.zip"
			fi
		done
	done
}

main
