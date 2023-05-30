#!/bin/bash
# Require root, since we want to create folders in a location where we maybe dont have access to
if [ "$(whoami)" != "root" ]; then
        echo "ERROR Install script must be run as root"
        exit -1
fi

executableName="deepeye-v0.0.1"
downloadUrl="https://github.com/aallali/DeepEye/releases/download/v0.0.1/deepeye-v0.0.1.tar.gz"
# TODO: wget path to executable from github repo
# ...
wget -O ${executableName}.tar.gz ${downloadUrl}

tar -xf ${executableName}.tar.gz


# delete old binary
rm -rf /usr/local/bin/deepeye

# # move new binary to its place
mv ${executableName} /usr/local/bin/deepeye

deepeye -u


echo """
execute this on your .zshrc or .bashrc in order to check for updates 

-------------------: [zsh] :-------------------
echo \"deepeye -u\" >> ~/.zshrc
source ~/.zshrc
-----------------------------------------------

or 

-------------------: [bash] :------------------
echo \"deepeye -u\" >> ~/.bashrc
source ~/.bashrc
-----------------------------------------------

"""