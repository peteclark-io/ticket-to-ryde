extends MarginContainer

var conversation = Array()
onready var text = $PanelContainer/VSplitContainer/HSplitContainer/MarginContainer/convo_text
var index = 0

func _ready():
	pass

func startConversation(txt):
	if txt.empty():
		print("Why are you trying to start a conversation with no text?!")
		return
		
	conversation = txt
	index = 0
	
	text.text = conversation.pop_front()
	show()

func _process(delta):
	if !Input.is_action_just_pressed("ui_accept"):
		return
	
	if conversation.empty():
		hide()
	else:
		text.text = conversation.pop_front()

