
var board;
var config = {
    position: 'start',
    onDrop: onDrop,
    draggable: true
}

var socket;

function onLoad() {
    socket = new WebSocket("ws://localhost:8080/ws");
    console.log("Attempting Connection...");
    
    socket.onopen = () => {
        console.log("Successfully Connected");
    };
    
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };
    
    socket.onmessage = event => {
        console.log("message event: ", event);
        splitString = String(event.data).split(" ");

        if (splitString[0] == "updatePosition") {
            fenString = splitString[1];
            board.position(fenString);
        }
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
    
    board = Chessboard('myBoard', config);
    document.getElementById("myButton").onclick = sendPosition;
    document.getElementById("gimmeNewPosition").onclick = gimmeNewPosition;

    document.getElementById("setRuyLopezBtn").onclick = setRuyLopezPosition;
    document.getElementById("setStartBtn").onclick = board.start;
}

function setRuyLopezPosition() {
    board.position('r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R');
}

function sendPosition() {
    console.log("sendPosition");
    fenString = board.fen();
    socket.send("userSentNewPosition " + fenString);
}

function gimmeNewPosition() {
    console.log("gimmeNewPosition");
    socket.send("gimmeNewPosition");
}

function onDrop (source, target, piece, newPos, oldPos, orientation) {
    console.log('onDrop');
    if (Chessboard.objToFen(newPos) == Chessboard.objToFen(oldPos)) {
        console.log("ignore drag drop - same position");
    }
    //socket.send("gimmeNewPosition");
}