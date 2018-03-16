extends Sprite

var activated

func _ready():
	$Area2D.connect("body_entered", self, "_on_body_entered")
	$Area2D.connect("body_exited", self, "_on_body_exited")

func _process(delta):
	if Input.is_action_just_pressed("ui_accept") && self.activated:
		print("Triggering " + self.get_parent().activity_name)
		get_parent()._on_marker_activated()

func get_player():
	return get_tree().get_nodes_in_group("player").front()

func _on_body_entered( body ):
	var player = get_player()
	if player != body:
		print(body.name)
		return
	
	if player != null:
		player.show_press_enter()
	self.activated = true

func _on_body_exited( body ):
	var player = get_player()
	if player != body:
		print(body.name)

	if player != null:
		player.hide_press_enter()
	self.activated = false