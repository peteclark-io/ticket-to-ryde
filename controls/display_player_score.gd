extends Label

signal update_score;

var score = 0;

func _ready():
	text = String(score) + " PTS"
	pass

func _on_update_score():
	score += 1;
	text = String(score) + " PTS"
	print(score)

func _on_activity(activity):
	pass