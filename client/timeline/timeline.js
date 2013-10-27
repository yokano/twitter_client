$(function() {
	var initialize = function() {
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
		var account = JSON.parse('{"id":283857182,"id_str":"283857182","name":"y.okano","screen_name":"yuta_okano","location":"","description":"\u30d7\u30ed\u30d5\u30a3\u30fc\u30eb\u30c6\u30b9\u30c8","url":null,"entities":{"description":{"urls":[]}},"protected":false,"followers_count":1,"friends_count":1,"listed_count":0,"created_at":"Mon Apr 18 04:42:51 +0000 2011","favourites_count":0,"utc_offset":32400,"time_zone":"Sapporo","geo_enabled":false,"verified":false,"statuses_count":3,"lang":"ja","status":{"created_at":"Sun Oct 27 11:19:55 +0000 2013","id":394423209908375552,"id_str":"394423209908375552","text":"\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc","source":"web","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"geo":null,"coordinates":null,"place":null,"contributors":null,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[],"symbols":[],"urls":[],"user_mentions":[]},"favorited":false,"retweeted":false,"lang":"ja"},"contributors_enabled":false,"is_translator":false,"profile_background_color":"B2DFDA","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_link_color":"93A644","profile_sidebar_border_color":"EEEEEE","profile_sidebar_fill_color":"FFFFFF","profile_text_color":"333333","profile_use_background_image":true,"default_profile":false,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false}');
		$('#profile_body').html(account.description + '<br><br>' + account.time_zone);
		$('#user_logo').css('background-image', 'url("' + account.profile_image_url + '")');
		$('#user_name').html(account.screen_name);
		
		// タイムラインの表示
		var timeLine = JSON.parse('[{"created_at":"Sun Oct 27 12:30:59 +0000 2013","id":394441092411559936,"id_str":"394441092411559936","text":"\u4eca\u65e5\u306f\u3068\u3093\u304b\u3064","source":"web","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":283857182,"id_str":"283857182","name":"y.okano","screen_name":"yuta_okano","location":"","description":"\u30d7\u30ed\u30d5\u30a3\u30fc\u30eb\u30c6\u30b9\u30c8","url":null,"entities":{"description":{"urls":[]}},"protected":false,"followers_count":1,"friends_count":1,"listed_count":0,"created_at":"Mon Apr 18 04:42:51 +0000 2011","favourites_count":0,"utc_offset":32400,"time_zone":"Sapporo","geo_enabled":false,"verified":false,"statuses_count":4,"lang":"ja","contributors_enabled":false,"is_translator":false,"profile_background_color":"B2DFDA","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_link_color":"93A644","profile_sidebar_border_color":"EEEEEE","profile_sidebar_fill_color":"FFFFFF","profile_text_color":"333333","profile_use_background_image":true,"default_profile":false,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"geo":null,"coordinates":null,"place":null,"contributors":null,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[],"symbols":[],"urls":[],"user_mentions":[]},"favorited":false,"retweeted":false,"lang":"ja"},{"created_at":"Sun Oct 27 11:19:55 +0000 2013","id":394423209908375552,"id_str":"394423209908375552","text":"\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc\u30c4\u30a4\u30c3\u30bf\u30fc","source":"web","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":283857182,"id_str":"283857182","name":"y.okano","screen_name":"yuta_okano","location":"","description":"\u30d7\u30ed\u30d5\u30a3\u30fc\u30eb\u30c6\u30b9\u30c8","url":null,"entities":{"description":{"urls":[]}},"protected":false,"followers_count":1,"friends_count":1,"listed_count":0,"created_at":"Mon Apr 18 04:42:51 +0000 2011","favourites_count":0,"utc_offset":32400,"time_zone":"Sapporo","geo_enabled":false,"verified":false,"statuses_count":4,"lang":"ja","contributors_enabled":false,"is_translator":false,"profile_background_color":"B2DFDA","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_link_color":"93A644","profile_sidebar_border_color":"EEEEEE","profile_sidebar_fill_color":"FFFFFF","profile_text_color":"333333","profile_use_background_image":true,"default_profile":false,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"geo":null,"coordinates":null,"place":null,"contributors":null,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[],"symbols":[],"urls":[],"user_mentions":[]},"favorited":false,"retweeted":false,"lang":"ja"},{"created_at":"Sun Oct 27 11:19:29 +0000 2013","id":394423099996643328,"id_str":"394423099996643328","text":"Hello World!","source":"web","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":283857182,"id_str":"283857182","name":"y.okano","screen_name":"yuta_okano","location":"","description":"\u30d7\u30ed\u30d5\u30a3\u30fc\u30eb\u30c6\u30b9\u30c8","url":null,"entities":{"description":{"urls":[]}},"protected":false,"followers_count":1,"friends_count":1,"listed_count":0,"created_at":"Mon Apr 18 04:42:51 +0000 2011","favourites_count":0,"utc_offset":32400,"time_zone":"Sapporo","geo_enabled":false,"verified":false,"statuses_count":4,"lang":"ja","contributors_enabled":false,"is_translator":false,"profile_background_color":"B2DFDA","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_link_color":"93A644","profile_sidebar_border_color":"EEEEEE","profile_sidebar_fill_color":"FFFFFF","profile_text_color":"333333","profile_use_background_image":true,"default_profile":false,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"geo":null,"coordinates":null,"place":null,"contributors":null,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[],"symbols":[],"urls":[],"user_mentions":[]},"favorited":false,"retweeted":false,"lang":"en"},{"created_at":"Sun Oct 27 11:19:06 +0000 2013","id":394423004488155136,"id_str":"394423004488155136","text":"Twitter\u30af\u30e9\u30a4\u30a2\u30f3\u30c8\u30c6\u30b9\u30c8","source":"web","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":283857182,"id_str":"283857182","name":"y.okano","screen_name":"yuta_okano","location":"","description":"\u30d7\u30ed\u30d5\u30a3\u30fc\u30eb\u30c6\u30b9\u30c8","url":null,"entities":{"description":{"urls":[]}},"protected":false,"followers_count":1,"friends_count":1,"listed_count":0,"created_at":"Mon Apr 18 04:42:51 +0000 2011","favourites_count":0,"utc_offset":32400,"time_zone":"Sapporo","geo_enabled":false,"verified":false,"statuses_count":4,"lang":"ja","contributors_enabled":false,"is_translator":false,"profile_background_color":"B2DFDA","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme13\/bg.gif","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/3538254770\/b619e383339170ef0667aeb9bc615f6f_normal.png","profile_link_color":"93A644","profile_sidebar_border_color":"EEEEEE","profile_sidebar_fill_color":"FFFFFF","profile_text_color":"333333","profile_use_background_image":true,"default_profile":false,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"geo":null,"coordinates":null,"place":null,"contributors":null,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[],"symbols":[],"urls":[],"user_mentions":[]},"favorited":false,"retweeted":false,"lang":"ja"}]');
		var ul = $('#tweets');
		var template = _.template($('#tweet_template').html());
		_.each(timeLine, function(tweet) {
			var date = new Date(tweet.created_at);
			tweet.created_at = date.getFullYear() + '年' + (date.getMonth() + 1) + '月' + date.getDate() + '日 - ' + date.getHours() + ':' + date.getMinutes();
			
			tweet.user.profile_image_url = tweet.user.profile_image_url.replace(/_normal/, '_bigger');
			
			ul.append(template(tweet));
		});
		console.log(timeLine[0]);
		
	};
	
	var getAccount = function() {
	
	};
	
	var getTimeline = function() {
	
	};
	
	
	initialize();
});