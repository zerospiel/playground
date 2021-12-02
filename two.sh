#!/bin/bash
echo '' > /usr/local/bin/seo-metrics-reports.pid
main_pid=$!
echo $main_pid >> /usr/local/bin/seo-metrics-reports.pid
sleep 30s &
b_pid=$!
echo seo-metrics-api syncing with pid process $b_pid >> /usr/local/bin/seo-metrics-reports.pid
sleep 3s &
b_pid=$!
echo seo-metrics-reports syncing with pid process $b_pid >> /usr/local/bin/seo-metrics-reports.pid
wait $b_pid
echo seo-metrics-api synced with status $? >> /usr/local/bin/seo-metrics-reports.pid
wait $b_pid
echo seo-metrics-reports synced with status $? >> /usr/local/bin/seo-metrics-reports.pid