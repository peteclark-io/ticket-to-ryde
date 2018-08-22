extends KinematicBody2D

signal moved_one_grid_space(player)

export (int) var SPEED

var velocity = Vector2()
var direction = Vector2()
var previous_direction = Vector2()

var screensize = Vector2()

var grid_size = 128.0
var grid_space = Vector2(0, 0) 

onready var sprite = $Sprite/anim

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
	previous_direction = direction
	velocity = Vector2()
	direction = Vector2()
	
	if Input.is_action_pressed("ui_right"):
		direction.x += 1
	if Input.is_action_pressed("ui_left"):
		direction.x -= 1
	if Input.is_action_pressed("ui_down"):
		direction.y += 1
	if Input.is_action_pressed("ui_up"):
		direction.y -= 1
	
	if direction.length() > 0:
		velocity = direction.normalized() * SPEED

	var before = grid_space()
	
	move_and_slide(velocity, Vector2(), 0.1, 4)
	
	position.x = clamp(position.x, 0, screensize.x)
	position.y = clamp(position.y, 0, screensize.y)
	
	animate(previous_direction, direction)
	
	var after = grid_space()
	if after != before:
		emit_signal("moved_one_grid_space", self)

func animate(prev, direction):
	if direction.length() == 0:
		$Sprite/anim.stop()
		$Sprite.frame = 0
		return
		
	var prev_animation = animation_for_vector(prev)
	var cur_animation = animation_for_vector(direction)
	
	if sprite.is_playing() && prev_animation == cur_animation:
		return
	
	sprite.play(cur_animation, 0.1)

func animation_for_vector(v):
	if v.y == 1:
		return "walk-forward"
	
	if v.y == -1:
		return "walk-backward"
	
	if v.x == -1:
		return "walk-left"
	
	if v.x == 1:
		return "walk-right"
	
	return "none"