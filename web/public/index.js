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

function createCanvas(id, mosaic) {

    var block = document.createElement('div');

    var title = document.createElement('p');
    block.appendChild(title);

    var canvas = document.createElement('canvas');
    block.appendChild(canvas);

    document.body.appendChild(block);

    canvas.id = id;
    canvas.width = document.body.clientWidth;
    canvas.height = document.body.clientHeight;

    if (mosaic) {
        block.classList.add('mosaic');
    }
}

function oneJSON(str) {
    var count = 0;

    var start = str.indexOf('{');
    if(start > -1)
        str = str.slice(start);

    for(var i=0; i<str.length; i++) {
        if (i !== 0 && count == 0)
            break;

        if(str[i] == '{') count++;
        else if(str[i] == '}') count--;
    }

    return str.slice(0, i);
}

function parseJSON(data) {
    var json = {event: "none"};

    try {
        json = JSON.parse(data.toString('utf8').trim());
    } catch(e) {
        // XXX : Count Hackula strikes again !
        // Si on n'arrive pas à parser la string, on a possiblement
        // reçu plusieurs objets en une seule string... FIXME
        console.error('JSON error : ', e, data.toString('utf8'));

        json = JSON.parse(oneJSON(data.toString('utf8').trim()));
    }

    return json;
}

function draw(tiles, room, lives, level) {
    var canvas = document.getElementById(room);
    var context = canvas.getContext('2d');

    var tileHeight = Math.round(canvas.height / tiles.length);
    var tileWidth = Math.round(canvas.width / tiles[0].length);

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

function redraw(tiles, room, lives, level) {
    var canvas = document.getElementById(room);
    var block = canvas.parentElement;
    var title = block.querySelector('p');

    var context = canvas.getContext('2d');

    var tileHeight = Math.round(canvas.height / tiles.length);
    var tileWidth = Math.round(canvas.width / tiles[0].length);

    var oldTiles = rooms[room];
    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            var tile = tiles[i][j];
            var oldTile = oldTiles[i][j];
            if (tile != oldTile) {
                if (tile !== '&' && tile !== '0' || oldTile === '$') {
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

    if(block.classList.contains('mosaic')) {
        title.innerHTML = room + " #" + level + ' ' + '♥' + lives;
    } else {
        title.innerHTML = 'Level ' + level + ' - ' + lives + " li" + (lives > 1 ? 'ves' : 'fe') + " left ";
    }
}
