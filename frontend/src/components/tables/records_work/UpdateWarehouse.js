import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { updateWarehouseDetail, fetchDrugs, fetchManufacturers } from '../../api';
import './Modal.css';

const EditWarehouseModal = ({ isOpen, onClose, detail, onUpdateDetail }) => {
    const [updatedDetail, setUpdatedDetail] = useState(detail);
    const [drugs, setDrugs] = useState([]);
    const [manufacturers, setManufacturers] = useState([]);

    useEffect(() => {
        if (detail) setUpdatedDetail(detail);
    }, [detail]);

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

    const handleChange = (e) => {
        const { name, value } = e.target;

        // При изменении лекарства обновляем ManufacturerID
        if (name === 'DrugID') {
            const selectedDrug = drugs.find((drug) => drug.DrugID === parseInt(value, 10));
            if (selectedDrug) {
                setUpdatedDetail((prev) => ({
                    ...prev,
                    DrugID: value,
                    ManufacturerID: selectedDrug.ManufacturerID,
                }));
                return;
            }
        }

        setUpdatedDetail((prev) => ({ ...prev, [name]: value }));
    };

    const handleDateChange = (date) => {
        setUpdatedDetail((prev) => ({ ...prev, Timestamp: date }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const preparedDetail = {
                ...updatedDetail,
                DrugID: parseInt(updatedDetail.DrugID, 10),
                ManufacturerID: parseInt(updatedDetail.ManufacturerID, 10),
                Quantity: parseInt(updatedDetail.Quantity, 10),
            };

            const response = await updateWarehouseDetail(preparedDetail.ID, preparedDetail);
            const warehouseDetail = response.data;

            onUpdateDetail(warehouseDetail);
            onClose();
        } catch (error) {
            console.error('Error updating warehouse detail:', error);
        }
    };

    if (!isOpen) return null;

    return (
        <div className="modal">
            <div className="modal-content">
                <button className="close-button" onClick={onClose}>
                    &times;
                </button>
                <form onSubmit={handleSubmit}>
                    <h2 className="title">Edit Warehouse Details</h2>
                    <div className="form-group">
                        <label>Drug Name</label>
                        <select
                            name="DrugID"
                            value={updatedDetail.DrugID}
                            onChange={handleChange}
                            className="select"
                            required
                        >
                            <option value="">Select a drug</option>
                            {drugs.map((drug) => (
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
                            value={updatedDetail.ManufacturerID}
                            onChange={handleChange}
                            className="select"
                            required
                            disabled
                        >
                            <option value="">Select a manufacturer</option>
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
                            value={updatedDetail.BatchNumber}
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
                            value={updatedDetail.Quantity}
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
                            value={updatedDetail.CertificateCode}
                            onChange={handleChange}
                            className="input"
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label>Date and Time</label>
                        <DatePicker
                            selected={new Date(updatedDetail.Timestamp)}
                            onChange={handleDateChange}
                            showTimeSelect
                            dateFormat="yyyy-MM-dd HH:mm:ss"
                            className="input"
                            required
                        />
                    </div>
                    <button type="submit" className="button">
                        Save Changes
                    </button>
                </form>
            </div>
        </div>
    );
};

export default EditWarehouseModal;
