package noderunnerai;

/**
 *
 */
public class Runner extends BasicRunner {

    public Runner() {
        /* TODO : Modifier avec les valeurs de votre nom d'équipe
         * et du niveau de départ que vous souhaitez tester.
         *
         * Notez que le niveau de départ sera 1 pour tout le monde
         * pendant la compétition */
        super("", 1); // Exemple : super("Ma super équipe !", 1);
    }

    @Override
    public void start(String[] grid) {

        System.out.println("Nouveau niveau ! Grille initiale reçue :");
        for (String row : grid) {
            System.out.println(row);
        }
    }

    @Override
    public Move next() {
        int direction = (int) (Math.random() * 4 + 1);
        Direction dir = Direction.fromInt(direction);

        return new Move(Event.MOVE, dir);
    }
}
