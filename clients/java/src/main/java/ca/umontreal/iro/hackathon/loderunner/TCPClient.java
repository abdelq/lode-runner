package ca.umontreal.iro.hackathon.loderunner;

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
    private DataOutputStream os;
    private PrintWriter pw;

    private BufferedReader reader;

    public TCPClient() {
        try {
            this.clientSocket = new Socket(IP, PORT);
            this.os = new DataOutputStream(clientSocket.getOutputStream());
            this.pw = new PrintWriter(os);

            this.reader = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
        } catch (IOException ex) {
            Logger.getLogger(TCPClient.class.getName()).log(Level.SEVERE, null, ex);
            System.exit(1);
        }
    }

    public void send(String event, Object... args) {
        String arg1 = "" + args[1];

        if (!args[0].equals("direction")) {
            arg1 = "\"" + arg1 + "\"";
        }

        String msg = "{"
                + "\"event\":\"" + event + "\","
                + "\"data\":{"
                + "\"" + args[0] + "\":" + arg1 + ",";

        for (int i = 2; i < args.length; i += 2) {
            msg += "\"" + args[i] + "\":";

            if (args[i + 1] instanceof String) {
                msg += "\"" + args[i + 1] + "\"";
            } else {
                msg += args[i + 1];
            }

            if (i != args.length - 2) {
                msg += ",";
            }
        }

        msg += "}}";

        System.out.println(msg);
        pw.println(msg);
        pw.flush();
    }

    public String readNext() throws IOException {
        return this.reader.readLine();
    }

    public void join(String name, String room, int level) {
        this.send("join", "name", name, "room", room, "level", level);
    }
}