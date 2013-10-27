$(function() {
	// 左側の閉じるボタン
	$('#close_button').on('click', function() {
		$('#left').hide();
		$(this).hide();
		$('#open_button').show();
		$('#main').addClass('wide');
	});
	
	// 左側の開くボタン
	$('#open_button').on('click', function() {
		$('#left').show();
		$(this).hide();
		$('#close_button').show();
		$('#main').removeClass('wide');
	});
	
	// アカウントプロフィールの表示
	$.ajax('/get_account', {
		dataType: 'json',
		success: function(account) {
			$('#profile_body').html(account.description + '<br><br>' + account.time_zone);
			$('#user_logo').css('background-image', 'url("' + account.profile_image_url + '")');
			$('#user_name').html(account.screen_name);
		},
		error: function() {
			console.log('アカウントの取得に失敗しました');
		}
	});
	
	// タイムラインの表示
	$.ajax('/get_timeline', {
		dataType: 'json',
		success: function(timeLine) {
			var ul = $('#tweets');
			var template = _.template($('#tweet_template').html());
			_.each(timeLine, function(tweet) {
				var date = new Date(tweet.created_at);
				tweet.created_at = date.getFullYear() + '年' + (date.getMonth() + 1) + '月' + date.getDate() + '日 - ' + date.getHours() + ':' + date.getMinutes();
				
				tweet.user.profile_image_url = tweet.user.profile_image_url.replace(/_normal/, '_bigger');
				
				ul.append(template(tweet));
			});
		},
		error: function() {
			console.log('タイムラインの取得に失敗しました');
		}
	});
}());