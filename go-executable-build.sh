#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

package_name=$package

platforms=("windows/amd64" "windows/386" "darwin/amd64")
cd client
yarn build
# shellcheck disable=SC2103
cd ..
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='./build/'$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done