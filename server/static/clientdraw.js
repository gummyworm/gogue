tilemap = CreateTilemap();

function CreateTilemap() {
  var tiles = [];
  for(var i=0; i<clientH; i++) {   //foreach row...
     tiles[i] = [];             //create a new array...
     for(var j=0; j<clientW; j++) {  //of 80 columns.
        tiles[i][j] = 0;
     }
  }
  return tiles;
}

// drawClientState renders the tilemap the server has provided it.
// This is a 80x25 map of tiles, which the client renders according to the
// spritesheet 'sprites'.
function drawClientState() {
    for(i = 0; i < clientH; i++) {
        for(j = 0; j < clientW; j++) {
            var x = (tilemap[i][j] % spritesPerRow) * spriteW;
            var y = Math.floor(tilemap[i][j] / spritesPerRow) * spriteH;
            context.drawImage(sprites,
                x, y, 
                clipW, clipH, 
                j*spriteW, i*spriteH, 
                spriteW, spriteH);
        }
    }
}

// redraw renders the client's screen.
function redraw() {
    drawClientState()
    drawSquare()
}
