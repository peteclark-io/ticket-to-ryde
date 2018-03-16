extends NinePatchRect

export (String) var activity_name = "Example"
export (String) var activity_hours = "All Day"
export (int) var activity_duration = 2

export (bool) var has_food = false
export (bool) var has_drink = false
export (bool) var has_bed = false

var visited = false

func _ready():
	set_material(get_material().duplicate())
	$area/CollisionShape2D.position = rect_size / 2
	$area/CollisionShape2D.shape.extents = rect_size / 2
	render()
	$bounce.play("bounce")

func render():
	if visited:
		activity_hours = "Already Visited!"
		material.set_shader_param("grayscale", true)
		$VBoxContainer.remove_child($VBoxContainer/activity_duration)
	else:
		$VBoxContainer/activity_duration.text = "Lasts: " + display_duration(activity_duration)
	
	$VBoxContainer/activity_name.text = activity_name
	$VBoxContainer/activity_hours.text = activity_hours

	if !has_food:
		$VBoxContainer/activity_types/food.hide()
	else:
		$VBoxContainer/activity_types/food.show()
		
	if !has_drink:
		$VBoxContainer/activity_types/drink.hide()
	else:
		$VBoxContainer/activity_types/drink.show()
		
	if !has_bed:
		$VBoxContainer/activity_types/bed.hide()
	else:
		$VBoxContainer/activity_types/bed.show()
		
func display_duration(duration):
	if duration < 4:
		return str(duration * 15) + "mins"
	
	var remainder = duration % 4
	var hours = duration / 4
	
	var hours_str = "hr"
	if hours > 1:
		hours_str = "hrs"
	
	if remainder == 0:
		return str(hours) + hours_str
	
	return str(duration / 4) + hours_str + " " + str(remainder * 15) + "mins"

func set_activity(activity):
	activity_name = activity.activity_name
	activity_hours = activity.activity_hours
	activity_duration = activity.activity_duration
	has_food = activity.has_food
	has_drink = activity.has_drink
	has_bed = activity.has_bed
	render()
	
func visited():
	if visited:
		return false
		
	visited = true
	render()
	
	return true

func _on_area_body_entered( body ):
	if body.name == "player":
		$bounce.play_backwards("fade")
		$bounce.queue("bounce")

func _on_area_body_exited( body ):
	$bounce.play("fade")
	$bounce.queue("bounce")
