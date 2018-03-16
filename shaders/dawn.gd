extends CanvasModulate

func _ready():
	show()

func _on_dawn(breaking_now):
	if breaking_now:
		$transition.play("dawn")
	else:
		show()

func _on_daytime(breaking_now):
	if breaking_now:
		$transition.play("daytime")
	else:
		show()

