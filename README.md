
Distributed Threat Detection and Response System

This project is a Distributed Threat Detection and Response System built using Python and Golang. The system is designed to monitor network traffic, detect potential security threats, and respond automatically. The project demonstrates basic packet capture, threat detection, and logging capabilities.

Features:
- Network Traffic Monitoring: Captures real-time network traffic using Python and the Scapy library.
- Threat Detection: A Golang server processes packet data and detects suspicious activity, such as known malicious IP addresses or certain traffic patterns (e.g., port scanning).
- Automatic Responses: Logs the detected threats and responds when suspicious activity is detected.

Technologies:
- Python: Used for client-side network traffic monitoring.
- Golang: Used for server-side threat detection and logging.
- Scapy: A Python library for packet manipulation and sniffing.
- TCP Communication: Python client communicates with the Go server over TCP.

Setup Instructions:

Prerequisites:
- Python 3.8+: Required for the Python client.
- Golang 1.18+: Required for the Go server.
- Scapy: Install the Scapy library for packet capture in Python.
- Git: Optional, for cloning the repository.

Clone the Repository:
1. Clone the project repository to your local machine:
   git clone https://github.com/<your-username>/distributed-threat-detection.git

2. Navigate into the project directory:
   cd distributed-threat-detection

Install Dependencies:

Python Client:
1. Navigate to the client directory:
   cd client

2. Install the required Python packages (especially Scapy):
   pip install -r requirements.txt

Go Server:
1. Navigate to the server directory:
   cd server

2. Install necessary Go packages:
   go get github.com/google/gopacket

Running the System:

1. Start the Go Server:
   cd server
   go run server.go

   This will start the server, listening on 127.0.0.1:65432.

2. Run the Python Client:
   cd client
   sudo python3 client.py

   You'll be prompted to either simulate a suspicious packet or capture live traffic.

Simulating a Suspicious Packet:
You can simulate a suspicious packet by choosing option 1 in the client. The Python client will send a fake packet with an IP that triggers the Go server’s threat detection.

Capturing Live Traffic:
You can choose option 2 to start capturing real network traffic on your machine. Open websites, use ping, or use curl to generate traffic for the Python client to capture and send to the Go server.

Example:
ping google.com

How It Works:
1. Packet Capture: The Python client captures packets from the network using Scapy.
2. Packet Transmission: The client sends packet data (IP addresses, protocols, etc.) to the Go server over TCP.
3. Threat Detection: The Go server checks the packets for suspicious activity (e.g., known malicious IPs or port scanning).
4. Logging: Detected threats are logged into a file (packets.log) and the server responds to the client with a status message.

Example Use Cases:
- Network Monitoring: Track incoming and outgoing network traffic and log potential threats.
- Threat Detection: Automatically detect suspicious activity based on traffic patterns.
- Learning Tool: Understand how basic network security systems work, including packet capture, threat detection, and logging.

Future Enhancements:
- Encryption: Add TLS/SSL for secure communication between the Python client and Go server.
- Real-time Dashboard: Build a web dashboard to monitor captured packets and detected threats in real time.
- Advanced Detection: Implement machine learning algorithms for anomaly-based detection of network traffic.

License:
This project is licensed under the MIT License - see the LICENSE file for details.
