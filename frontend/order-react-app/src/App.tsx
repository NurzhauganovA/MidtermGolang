import React from 'react';
import './App.css';

import {BrowserRouter as Router, Routes, Route} from "react-router-dom";

import Nav from "./components/Nav";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {

    return (
        <div className="App">
            <Router>
                <Nav/>
                <main className="form-signin w-100 m-auto">
                    <Routes>
                        <Route path="/" element={<Home/>} />
                        <Route path="/authorization/login" element={<Login/>} />
                        <Route path="/authorization/register" element={<Register/>} />
                    </Routes>
                </main>
            </Router>
        </div>
    );
}

export default App;
