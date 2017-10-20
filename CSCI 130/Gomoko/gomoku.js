var turn;
var matrix;
var victor;
var time;
var win = false;

function startGame() {
    var bCol = document.getElementById("boardcolor").value;
    var p1Col = document.getElementById("player1").value;
    var p2Col = document.getElementById("player2").value;
    (document.getElementById("15").checked ? gameArea.start(15, bCol, p1Col, p2Col): gameArea.start(19, bCol, p1Col, p2Col));
}

var gameArea = {
    canvas : document.createElement("canvas"),
    start : function(size, bc, p1, p2) {
        this.canvas.width = 600;
        this.canvas.height = 600;
        this.context = this.canvas.getContext("2d");
        this.playCount = 0; // Counts pieces played
        this.size = size; // Size of board i.e. number of intersections
        this.p1Col = p1; // Color of player 1
        this.p2Col = p2; // Color of player 2
        this.boardCol = bc; // Board color
        turn = p1; // Sets current player to player 1
        time = 0; 

        console.log("player 1:" + this.p1Col);
        console.log("player 2:" + this.p2Col);
        // Starts timer.
        this.timer = window.setInterval(timer, 1000);
        matrix = [];

        // Zero the starting matrix
        for(var i=0; i< size; i++) {
            matrix[i] = [];
            for(var j = 0; j < size; j++) {
                matrix[i][j] = 0;
            }
        }

        // Draw the board
        this.context.fillStyle = this.boardCol;
        this.context.fillRect(0, 0, 600, 600);

        this.x = ((Number(this.canvas.width)-50) / size);
        this.y = ((Number(this.canvas.height)-50) / size);
    
        this.context.strokeStyle = invertColor(this.boardCol);
        // Draws the verticle lines
        for(var i = 0; i < size; i++) {
            this.context.moveTo(this.x * (i + 1), 0);
            this.context.lineTo(this. x * (i + 1), Number(this.canvas.height) - 50 + this.y);
            this.context.stroke();
        }
    
        // Draws the verticle lines
        for(var j = 0; j < size; j++)
        {
            this.context.moveTo(0, this.y * (j + 1));
            this.context.lineTo(Number(this.canvas.width) - 50 + this.x, this.y*(j + 1));
            this.context.stroke();
        }
        document.getElementById("board").appendChild(this.canvas);

        //Handler for clicking on the grid
        this.canvas.addEventListener('click', function (e) {
            var pos = getMousePos(gameArea.canvas, e);
            if (!win && findClosest(pos.x, pos.y)) {
                gameArea.playCount++;
                console.log("Checking for win");
                if (fiveInARow()) {
                    console.log(turn + " is the winner");
                    gameArea.clear();
                    var mygrad = gameArea.context.createLinearGradient(0,0,170,0);
                    mygrad.addColorStop(0, "magenta");
                    mygrad.addColorStop(0.5, "blue");
                    mygrad.addColorStop(1, "red");
                    gameArea.context.strokeStyle = mygrad;
                    gameArea.context.font = "50px Verdana";
                    gameArea.context.strokeText((turn == gameArea.p1Col?"Player 1 wins":"Player 2 wins"),50,50);
                    gameArea.context.font = "20px Verdana";
                    gameArea.context.strokeText(fmtTime(time),80,90);
                    gameArea.context.strokeText((turn == gameArea.p1Col?Math.ceil(gameArea.playCount/2):Math.floor(gameArea.playCount/2)) + " rounds played", 80, 120);
                    gameArea.canvas.removeEventListener('click', function(e){});
                    win = true;
                    clearInterval(gameArea.timer);
                }

                turn = (turn == gameArea.p1Col ? gameArea.p2Col:gameArea.p1Col);
                plays = document.getElementById("plays");
                plays.innerHTML = "Pieces played<br>Player 1- " + Math.ceil(gameArea.playCount/2) + "  Player 2- "+Math.floor(gameArea.playCount/2);
            }  
        })
    }, 
    clear : function(){
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }
}

function getMousePos(canvas, evt) {
    var rect = canvas.getBoundingClientRect();
    return {
      x: evt.clientX - rect.left,
      y: evt.clientY - rect.top
    };
}

// Finds the closest board interesection from a coordinate and places the appropriate piece there
function findClosest(x, y) {
    var closestDist = gameArea.canvas.width + gameArea.canvas.height;
    var closestx;
    var closesty;
    var matInd;
    var xInd;
    var yInd;

    for(var i = 0; i < gameArea.size; i++) {
        for(j = 0; j < gameArea.size; j++) {
            var temp = distFrom(x,y,gameArea.x*(i+1),gameArea.y*(j+1));
            if(temp < closestDist) {
                closestDist = temp;
                closestx = gameArea.x*(i+1);
                closesty = gameArea.y*(j+1);
                xInd = i;
                yInd = j;
            }

        }
    }
    if(matrix[xInd][yInd] == 0) {
        drawPiece(12,turn, closestx, closesty);
        matrix[xInd][yInd] = (turn == gameArea.p1Col ? 1 : -1);
        console.log(matrix);
        return true;
    }
    return false;
}

