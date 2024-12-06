import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { addWarehouseDetail, fetchDrugs, fetchManufacturers } from '../../api';
import './Modal.css';

const AddWarehouseModal = ({ isOpen, onClose, onAddWarehouse }) => {
    const [newDetail, setNewDetail] = useState({
        DrugID: 0,
        ManufacturerID: 0,
        BatchNumber: '',
        Quantity: 0,
        CertificateCode: '',
        Timestamp: new Date(),
    });

    const [drugs, setDrugs] = useState([]);
    const [manufacturers, setManufacturers] = useState([]);
    const [filteredDrugs, setFilteredDrugs] = useState([]);
    const [user, setUser] = useState(null);

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            setUser(JSON.parse(storedUser));
        }
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const drugsResponse = await fetchDrugs();
                setDrugs(drugsResponse.data);

                const manufacturersResponse = await fetchManufacturers();
                setManufacturers(manufacturersResponse.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    useEffect(() => {
        if (user?.DataAdditional?.ManufacturerName) {
            const manufacturer = manufacturers.find(
                (m) => m.ManufacturerName === user.DataAdditional.ManufacturerName
            );
            if (manufacturer) {
                const filtered = drugs.filter(
                    (drug) => drug.ManufacturerID === manufacturer.ManufacturerID
                );
                setFilteredDrugs(filtered);
            }
        } else {
            setFilteredDrugs(drugs);
        }
    }, [drugs, manufacturers, newDetail.ManufacturerID, user]);

    const handleChange = (e) => {
        const { name, value } = e.target;

        if (name === 'DrugID') {
            const selectedDrug = drugs.find((drug) => drug.DrugID === parseInt(value, 10));
            if (selectedDrug) {
                setNewDetail((prev) => ({
                    ...prev,
                    DrugID: value,
                    ManufacturerID: selectedDrug.ManufacturerID,
                }));
                return;
            }
        }

        setNewDetail((prev) => ({ ...prev, [name]: value }));
    };

    const handleDateChange = (date) => {
        setNewDetail((prev) => ({ ...prev, Date: date }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const preparedDetail = {
                ...newDetail,
                DrugID: parseInt(newDetail.DrugID, 10),
                ManufacturerID: parseInt(newDetail.ManufacturerID, 10),
                Quantity: parseInt(newDetail.Quantity, 10),
            };

            const response = await addWarehouseDetail(preparedDetail);
            const warehouseDetail = response.data;

            onAddWarehouse(warehouseDetail);
        } catch (error) {
            console.error('Error adding warehouse detail:', error);
        }
    };

    const handleModalClose = () => {
        onClose();
    };

    if (!isOpen) return null;

    return (
        <div className="modal">
            <div className="modal-content">
                <button className="close-button" onClick={handleModalClose}>
                    &times;
                </button>
                <form onSubmit={handleSubmit}>
                    <h2 className="title">Add Warehouse Details</h2>
                    <div className="form-group">
                        <label>Drug Name</label>
                        <select
                            name="DrugID"
                            value={newDetail.DrugID}
                            onChange={handleChange}
                            className="select"
                            required
                        >
                            <option value="">Select a drug</option>
                            {filteredDrugs.map((drug) => (
                                <option key={drug.DrugID} value={drug.DrugID}>
                                    {drug.DrugName}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="form-group">
                        <label>Manufacturer Name</label>
                        <select
                            name="ManufacturerID"
                            value={newDetail.ManufacturerID}
                            onChange={handleChange}
                            className="select"
                            required
                            disabled
                        >
                            {manufacturers.map((manufacturer) => (
                                <option
                                    key={manufacturer.ManufacturerID}
                                    value={manufacturer.ManufacturerID}
                                >
                                    {manufacturer.ManufacturerName}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="form-group">
                        <label>Batch Number</label>
                        <input
                            type="text"
                            name="BatchNumber"
                            value={newDetail.BatchNumber}
                            onChange={handleChange}
                            className="input"
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label>Quantity</label>
                        <input
                            type="number"
                            name="Quantity"
                            value={newDetail.Quantity}
                            onChange={handleChange}
                            className="input"
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label>Certificate</label>
                        <input
                            type="text"
                            name="CertificateCode"
                            value={newDetail.CertificateCode}
                            onChange={handleChange}
                            className="input"
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label>Date and Time</label>
                        <DatePicker
                            selected={newDetail.Date}
                            onChange={handleDateChange}
                            showTimeSelect
                            dateFormat="yyyy-MM-dd HH:mm:ss"
                            className="input"
                            required
                        />
                    </div>
                    <button type="submit" className="button">
                        Add Warehouse
                    </button>
                </form>
            </div>
        </div>
    );
};

export default AddWarehouseModal;
