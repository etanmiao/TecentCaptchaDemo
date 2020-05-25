window.callback = function(res){
  console.log(res)
  // res（用户主动关闭验证码）= {ret: 2, ticket: null}
  // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
  if(res.ret === 0){
              //alert(res.ticket)
      this.console.log(res.ticket);
  }
  var xhr = new this.XMLHttpRequest();
  var param = 'ticket='.concat(res.ticket).concat("&randstr=").concat(res.randstr)
  xhr.open('GET',"/captcha/authresult?".concat(param))
  var captchaCode
  xhr.send(null)
  xhr.onload = function(){
    console.log(this.responseText)
    console.log(typeof(this.responseText))
    var obj = eval('('+this.responseText+')')
    captchaCode = obj.CaptchaCode
    var username = document.getElementById("username").value
    var password = document.getElementById("password").value
    if (captchaCode != 1){
       alert("验证码验证失败")
       return
    } else {
       alert("验证码验证成功，准备验证账号密码")	
}
    if (username == "") {
       alert("用户名非法")
       return;
    }
    if (password == "" ){
      alert("密码非法")
      return
    }
    if (username=="admin" && password=="123"){
      alert("验证码验证成功，登录成功")
    } else {
        alert("密码错误")
    }
  }
}
