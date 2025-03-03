#!/bin/bash

set -o nounset

APP_NAME=bilkoll
REPO_DIR=/home/bilkoll/$APP_NAME
DEPL_DIR=$REPO_DIR-depl

echo "Starting deploy"
echo "APP_NAME $APP_NAME"
echo "REPO_DIR $REPO_DIR"
echo "DEPL_DIR $DEPL_DIR"

rm -rf $DEPL_DIR
mkdir $DEPL_DIR
mkdir $DEPL_DIR/rss

cp -r $REPO_DIR/src/assets $DEPL_DIR/assets
cp $REPO_DIR/src/rss/template.xml $DEPL_DIR/rss/.

cd $REPO_DIR/src

GOOS=linux 
GOARCH=amd64 

/usr/local/go/bin/go build -o $DEPL_DIR/$APP_NAME

sudo /usr/bin/systemctl daemon-reload
sudo /usr/bin/systemctl restart $APP_NAME.service
sudo /usr/bin/systemctl enable $APP_NAME.service
sudo /usr/bin/systemctl status $APP_NAME.service
