extends Node

signal new_score(score)

var score = 0

func _ready():
	pass

func update_score(delta):
	score += delta
	emit_signal("new_score", score)

func _on_map_player_entered_activity( player, activity ):
	update_score(activity.points)
