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
		dataType: 'text',
		success: function(account) {
			account = account.replace(/"source":[^,]*,/g, '');
			account = JSON.parse(account);
			$('#profile_body').html(account.description + '<br><br>' + account.time_zone);
			$('#user_logo').css('background-image', 'url("' + account.profile_image_url + '")');
			$('#user_name').html(account.screen_name);
			
			// タイムラインの表示
			$.ajax('get_timeline', {
				data: {
					screen_name: account.screen_name
				},
				dataType: 'text',
				success: function(timeline) {
					timeline = timeline.replace(/"source":[^,]*,/g, '');  // source に "" がネストされているため
					timeline = timeline.replace(/[\r\n]/g, '');  // 改行を削除
					timeline = JSON.parse(timeline);
					var ul = $('#tweets');
					var template = _.template($('#tweet_template').html());
					_.each(timeline, function(tweet) {
						var date = new Date(tweet.created_at);
						tweet.created_at = date.getFullYear() + '年' + (date.getMonth() + 1) + '月' + date.getDate() + '日 - ' + date.getHours() + ':' + date.getMinutes();
						
						tweet.user.profile_image_url = tweet.user.profile_image_url.replace(/_normal/, '_bigger');
						
						ul.append(template(tweet));
					});
				},
				error: function() {
					alert('タイムラインの取得に失敗しました');
				}
			})
		},
		error: function() {
			alert('アカウントの取得に失敗しました');
		}
	});
}());
