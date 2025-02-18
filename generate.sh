PROTO_DIR=proto
RPC_DIR=rpc
SWAGGER_DIR=swagger

GRPC_GATEWAY="$(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2)"

CURRENT_DIR=$(pwd)

PROTOC_INCLUDES=.:"$CURRENT_DIR/$PROTO_DIR":"$GRPC_GATEWAY/third_party/googleapis"

RED='\033[0;31m'
ORANGE='\033[0;33m'
GREEN='\033[0;32m'
NO_COLOR='\033[0m'

generate() {
    echo "=========================================================="
    echo "${ORANGE}Generating: $1/$2${NO_COLOR}"
    echo "----------------------------------------------------------"

    mkdir -p "$RPC_DIR/$1"
    mkdir -p "$SWAGGER_DIR/$1"
    cd "$CURRENT_DIR/$PROTO_DIR/$1"

    protoc -I$PROTOC_INCLUDES \
        --go_out=paths=source_relative:"$CURRENT_DIR/$RPC_DIR/$1" \
        --go-grpc_out=paths=source_relative:"$CURRENT_DIR/$RPC_DIR/$1" \
        --grpc-gateway_out=logtostderr=true,paths=source_relative:"$CURRENT_DIR/$RPC_DIR/$1" \
        --openapiv2_out=logtostderr=true:"$CURRENT_DIR/$SWAGGER_DIR/$1" \
        $2
    if [ $? -ne 0 ]; then
        echo "----------------------------------------------------------"
        echo "${RED}ERROR while generating: $1/$2${NO_COLOR}"
    else
        echo "${GREEN}Generated: $1/$2${NO_COLOR}"
    fi

    cd "$CURRENT_DIR"
}

get_directory() {
    dirname $(echo $1 | sed "s/^$PROTO_DIR\///g")
}

get_proto_file() {
    basename $(echo $1 | sed "s/^$PROTO_DIR\///g")
}

PROTO_FILES=$(find $PROTO_DIR -name "*.proto")

rm -rf $RPC_DIR
rm -rf $SWAGGER_DIR
for file in $PROTO_FILES; do
    generate $(get_directory $file) $(get_proto_file $file)
done
echo "=========================================================="
