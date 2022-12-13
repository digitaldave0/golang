#!/bin/sh
sudo --preserve-env=HOME /usr/bin/env
URL='https://golang.org/dl/go1.15.5.linux-amd64.tar.gz'
FILE=${URL##*/}
if [ -f "/tmp/$FILE" ]; then
    echo "File, $FILE found will delete as old and run"
    rm -f "/tmp/$FILE"
else
    echo "installing go $"
    wget  $URL  -O /tmp/$FILE
    sudo tar -xzvf /tmp/$FILE -C /usr/local
fi
echo "here"
echo -e "export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/$USER/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN" >>$HOME/.profile
echo "end"