#!/bin/bash

REPO_DIR=/home/bilkoll/bilkoll
DEPLOY=$REPO_DIR/deployment/deploy.sh

echo "$(date): Autodeploy started"
echo "Running as user $(whoami)"

cd $REPO_DIR || exit

/usr/bin/git fetch origin

LOCAL=$(/usr/bin/git rev-parse HEAD)
REMOTE=$(/usr/bin/git rev-parse origin/master)

echo "Local: $LOCAL"
echo "Remote: $REMOTE"

if [ "$LOCAL" != "$REMOTE" ]; then
    echo "$(date) Out of date, updating."
    /usr/bin/git pull
    $DEPLOY
    echo "$(date) Autodeploy done."
else
    echo "$(date) No changes detected."
fi
