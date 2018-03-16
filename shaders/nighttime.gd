extends CanvasModulate

func _ready():
	hide()

func _on_nighttime(falling_now):
	if falling_now:
		$transition.play("nighttime")
	else:
		show()

func _on_dawn(breaking_now):
	hide()
