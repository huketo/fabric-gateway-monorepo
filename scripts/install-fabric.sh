#!/bin/bash

# Log function for debugging
log() {
    echo "===> $(date '+%Y-%m-%d %H:%M:%S') $1"
}

die() {
    local _ret="${2:-1}"
    test "${_PRINT_HELP:-no}" = yes && print_help >&2
    echo "$1" >&2
    exit "${_ret}"
}

parse_commandline() {
    log "Parsing command line arguments"
    _positionals_count=0
    while test $# -gt 0; do
        _key="$1"
        case "$_key" in
        -f | --fabric-version)
            test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
            _arg_fabric_version="$2"
            shift
            ;;
        -c | --ca-version)
            test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
            _arg_ca_version="$2"
            shift
            ;;
        -h | --help)
            print_help
            exit 0
            ;;
        *)
            _positionals+=("$1")
            _positionals_count=$((_positionals_count + 1))
            ;;
        esac
        shift
    done
}

# Get the absolute path of the script
SCRIPT_DIR=$(dirname "$(realpath "${BASH_SOURCE[0]}")")
FABRIC_DEST_DIR=$(realpath "${SCRIPT_DIR}/../libs/fabric-network")

log "Script directory set to ${SCRIPT_DIR}"
log "Fabric destination directory set to ${FABRIC_DEST_DIR}"

# Define OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')
ARCH=$(uname -m | sed 's/x86_64/amd64/g' | sed 's/aarch64/arm64/g')
PLATFORM="${OS}-${ARCH}"

REGISTRY=${FABRIC_DOCKER_REGISTRY:-docker.io/hyperledger}

log "Platform is set to ${PLATFORM}"

singleImagePull() {
    three_digit_image_tag=$1
    shift
    two_digit_image_tag=$(echo "$three_digit_image_tag" | cut -d'.' -f1,2)
    while [[ $# -gt 0 ]]; do
        image_name="$1"
        echo "====>  ${REGISTRY}/fabric-$image_name:$three_digit_image_tag"
        ${CONTAINER_CLI} pull "${REGISTRY}/fabric-$image_name:$three_digit_image_tag"
        ${CONTAINER_CLI} tag "${REGISTRY}/fabric-$image_name:$three_digit_image_tag" "${REGISTRY}/fabric-$image_name"
        ${CONTAINER_CLI} tag "${REGISTRY}/fabric-$image_name:$three_digit_image_tag" "${REGISTRY}/fabric-$image_name:$two_digit_image_tag"
        shift
    done
}

# Function to download and extract files
download() {
    local BINARY_FILE=$1
    local URL=$2
    log "Starting download: ${URL}"
    log "Unpacking to: ${FABRIC_DEST_DIR}"
    curl -L --retry 5 --retry-delay 3 "${URL}" | tar xz -C "${FABRIC_DEST_DIR}" || rc=$?
    if [ -n "$rc" ]; then
        log "Error downloading the binary file."
        return 22
    else
        log "Download completed successfully."
    fi
}

# Function to pull binaries
pullBinaries() {
    log "Pulling Hyperledger Fabric binaries for version ${VERSION}"
    download "${BINARY_FILE}" "https://github.com/hyperledger/fabric/releases/download/v${VERSION}/${BINARY_FILE}"
    if [ $? -eq 22 ]; then
        log "Fabric binaries for version ${VERSION} are not available for download."
        exit 1
    fi

    log "Pulling Fabric CA client binaries for version ${CA_VERSION}"
    download "${CA_BINARY_FILE}" "https://github.com/hyperledger/fabric-ca/releases/download/v${CA_VERSION}/${CA_BINARY_FILE}"
    if [ $? -eq 22 ]; then
        log "Fabric CA client binaries for version ${CA_VERSION} are not available for download."
        exit 1
    fi
}

# Function to pull Docker images
pullImages() {
    command -v ${CONTAINER_CLI} >&/dev/null
    NODOCKER=$?
    if [ "${NODOCKER}" == 0 ]; then
        FABRIC_IMAGES=(peer orderer ccenv tools baseos)
        FABRIC_TAG=$VERSION
        CA_TAG=$CA_VERSION

        echo "FABRIC_IMAGES:" "${FABRIC_IMAGES[@]}"
        echo "===> Pulling fabric Images"
        singleImagePull "${FABRIC_TAG}" "${FABRIC_IMAGES[@]}"
        echo "===> Pulling fabric ca Image"
        CA_IMAGE=(ca)
        singleImagePull "${CA_TAG}" "${CA_IMAGE[@]}"
        echo "===> List out hyperledger images"
        ${CONTAINER_CLI} images | grep hyperledger
    else
        echo "========================================================="
        echo "${CONTAINER_CLI} not installed, bypassing download of Fabric images"
        echo "========================================================="
    fi
}

# Main code starts here
log "Starting the install-fabric.sh script"

# Set default values if not provided
_arg_fabric_version="${_arg_fabric_version:-2.5.10}"
_arg_ca_version="${_arg_ca_version:-1.5.13}"

parse_commandline "$@"

VERSION=$_arg_fabric_version
CA_VERSION=$_arg_ca_version

log "Fabric version is ${VERSION}, CA version is ${CA_VERSION}"

BINARY_FILE=hyperledger-fabric-${PLATFORM}-${VERSION}.tar.gz
CA_BINARY_FILE=hyperledger-fabric-ca-${PLATFORM}-${CA_VERSION}.tar.gz

log "Binary file name: ${BINARY_FILE}"
log "CA binary file name: ${CA_BINARY_FILE}"

# If no components are passed, default to binary, docker, and samples
if [ ${#_positionals[@]} -eq 0 ]; then
    log "No components passed, defaulting to binary, docker, and samples"
    _arg_comp=('binary' 'docker')
fi

if [[ "${_arg_comp[@]}" =~ (^| |,)b(inary)? ]]; then
    log "Pulling binaries"
    pullBinaries
fi

if [[ "${_arg_comp[@]}" =~ (^| |,)d(ocker)? ]]; then
    log "Pulling docker images"
    CONTAINER_CLI=docker
    pullImages
fi
