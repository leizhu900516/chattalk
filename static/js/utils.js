//自己发送的信息插入聊天框函数
function generateSelfSendMsgHtml(message) {
    return `<div class="msg-item d-flex flex-row flex-row-reverse">
                                <img src="/static/images/online.png" class="user-photo rounded d-block ml-1">
                                <div >
                                    <p class="chat-msg rounded p-1">`+message+`</p>
                                </div>
                            </div>`
}
//好友发送的信息插入聊天框函数
function generateFriendSendMsgHtml(message) {
    return `<div class="msg-item d-flex flex-row">
            <img src="https://upyun.bao361.cn/image_upload_api/945886ff2cf76308e237186fb57ad9530.983230097368" class="user-photo rounded d-block mr-1">
            <div >
                <p class="chat-msg-dest rounded p-1">`+message+`</p>
            </div>
        </div>`
}
//生成客服左侧的聊天用户列表
function  generateFriendListHtml(userid,unreadMsgCount,username) {
    return `<div class="chc-users d-flex flex-row align-items-center " data-cid="`+userid+`" onclick="chatClick(this)"  data-online="yes" >
                                <img src="/static/images/touxiang.jpeg" class="mr-1">
                                <div class="consumer-info mr-1">
                                    <div class="name-unread-status">
                                        <span class="">`+username+`</span>
                                        <span class="badge badge-danger">`+unreadMsgCount+`</span>
                                        <span class="badge badge-success u-online" >在线</span>
                                    </div>
                                    <div class="from-city">
                                        <span>PC</span>
                                        <span>河南 郑州</span>
                                    </div>
                                </div>
                            </div>`
}

function appendMsg(dom,item){
    var doScroll = dom.scrollTop > dom.scrollHeight - dom.clientHeight - 1;
    dom.append(item);
    if (doScroll) {
        dom.scrollTop = dom.scrollHeight - dom.clientHeight;
    }
}
function prependFriendslist(dom,item){
    var doScroll = dom.scrollTop > dom.scrollHeight - dom.clientHeight - 1;
    dom.prepend(item);
    if (doScroll) {
        dom.scrollTop = dom.scrollHeight - dom.clientHeight;
    }
}
//自动生成n为数字组合
function RndNum(n){
    var rnd="";
    for(var i=0;i<n;i++)
        rnd+=Math.floor(Math.random()*10);
    return rnd;
}

//时间格式化工具
function formatDate(date,format){
    var paddNum = function(num){
        num += "";
        return num.replace(/^(\d)$/,"0$1");
    };
    //指定格式字符
    var cfg = {
        yyyy : date.getFullYear() //年 : 4位
        ,yy : date.getFullYear().toString().substring(2)//年 : 2位
        ,M  : date.getMonth() + 1  //月 : 如果1位的时候不补0
        ,MM : paddNum(date.getMonth() + 1) //月 : 如果1位的时候补0
        ,d  : date.getDate()   //日 : 如果1位的时候不补0
        ,dd : paddNum(date.getDate())//日 : 如果1位的时候补0
        ,hh : date.getHours()  //时
        ,mm : date.getMinutes() //分
        ,ss : date.getSeconds() //秒
    };
    format || (format = "yyyy-MM-dd hh:mm:ss");
    return format.replace(/([a-z])(\1)*/ig,function(m){return cfg[m];});
}
//test
formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");


//设置cookies
function setCookie(name,value)
{
    var Days = 30;
    var exp = new Date();
    exp.setTime(exp.getTime() + Days*24*60*60*1000);
    document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
}
function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
    if(arr=document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

function delCookie(name)
{
    var exp = new Date();
    exp.setTime(exp.getTime() - 1);
    var cval=getCookie(name);
    if(cval!=null)
        document.cookie= name + "="+cval+";expires="+exp.toGMTString();
}