extends CanvasLayer

var showing_activity_complete = false

func _ready():
	pass

func _process(delta):
	if showing_activity_complete && Input.is_action_just_pressed("ui_accept"):
		showing_activity_complete = false
		$activity_complete_container.hide()
		$fade_to_black._fade_in()

func _on_game_clock_time_increased(time):
	$display_time.text = time

func _on_player_new_score( score ):
	$display_player_score.text = String(score)+" PTS"

func _on_activity(player, activity):
	$fade_to_black._fade_out()
	
	$activity_complete_container/fade_in_timer.start()
	$activity_complete_container/activity_blurb.text = activity.get_activity_blurb(null)

func _on_fade_in_timer_timeout():
	$activity_complete_container.show()
	showing_activity_complete = true
