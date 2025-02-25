#!/bin/bash

set -o nounset

APP_NAME=bilkoll
HOME_DIR=/home/$APP_NAME
REPO_DIR=$HOME_DIR/$APP_NAME
DEPL_DIR="$HOME_DIR/$APP_NAME-dist"
OVERRIDE_CONF=/etc/systemd/system/$APP_NAME.service.d/override.conf

echo "Repo dir ${REPO_DIR}"
echo "Deployment dir ${DEPL_DIR}"
echo "Service file: $REPO_DIR/deployment/$APP_NAME.service"
echo "sudo ln -s $REPO_DIR/deployment/$APP_NAME.service /etc/systemd/system/$APP_NAME.service"

sudo ln -s "$REPO_DIR/deployment/$APP_NAME.service" "/etc/systemd/system/$APP_NAME.service"
        
sudo touch $OVERRIDE_CONF

if [ -d "/etc/systemd/system/$APP_NAME.service.d" ]; then
    echo 'Env dir already created'
else
    sudo mkdir /etc/systemd/system/$APP_NAME.service.d
    echo "[Service]" | sudo tee -a $OVERRIDE_CONF
    echo "Created file '$OVERRIDE_CONF'. Set the environment variables in this file."
fi

