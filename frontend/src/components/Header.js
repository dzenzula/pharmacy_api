import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Login from './Login';
import Register from './Register';

const Header = ({ setIsAuthenticated }) => {
    const [isAuthenticatedLocal, setIsAuthenticatedLocal] = useState(false);
    const [username, setUsername] = useState('');
    const [permissions, setPermissions] = useState([]);
    const [role, setRole] = useState('');
    const [showLogin, setShowLogin] = useState(false);
    const [showRegister, setShowRegister] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const storedAuth = localStorage.getItem('isAuthenticated');
        if (storedAuth === 'true') {
            setIsAuthenticatedLocal(true);
            setIsAuthenticated(true);
    
            const storedUsername = localStorage.getItem('username');
            setUsername(storedUsername || '');
    
            const storedPermissions = localStorage.getItem('permissions');
            setPermissions(storedPermissions);
    
            const storedRole = localStorage.getItem('role');
            setRole(storedRole || '');
        }
    }, [setIsAuthenticated]);

    const handleOpenLogin = () => {
        setShowLogin(true);
        setShowRegister(false);
    };

    const handleOpenRegister = () => {
        setShowRegister(true);
        setShowLogin(false);
    };

    const handleLoginClose = () => {
        setShowLogin(false);
    };

    const handleRegisterClose = () => {
        setShowRegister(false);
    }

    const handleLoginSuccess = (user) => {
        const permissions = user.Role.Permissions.map(permission => permission.PermissionName);
        setIsAuthenticatedLocal(true);
        setIsAuthenticated(true);
        setUsername(user.Username);
        setPermissions(permissions);
        setRole(user.Role.RoleName);

        localStorage.setItem('isAuthenticated', 'true');
        localStorage.setItem('username', user.Username);
        localStorage.setItem('permissions', permissions);
        localStorage.setItem('role', user.Role.RoleName);
    
        handleLoginClose();
        navigate('/dashboard');
    };

    const handleRegisterSuccess = () => {
        handleRegisterClose();
        setShowRegister(false);
    };

    const handleLogout = () => {
        setIsAuthenticatedLocal(false);
        setIsAuthenticated(false);
        setUsername('');
        setPermissions([]);
        setRole('');

        localStorage.removeItem('token');
        localStorage.removeItem('isAuthenticated');
        localStorage.removeItem('username');
        localStorage.removeItem('user');
        localStorage.removeItem('permissions');
        localStorage.removeItem('role');

        navigate('/');
    };

    const hasAddUserPermission = permissions.includes('add_user');

    return (    
        <div style={styles.header}>
            <h2 style={styles.siteTitle}>Pharmacy Site</h2>
            {showLogin && (
                <div style={styles.modal}>
                    <div style={styles.modalContent}>
                        <button style={styles.closeButton} onClick={handleLoginClose}>
                            &times;
                        </button>
                        <Login onLoginSuccess={handleLoginSuccess} />
                    </div>
                </div>
            )}
            {showRegister && (
                <div style={styles.modal}>
                    <div style={styles.modalContent}>
                        <button style={styles.closeButton} onClick={handleRegisterClose}>
                            &times;
                        </button>
                        <Register onRegisterSuccess={handleRegisterSuccess} />
                    </div>
                </div>
            )}
            <div style={styles.authSection}>
                {isAuthenticatedLocal ? (
                    <>
                        <span style={styles.username}>{username}</span>
                        {hasAddUserPermission && (
                            <button style={styles.button} onClick={handleOpenRegister}>
                                Add User
                            </button>
                        )}
                        <button style={styles.button} onClick={handleLogout}>
                            Logout
                        </button>
                    </>
                ) : (
                    <button style={styles.button} onClick={handleOpenLogin}>
                        Login
                    </button>
                )}
            </div>
        </div>
    );
};

const styles = {
    header: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '20px 40px',
        backgroundColor: '#f8f9fa',
        borderBottom: '2px solid #dee2e6',
        fontFamily: "'Roboto', sans-serif"
    },
    siteTitle: {
        fontFamily: "'Lobster', cursive", // Custom font for title
        fontSize: '2rem',
        color: '#007bff',
        margin: 0
    },
    button: {
        marginLeft: '10px',
        padding: '8px 15px',
        cursor: 'pointer',
        backgroundColor: '#007bff',
        color: '#fff',
        border: 'none',
        borderRadius: '5px',
        fontWeight: 'bold'
    },
    modal: {
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        zIndex: 1000
    },
    modalContent: {
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '10px',
        position: 'relative',
        width: '400px',
        maxWidth: '90%',
        boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)'
    },
    closeButton: {
        position: 'absolute',
        top: '10px',
        right: '10px',
        background: 'none',
        border: 'none',
        fontSize: '20px',
        cursor: 'pointer',
        color: '#007bff'
    },
    authSection: {
        display: 'flex',
        alignItems: 'center'
    },
    username: {
        fontWeight: 'bold',
        marginRight: '15px',
        fontSize: '1.2rem',
        color: '#28a745'
    }
};

export default Header;
