$(document).ready(function() {
	$('.content__top-iconsleft').click(function(){
		$('.content__middle-man').addClass(' content__middle-wooman')
		$('.content__middle-man').removeClass('content__middle-man');
		$('.content__class').text('');
		$('.content__class').text('ЛЕКАРЬ');
		localStorage.setItem("currentClass", "1")
	});
	$('.content__top-iconsright').click(function(){
		$('.content__middle-wooman').addClass('content__middle-man');
		$('.content__middle-wooman').removeClass(' content__middle-wooman')
		$('.content__class').text('');
		$('.content__class').text('ВОИН');
		localStorage.setItem("currentClass", "0")
	});

	$('.friend').hide();
	$('.friends').click(function(){
		$('.friend').toggle('slide');
	});

	$('.content__top-iconsright').click(function () {
        var audio = {};
        audio["walk"] = new Audio();
        audio["walk"].src = "http://www.rangde.org/static/bell-ring-01.mp3"
        audio["walk"].addEventListener('load', function () {
            audio["walk"].play();
        });
    });


	// setTimeout(function(){
	// $('.info__chat-send').trigger('click');
	// },1000);

    $('.info__chat-send').click(function(){
    	var textarea = $('.info__chat-windows').val();
    	console.log(textarea);

		var set = {
		  "url": "http://10.0.31.45:8080/api/chat/data",
		  "method": "POST",
		  "headers": {
		    "Content-Type": "application/json",
		    "Access-Control-Allow-Origin":"*"
		  },
		}

		var get = {
		  "url": "http://10.0.31.45:8080/api/chat/create",
		  "method": "POST",
		  "headers": {
		    "Content-Type": "application/json",
		    "Access-Control-Allow-Origin":"*"
		  },
		  "data": '{"Message": "'+textarea+'"}'
		}

		  $('.info__chat-sms').html('');
		$.ajax(set).done(function (response) {
		  var name = $('.content__bottom-name').val();
		  localStorage.setItem('cureentUser', name);
		  $('.info__chat-windows').val('');
		  //$('.info__chat-sms').append('<span style="color: lime;">'+name+'</span>: '+response.Chat+'<br><br>');
		  // $('.info__chat-windows').val(response.Chat);
		});

		var name = $('.content__bottom-name').val();
		var textarea = $('.info__chat-windows').val();
		$('.info__chat-sms').append('<span style="color: lime;">'+name+'</span>: '+textarea+'<br><br>');
	});

	// строка с параметрами для отправки
	// var body = "name=" + user.name + "&age="+user.age;
	// request.open("GET", "http://localhost:8080/postdata.php?"+body);
	// request.onreadystatechange = reqReadyStateChange;
	// request.send();

	$('.info__control-help').click(function(){
		$('.blackscreen').fadeIn();
	});
	$('.help__btn').click(function(){
		$('.blackscreen').fadeOut();
	});
});