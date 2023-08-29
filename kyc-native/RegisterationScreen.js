import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet, Image } from 'react-native';

export default function RegistrationScreen() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleRegister = async () => {
    const jsonData = {
      username,
      password,
    };

    try {
      const response = await fetch('http://localhost:5000/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(jsonData),
      });

      if (!response.ok) {
        throw new Error('Failed to register user');
      }

      const data = await response.json();
      console.log(data);
      alert('User registered successfully!');
      // Navigate to the login screen or dashboard
    } catch (error) {
      console.error('Error:', error);
      alert('Failed to register user. Please try again later.');
    }
  };

  return (
    <View style={styles.container}>
      <View style={styles.logoContainer}>
      <Image source={require('../static/images/Kealogo.jpeg')} style={styles.brandLogo} />
      </View>
      <Text style={styles.title}>Sign Up with Kea</Text>
      <TextInput
        style={styles.input}
        placeholder="Username"
        value={username}
        onChangeText={setUsername}
      />
      <TextInput
        style={styles.input}
        placeholder="Password"
        secureTextEntry
        value={password}
        onChangeText={setPassword}
      />
      <TouchableOpacity style={styles.button} onPress={handleRegister}>
        <Text style={styles.buttonText}>Register</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#091723',
    alignItems: 'center',
    justifyContent: 'center',
  },
  logoContainer: {
    marginBottom: 20,
  },
  brandLogo: {
    width: 150,
    height: 150,
  },
  title: {
    fontSize: 24,
    marginBottom: 20,
    color: '#F2F8F3',
  },
  input: {
    width: '80%',
    padding: 10,
    marginBottom: 15,
    backgroundColor: '#F2F8F3',
    borderRadius: 8,
  },
  button: {
    width: '80%',
    padding: 15,
    backgroundColor: '#F2F8F3',
    borderRadius: 8,
    alignItems: 'center',
  },
  buttonText: {
    color: '#091723',
    fontSize: 18,
  },
});

