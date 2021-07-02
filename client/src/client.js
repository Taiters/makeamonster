import messages from './proto/messages_pb.js';

class Client {
    constructor() {
        this.socket = null;
    }

    onWelcome(f) {
        this.onWelcome = f;
    }

    onClientUpdate(f) {
        this.onClientUpdate = f;
    }

    onClientJoin(f) {
        this.onClientJoin = f;
    }

    onClientLeave(f) {
        this.onClientLeave = f;
    }

    create() {
        this.socket = new WebSocket('ws://localhost:5000/create');
        this.socket.onmessage = (event) => this.handleMessage(event);
    }

    join(id) {
        this.socket = new WebSocket('ws://localhost:5000/join/' + id)
        this.socket.onmessage = (event) => this.handleMessage(event);
    }

    handleMessage(event) {
        console.log(event);
        console.log(event.data.text());
        event.data.arrayBuffer().then((buffer) => {
            const test = messages.Event.deserializeBinary(new Uint8Array(buffer));
            console.log(test.getClientId());
        })
        const message = JSON.parse(event.data);
        switch (message.MessageType) {
            case "welcome":
                this.onWelcome(message.Content)
                break;

            case "clientUpdate":
                this.onClientUpdate(message.Content)
                break;

            case "clientJoined":
                this.onClientJoin(message.Content)
                break;

            case "clientLeft":
                this.onClientLeave(message.Content)
                break;
        
            default:
                break;
        }
    }

    sendClientUpdate(content) {
        const data = JSON.stringify({
            MessageType: "clientUpdate",
            Content: content,
        });
        this.socket.send(data);
    }
}

export default Client