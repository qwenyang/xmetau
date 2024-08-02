
rm -rf /home/qwenyang/bin/unigateway

cp ./unigateway /home/qwenyang/bin

killall unigateway

nohup /home/qwenyang/bin/unigateway > /data/game/log/unigateway.log 2>&1 &
