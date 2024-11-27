import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Login from './Login';
import Register from './Register';

const Header = ({ setIsAuthenticated }) => {
    const [isAuthenticatedLocal, setIsAuthenticatedLocal] = useState(false);
    const [username, setUsername] = useState('');
    const [showLogin, setShowLogin] = useState(false); // Управление отображением формы входа
    const [showRegister, setShowRegister] = useState(false); // Управление отображением формы регистрации
    const navigate = useNavigate();

    useEffect(() => {
        const storedAuth = localStorage.getItem('isAuthenticated');
        if (storedAuth) {
            setIsAuthenticatedLocal(true);
            setIsAuthenticated(true);
            setUsername(localStorage.getItem('username'));
        }
    }, [setIsAuthenticated]);

    const handleOpenLogin = () => {
        setShowLogin(true);
        setShowRegister(false); // Закрываем форму регистрации, если открыта
    };

    const handleOpenRegister = () => {
        setShowRegister(true);
        setShowLogin(false); // Закрываем форму входа, если открыта
    };

    const handleClose = () => {
        setShowLogin(false);
        setShowRegister(false);
    };

    const handleLoginSuccess = (user) => {
        setIsAuthenticatedLocal(true);
        setIsAuthenticated(true);
        setUsername(user.username);
        localStorage.setItem('isAuthenticated', 'true');
        localStorage.setItem('username', user.username);
        handleClose();
        navigate('/dashboard');
    };

    const handleRegisterSuccess = (user) => {
        setIsAuthenticatedLocal(true);
        setIsAuthenticated(true);
        setUsername(user.username);
        localStorage.setItem('isAuthenticated', 'true');
        localStorage.setItem('username', user.username);
        handleClose();
        navigate('/dashboard');
    };

    const handleLogout = () => {
        setIsAuthenticatedLocal(false);
        setIsAuthenticated(false);
        setUsername('');
        localStorage.removeItem('isAuthenticated');
        localStorage.removeItem('username');
        navigate('/');
    };

    return (
        <div style={styles.header}>
            <h2>Pharmacy Site</h2>
            <div>
                {isAuthenticatedLocal ? (
                    <>
                        <span>{username}</span>
                        <button style={styles.button} onClick={handleLogout}>
                            Logout
                        </button>
                    </>
                ) : (
                    <>
                        <button style={styles.button} onClick={handleOpenLogin}>
                            Login
                        </button>
                        <button style={styles.button} onClick={handleOpenRegister}>
                            Registration
                        </button>
                    </>
                )}
            </div>

            {/* Модальное окно для входа */}
            {showLogin && (
                <div style={styles.modal}>
                    <div style={styles.modalContent}>
                        <button style={styles.closeButton} onClick={handleClose}>
                            &times;
                        </button>
                        <Login onLoginSuccess={handleLoginSuccess} />
                    </div>
                </div>
            )}

            {/* Модальное окно для регистрации */}
            {showRegister && (
                <div style={styles.modal}>
                    <div style={styles.modalContent}>
                        <button style={styles.closeButton} onClick={handleClose}>
                            &times;
                        </button>
                        <Register onRegisterSuccess={handleRegisterSuccess} />
                    </div>
                </div>
            )}
        </div>
    );
};

const styles = {
    header: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '10px 20px',
        backgroundColor: '#f8f9fa',
        borderBottom: '1px solid #dee2e6'
    },
    button: {
        marginLeft: '10px',
        padding: '5px 10px',
        cursor: 'pointer',
        backgroundColor: '#007bff',
        color: '#fff',
        border: 'none',
        borderRadius: '5px'
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
    }
};

export default Header;
