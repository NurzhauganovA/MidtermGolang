import React, {useState, SyntheticEvent} from 'react';
import { Navigate } from 'react-router-dom';


const Login = () => {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    function getCookie(name: string): string {
        const nameLenPlus = (name.length + 1);
        // @ts-ignore
        return document.cookie.split(';').map(c => c.trim()).filter(cookie => {
            return cookie.substring(0, nameLenPlus) === `${name}=`;
        }).map(cookie => {
            return decodeURIComponent(cookie.substring(nameLenPlus));
        })[0] || null;
    }

    const csrftoken = getCookie('csrftoken');

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        let response = await fetch('http://localhost:8080/auth/sign-in', {
            method: "POST",
            headers: new Headers({
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'X-CSRFToken': csrftoken,
            }),
            credentials: 'include',
            body: JSON.stringify({
                username,
                password
            })
        });
        console.log(await response.json())
        setRedirect(true);
    }

    if(redirect) {
        return <Navigate to="/" />
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Sign In</h1>

            <input type="email" className="form-control" placeholder="Username" required
                   onChange={e => setUsername(e.target.value)}
            />
            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign In</button>
        </form>
    )
};

export default Login;