// Calls helper functions to test if there are 5 pieces in a row of the same color
function fiveInARow() {
   return (testRow() || testCol() || testTLBRDiag() || testTRBLDiag());
}

// Tests each row for 5 in a row
function testRow() {
    var color;
    var inARow;

    for(var j = 0; j < gameArea.size; j++) {
        color = 0;
        inARow = 0;
        for(var i = 0; i < gameArea.size; i++)
            if(matrix[i][j]==color && matrix[i][j] != 0) {
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a row by " + color);
                    console.log(i+":"+j);
                    console.log(matrix);
                    return true;
                }
            }
            else {
                inARow = 1;
                color = matrix[i][j];
            }
    }

    return false;
}

// Tests each column
function testCol() {
    var color;
    var inARow;

    for(var i = 0; i < gameArea.size; i++) {
        color = 0;
        inARow = 0;
        for(var j = 0; j < gameArea.size; j++) {
            if(matrix[i][j]==color && matrix[i][j] != 0) {
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a column by " + color);
                    return true;
                }
            }
            else {
                inARow = 1;
                color = matrix[i][j];
            }
        }
    }

    return false;
}

// Tests each top right to bottom left diagonal
function testTRBLDiag() {
    var color;
    var inARow;

    for(var i = 4; i < gameArea.size; i++) {
        color = 0;
        inARow = 0;
        var tempI = i;
        for(var j = 0; j < gameArea.size; j++)
        {
            inARow = 0;
            color = matrix[i][j];
            var tempI = i;
            var tempJ = j;
            while(tempJ < gameArea.size && matrix[tempI][tempJ]==color && matrix[tempI][tempJ] != 0) {
                //console.log("left to right check: " + tempI +":"+tempJ);
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a right to left diagonal by " + color);
                    console.log(i +":"+j + " to " + tempI+":"+tempJ);
                    return true;
                }
                tempI--;
                tempJ++;
            }

        }
    }

    return false;
}

// Tests each top left to bottom right diagonal
function testTLBRDiag() {
    var color;
    var inARow;

    for(var i = 0; i < gameArea.size - 4; i++) {
        for(var j = 0; j < gameArea.size ; j++)
        {
            inARow = 0;
            color = matrix[i][j];
            var tempI = i;
            var tempJ = j;
            while(tempJ < gameArea.size && matrix[tempI][tempJ]==color && matrix[tempI][tempJ] != 0) {
                //console.log("left to right check: " + tempI +":"+tempJ);
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a right to left diagonal by " + color);
                    console.log(i +":"+j + " to " + tempI+":"+tempJ);
                    return true;
                }
                tempI++;
                tempJ++;
            }
            
        }
    }

    return false;
}

// Calculate the distance two points are from each other.
function distFrom(x1, y1, x2, y2    ) {
    var a = Math.abs(x1 - x2);
    var b = Math.abs(y1 - y2);
    return Math.sqrt(a*a + b*b);
}

// Draws a piece of color at the passed in coordinates
function drawPiece(radius, color, x, y) {   
    ctx = gameArea.context;
    //console.log(color);
    ctx.beginPath();
    ctx.fillStyle = color;
    ctx.arc(x,y,radius,0,2*Math.PI);
    ctx.fill();

    ctx.beginPath();
    ctx.arc(x,y,radius+1, 0, 2*Math.PI);
    ctx.stroke();
}

// timer function that updates every second
function timer() {
    time++;
    var div = document.getElementById("timer");
    div.innerHTML = fmtTime(time);
}

function fmtTime(t) {
    var min = Math.floor(t / 60);
    var sec = t - min * 60;
    if (sec < 10) sec = "0" + sec;
    return min +":"+sec;
}
// Function to invert colors to guarantee that colors don't blend in
function invertColor(hex) {
    hex = hex.slice(1);
    // invert color components
    var r = (255 - parseInt(hex.slice(0, 2), 16)).toString(16),
        g = (255 - parseInt(hex.slice(2, 4), 16)).toString(16),
        b = (255 - parseInt(hex.slice(4, 6), 16)).toString(16);
    // pad each with zeros and return
    return '#' + padZero(r) + padZero(g) + padZero(b);
}

function padZero(str, len) {
    len = len || 2;
    var zeros = new Array(len).join('0');
    return (zeros + str).slice(-len);
}