import { useState } from 'react';

const NameForm = ({onName}) => {
    const [name, setName] = useState('');

    const handleName = (event) => {
        event.preventDefault();
        onName(name);
    };

    return (
        <div className="NameForm">
            <form onSubmit={handleName}>
                <input type="text" value={name} onChange={(e) => setName(e.target.value)}/>
                <input type="submit" value="Join" />
            </form>
        </div>
    )
}

export default NameForm;