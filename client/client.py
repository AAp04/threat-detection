import socket
from scapy.all import sniff, IP

# Server details
SERVER_HOST = '127.0.0.1'  # Server is running locally
SERVER_PORT = 65432        # Port the Go server is listening on

# Function to send packet data to the server
def send_packet_to_server(packet_data):
    try:
        # Create a TCP socket connection to the server
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((SERVER_HOST, SERVER_PORT))
            # Send the packet info to the server
            s.sendall(packet_data.encode('utf-8'))

            # Receive and print the server's response
            data = s.recv(1024)
            print(f"Server response: {data.decode('utf-8')}")
    except Exception as e:
        print(f"Failed to send packet to server: {e}")

# Function to process and capture live packets
def capture_packets(packet):
    # Extract IP layer details if available
    if IP in packet:
        packet_info = f"IP: {packet[IP].src} -> {packet[IP].dst}, Protocol: {packet[IP].proto}"
        print(f"Captured: {packet_info}")
        send_packet_to_server(packet_info)
    else:
        print("Non-IP packet captured")

# Function to simulate a suspicious packet for testing
def send_fake_packet_to_server():
    # Simulate a suspicious packet with a specific source IP
    fake_packet_info = "IP: 192.168.1.1 -> 216.58.217.206, Protocol: 6"
    print("Sending suspicious packet to server...")
    send_packet_to_server(fake_packet_info)

# Function to capture live network traffic
def start_packet_capture():
    print("Capturing packets...")
    sniff(filter="ip", prn=capture_packets, store=0)

if __name__ == "__main__":
    # Option to either simulate suspicious traffic or capture live packets
    choice = input("Enter '1' to simulate a suspicious packet, '2' to capture live packets: ")

    if choice == '1':
        # Simulate sending a suspicious packet
        send_fake_packet_to_server()
    elif choice == '2':
        # Start capturing live packets
        start_packet_capture()
    else:
        print("Invalid choice. Exiting.")
