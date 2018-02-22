package noderunnerkeyboard;

import java.awt.BorderLayout;
import java.awt.Dimension;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.SwingConstants;

public class NodeRunnerKeyboard {

    // TODO : Remplacer ceci par votre nom d'équipe
    public static final String ROOM = "";

    private static TCPClient client;

    public static void main(String[] args) {
        JFrame frame = new JFrame("Node Runner");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        JLabel textLabel = new JLabel("Utilisez WASD pour contrôller le Runner", SwingConstants.CENTER);
        textLabel.setPreferredSize(new Dimension(300, 100));

        frame.getContentPane().add(textLabel, BorderLayout.CENTER);
        frame.setLocationRelativeTo(null);
        frame.pack();

        client = new TCPClient();

        client.join(ROOM);

        frame.setVisible(true);

        frame.addKeyListener(new KeyListener() {
            @Override
            public void keyTyped(KeyEvent e) {
            }

            @Override
            public void keyPressed(KeyEvent e) {
                switch (e.getKeyChar()) {
                    case 'q':
                        System.exit(-1);
                        break;
                    case 'w':
                        client.send("move", "direction", "1", "room", ROOM);
                        break;
                    case 'a':
                        client.send("move", "direction", "2", "room", ROOM);
                        break;
                    case 's':
                        client.send("move", "direction", "3", "room", ROOM);
                        break;
                    case 'd':
                        client.send("move", "direction", "4", "room", ROOM);
                        break;
                    case 'z':
                        client.send("dig", "direction", "2", "room", ROOM);
                        break;
                    case 'c':
                        client.send("dig", "direction", "4", "room", ROOM);
                        break;
                }
            }

            @Override
            public void keyReleased(KeyEvent e) {
            }
        });
    }
}
