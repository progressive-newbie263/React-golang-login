import React, { useState, SyntheticEvent } from "react";
import { useNavigate } from "react-router-dom";

const Login = (props: { setName: (name: string) => void}) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  //const [isRegistered, setIsRegistered] = useState(false);
  const navigate = useNavigate();


  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    const response = await fetch('http://localhost:8000/api/login', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
      body: JSON.stringify({
        email,
        password,
      }),
    })

    const returnAccount = await response.json();

    props.setName(returnAccount.name);

    console.log(returnAccount);

    //
    navigate('/')
  }

  return (
    <main className='form-signin' onSubmit={submit}>
      <form>
        <h1 className='h3 mb-3 fw-normal text-center'>Log in</h1>
        
        <input type='email' className='form-control' placeholder='abc@gmail.com' required autoFocus
        onChange={e => setEmail(e.target.value)}
        />

        <input type='password' className='form-control' placeholder='Password' required autoFocus
        onChange={e => setPassword(e.target.value)}
        />

        <button className='w-100 btn btn-lg btn-primary' type='submit'>Sign in</button>
      </form>
    </main>
  )
}

export default Login