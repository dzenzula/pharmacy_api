import React, { useState, useEffect } from 'react';
import { fetchShipments } from '../api'; // Импортируем API для загрузки данных о поставках
import AddShipmentModal from './records_work/AddShipment'; // Импортируем компонент модального окна
import EditShipmentModal from './records_work/UpdateShipment'; // Импортируем компонент модального окна для редактирования
import './Tables.css'; // Импортируем CSS-файл для таблиц

const ShipmentsTable = () => {
    const [shipments, setShipments] = useState([]);
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [editingRowId, setEditingRowId] = useState(null);
    const [editingRowData, setEditingRowData] = useState({});

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetchShipments();
                setShipments(response.data);
            } catch (error) {
                console.error('Error fetching shipments:', error);
            }
        };

        fetchData();
    }, []);

    const openAddModal = () => {
        setIsAddModalOpen(true);
    };

    const closeAddModal = () => {
        setIsAddModalOpen(false);
    };

    const openEditModal = (shipment) => {
        setEditingRowId(shipment.ID);
        setEditingRowData({ ...shipment });
        setIsEditModalOpen(true);
    };

    const closeEditModal = () => {
        setIsEditModalOpen(false);
        setEditingRowId(null);
    };

    const handleAddShipment = (newShipment) => {
        setShipments(prevShipmens => [...prevShipmens, newShipment]);
    };

    const handleUpdateShipment = (updatedShipment) => {
        setShipments((prevShipments) =>
            prevShipments.map((shipment) =>
                shipment.ID === updatedShipment.ID ? updatedShipment : shipment
            )
        );
        closeEditModal();
    };

    return (
        <div className="table-container">
            <div className="header">
                <h2 className="title">Shipments Table</h2>
                <button className="button add-button" onClick={openAddModal}>
                    Add Shipment
                </button>
            </div>
            {isAddModalOpen && (
                <AddShipmentModal
                    isOpen={isAddModalOpen}
                    onClose={closeAddModal}
                    onAddShipment={handleAddShipment}
                />
            )}
            {isEditModalOpen && (
                <EditShipmentModal
                    isOpen={isEditModalOpen}
                    onClose={closeEditModal}
                    shipment={editingRowData}
                    onUpdateShipment={handleUpdateShipment}
                />
            )}
            <div className="scrollable-table">
                <table className="table">
                    <thead>
                        <tr>
                            <th className="th">Batch</th>
                            <th className="th">Status</th>
                            <th className="th">From City</th>
                            <th className="th">From Warehouse</th>
                            <th className="th">To City</th>
                            <th className="th">To Warehouse</th>
                            <th className="th">To Country</th>
                            <th className="th">Timestamp</th>
                            <th className="th">Last Update</th>
                            <th className="th">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {shipments.length > 0 ? (
                            shipments.map((shipment) => (
                                <tr key={shipment.ID}>
                                    <td className="td">{shipment.BatchID}</td>
                                    <td className="td">{shipment.Status && shipment.Status.Status ? shipment.Status.Status : ''}</td>
                                    <td className="td">{shipment.FromCity}</td>
                                    <td className="td">{shipment.FromWarehouseAddress}</td>
                                    <td className="td">{shipment.ToCity}</td>
                                    <td className="td">{shipment.ToWarehouseAddress}</td>
                                    <td className="td">{shipment.ToCountry}</td>
                                    <td className="td">{shipment.Timestamp}</td>
                                    <td className="td">{shipment.LastUpdate}</td>
                                    <td className="td">
                                        {shipment.Status && shipment.Status.Status !== 'Cancelled' && (
                                            <button className="button small-button" onClick={() => openEditModal(shipment)}>
                                                Edit
                                            </button>
                                        )}
                                    </td>
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan="10" className="td">No data available</td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default ShipmentsTable;
