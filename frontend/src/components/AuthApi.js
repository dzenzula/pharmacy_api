import axios from 'axios';

const API = axios.create({
    baseURL: 'http://192.168.0.101:3000',
});

export const fetchManufacturers = () => API.get('/manufacturers');
export const fetchRoles = () => API.get('/roles');
export const loginUser = (username, password) => API.post('/login', { username, password });
export const addUser = (username, password, email, role, manufacturerName) => API.post('/register', { username, password, email, role, manufacturerName });