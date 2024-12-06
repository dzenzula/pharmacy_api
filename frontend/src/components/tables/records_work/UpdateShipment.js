import React, { useState, useEffect } from 'react';
import { updateShipment, fetchDrugs, addShipmentDetails, fetchShipments } from '../../api';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import './Modal.css';
import '../Tables.css';

const EditShipmentModal = ({ isOpen, onClose, shipment, onUpdateShipment }) => {
    const [editedShipment, setEditedShipment] = useState(shipment);
    const [drugs, setDrugs] = useState([]);
    const [newDetail, setNewDetail] = useState({
        DrugID: '',
        ManufacturerID: '',
        ManufacturerName: '',
        Quantity: '',
        Price: '',
        ManufacturingDate: ''
    });

    useEffect(() => {
        setEditedShipment(shipment);

        const fetchData = async () => {
            try {
                const response = await fetchDrugs();
                setDrugs(response.data);
            } catch (error) {
                console.error('Error fetching drugs:', error);
            }
        };

        fetchData();
    }, [shipment]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setEditedShipment(prevShipment => ({
            ...prevShipment,
            Status: {
                ...prevShipment.Status,
                Status: value
            }
        }));
    };

    const handleNewDetailChange = (e) => {
        const { name, value } = e.target;
        setNewDetail({
            ...newDetail,
            [name]: value,
        });

        if (name === 'DrugID') {
            const selectedDrug = drugs.find(drug => drug.DrugID === parseInt(value, 10));
            setNewDetail({
                DrugID: selectedDrug.DrugID,
                ManufacturerID: selectedDrug.Manufacturer.ManufacturerID,
                ManufacturerName: selectedDrug.Manufacturer.ManufacturerName,
                DrugName: selectedDrug.DrugName
            });
        }
    };

    const handleAddDetail = async (e) => {
        const detail = {
            BatchID: editedShipment.BatchID,
            DrugID: parseInt(newDetail.DrugID, 10),
            ManufacturerID: parseInt(newDetail.ManufacturerID, 10),
            Quantity: parseInt(newDetail.Quantity, 10),
            Price: parseFloat(newDetail.Price),
            ManufacturingDate: newDetail.ManufacturingDate,
        };

        try {
            await addShipmentDetails(detail);
        } catch (error) {
            console.error('Error adding details:', error);
        }

        const response = await fetchShipments();
        setEditedShipment(response.data.find(s => s.ID === editedShipment.ID));

        setNewDetail({
            DrugID: '',
            ManufacturerID: '',
            ManufacturerName: '',
            Quantity: '',
            Price: '',
            ManufacturingDate: ''
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await updateShipment(editedShipment.ID, editedShipment);
            onUpdateShipment(editedShipment);
            await fetchShipments();
            onClose();
        } catch (error) {
            console.error('Error updating shipment:', error);
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
                    <h2 className="title">Edit Shipment</h2>
                    <div className="form-group">
                        <label>Status</label>
                        <select
                            name="Status"
                            value={editedShipment.Status.Status}
                            onChange={handleChange}
                            className="select"
                        >
                            <option value="Accepted">Accepted</option>
                            <option value="Assembling">Assembling</option>
                            <option value="In Transit">In Transit</option>
                            <option value="Delivered">Delivered</option>
                            <option value="Cancelled">Cancelled</option>
                        </select>
                    </div>
                    <div className="form-group">
                        <label>From City</label>
                        <input
                            type="text"
                            name="FromCity"
                            value={editedShipment.FromCity}
                            onChange={handleChange}
                            className="input"
                        />
                    </div>
                    <div className="form-group">
                        <label>From Warehouse</label>
                        <input
                            type="text"
                            name="FromWarehouseAddress"
                            value={editedShipment.FromWarehouseAddress}
                            onChange={handleChange}
                            className="input"
                        />
                    </div>
                    <div className="form-group">
                        <label>To City</label>
                        <input
                            type="text"
                            name="ToCity"
                            value={editedShipment.ToCity}
                            onChange={handleChange}
                            className="input"
                        />
                    </div>
                    <div className="form-group">
                        <label>To Warehouse</label>
                        <input
                            type="text"
                            name="ToWarehouseAddress"
                            value={editedShipment.ToWarehouseAddress}
                            onChange={handleChange}
                            className="input"
                        />
                    </div>
                    <div className="form-group">
                        <label>To Country</label>
                        <input
                            type="text"
                            name="ToCountry"
                            value={editedShipment.ToCountry}
                            onChange={handleChange}
                            className="input"
                        />
                    </div>
                    <div className="form-group">
                        <label>Departure Date and Time</label>
                        <DatePicker
                            selected={editedShipment.Status.DepartureDate ? new Date(editedShipment.Status.DepartureDate) : null}
                            onChange={(date) =>
                                setEditedShipment((prevShipment) => ({
                                    ...prevShipment,
                                    Status: {
                                        ...prevShipment.Status,
                                        DepartureDate: date.toISOString(),
                                    },
                                }))
                            }
                            showTimeSelect
                            dateFormat="yyyy-MM-dd HH:mm:ss"
                            className="input"
                            required={editedShipment.Status.Status === 'In Transit'}
                        />
                    </div>
                    <div className="form-group">
                        <label>Details</label>
                            <table className="details-table">
                                <thead>
                                    <tr>
                                        <th className="th">Drug Name</th>
                                        <th className="th">Manufacturer Name</th>
                                        <th className="th">Quantity</th>
                                        <th className="th">Price</th>
                                        <th className="th">Manufacturing Date</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {editedShipment.Details.map((detail, index) => (
                                        <tr key={index}>
                                            <td className="td">{detail.Drug ? detail.Drug.DrugName : ''}</td>
                                            <td className="td">{detail.Manufacturer ? detail.Drug.Manufacturer.ManufacturerName : ''}</td>
                                            <td className="td">{detail.Quantity}</td>
                                            <td className="td">{detail.Price}</td>
                                            <td className="td">{detail.ManufacturingDate ? new Date(detail.ManufacturingDate).toLocaleDateString() : ''}</td>
                                        </tr>
                                    ))}
                                    {!['In Transit', 'Delivered', 'Cancelled'].includes(shipment.Status.Status) && (
                                        <tr>
                                            <td>
                                                <select name="DrugID" value={newDetail.DrugID} onChange={handleNewDetailChange}>
                                                    <option value="">Select a drug</option>
                                                    {drugs.map(drug => (
                                                        <option key={drug.DrugID} value={drug.DrugID}>
                                                            {drug.DrugName}
                                                        </option>
                                                    ))}
                                                </select>
                                            </td>
                                            <td>
                                                <input
                                                    type="text"
                                                    name="ManufacturerName"
                                                    value={newDetail.ManufacturerName}
                                                    readOnly
                                                    className="readonly"
                                                    required
                                                />
                                            </td>
                                            <td>
                                                <input
                                                    type="number"
                                                    name="Quantity"
                                                    value={newDetail.Quantity}
                                                    onChange={handleNewDetailChange}
                                                    required
                                                />
                                            </td>
                                            <td>
                                                <input
                                                    type="decimal"
                                                    name="Price"
                                                    value={newDetail.Price}
                                                    onChange={handleNewDetailChange}
                                                    required
                                                />
                                            </td>
                                            <td>
                                                <input
                                                    type="date"
                                                    name="ManufacturingDate"
                                                    value={newDetail.ManufacturingDate}
                                                    onChange={handleNewDetailChange}
                                                    required
                                                />
                                            </td>
                                        </tr>
                                    )}
                                </tbody>
                            </table>
                    </div>
                    {!['In Transit', 'Delivered', 'Cancelled'].includes(shipment.Status.Status) && (
                        <button type="button" className="add-button" onClick={handleAddDetail}>Add drug</button>
                    )}
                    <p></p>
                    <button type="submit" className="button">Save Changes</button>
                </form>
            </div>
        </div>
    );
};

export default EditShipmentModal;
