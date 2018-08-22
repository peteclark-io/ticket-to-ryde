extends KinematicBody2D

export (String) var NAME
const MOVING = "MOVING"
const WAITING_FOR_INTERACTION = "WAITING_FOR_INTERACTION"
const IDLE = "IDLE"

var speed = 0

var velocity = Vector2()
var screensize = Vector2()
var move_to = null

var grid_size = 128.0
var grid_space = Vector2(0, 0)

var current_state = IDLE

func _ready():
	add_to_group("characters")
	set_physics_process(true)

func show_press_enter():
	$press_enter.show()

func hide_press_enter():
	$press_enter.hide()

func transition_state(to_state):
	if to_state == WAITING_FOR_INTERACTION:
		show_press_enter()
	elif current_state == WAITING_FOR_INTERACTION:
		hide_press_enter()

	current_state = to_state

func move_to(vector, s):
	print(NAME, " is moving to ", vector, s)
	move_to = vector
	speed = s
	transition_state(MOVING)
	print(move_to, speed)

func _physics_process(delta):
	if current_state == IDLE:
		return

	if current_state == MOVING:
		process_movement()

func process_movement():
	if move_to == null || move_to == position:
		return

	var velocity = Vector2()
	var s = speed

	var distance = (move_to - position).length()
	print(distance)
	if distance <= speed:
		print("decelerating", distance)
		speed = distance

	if position.x != move_to.x:
		if position.x > move_to.x:
			velocity.x = velocity.x - 1
		else:
			velocity.x = velocity.x + 1
	elif position.y != move_to.y:
		if position.y > move_to.y:
			velocity.y = velocity.y - 1
		else:
			velocity.y = velocity.y + 1

	if velocity.length() > 0:
		velocity = velocity.normalized() * s

	move_and_slide(velocity, Vector2(), 0.1, 4)
