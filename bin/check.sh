unidaonum=`ps -ef | grep unidao | grep -v grep | wc -l`

if [ "$unidaonum" -lt 1 ];then
    killall unidao
    nohup /www/wwwroot/xmetau/bin/unidao > /www/wwwroot/xmetau/log/unidao.log 2>&1 &
    echo "restart unidao"
fi

unigatewaynum=`ps -ef | grep unigateway | grep -v grep | wc -l`

if [ "$unigatewaynum" -lt 1 ];then
    killall unigateway
    nohup /www/wwwroot/xmetau/bin/unigateway > /www/wwwroot/xmetau/log/unigateway.log 2>&1 &
    echo "restart unigateway"
fi