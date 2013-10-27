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
});