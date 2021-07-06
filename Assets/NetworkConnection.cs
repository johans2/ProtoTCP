using System;
using System.Collections.Concurrent;
using System.IO;
using System.Net.Sockets;
using System.Threading;
using Google.Protobuf;
using Networking.Messages;
using UnityEngine;

public class NetworkConnection {

    
    public volatile bool connected;
    public ConcurrentQueue<TestMessage> MessageQueue { get; private set; } = new ConcurrentQueue<TestMessage>();
    public ConcurrentQueue<TestMessage> TestQueue { get; private set; } = new ConcurrentQueue<TestMessage>();
    
    
    private Thread networkingThread;
    private volatile bool listen = false;
 
    private NetworkStream netStream;
    private TcpClient client;

    private string server;
    private int port;
    
    public NetworkConnection(string server, int port) {
        this.server = server;
        this.port = port;
        this.connected = false;
    }

    public void SendMessage(TestMessage msg) {
        Debug.Log("size: " + msg.CalculateSize());
        msg.WriteTo(netStream);
    }

    public void Connect() {
        client = new TcpClient(server, port);
        netStream = client.GetStream();
        networkingThread = new Thread(ConnectAndListen);
        networkingThread.IsBackground = true;
        networkingThread.Start();
        listen = true;
        connected = true;
    }

    private void ConnectAndListen() {
        try {
            while (listen) {
                if (!netStream.DataAvailable) {
                    continue;
                }
                
                int msgSize = netStream.ReadByte();

                if (msgSize == 0) {
                    continue;
                }

                var msgBuffer = new byte[msgSize];
                netStream.Read(msgBuffer, 0, msgSize);
                
                
                
                //TestMessage t = TestMessage.Parser.ParseFrom(netStream);
                //TestMessage t = TestMessage.Parser.ParseDelimitedFrom(netStream);
                TestMessage t = TestMessage.Parser.ParseFrom(msgBuffer);
                MessageQueue.Enqueue(t);
            }
        }
        catch (Exception e) {
            connected = false;
            client.Close();
            netStream.Dispose();
        }
    }

}
