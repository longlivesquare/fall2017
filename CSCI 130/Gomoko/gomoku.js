var turn = "white";
var matrix;
var victor;

function startGame() {
    gameArea.start(15);
}

var gameArea = {
    canvas : document.createElement("canvas"),
    turn : "white",
    start : function(size) {
        this.canvas.width = 600;
        this.canvas.height = 600;
        this.context = this.canvas.getContext("2d");
        this.size = size;
        matrix = [];
        for(var i=0; i< size; i++) {
            matrix[i] = [];
            for(var j = 0; j < size; j++) {
                matrix[i][j] = 0;
            }
        }

        this.x = ((Number(this.canvas.width)-50) / size);
        this.y = ((Number(this.canvas.height)-50) / size);
    
        for(var i = 0; i < size; i++) {
            this.context.moveTo(this.x * (i + 1), 0);
            this.context.lineTo(this. x * (i + 1), Number(this.canvas.height) - 50 + this.y);
            this.context.stroke();
        }
    
        for(var j = 0; j < size; j++)
        {
            this.context.moveTo(0, this.y * (j + 1));
            this.context.lineTo(Number(this.canvas.width) - 50 + this.x, this.y*(j + 1));
            this.context.stroke();
        }
        document.body.insertBefore(this.canvas, document.body.childNodes[0]);
        //this.interval = setInterval(updateGameArea, 20);
        window.addEventListener('click', function (e) {
            //drawPiece(15, turn, e.clientX-7, e.clientY-7);
            if (findClosest(e.x, e.y)) {
                console.log("Checking for win");
                if (fiveInARow()) {
                    console.log(turn + " is the winner");
                }
                turn = (turn == "black" ? "white":"black");
            }  
        })
    }, 
    clear : function(){
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }
}

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
        console.log(turn);
        matrix[xInd][yInd] = (turn == "white" ? 1 : -1);
        console.log(matrix);
        return true;
    }
    console.log(closestDist);
    console.log(closestx+":"+closesty);
    console.log(matInd);
    console.log(matrix);
    return false;
}


function fiveInARow() {
    var color;
    var inARow;
    
    //Test Row

    for(var j = 0; j < gameArea.size; j++) {
        color = 0;
        inARow = 0;
        for(var i = 0; i < gameArea.size; i++)
            if(matrix[i][j]==color && matrix[i][j] != 0) {
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a row by " + color);
                    return true;
                }
            }
            else {
                inARow = 1;
                color = matrix[i][j];
            }
    }

    // Test Column
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

    // Test top-left to bottom right diagonal
    for(var i = 0; i < gameArea.size - 4; i++) {
        color = 0;
        inARow = 0;
        var tempI = i;
        for(var j = 0; j < gameArea.size ; j++)
        {
            if(matrix[tempI][j]==color && matrix[tempI][j] != 0) {
                inARow++;
                if(inARow == 5) {
                    console.log("5 in a left to right diagonal by " + color);
                    console.log(i +":"+(j-5) + " to " + tempI+":"+j);
                    return true;
                }
                
                if(tempI < gameArea.size-1) {
                    tempI++;
                }
            }
            else {
                inARow = 1;
                color = matrix[tempI][j];
                tempI = i;
            }
            if(tempI < gameArea.size - 1) {
                tempI++;
            }
            else {
                tempI = i;
            }
        }
    }

    // Test top-right to bottom-left diagonal
    for(var i = 4; i < gameArea.size; i++) {
        color = 0;
        inARow = 0;
        var tempI = i;
        for(var j = 0; j < gameArea.size; j++)
        {
            if(matrix[tempI][j]==color && matrix[tempI][j] != 0) {
                inARow++;
                if(inARow == 5) {
                    return true;
                }
            }
            else {
                inARow = 1;
                color = matrix[tempI][j];
            }
            if(tempI > 0) {
                tempI--;
            }
            else {
                tempI = i;
            }

        }
    }

    return false;
   
}

function distFrom(x1, y1, x2, y2    ) {
    var a = Math.abs(x1 - x2);
    var b = Math.abs(y1 - y2);
    return Math.sqrt(a*a + b*b);
}

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