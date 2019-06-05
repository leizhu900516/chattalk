window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = $("#log");
    var localuid = RndNum(11);

    document.getElementById("send-id").onclick = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        var datetime = formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");
        var message = {"userid":localuid,"destid":'1000',"content":msg.value,"addtime":datetime};
        //conn.send(msg.value);
        conn.send(JSON.stringify(message));

        appendMsg(log,generateSelfSendMsgHtml(msg.value));
        msg.value = "";
        return false;
    };
    //绑定回车事件
    $("#msg").keydown(function (event) {
        if(event.keyCode=='13'){
            if (!conn) {
                return false;
            }
            var selfmessage = msg.value;
            if (!selfmessage) {
                return false;
            }
            // alert("aaa");
            var datetime = formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");
            var message = {"userid":localuid,"destid":'1000',"content":selfmessage,"addtime":datetime,msgtype:1024};
            conn.send(JSON.stringify(message));
            appendMsg(log,generateSelfSendMsgHtml(selfmessage));
            msg.value = "";
            return false;
        }
    });

    if (window["WebSocket"]) {
        //test
        // conn = new WebSocket("ws://" + document.location.host + "/ws?userid="+localuid+"&destid=1000");
        //formal
        conn = new WebSocket("ws://wss.baidu.cn/ws?userid="+localuid+"&destid=1000");
        //第一次进入初始化,建立链接的
        if(conn.readyState==1){
            console.log("connection success")
        }else {
            console.log("socket正在建立连接")
        }
        conn.onopen = function(evt){
          //启用心跳 10/s
        };
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendMsg(log,item);
        };
        //接收到消息
        conn.onmessage = function (evt) {
            console.log(JSON.parse(evt.data));
            var messages = JSON.parse(evt.data);
            if (messages.destid==localuid){
                appendMsg(log,generateFriendSendMsgHtml(messages['content']));
            }
        };

    } else {
        console.log("不支持websocket");
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};
