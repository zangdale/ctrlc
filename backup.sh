#!/bin/bash

DBPATH=$1
HOMEDIR=$2
DATE=`date '+%Y-%m-%d %H:%M:%S'`

echo "${DATE} >>>> ${DBPATH} ${HOMEDIR}" >> backup.log