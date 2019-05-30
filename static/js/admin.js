window.onload = function () {
    var conn;
    var msg = $("#msg");
    log = $("#messagelog");
    //所有的消息对象集合
    //格式{
    //     "客户id":
    //      {
    //          unread:12:
    //            message:[
    //                {"msg":"ms1","obj":"me","date":"2017-05-12 12:00"},
    //                {"msg":"ms1","obj":"he","date":"2017-05-12 12:00"},
     //               ]
    //      }，
    //      ...
    // }
    allMsgObj = {};
    localWindow = '';


    //绑定回车事件
    $("#msg").keydown(function (event) {
        if(event.keyCode=='13'){
            if (!conn) {
                return false;
            }
            var selfmessage = msg.text();
            if (!selfmessage) {
                return false;
            }
            var datetime = formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");
            var message = {"userid":'1000',"destid":localWindow,"content":selfmessage,"addtime":datetime};
            conn.send(JSON.stringify(message));
            appendMsg(log,generateSelfSendMsgHtml(selfmessage));
            updateAllMsgData(localWindow,selfmessage,"me",addtime=datetime);
            msg.empty();
            return false;
        }
    });

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws?userid=1000&destid=null");
        if(conn.readyState==1){
            console.log("socket成功建立连接")
        }else {
            console.log("socket正在建立连接")
        }
        //服务端关闭了连接
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            alert("server close");
        };
        conn.onmessage = function (evt) {
            console.log(JSON.parse(evt.data));
            var messages = JSON.parse(evt.data);
            recvMsgreloadChatPage(messages);

        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }


    //当有消息发送时，页面重载函数
    function recvMsgreloadChatPage(message) {
        //当前用户集合
        var userlist = [];
        $("#chat-user-list>div").each(function(){
            userlist.push($(this).attr("data-cid"));
        });
        //1、判断此条消息的用户是否在用户列表中
        var destid = message.userid;
        var msg = message.content;
        var recvMsgDate = message.addtime;
        if (allMsgObj[destid]){
            //1.1当前窗口是不是这条消息的用户
            if(localWindow == destid){
                appendMsg(log,generateFriendSendMsgHtml(msg));
            }else{
                //1.2不是当前用户
                //jquery动态更新气泡
                $("#chat-user-list>div").each(function(){
                    if($(this).attr("data-cid")==destid){
                        var unreadCount = allMsgObj[destid]['unread']+1;
                        $(this).children("div").children("span:last").html(unreadCount);
                    }
                });
                //todo
                //更新总数据结构
                allMsgObj[destid]['unread']= allMsgObj[destid]['unread']+1;
                updateAllMsgData(destid,msg,"he",addtime=recvMsgDate);
            }
        }else{
            //2、不在用户列表
            //2.1新增div，置顶，气泡
            var friendslitdom = $("#chat-user-list");
            prependFriendslist(friendslitdom,generateFriendListHtml(destid,1,"访"+destid));
            //2.2更新总数据
            allMsgObj[destid]={};
            allMsgObj[destid]['unread']=1;
            allMsgObj[destid]['message']=[];
            allMsgObj[destid]['message'].push({"msg":msg,"obj":"he","date":recvMsgDate},);
        }
        console.log(allMsgObj);
    }

};
//重载聊天面板
function chatpanelreload(messagedata) {
    $("#messagelog").empty();
    $.each(messagedata,function (i, j) {
        var msg = j.msg;
        if(j.obj == "me"){
            appendMsg(log,generateSelfSendMsgHtml(msg));
        }else if(j.obj == "he"){
            appendMsg(log,generateFriendSendMsgHtml(msg));
        }
    })
}


//点击用户列表的某个用户时，重载函数
function chatClick(e) {
    console.log(allMsgObj);
    //获取div的用户信息
    var userid = $(e).attr("data-cid");
    console.log(userid);
    //气泡为0
    if(allMsgObj[userid]){
        allMsgObj[userid]['unread'] = 0;
        $(e).children("div").children("span:last").html("");
        //去除兄弟元素的chc-chating-this效果
        $(e).siblings().removeClass("chc-chating-this");
        $(e).addClass("chc-chating-this");
        $("#local-username").html(userid);
        //定位当前聊天用户
        localWindow = userid;
        //重载聊天面板
        chatpanelreload(allMsgObj[userid]['message']);
    }else{
        alert("用户不存在");
    }
    console.log(allMsgObj);
}

function updateAllMsgData(uid,msg,who,addtime=null) {
    if(datetime){
        var datetime = addtime;
    }else{
        var datetime = formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");
    }

    allMsgObj[uid]['message'].push({"msg":msg,"obj":who,"date":datetime},);
}