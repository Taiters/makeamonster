import { useReducer, useEffect } from 'react';
import './App.css';
import Client from './client';

import Home from './components/Home.js';
import NameForm from './components/NameForm.js';
import RoomService from './roomService';


const roomService = new RoomService(new Client());
const roomStateReducer = (state, action) => {
  switch (action.type) {
    case 'welcome':
      const clientId = action.welcome.ClientId;
      const room = action.welcome.Room;
      return {
        localClient: room.Clients[clientId],
        room: room,
      }
    case 'clientJoin':
      state.room.Clients[action.client.Id] = action.client;
      return {...state};
    case 'clientLeave':
      delete(state.room.Clients[action.client.Id])
      return {...state};
    case 'clientUpdate':
      state.room.Clients[action.client.Id] = action.client;
      return {...state};
    case 'nameUpdate':
      state.localClient.Name = action.name;
      return {...state};
    default:
      console.log('Unexpected action');
  }
};

const App = () => {
  const [roomState, dispatch] = useReducer(roomStateReducer, null);

  const handleCreate = () => roomService.create();
  const handleJoin = (roomId) => roomService.join(roomId);
  const handleName = (name) => roomService.setName(name);


  useEffect(() => {
    roomService.setDispatch(dispatch);
    console.log('in here');
  }, [dispatch]);
  
  const renderApp = () => {
    if (roomState !== null) {
      const clients = Object.values(roomState.room.Clients).map((client) => (
        <p>{client.Id}: {client.Name != null ? client.Name : '...'}</p>
      ));
      return (
        <div>
          <NameForm onName={handleName} />
          <p>Room: {roomState.room.Id}</p>
          <p>You are: {roomState.localClient.Id} {roomState.localClient.Name}</p>
          {clients}
        </div>
      );
    }

    return <Home
      onCreate={handleCreate}
      onJoin={handleJoin}
    />
  }

  return renderApp()
}

export default App;
