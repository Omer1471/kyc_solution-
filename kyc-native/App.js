// App.js
import React, { useState } from 'react';
import LoginScreen from './LoginScreen';
import RegistrationScreen from './RegistrationScreen';

export default function App() {
  const [isLogin, setIsLogin] = useState(true);

  const switchToLogin = () => {
    setIsLogin(true);
  };

  const switchToRegister = () => {
    setIsLogin(false);
  };

  return (
    <>
      {isLogin ? (
        <LoginScreen switchToRegister={switchToRegister} />
      ) : (
        <RegistrationScreen switchToLogin={switchToLogin} />
      )}
    </>
  );
}


