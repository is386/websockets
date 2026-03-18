# Websockets

## WebSocket Server — Implementation Checklist

- [x] **TCP listener** — bind to a port, accept incoming connections
- [ ] **HTTP request parser** — read raw bytes, parse request line and headers (just enough to validate an upgrade request)
- [ ] **WebSocket handshake** — extract `Sec-WebSocket-Key`, SHA-1 hash it with the magic UUID, base64 encode it, send back the HTTP 101 response
- [ ] **Frame parser** — read the binary frame format: FIN bit, opcode, payload length (3 possible sizes), masking bit, masking key, and unmasked payload
- [ ] **Frame writer** — construct and send frames back to the client (server frames are never masked)
- [ ] **Opcode handling** — text frames, binary frames, ping/pong, and close frames each need distinct behavior
- [ ] **Connection close handshake** — when either side wants to close, a specific close frame exchange must happen before the TCP connection drops
