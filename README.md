# gogue

Gogue is a roguelike engine that facilitates the creation and serving of traditional roguelike-type games on the web.  The client-server relationship in Gogue sees the server doing virtually all the work while the client acts as a dumb terminal.  The client communicates to the server with a set of commands (extendable to meet the needs of the particular game) an array of tiles that the client wishes this command to be performed upon.  For example, the client might send the 's' (select) command followed by the tiles [(0,0), (1,0), (0,1), (1,1), (50,6)].  Note that this does not imply such an action is valid.  The server will decide how to handle the request.  's' is a somewhat special command as it determines the operators of all future commands until the units are deselected (explicitly or as the result of, for example, the client denying the selection).

The way the server determines the validity of these requests is by a type assertion to an interface associated with the provided command.  If the user chooses to attack, all units in the client's selection will attempt to attack the target(s) provided by the client.

For each command, there should be one corresponding interface to dispatch the action it represents.


