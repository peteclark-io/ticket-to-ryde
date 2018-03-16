extends Area2D

signal increase_time

func _ready():
	pass

func _on_movement_area_body_exited( body ):
	emit_signal("increase_time")
