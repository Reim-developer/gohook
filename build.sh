#!/usr/bin/bash



function tool_check() {
    if ! command -v go > /dev/null; then 
        echo "Could not find Go in your operating system. Aborting."
        exit 1
    fi 
}

function build() {
    local OS_LIST=("linux" "darwin" "windows")
    local OS_ARCH="amd64"
    local DIST="dist"

    if [ ! -d "$DIST" ]; then 
        mkdir -p "$DIST"
    fi 

    for OS in "${OS_LIST[@]}"; do 
        local OUTPUT="gohook-${OS}-${OS_ARCH}"

        if [[ "$OS" == "windows" ]]; then
            OUTPUT="${OUTPUT}.exe"
        fi 

        echo "[!] Building for $OS/$ARCH -> $OUTPUT"
        GOOS=$OS GOARCH=$OS_ARCH go build -o "$DIST/$OUTPUT"
    done

    echo "Build success."
}

function main() {
    tool_check
    build
}

main