#!/bin/bash

# An example hook script to verify what is about to be pushed.  Called by "git
# push" after it has checked the remote status, but before anything has been
# pushed.  If this script exits with a non-zero status nothing will be pushed.
#
# This hook is called with the following parameters:
#
# $1 -- Name of the remote to which the push is being done
# $2 -- URL to which the push is being done
#
# If pushing without using a named remote those arguments will be equal.
#
# Performs static go linting on sources 

remote="$1"
url="$2"

echo -e "Performing linting on go sources..."
printf "%0.s-" {1..100}
echo

cd Proj1
docker run -v $(pwd)/src:/src bkstud/ebizness:latest /bin/bash -c '/root/go/bin/golangci-lint run $(find /src -type f)'
output=$?
printf "%0.s-" {1..100}
if [ $output -ne 0 ]
then
	echo -en "\nLinting failed!"
	echo
	exit 1
fi

echo
exit 0
