#!/usr/bin/bash
function tool_check() {
    if ! command -v go > /dev/null; then 
        echo "Could not find Go in your operating system. Aborting."
        exit 1
    fi

    if ! command -v zip > /dev/null; then
        echo "Could not find zip in  your operating system. Aborting."
        exit 1
    fi

    if ! command -v tar > /dev/null; then
        echo "Could not find tar in your operating system. Aborting."
        exit 1
    fi 
}

function build() {
    local OS_LIST=("linux" "darwin" "windows")
    local OS_ARCH="amd64"
    local DIST="dist"
    local ZIP_DIST="$DIST/zip"
    local CHECKSUM_FILE="$ZIP_DIST/checksum.txt"

    if [ ! -d "$DIST" ]; then 
        mkdir -p "$DIST"
    fi 

    if [ ! -d "$ZIP_DIST" ]; then
        mkdir -p "$ZIP_DIST"
    fi

    : > "$CHECKSUM_FILE"

    for OS in "${OS_LIST[@]}"; do 
        local OUTPUT="gohook-${OS}-${OS_ARCH}"

        if [[ "$OS" == "windows" ]]; then
            OUTPUT="${OUTPUT}.exe"
        fi 

        echo "[!] Building for $OS/$OS_ARCH -> $OUTPUT"

        GOOS=$OS GOARCH=$OS_ARCH go build -o "$DIST/$OUTPUT"
        
        local ARCHIVE_NAME=""
        if [[ "$OS" == "windows" ]];then
            ARCHIVE_NAME="GoHook_Windows.zip"
            zip -j "$ZIP_DIST/$ARCHIVE_NAME" "$DIST/$OUTPUT"

        else
            ARCHIVE_NAME="GoHook_$OS.tar.gz"
            tar -czvf "$ZIP_DIST/GoHook_$OS.tar.gz" "$DIST/$OUTPUT"
        
        fi
        sha256sum "$ZIP_DIST/$ARCHIVE_NAME" >> "$CHECKSUM_FILE"
    done

    echo "Build success."
}

function main() {
    tool_check
    build
}

main