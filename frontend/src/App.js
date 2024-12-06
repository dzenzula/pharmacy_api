import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';
import DrugsTable from './components/tables/DrugsTable';
import ShipmentsTable from './components/tables/ShipmentsTable';
import './App.css'
import WarhouseTable from './components/tables/WarehouseTable';

const App = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        const storedAuth = localStorage.getItem('isAuthenticated');
        if (storedAuth) {
            setIsAuthenticated(true);
        }
    }, []);

    return (
        <Router>
            <Header setIsAuthenticated={setIsAuthenticated} />
            <div style={styles.main}>
                <Sidebar isAuthenticated={isAuthenticated} />
                <div style={styles.content}>
                    <Routes>
                        <Route path="/" element={<Dashboard />} />
                        <Route path="/login" element={<Login />} />
                        <Route path="/register" element={<Register />} />
                        <Route path="/dashboard" element={<Dashboard />} />
                        <Route path="/drugs" element={<DrugsTable />} />
                        <Route path="/shipments" element={<ShipmentsTable />} />
                        <Route path="/warehouse" element={<WarhouseTable />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
};

const styles = {
    main: {
        display: 'flex',
        height: '100%',
    },
    content: {
        flex: 1,
        padding: '20px',
        overflow: 'hidden',
    }
};

export default App;
