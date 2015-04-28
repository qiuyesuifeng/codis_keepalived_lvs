#!/bin/bash
LOG_FILE="/var/log/codis_keepalived_lvs.log"
DATE=`date`
echo "[fault]${DATE}" >> ${LOG_FILE}