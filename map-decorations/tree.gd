extends StaticBody2D

export (String) var TreeType

func _ready():
	for c in get_children():
		if c.get_class() == "Sprite":
			c.hide()
	
	if TreeType == "thin_tree_light":
		$thin_tree_light.show()
		
	if TreeType == "thin_tree_brown":
		$thin_tree_brown.show()
		
	if TreeType == "thin_tree_dark":
		$thin_tree_dark.show()
	
	if TreeType == "tree_light":
		$tree_light.show()
		
	if TreeType == "tree_brown":
		$tree_brown.show()
		
	if TreeType == "tree_dark":
		$tree_dark.show()
	
	if TreeType == "fat_tree_light":
		$fat_tree_light.show()
		
	if TreeType == "fat_tree_brown":
		$fat_tree_brown.show()
		
	if TreeType == "fat_tree_dark":
		$fat_tree_dark.show()
