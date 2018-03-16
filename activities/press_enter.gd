extends Label

func _ready():
	pass

func _on_marker_stepped_on():
	show()

func _on_marker_stepped_off():
	hide()
