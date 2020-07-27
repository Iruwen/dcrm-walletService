#!/bin/bash

set -e

BASEDIR=/dcrm-walletservice
CONFDIR=$BASEDIR/conf
DATADIR=$BASEDIR/data
LOGDIR=$BASEDIR/log

mkdir -p -m 750 $CONFDIR $DATADIR $LOGDIR
chmod -R o-rwx $CONFDIR $DATADIR $LOGDIR

touch $LOGDIR/dcrm-walletservice.log
chmod 640 $LOGDIR/dcrm-walletservice.log

[ -e $CONFDIR/dcrm.privatekey ] || gdcrm --genkey $CONFDIR/dcrm.privatekey
chmod 640 $CONFDIR/dcrm.privatekey

exec gdcrm --nodekey $CONFDIR/dcrm.privatekey --datadir $DATADIR --log $LOGDIR/dcrm-walletservice.log
