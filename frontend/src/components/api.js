import axios from 'axios';

const API = axios.create({
    baseURL: 'http://localhost:3000/pharmacy', // URL бекенда
    headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
    },
});

export const fetchDrugs = () => API.get('/drugs');
export const fetchManufacturers = () => API.get('/manufacturers');
export const fetchShipments = () => API.get('/shipments');