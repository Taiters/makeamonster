import { useState } from 'react';

const validateRoomId = (roomId) => {
    if (roomId.length !== 5) {
        return 'room id must be 5 characters';
    }
    return null;
}
const Home = ({onJoin, onCreate}) => {
    const [roomId, setRoomId] = useState('');

    const handleJoin = (event) => {
        event.preventDefault();

        const error = validateRoomId(roomId);
        if (error !== null) {
            alert(error);
            return;
        }

        onJoin(roomId.toUpperCase());
    }

    return (
        <div className="Home">
            <h1>Make-a-Monster</h1>
            <div>
                <form onSubmit={handleJoin}>
                    <input type="text" value={roomId} onChange={(e) => setRoomId(e.target.value)}/>
                    <input type="submit" value="Join" />
                </form>
            </div>
            <button onClick={onCreate}>Create</button>
        </div>
    )
}

export default Home;