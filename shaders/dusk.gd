extends CanvasModulate

func _ready():
	hide()

func _on_dusk(falling_now):
	if falling_now:
		$transition.play("dusk")
	else:
		show()

func _on_nighttime(falling_now):
	hide()
