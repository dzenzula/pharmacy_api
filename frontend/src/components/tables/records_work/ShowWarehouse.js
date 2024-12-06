import React from 'react';
import './Modal.css';

const ShowWarehouseModal = ({ isOpen, onClose, detail }) => {
    console.log('ShowWarehouseModal props:', { isOpen, detail });

    if (!isOpen || !detail) return null;

    return (
        <div className="modal">
            <div className="modal-content show-modal-content">
                <button className="close-button" onClick={onClose}>
                    &times;
                </button>
                <h2 className="title">Warehouse Details</h2>
                {detail ? (
                    <div>
                        <div className="form-group">
                            <label>Drug Name</label>
                            <p>{detail.Drug ? detail.Drug.DrugName : 'N/A'}</p>
                        </div>
                        <div className="form-group">
                            <label>Manufacturer</label>
                            <p>{detail.Manufacturer ? detail.Manufacturer.ManufacturerName : 'N/A'}</p>
                        </div>
                        <div className="form-group">
                            <label>Batch Number</label>
                            <p>{detail.BatchNumber}</p>
                        </div>
                        <div className="form-group">
                            <label>Quantity</label>
                            <p>{detail.Quantity}</p>
                        </div>
                        <div className="form-group">
                            <label>Certificate</label>
                            <p>{detail.CertificateCode}</p>
                        </div>
                        <div className="form-group">
                            <label>Date</label>
                            <p>{detail.Timestamp}</p>
                        </div>
                    </div>
                ) : (
                    <p>No warehouse details available.</p>
                )}
            </div>
        </div>
    );
};

export default ShowWarehouseModal;
