if (window.WebSocket === undefined) {
    alert("Your trash browser doesn't support WebSockets");
}
var ws = new WebSocket("ws://127.0.0.1:8080/ws");
ws.binaryType = "arraybuffer";
ws.onopen = function() {
    //socket is connected 
    alert("connected")
    SendCommand('l', 0, 0)
}
ws.onmessage = function(evt) {
    var buff = new Uint8Array(evt.data)
    for(i = 0; i < clientH; i++) {
        for(j = 0; j < clientW; j++) {
            tilemap[i][j] = buff[i*clientW + j];
        }
    }
    redraw()
}
ws.onclose = function(evt) {
    alert("disconnected")
}
