using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Net.Sockets;
using Google.Protobuf.Examples.AddressBook;
using UnityEngine;
using Google.Protobuf;
using Networking.Messages;

public class ClientBehaviour : MonoBehaviour {
    public string server = "localhost";
    public int port = 8001;

    private TcpClient client;
    private NetworkStream stream;

    private NetworkConnection networkConnection;

    private void Awake() {
        networkConnection = new NetworkConnection(server, port);
    }

    void Update() {
        if (Input.GetKeyDown(KeyCode.A)) {
            Debug.Log("Connecting..");
            StartCoroutine(ConnectAndSendSomeMessages());
        }

        if (Input.GetKeyDown(KeyCode.C)) {
            networkConnection.Connect();

            for (int i = 0; i < 1; i++) {
                TestMessage tmsg = new TestMessage {Number = i, Greeting = "Hello from the client!"};

                tmsg.Pl1 = i switch {
                    0 => new PayLoad1() {Msg = "payload1hello!"},
                    _ => tmsg.Pl1
                };

                networkConnection.SendMessage(tmsg);
            }
        }

        Debug.Log($"Connected: {networkConnection.connected} MessageQueueCount: {networkConnection.MessageQueue.Count}");
        if (networkConnection.connected && networkConnection.MessageQueue.TryDequeue(out TestMessage msg)) {
            switch (msg.PayloadCase) {
                case TestMessage.PayloadOneofCase.Pl1:
                    Debug.Log(msg.Pl1.Msg);
                    break;
                case TestMessage.PayloadOneofCase.Pl2:
                    Debug.Log(msg.Pl2.Number);
                    break;
                case TestMessage.PayloadOneofCase.Pl3:
                    Debug.Log(msg.Pl3.Percent);
                    break;
            }
            
        }

    }

    IEnumerator ConnectAndSendSomeMessages() {
        Debug.Log("Connecting..");
        client = new TcpClient(server, port);
        stream = client.GetStream();

        for (int i = 0; i < 3; i++) {
            yield return new WaitForSeconds(0.5f);
            string msg = "message" + i + "\n";

            TestMessage tmsg = new TestMessage {Number = i, Greeting = $"Msg {i} from the client::"};

            switch (i) {
                case 0:
                    tmsg.Pl1 = new PayLoad1() { Msg = "payload1hello!" };
                    break;
                case 1:
                    tmsg.Pl2 = new PayLoad2() { Number = 111 };
                    break;
                case 2:
                    tmsg.Pl3 = new PayLoad3() { Percent = 0.90f };
                    break;
            }

            Debug.Log("Sending message: " + msg);
            Byte[] data = System.Text.Encoding.ASCII.GetBytes(msg);
            //stream.Write(data, 0, data.Length);
            
            tmsg.WriteTo(stream);
            
            data = new Byte[256];

            String responseData = String.Empty;
            // Read the first batch of the TcpServer response bytes.
            Int32 bytes = stream.Read(data, 0, data.Length);
            responseData = System.Text.Encoding.ASCII.GetString(data, 0, bytes);
            Debug.LogFormat("Received: {0}", responseData);
        }
        
        // Close everything.
        stream.Close();
        client.Close();
        
        Debug.Log("Connection closed.");
    }
}
