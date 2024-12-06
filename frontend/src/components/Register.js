import React, { useState, useEffect } from 'react';
import { fetchManufacturers, fetchRoles, addUser } from './AuthApi';
import axios from 'axios';

const Register = ({ onRegisterSuccess }) => {
    const [username, setUsername] = useState('');
    const [role, setRole] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [manufacturerName, setManufacturerName] = useState('');
    const [error, setError] = useState('');
    const [roles, setRoles] = useState([]);
    const [manufacturers, setManufacturers] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const resRoles = await fetchRoles();
                setRoles(resRoles.data);

                const resManuf = await fetchManufacturers();
                setManufacturers(resManuf.data);
            } catch (error) {
                console.error('Error fetching roles:', error);
            }
        };

        fetchData();
    }, []);

    const handleRegister = async () => {
        try {
            await addUser(
                username,
                password,
                email,
                role,
                manufacturerName
            );
            const user = { username };
            onRegisterSuccess(user);
        } catch (error) {
            if (error.response && error.response.data) {
                setError(error.response.data);
            } else {
                setError('Registration error: ' + error);
            }
        }
    };

    return (
        <div style={styles.container}>
            <h2 style={styles.title}>Registration</h2>
            {error && <p style={styles.error}>{error.error}</p>}
            <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                style={styles.input}
            />
            <select 
                value={role}
                onChange={(e) => setRole(e.target.value)}
                style={styles.input}
            >
                <option value="" disabled>Select Role</option>
                {roles.map(r => (
                    <option key={r.RoleID} value={r.RoleName}>
                        {r.RoleName}
                    </option>
                ))}
            </select>
            {role === 'manufacturer' && (
                <select
                    value={manufacturerName}
                    onChange={(e) => setManufacturerName(e.target.value)}
                    style={styles.input}
                >
                    <option value="" disabled>Select Manufacturer</option>
                    {manufacturers.map(m => (
                        <option key={m.ManufacturerID} value={m.ManufacturerName}>
                            {m.ManufacturerName}
                        </option>
                    ))}
                </select>
            )}
            <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                style={styles.input}
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                style={styles.input}
            />
            <button onClick={handleRegister} style={styles.button}>Register</button>
        </div>
    );
};

const styles = {
    container: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
    },
    title: {
        marginBottom: '20px',
        color: '#333'
    },
    input: {
        width: '100%',
        padding: '10px',
        margin: '10px 0',
        border: '1px solid #ccc',
        borderRadius: '5px'
    },
    button: {
        width: '100%',
        padding: '10px',
        backgroundColor: '#007bff',
        color: '#fff',
        border: 'none',
        borderRadius: '5px',
        cursor: 'pointer'
    },
    error: {
        color: 'red',
        marginBottom: '10px'
    }
};

export default Register;
