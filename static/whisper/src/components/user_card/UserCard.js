import React, {useState, useEffect} from 'react';
import {User} from "./user";

function UserCard() {
    const [user, setUser] = useState(User.init());
    useEffect(() => {
        setUser({
            ...user,
            age: user.age || 21,
            name: user.name || "test"
        });
    }, [user]);

    const onClick = () => {
        setUser({
            ...user,
            age: user.age + 1,
            name: "Changed Name"
        });
    };


    return (
        <div>
            User age: {user.age}
            <div/>
            User name: {user.name}
            <div/>
            <button onClick={onClick}>Add age by 1</button>
        </div>
    )
}

export default UserCard;
