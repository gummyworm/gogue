var Commands = (
		SELECT,
		DESELECT
);

//Command format 
var command = {
	id: 0,
	selection: []	//selected tiles (tuples of (x,y))
};

//SendCommand sends the current command 
SendCommand = function() {
	cmd = new Uint8Array(selection.length + 1);
	cmd[0] = id;
	for(var i = 0; i < selection.length; i++) {
		cmd[i+1] = selection[i];
	}
	ws.send(cmd)
}

//Select adds the provided selection to this player's selection
Select = function(sel) {
	update = new Uint8Array(sel.length + 1);
	update[0] = Commands.SELECT
	for(var i = 0; i < sel.length; i++) {
		update[i] = sel[i]
	}
}

//Unselect removes the provided positions from the player's selection
Unselect = function(sel) {
	for(var i = 0; i < sel.length; i += 2) {
		if(sel[i].x == 
		update[i] = sel[i]
	}
	update.selection = new Uint8Array(sel.length + 1);
	update[0] = Commands.DESELECT
}

//Take tells the selection units to take the item @ (x, y)
Take = function(x,y) {
}
