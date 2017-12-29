var rooms = {};
var images = {
    " ": "img/empty.png",
    "&": "img/runner.png",
    "0": "img/guard.png",
    "#": "img/brick.png",
    "@": "img/block.png",
    "X": "img/trap.png",
    "H": "img/ladder.png",
    "S": "img/hladder.png",
    "-": "img/rope.png",
    "$": "img/gold.png"
};

Object.keys(images).forEach(function (tile) {
    var img = new Image();
    img.src = images[tile];
    images[tile] = img;
});

function createCanvas(id) {
    var canvas = document.createElement('canvas');
    document.body.appendChild(canvas);
    canvas.id = id;
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;
}

function draw(tiles, room) {
    var canvas = document.getElementById(room);
    var context = canvas.getContext('2d');

    var tileHeight = Math.floor(canvas.height / tiles.length);
    var tileWidth = Math.floor(canvas.width / tiles[0].length);

    context.clearRect(0, 0, canvas.width, canvas.height);
    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            context.drawImage(
                images[tiles[i][j]],
                j * tileWidth, i * tileHeight,
                tileWidth, tileHeight
            );
        }
    }
}

function redraw(tiles, room) {
    var canvas = document.getElementById(room);
    var context = canvas.getContext('2d');

    var tileHeight = Math.floor(canvas.height / tiles.length);
    var tileWidth = Math.floor(canvas.width / tiles[0].length);

    var oldTiles = rooms[room];
    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            var tile = tiles[i][j];
            var oldTile = oldTiles[i][j];
            if (tile != oldtile) {
                if (tile !== '&' && tile !== '0' || oldtile === '$') {
                    context.clearRect(
                        j * tileWidth, i * tileHeight,
                        tileWidth, tileHeight
                    );
                }
                context.drawImage(
                    images[tile],
                    j * tileWidth, i * tileHeight,
                    tileWidth, tileHeight
                );
            }
        }
    }
}
