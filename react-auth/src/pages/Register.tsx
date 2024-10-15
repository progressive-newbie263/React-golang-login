import React, {SyntheticEvent, useState} from "react";
import { useNavigate } from "react-router-dom";
//Redirect is an error. From GPT, it said Rediredct has been removed from REact recently.
//import { Redirect } from "react-router-dom";


const Register = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    const response = await fetch('http://localhost:8000/api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name,
        email,
        password,
      }),
    })

    const returnAccount = await response.json();

    console.log(returnAccount);

    //
    navigate('/login')
  }


  return (
    <form className="form-register" onSubmit={submit}>
      <h1 className="h3 nb-3 fw-normal text-center">Sign up</h1>

      <input type="text" className="form-control" placeholder="username" required 
      onChange={e => setName(e.target.value)}
      />
      
      <input type="email" className="form-control" placeholder="email" required
      onChange={e => setEmail(e.target.value)}
      />
      
      <input type="password" className="form-control" placeholder="password" required
      onChange={e => setPassword(e.target.value)}
      />

      <button className="w-100 btn btn-lg btn-primary" type="submit">
        Register
      </button>
    </form>
  )
}

export default Register;