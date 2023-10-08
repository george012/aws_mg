#!/bin/bash
set -e

product_name=aws_mg
build_path=./build
run_mode=release
OS_TYPE="Unknown"
GetOSType() {
    uNames=`uname -s`
    osName=${uNames: 0: 4}
    if [ "$osName" == "Darw" ] # Darwin
    then
        OS_TYPE="Darwin"
    elif [ "$osName" == "Linu" ] # Linux
    then
        OS_TYPE="Linux"
    elif [ "$osName" == "MING" ] # MINGW, windows, git-bash
    then
        OS_TYPE="Windows"
    else
        OS_TYPE="Unknown"
    fi
}
GetOSType

function toBuild() {
    if [[ "$OS_TYPE" != "Linux" && "$run_mode" == "release" ]]; then
        echo "release build must to linux OS"
        exit 0
    fi

    rm -rf ${build_path}/${run_mode}
    mkdir -p ${build_path}/${run_mode}

    go_version=$(go version | awk '{print $3}')
    commit_hash=$(git show -s --format=%H)
    commit_date=$(git show -s --format="%ci")
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        formatted_date=$(date -u -jf "%Y-%m-%d %H:%M:%S %z" "${commit_date}" "+%Y-%m-%d_%H:%M:%S")
    else
        # Linux
        formatted_date=$(date -u -d "${commit_date}" "+%Y-%m-%d_%H:%M:%S")
    fi

    build_time=$(date +"%Y-%m-%d_%H:%M:%S")

    ld_flag_master="-X main.mGitCommitHash=${commit_hash} -X main.mGitCommitDate=${formatted_date} -X main.mGoVersion=${go_version} -X main.mPackageOS=${OS_TYPE} -X main.mPackageTime=${build_time} -X main.mRunMode=${run_mode} -s -w"

    out_file=${build_path}/${run_mode}/${product_name}
    if [[ "$OS_TYPE" == "Windows" ]]; then
        out_file=${build_path}/${run_mode}/${product_name}.exe
    fi

    echo "buid file with " ${out_file}

    go build -o ${out_file} -trimpath -ldflags "${ld_flag_master}" main.go

    package_files

}

function package_files(){
    cd $build_path \
    && if [[ "$OS_TYPE" == "Windows" ]]; then
          7z a ./${product_name}_${run_mode}.zip ./${run_mode} >/dev/null 2>&1
        else
            zip -r ./${product_name}_${run_mode}.zip ./${run_mode}
        fi \
    && cd ../
}


function handlerunMode() {
    if [[ "$1" == "release" || "$1" == "" ]]; then
        run_mode=release
    elif [[ "$1" == "test" ]]; then
        run_mode=test
    else
        echo "Usage: bash build.sh [release|test],default with:release"
        exit 0
    fi
}


handlerunMode "$1" && toBuild

