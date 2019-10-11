#!/bin/bash
#
# Build:
#   - git clone -b dev https://github.com/zu1k/clashr
#   - cd clashr
#   - ANDROID_NDK=/path/to/android/ndk /path/to/this/script
#

export ANDROID_NDK=/home/king/Android/Sdk/ndk/20.0.5594570
# export GOPATH=/usr/lib/go

NAME=clashr
BINDIR=bin
VERSION=$(git describe --tags || echo "unknown version")
BUILDTIME=$(LANG=C date -u)
cd ..

ANDROID_CC=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang
ANDROID_CXX=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang++
ANDROID_LD=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android-ld
export GOARCH=arm64 
export GOOS=android 
export CXX=$ANDROID_CXX
export CC=$ANDROID_CC 
export LD=$ANDROID_LD 
export CGO_ENABLED=1
go build -ldflags "-X \"github.com/zu1k/clashr/constant.Version=$VERSION\" -X \"github.com/zu1k/clashr/constant.BuildTime=$BUILDTIME\" -w -s" \
            -o "build4android/clashr_arm64"


ANDROID_CC=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi21-clang
ANDROID_CXX=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi21-clang++
ANDROID_LD=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-android-ld
export GOARCH=arm
export GOOS=android 
export CXX=$ANDROID_CXX
export CC=$ANDROID_CC 
export LD=$ANDROID_LD 
export CGO_ENABLED=1
go build -ldflags "-X \"github.com/zu1k/clashr/constant.Version=$VERSION\" -X \"github.com/zu1k/clashr/constant.BuildTime=$BUILDTIME\" -w -s" \
            -o "build4android/clashr_armv7a"


ANDROID_CC=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android21-clang
ANDROID_CXX=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android21-clang++
ANDROID_LD=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android-ld
export GOOS=android 
export CXX=$ANDROID_CXX
export CC=$ANDROID_CC 
export LD=$ANDROID_LD 
export CGO_ENABLED=1
export GOARCH=386
go build -ldflags "-X \"github.com/zu1k/clashr/constant.Version=$VERSION\" -X \"github.com/zu1k/clashr/constant.BuildTime=$BUILDTIME\" -w -s" \
            -o "build4android/clashr_x86"


ANDROID_CC=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android21-clang
ANDROID_CXX=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android21-clang++
ANDROID_LD=$ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android-ld
export GOOS=android 
export CXX=$ANDROID_CXX
export CC=$ANDROID_CC 
export LD=$ANDROID_LD 
export CGO_ENABLED=1
export GOARCH=amd64
go build -ldflags "-X \"github.com/zu1k/clashr/constant.Version=$VERSION\" -X \"github.com/zu1k/clashr/constant.BuildTime=$BUILDTIME\" -w -s" \
            -o "build4android/clashr_amd64"



