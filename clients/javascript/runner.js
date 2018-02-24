// Possible moves
const moves = Object.freeze({
    NONE: 0, UP: 1, LEFT: 2, DOWN: 3, RIGHT: 4
});

module.exports = class Runner {
    constructor() {
        this.name = 'name';
        this.room = 'room';
        this.level = 1;
    }

    start(grid) {
        console.log(grid);
    }

    next(x, y) {
        console.log(`Coordonnées : ${x}, ${y}`);

        // Random action (move or dig)
        if (Math.random() < .5)
            this.move(Math.floor(Math.random() * 4) + 1);
        else
            this.dig(Math.random() < .5 ? moves.LEFT : moves.RIGHT);
    }
}
