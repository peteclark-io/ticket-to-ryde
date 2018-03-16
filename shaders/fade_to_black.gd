extends AnimationPlayer

func _ready():
	pass

func _fade_out():
	play("fade")
	
func _fade_in():
	play_backwards("fade")

