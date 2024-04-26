#!/bin/bash
echo 'Compiling for Kazan'
GOOS=linux GOARCH=arm  go build
if [ $? -ne 0 ]; then
	echo 'An error has occurred! Aborting the script execution...'
	exit 1
fi
echo 'Copy radar to device Kazan'
# tar -czvf potop.tar.gz potop
scp -P 222 radar root@185.27.195.194:/cache 
