package ca.umontreal.iro.hackathon.loderunner;

/**
 *
 */
public class Runner extends BasicRunner {

    // TODO : Remplacer ceci par votre clé secrète
    public static final String ROOM = "";

    /* Utilisez cette variable pour choisir le niveau de départ
     *
     * Notez: le niveau de départ sera 1 pour tout
     * le monde pendant la compétition :v)
     */
    public static final int START_LEVEL = 1;

    public Runner() {
        super(ROOM, START_LEVEL);
    }

    @Override
    public void start(String[] grid) {
        System.out.println("Nouveau niveau ! Grille initiale reçue :");

        for (int i=0; i<grid.length; i++) {
            String ligne = grid[i];

            System.out.println(ligne);
        }
    }

    @Override
    public Move next(int x, int y) {
        System.out.println("Position du runner : (" + x + ", " + y + ")");

        int direction = (int) (Math.random() * 4 + 1);

        Direction dir = Direction.fromInt(direction);

        return new Move(Event.MOVE, dir);
    }
}
