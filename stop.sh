#!/bin/bash
killall fil-admin # kill fil-admin service
echo "stop fil-admin success"
ps -aux | grep go-admin