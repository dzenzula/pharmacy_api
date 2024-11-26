import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = ({ isAuthenticated }) => {
    const location = useLocation();

    return (
        <div style={styles.sidebar}>
            <ul style={styles.list}>
                {isAuthenticated && (
                    <>
                        <li>
                            <Link to="/table1" style={location.pathname === '/table1' ? styles.linkActive : styles.link}>Test1</Link>
                        </li>
                        <li>
                            <Link to="/table2" style={location.pathname === '/table2' ? styles.linkActive : styles.link}>Test2</Link>
                        </li>
                        <li>
                            <Link to="/table3" style={location.pathname === '/table3' ? styles.linkActive : styles.link}>Test3</Link>
                        </li>
                        <li>
                            <Link to="/table4" style={location.pathname === '/table4' ? styles.linkActive : styles.link}>Test4</Link>
                        </li>
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
