var startX, endX, startY, endY; //the start/end points of the selection rect
var mouseIsDown = false;        //mouse state flag

function mouseUp(eve) {
    if (!mouseIsDown) {
        var pos = getMousePos(canvas, eve);
        endX = pos.x;
        endY = pos.y;
    }
    mouseIsDown = false;
    //select what was selected
    // creating a square
    var w = endX - startX;
    var h = endY - startY;
    var offsetX = (w < 0) ? w : 0;
    var offsetY = (h < 0) ? h : 0;
    var width = Math.abs(w);
    var height = Math.abs(h);

    //convert from pixel bounds to tile bounds
    offsetX = Math.floor(offsetX/spriteW)
    offsetY = Math.floor(offsetX/spriteH)
    width = Math.floor(width/spriteW)
    height = Math.floor(height/spriteH)

    // TODO: send the message of the selected tiles
    redraw();
}

function mouseDown(eve) {
    mouseIsDown = true;
    var pos = getMousePos(canvas, eve);
    startX = endX = pos.x;
    startY = endY = pos.y;

    redraw();
    drawSquare(); //update
}

function mouseXY(eve) {
    if (mouseIsDown) {
        var pos = getMousePos(canvas, eve);
        endX = pos.x;
        endY = pos.y;
        redraw();
        drawSquare(); //update
    }
}

function drawSquare() {
    if(!mouseIsDown) {
        return;
    }
    // creating a square
    var w = endX - startX;
    var h = endY - startY;
    var offsetX = (w < 0) ? w : 0;
    var offsetY = (h < 0) ? h : 0;
    var width = Math.abs(w);
    var height = Math.abs(h);

    context.beginPath();
    context.rect(startX + offsetX, startY + offsetY, width, height);
    context.lineWidth = 3;
    context.strokeStyle = 'red';
    context.stroke();
}

function getMousePos(canvas, evt) {
    var rect = canvas.getBoundingClientRect();
    return {
        x: evt.clientX - rect.left,
        y: evt.clientY - rect.top
    };
}
canvas.addEventListener("mousedown", mouseDown, false);
canvas.addEventListener("mousemove", mouseXY, false);
canvas.addEventListener("mouseup", mouseUp, false);
