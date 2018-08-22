extends YSort

signal player_moved_one_space(player)
signal player_entered_activity(player, activity)

export (Vector2) var screensize = Vector2(3840, 3840)

func _ready():
	$player.screensize = screensize

func _dawn(breaking_now):
	$nighttime._on_dawn(breaking_now)
	$dawn._on_dawn(breaking_now)
	get_tree().call_group("lights", "turn_light_off")

func _daytime(breaking_now):
	$dawn._on_daytime(breaking_now)

func _dusk(falling_now):
	$dusk._on_dusk(falling_now)

func _nighttime(falling_now):
	$dusk._on_nighttime(falling_now)
	$nighttime._on_nighttime(falling_now)
	get_tree().call_group("lights", "turn_light_on")

func _on_any_activity_entered( activity ):
	emit_signal("player_entered_activity", $player, activity)

func _on_player_moved_one_grid_space( player ):
	emit_signal("player_moved_one_space", player)

