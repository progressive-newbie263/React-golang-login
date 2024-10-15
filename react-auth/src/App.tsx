import React, { useState, useEffect } from 'react';
import './styles/main.css';
import './styles/navbar.css';
import Login from './pages/Login';
import Home from './pages/Home';
import Register from './pages/Register';
import Navbar from './Component/Navbar';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

function App() {
  const [name, setName] = useState('');

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
        });

        const content = await response.json();

        // If the user is not logged in, ensure name is empty
        if (response.ok && content.name) {
          setName(content.name);
        } else {
          setName('');  // Ensure this sets to empty if the user is not authenticated
        }
      } catch (error) {
        console.error('Error fetching user:', error);
        setName('');  // Handle fetch errors by resetting the name
      }
    };

    fetchUser();
  }, []); // Add an empty dependency array to ensure this runs only once

  return (
    <>
      <BrowserRouter>
        <Navbar name={name} setName={setName} />
        <Routes>
          <Route path="/" element={<Home name={name} />} />
          <Route path="/login" element={<Login setName={setName} />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
