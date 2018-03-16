extends StaticBody2D

func _ready():
	add_to_group("lights")
	turn_light_off()

func turn_light_on():
	$light.show()

func turn_light_off():
	$light.hide()
