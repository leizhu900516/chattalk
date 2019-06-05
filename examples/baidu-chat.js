var utils  = require('../../utils/util.js');
Page({
    data: {
        sid:1000,//客服id
        msg:[
            {"msg":"您好，需要咨询什么？","who":"he"},
            //{"msg":"这是聊天内容","who":"self"},
        ],
        inputvalue:"",
        uid:'',
    },
    onLoad: function () {
        var uid = utils.RndNum(10);
        var that = this;
        this.setData("uid",uid);
        console.log(uid);
        // 监听页面加载的生命周期函数
        swan.connectSocket({
            url: 'wss://wss.bao361.cn/ws?userid='+uid+'&destid=1000',
            header:{},
            protocols:[''],
            success: function (res) {
                console.log('connectSocket success', res);
            },
            fail: function (err) {
                console.log('connectSocket fail', err);
            }
        });
        swan.onSocketOpen(function (res) {
            console.log('WebSocket连接已打开！', res);
            // 发送信息样例
            // swan.sendSocketMessage({
            //     data: 'baidu',
            //     success: function (res) {
            //         console.log('WebSocket发送数据成功', res);
            //     },
            //     fail: function (err) {
            //         console.log('WebSocket发送数据失败', err);
            //     }
            // });
        });
        swan.onSocketError(function (res) {
            console.log('WebSocket连接打开失败，请检查！', res);
        });
        swan.onSocketMessage(function (res) {
            console.log('收到服务器内容：', res.data);
            
            var datas = JSON.parse(res.data)
            var msg = datas.content;
            console.log('收到服务器内容：', msg);
            that.data.msg.push({"msg":msg,"who":"he"})
            that.setData('msg',that.data.msg)
        });

    },
    onReady: function() {
        // 监听页面初次渲染完成的生命周期函数
    },
    onShow: function() {
        // 监听页面显示的生命周期函数
    },
    onHide: function() {
        // 监听页面隐藏的生命周期函数
    },
    onUnload: function() {
        // 监听页面卸载的生命周期函数
    },
    onPullDownRefresh: function() {
        // 监听用户下拉动作
    },
    onReachBottom: function() {
        // 页面上拉触底事件的处理函数
    },
    onShareAppMessage: function () {
        // 用户点击右上角转发
    },
    formSubmit: function(e) {
        console.log('form发生了submit事件，携带数据为：', e.detail.value);
        var that = this;
        var msg = e.detail.value.msg;
        var datetime = utils.formatDate((new Date()),"yyyy-MM-dd hh:mm:ss");
        //发送信息
        swan.sendSocketMessage({
            data: JSON.stringify({"userid":that.data.uid,"destid":'1000',"content":msg,"addtime":datetime,msgtype:1024}),
            success: function (res) {
                console.log('WebSocket发送数据成功', res);
            },
            fail: function (err) {
                console.log('WebSocket发送数据失败', err);
            }
        });
        that.data.msg.push({"msg":msg,"who":"self"})
        that.setData('msg',that.data.msg)
        that.setData('inputvalue',' ');
        that.setData('inputvalue','');
    },

});