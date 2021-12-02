#!/bin/bash
echo '' > /usr/local/bin/bx.pid
main_pid=$!
echo $main_pid >> /usr/local/bin/bx.pid
sleep 15s &
business_static_pid=$!
echo business-static syncing with pid process $business_static_pid >> /usr/local/bin/bx.pid
sleep 10s &
seo_admin_report_pid=$!
echo seo-admin-report syncing with pid process $seo_admin_report_pid >> /usr/local/bin/bx.pid
sleep 5s &
slack_pid=$!
echo slack syncing with pid process $slack_pid >> /usr/local/bin/bx.pid
sleep 3s &
static_pid=$!
echo static syncing with pid process $static_pid >> /usr/local/bin/bx.pid
wait $business_static_pid
echo business-static synced with status $? >> /usr/local/bin/bx.pid
wait $seo_admin_report_pid
echo seo-admin-report synced with status $? >> /usr/local/bin/bx.pid
wait $slack_pid
echo slack synced with status $? >> /usr/local/bin/bx.pid
wait $static_pid
echo static synced with status $? >> /usr/local/bin/bx.pid