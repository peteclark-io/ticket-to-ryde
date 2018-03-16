extends Node

signal time_increased
signal dawn(breaking_now)
signal daytime(breaking_now)
signal dusk(falling_now)
signal nighttime(falling_now)

var time
var day

func _ready():
	time = 35
	day = 1
	increase_time(1)

func get_time_display(time):
	var minute = time % 4
	var hour = str((time / 4) % 24).pad_zeros(2) 
	var day_display = "DAY " + str(day).pad_zeros(2)
	
	if minute == 0:
		return day_display + " " + hour + ":00"
	if minute == 1:
		return day_display + " " + hour + ":15"
	if minute == 2:
		return day_display + " " + hour + ":30"
	if minute == 3:
		return day_display + " " + hour + ":45"
	return "err"

func increase_time(delta):
	if (time + delta) >= 96:
		day += 1
	 
	time = (time+delta)%96
	emit_signal("time_increased", get_time_display(time))
	
	if time >= 28:
		if time == 28:
			emit_signal("dawn", true)
		else:
			emit_signal("dawn", false)
	
	if time >= 36:
		if time == 36:
			emit_signal("daytime", true)
		else: 
			emit_signal("daytime", false)
	
	if time >= 80:
		if time == 80:
			emit_signal("dusk", true)
		else: 
			emit_signal("dusk", false)
	
	if time >= 88 || time < 28:
		if time == 88:
			emit_signal("nighttime", true)
		else: 
			emit_signal("nighttime", false)

func _on_player_moved_one_space( player ):
	increase_time(1)

func _on_map_player_entered_activity( player, activity ):
	print("Increasing time for activity " + activity.activity_name)
	increase_time(activity.activity_duration)
