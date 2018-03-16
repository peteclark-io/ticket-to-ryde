extends Node

func _ready():
	for c in get_children():
		get_parent().add_child(c)