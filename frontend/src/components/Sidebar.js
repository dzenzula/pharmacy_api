import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = ({ isAuthenticated }) => {
    const location = useLocation();

    const hasManufacturerRole = localStorage.getItem('role') === 'manufacturer'

    return (
        <div style={styles.sidebar}>
            <ul style={styles.list}>
                {isAuthenticated && (
                    <>
                        <li><Link to="/drugs" style={location.pathname === '/drugs' ? styles.linkActive : styles.link}>Drugs</Link></li>
                        <li><Link to="/warehouse" style={location.pathname === '/warehouse' ? styles.linkActive : styles.link}>Warehouse</Link></li>
                        {!hasManufacturerRole && <li><Link to="/shipments" style={location.pathname === '/shipments' ? styles.linkActive : styles.link}>Shipments</Link></li>}
                    </>
                )}
            </ul>
        </div>
    );
};

const styles = {
    sidebar: {
        width: '200px',
        backgroundColor: '#343a40',
        color: '#fff',
        height: '100vh',
        paddingTop: '20px'
    },
    list: {
        listStyleType: 'none',
        padding: 0
    },
    link: {
        color: '#fff',
        textDecoration: 'none',
        padding: '10px',
        display: 'block',
        borderBottom: '1px solid #495057',
        transition: 'background-color 0.3s',
    },
    linkActive: {
        color: '#fff',
        textDecoration: 'none',
        padding: '10px',
        display: 'block',
        borderBottom: '1px solid #495057',
        backgroundColor: '#495057',
    }
};

export default Sidebar;
