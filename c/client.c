#include "client.h"

int main(int argc, char *argv[]) {
    int sock;
    struct sockaddr_in server_addr;
    char buffer[BUFFER_SIZE];

    if (argc != 2) {
        fprintf(stderr, "Usage: %s <server_ip>\n", argv[0]);
        exit(EXIT_FAILURE);
    }

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        perror("Failed to create socket");
        exit(EXIT_FAILURE);
    }

    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(PORT);

    if (inet_pton(AF_INET, argv[1], &server_addr.sin_addr) <= 0) {
        perror("Invalid IP-adres");
        close(sock);
        exit(EXIT_FAILURE);
    }

    // Connect with the server
    if (connect(sock, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
        perror("Failed to connect to server");
        close(sock);
        exit(EXIT_FAILURE);
    }

    // Communicate with the server
    while (1) {
        printf("Message: ");
        if (fgets(buffer, BUFFER_SIZE, stdin) == NULL) {
            break;
        }

        send(sock, buffer, strlen(buffer), 0);

        memset(buffer, 0, BUFFER_SIZE);
        if (read(sock, buffer, BUFFER_SIZE) <= 0) {
            break;
        }

        printf("Server answer: %s\n", buffer);
    }

    close(sock);
    return 0;
}