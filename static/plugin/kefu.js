
window.onload=function(){
  // alert("aaaa");
  $('<div class="chat-icon" style="position: fixed;bottom: 10%;right: 3%;width: 100px;height: 40px;padding: 7px;border-radius: 20px;background: linear-gradient(to left,#08d9d6, #10eaf0);color: #fff!important;cursor: pointer">\n' +
      '    <div>\n' +
      '        <svg class="icon" aria-hidden="true" style="color: #fff!important;">\n' +
      '            <use xlink:href="#icon-kefu"></use>\n' +
      '        </svg>\n' +
      '        <cite style="font-style: normal" id="online-kefu">在线咨询</cite>\n' +
      '    </div>\n' +
      '</div>').appendTo('body');


      var online_kefu = document.getElementById('online-kefu');
      online_kefu.onclick = function () {
        //生成聊天窗口-iframe
        alert('点击了click');
      }

};
