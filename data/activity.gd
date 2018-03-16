extends Node

signal activity_entered(activity)

export (String) var activity_hours
export (String) var activity_name
export (String) var activity_blurb

export (int) var points
export (int) var activity_duration

export (bool) var has_food
export (bool) var has_bed
export (bool) var has_drink

func _ready():
	$activity_panel.set_activity(self)

func _on_marker_activated():
	var changed = $activity_panel.visited()
	if !changed:
		return
	
	emit_signal("activity_entered", self)

func get_activity_blurb(time):
	return activity_blurb