import React from 'react';  
import { Link } from "react-router-dom";

const Navbar = (props: { name: string, setName: (name:string) => void }) => {
  const logout = async () => {
    await fetch('http://localhost:8000/api/logout', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
    });

    props.setName('');
  }

  let menu;

  if (props.name === '') {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0 flex flex-row">
        <li className="nav-item active mx-3">
          <Link to="/login" className="nav-link" aria-current="page">Login</Link>
        </li>

        <li className="nav-item active">
          <Link to="/register" className="nav-link" aria-current="page">Register</Link>
        </li>
      </ul>
    )
  } else {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0 flex flex-row">
        <li className="nav-item active mx-3">
          <Link to="/login" className="nav-link" aria-current="page" onClick={logout}>
            Logout
          </Link>
        </li>

        <li className="nav-item active mx-3">
          {/* Display the logged-in user's name */}
          <div className="nav-link">
            <span>Welcome, {props.name}</span>
          </div>
        </li>
      </ul>
    )
  }

  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <Link to="/" className="navbar-brand">Home</Link>
        
        <div className='nb-right-section'>
          {menu}
         
          {/* <form className='flex flex-row'>
            <input className='nb-form-control me-2' type='search' placeholder='Search' aria-label="Search"/>
            <button className='btn btn-outline-success' type='submit'>Search</button>
          </form> */}
        </div>
      </div>
    </nav>
  )
}

export default Navbar;