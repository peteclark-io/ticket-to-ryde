extends KinematicBody2D

signal moved_one_grid_space(player)

export (int) var SPEED

var velocity = Vector2()
var screensize = Vector2()

var grid_size = 128.0
var grid_space = Vector2(0, 0) 

func _ready():
	add_to_group("player")
	set_physics_process(true)

func grid_space():
	grid_space.x = int(position.x / grid_size)
	grid_space.y = int(position.y / grid_size)
	return Vector2(grid_space.x, grid_space.y)

func show_press_enter():
	$press_enter.show()

func hide_press_enter():
	$press_enter.hide()

func _physics_process(delta):
	var velocity = Vector2()
	if Input.is_action_pressed("ui_right"):
		velocity.x += 1
	if Input.is_action_pressed("ui_left"):
		velocity.x -= 1
	if Input.is_action_pressed("ui_down"):
		velocity.y += 1
	if Input.is_action_pressed("ui_up"):
		velocity.y -= 1
	
	if velocity.length() > 0:
		velocity = velocity.normalized() * SPEED
	
	var before = grid_space()
	
	move_and_slide(velocity, Vector2(), 0.1, 4)
	
	position.x = clamp(position.x, 0, screensize.x)
	position.y = clamp(position.y, 0, screensize.y)
	
	var after = grid_space()
	if after != before:
		emit_signal("moved_one_grid_space", self)
	