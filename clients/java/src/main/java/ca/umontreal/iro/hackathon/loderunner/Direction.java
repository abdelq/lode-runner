package ca.umontreal.iro.hackathon.loderunner;

public enum Direction {
    NONE(0), UP(1), LEFT(2), DOWN(3), RIGHT(4);
    
    private final int id;

    Direction(int id) {
        this.id = id;
    }

    public int getValue() {
        return id;
    }
    
    public static Direction fromInt(int n) {
        switch(n) {
            case 1:
                return UP;
            case 2:
                return LEFT;
            case 3:
                return DOWN;
            case 4:
                return RIGHT;
        }

        return NONE;
    }
}
