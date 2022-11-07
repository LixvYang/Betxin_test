#!/bin/sh
while [ true ]; do
/bin/sleep 1
cat /proc/1060967/status  | grep VmRSS >> ./cen.txt 
done
