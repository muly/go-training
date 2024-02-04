#!/bin/sh
set -x
set -e


sleep 1 && open http://127.0.0.1:3999 &
present 
