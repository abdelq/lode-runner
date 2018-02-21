package noderunnerkeyboard;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.util.logging.Level;
import java.util.logging.Logger;

public class TCPClient {

    private static final String IP = "localhost";
    private static final int PORT = 1337;

    private Socket clientSocket;
    private DataOutputStream outToServer;
    private PrintWriter pw;

    public TCPClient() {
        try {
            this.clientSocket = new Socket(IP, PORT);
            this.outToServer = new DataOutputStream(clientSocket.getOutputStream());
            pw = new PrintWriter(outToServer);
        } catch (IOException ex) {
            Logger.getLogger(TCPClient.class.getName()).log(Level.SEVERE, null, ex);
            System.exit(1);
        }
    }

    public void send(String event, String... args) {
        String arg1 = "" + args[1];

        if (!args[0].equals("direction")) {
            arg1 = "\"" + arg1 + "\"";
        }

        String msg = "{"
                + "\"event\":\"" + event + "\","
                + "\"data\":{"
                + "\"" + args[0] + "\":" + arg1 + ","
                + "\"" + args[2] + "\":\"" + args[3] + "\""
                + "}"
                + "}";

        System.out.println(msg);

        pw.println(msg);
        pw.flush();
    }

    public void join(String room) {
        this.send("join", "name", "runner", "room", room);
    }

    public void doStuff() throws Exception {
        BufferedReader inFromServer = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
    }
}
