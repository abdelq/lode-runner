package ca.umontreal.iro.hackathon.loderunner;

/**
 *
 */
public class Move {

    public Event event;
    public Direction direction;

    public Move(Event event, Direction direction) {
        this.event = event;
        this.direction = direction;
    }
}
