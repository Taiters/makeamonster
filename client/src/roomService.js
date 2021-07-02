class RoomService {
    constructor(client) {
        this.client = client;
        this.client.onWelcome((welcome) => this.handleWelcome(welcome));
        this.client.onClientJoin((client) => this.handleClientJoin(client));
        this.client.onClientLeave((client) => this.handleClientLeave(client));
        this.client.onClientUpdate((client) => this.handleClientUpdate(client));
    }

    join(roomId) {
        this.client.join(roomId);
    }

    create() {
        this.client.create();
    }

    setName(name) {
        this.client.sendClientUpdate({
            Name: name
        });
        this.dispatch({
            type: 'nameUpdate',
            name,
        });
    }

    setDispatch(dispatch) {
        this.dispatch = dispatch;
    }

    handleWelcome(welcome) {
        this.dispatch({
            type: 'welcome',
            welcome,
        });
    }

    handleClientJoin(client) {
        this.dispatch({
            type: 'clientJoin',
            client,
        })
    }

    handleClientLeave(client) {
        this.dispatch({
            type: 'clientLeave',
            client,
        })
    }

    handleClientUpdate(client) {
        this.dispatch({
            type: 'clientUpdate',
            client,
        })
    }
}

export default RoomService;