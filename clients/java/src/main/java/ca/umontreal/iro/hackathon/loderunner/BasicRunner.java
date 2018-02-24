package ca.umontreal.iro.hackathon.loderunner;

import java.io.IOException;
import java.util.List;
import org.json.JSONArray;
import org.json.JSONObject;

/**
 *
 */
public abstract class BasicRunner {

    protected String name;
    protected String room;
    protected int level;

    private TCPClient client;

    public BasicRunner(String room, int level) {
        this.name = room;
        this.room = room;
        this.level = level;

        this.client = new TCPClient();
    }

    public BasicRunner(String room) {
        this(room, 1);
    }

    private void move(int direction) {
        client.send("move", "direction", "" + direction, "room", room);
    }

    private void dig(int direction) {
        client.send("dig", "direction", "" + direction, "room", room);
    }

    public void run() {
        this.client.join(name, room, level);

        try {
            String next = client.readNext();

            while (next != null) {

                JSONObject obj = new JSONObject(next);
                
                String event = obj.getString("event");
                
                if(event.equals("error")) {
                    System.err.println("Error : " + obj.getString("data"));
                    System.exit(-1);
                }
                
                JSONObject data = obj.getJSONObject("data");
                                
                switch(event) {
                    
                    case "start":
                        JSONArray tiles = data.getJSONArray("tiles");
                        
                        List<Object> list = tiles.toList();

                        String[] arr = list.toArray(new String[list.size()]);

                        start(arr);
                        break;

                    case "next":
                        JSONObject runner = data.getJSONObject("runner");

                        sendNext(runner.getInt("x"), runner.getInt("y"));

                        break;
                }

                next = client.readNext();
            }
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }

    private void sendNext(int x, int y) {
        Move m = this.next(x, y);

        if (m.event == Event.MOVE) {
            move(m.direction.getValue());
        } else if (m.event == Event.DIG) {
            dig(m.direction.getValue());
        }
    }

    public abstract void start(String[] grid);

    public abstract Move next(int x, int y);
}
