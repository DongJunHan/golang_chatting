$(function(){
	if(!window.EventSource){
		alert("No EventSource!");
		return;
	}

	var $chatlog = $('#chat-log');
	var $chatmsg = $('#chat-msg');
	
	var isBlank = function(string){
		//===는 값뿐만 아니라 타입까지 검사함.
		return string == null || string.trim() === "";
	};

	var username;
	while(isBlank(username)){
		username = prompt("What's your name?");
		console.log(username);
		if(!isBlank(username)){
			$('#user-name').html('<b>' + username + '</b>');
		}
	}
	$('#input-form').on('submit', function(e){
		$.post("/messages", {
			msg : $chatmsg.val(),
			name : username
		});

		$chatmsg.val("");
		$chatmsg.focus();
	});
});
