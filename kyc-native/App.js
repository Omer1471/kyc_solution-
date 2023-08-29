import React, { useState } from 'react';
import LoginScreen from './LoginScreen';
import RegistrationScreen from './RegisterationScreen';

export default function App() {
  const [isLogin, setIsLogin] = useState(true);

  return (
    <>
      {isLogin ? <LoginScreen /> : <RegistrationScreen />}
      <button onClick={() => setIsLogin(!isLogin)}>
        {isLogin ? "Go to Registration" : "Go to Login"}
      </button>
    </>
  );
}

