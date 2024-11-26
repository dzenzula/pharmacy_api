import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';
import Table1 from './components/tables/Table1';
import Table2 from './components/tables/Table2';
import Table3 from './components/tables/Table3';
import Table4 from './components/tables/Table4';

const App = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);

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
                        <Route path="/table1" element={<Table1 />} />
                        <Route path="/table2" element={<Table2 />} />
                        <Route path="/table3" element={<Table3 />} />
                        <Route path="/table4" element={<Table4 />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
};

const styles = {
    main: {
        display: 'flex'
    },
    content: {
        flex: 1,
        padding: '20px'
    }
};

export default App;
