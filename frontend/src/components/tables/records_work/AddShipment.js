import React, { useState } from 'react';
import { QRCodeSVG } from 'qrcode.react';
import { addShipments } from '../../api';
import './Modal.css';

const AddShipmentModal = ({ isOpen, onClose, onAddShipment }) => {
    const [shipment, setShipment] = useState({
        FromCity: '',
        FromWarehouseAddress: '',
        ToCity: '',
        ToWarehouseAddress: '',
        ToCountry: '',
    });

    const [isQRCodeVisible, setIsQRCodeVisible] = useState(false);
    const [qrData, setQrData] = useState(null);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setShipment({
            ...shipment,
            [name]: value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await addShipments(shipment);
            console.log(response.data);
            const addedShipment = response.data;
            onAddShipment(addedShipment.shipment);

            setQrData(JSON.stringify(addedShipment));
            setIsQRCodeVisible(true);
        } catch (error) {
            console.error('Error adding shipment:', error);
        }
    };

    const handleModalClose = () => {
        setIsQRCodeVisible(false); 
        setQrData(''); 
        onClose();
    };

    if (!isOpen) return null;

    return (
        <div className="modal">
            <div className="modal-content">
                <button className="close-button" onClick={handleModalClose}>
                    &times;
                </button>
                {!isQRCodeVisible ? (
                    <form onSubmit={handleSubmit}>
                        <h2 className="title">Add Shipment</h2>
                        <div className="form-group">
                            <label>From City</label>
                            <input
                                type="text"
                                name="FromCity"
                                value={shipment.FromCity}
                                onChange={handleChange}
                                className="input"
                            />
                        </div>
                        <div className="form-group">
                            <label>From Warehouse</label>
                            <input
                                type="text"
                                name="FromWarehouseAddress"
                                value={shipment.FromWarehouseAddress}
                                onChange={handleChange}
                                className="input"
                            />
                        </div>
                        <div className="form-group">
                            <label>To City</label>
                            <input
                                type="text"
                                name="ToCity"
                                value={shipment.ToCity}
                                onChange={handleChange}
                                className="input"
                            />
                        </div>
                        <div className="form-group">
                            <label>To Warehouse</label>
                            <input
                                type="text"
                                name="ToWarehouseAddress"
                                value={shipment.ToWarehouseAddress}
                                onChange={handleChange}
                                className="input"
                            />
                        </div>
                        <div className="form-group">
                            <label>To Country</label>
                            <input
                                type="text"
                                name="ToCountry"
                                value={shipment.ToCountry}
                                onChange={handleChange}
                                className="input"
                            />
                        </div>
                        <button type="submit" className="button">Add Shipment</button>
                    </form>
                ) : (
                    <div className="qr-code-container">
                        <h3>Add Shipment</h3>
                        <div className="qr-code">
                            <QRCodeSVG value={qrData} size={200} />
                        </div>
                        <div className="modal-footer">
                            <button className="ok-button" onClick={onClose}>
                            OK
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};

export default AddShipmentModal;