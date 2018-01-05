# Notes

A game has players
A game has a score
A game has 10 days, consisting of 2 rounds
A game has a board

A player has a name
A player has perks

A round is either day or night
Each player has a turn during a round
A round is complete when all players have had their turn

A turn contains 5 action points
A turn is completed when action points have been spent or discarded

Action points can be spent on movement
AP can be spent on an activity
1 AP can be spent to pick up a card
1 AP can be spent to play a card

A board has districts
A board has ferry ports
A board has train stations

A district has activities

An activity has rewards
An activity can be during the day, the night or either
An activity can be visited once per round for rewards
An activity revisited yields rewards at 10% of the original value
An activity can contain a bus stop
An activity can be a special event

A reward adds happiness points (HP) to the players score
A reward can add an experience to the players score
A reward can either be labelled food, entertainment or an experience

A special event is either the Whale watching or Dragon watching activities
A special event costs 5AP
A special event gives a 1/10 chance of an excellent reward
A special event gives a 3/10 chance of an good reward
A special event gives a 6/10 chance of no reward
Each time a special event is used, the reward that is won is removed from the pool
A special event can only be revisited by a player who did not win a reward
Revisiting a special event does not decrease the value of the reward
A special event otherwise follows the same rules for revisiting regular activities

A card contains a perk or a card action

A perk is applied to one or many future turns for the player who used it
A perk can increase the number of APs on a future turn
A perk can reduce the AP cost of specific activities if visited in a future turn
A perk can reduce the AP cost of ferries, trains and the bus service

A card action can be applied to a selected player or players depending on criteria
A card action can be of benefit or of detriment to the targeted player
A card action can be applied to the player who triggered the action

A criteria can include players in specific locations, activities containing types of reward, or districts
A criteria applies to *all* players

A train station costs 4 AP to use
A train station moves the player to another train station of their choice

A ferry costs the players remaining AP for the turn
A ferry moves the player to another ferry port, or a specific (connected) activity

A bus stop costs 3 AP to use
A bus moves the player to the train station within the district

A score consists of the players HP total plus the number of experiences collected
The player with at least 3 experiences, and the highest total HP wins
If HP scores are tied, tiebreakers are the player with the most experiences, the player with the fewest revisited activities, then the player with the most activities. Otherwise a draw
