#!/bin/bash
# Require root, since we want to create folders in a location where we maybe dont have access to
if [ "$(whoami)" != "root" ]; then
        echo "ERROR Install script must be run as root"
        exit -1
fi

# TODO: wget path to executable from github repo
# ...

# delete old binary
rm -rf /usr/local/bin/deepeye

# move new binary to its place
mv deepeye /usr/local/bin/