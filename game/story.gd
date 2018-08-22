extends Node

var move = false

func _ready():
	pass

func _process(delta):
	if !move:
		get_tree().call_group("characters", "move_to", Vector2(43, 3674), 140)
		move = true
	pass
