import React, {SyntheticEvent, useState} from 'react';
import { Navigate } from 'react-router-dom';


const Register = () => {

    const [name, setName] = useState('');
    const [surname, setSurname] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch('http://127.0.0.1:8000/auth/sign-up', {
            method: 'POST',
            headers: {'Content-type': 'application/json'},
            body: JSON.stringify({
                name,
                surname,
                username,
                email,
                password,
            })
        })

        setRedirect(true);
    }

    if(redirect) {
        return <Navigate to="/authorization/login" />
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Sign Up</h1>

            <input type="text" className="form-control" placeholder="Name"
                   onChange={e => setName(e.target.value)}
            /><br/>
            <input type="text" className="form-control" placeholder="Surname"
                   onChange={e => setSurname(e.target.value)}
            /><br/>
            <input type="text" className="form-control" placeholder="Username"
                   onChange={e => setUsername(e.target.value)}
            /><br/>
            <input type="text" className="form-control" placeholder="Email"
                   onChange={e => setEmail(e.target.value)}/>
            <input type="password" className="form-control" placeholder="Password"
                   onChange={e => setPassword(e.target.value)}
            /><br/>

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    );
};

export default Register;