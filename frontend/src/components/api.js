import axios from 'axios';

const API = axios.create({
    baseURL: 'http://192.168.0.101:3000/pharmacy',
});

API.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export const fetchDrugs = () => API.get('/drugs');
export const fetchManufacturers = () => API.get('/manufacturers');
export const fetchShipments = () => API.get('/shipments');
export const fetchWarehouse = () => API.get('/warehouse');
export const addShipments = (shipment) => API.post('/shipments', shipment);
export const updateShipment = (editingRowDataID, editingRowData) => API.put(`/shipments/${editingRowDataID}`, editingRowData);
export const updateWarehouseDetail = (editingRowDataID, editingRowData) => API.put(`/warehouse/${editingRowDataID}`, editingRowData);
export const addShipmentDetails = (details) => API.post('/shipment-details', details)
export const addWarehouseDetail = (details) => API.post('/warehouse', details);