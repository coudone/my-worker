<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>无标题文档</title>
<style>
*{margin:0;padding:0;list-style:none;}
_a{ text-decoration:none;}
#icon{border:solid 1px red;width:60px;height:30px;line-height:30px;text-align:center;position:absolute;top:20px;right:50px;}
#div1{width:300px;border:solid 1px #999;position:absolute;top:50%;left:50%;margin-left:-150px;margin-top:-230px;display:none;}
#div2{height:60px;background:blue;cursor:move;}
#div2 a{float:right;font-size:16px;color:red;line-height:60px;}
span{display:block;height:400px;font-size:24px;line-height:400px;text-align:center;background:#666}
</style>
<script>

	function getPos(obj){
		var l=0;
		var t=0;
		while(obj){
			l+=obj.offsetLeft;
			t+=obj.offsetTop;
			obj=obj.offsetParent;
		}
		return {left:l,top:t}
	};

window.onload=function(){
	var oDiv2=document.getElementById('icon');
	
	
	oDiv2.onmousedown=function(ev){
		var oEvent=ev||event;
		var x1=oEvent.clientX-getPos(oDiv2).left;
		var y1=oEvent.clientY-getPos(oDiv2).top;
		
		document.onmousemove=function(ev){
			var oEvent=ev||event;
			var l=oEvent.clientX-x1;
			var t=oEvent.clientY-y1;
			
			oDiv2.style.left=l+'px';
			oDiv2.style.top=t+'px';

		};
	
	};
	
};
</script>
</head>

<body>
<div id="icon">登录</div>

</body>
</html>
