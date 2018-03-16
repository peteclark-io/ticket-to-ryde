extends Node

signal any_activity_entered(activity)

func _ready():
	for child in get_children():
		child.connect("activity_entered", self, "_on_any_activity_entered")

func _on_any_activity_entered(activity):
	emit_signal("any_activity_entered", activity)