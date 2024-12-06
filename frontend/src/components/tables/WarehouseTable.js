import React, { useState, useEffect } from 'react';
import { fetchWarehouse } from '../api';
import AddWarehouseModal from './records_work/AddWarehouse';
import EditWarehouseModal from './records_work/UpdateWarehouse';
import ShowWarehouseModal from './records_work/ShowWarehouse';
import './Tables.css';

const WarhouseTable = () => {
    const [warehouses, setWarehouses] = useState([]);
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [isShowModalOpen, setIsShowModalOpen] = useState(false);
    const [editingRowId, setEditingRowId] = useState(null);
    const [editingRowData, setEditingRowData] = useState({});
    const [selectedWarehouse, setSelectedWarehouse] = useState({});

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetchWarehouse();
                setWarehouses(response.data);
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

    const openShowModal = (detail) => {
        setSelectedWarehouse(detail);
        setIsShowModalOpen(true);
    };

    const closeShowModal = () => {
        setSelectedWarehouse({});
        setIsShowModalOpen(false);
    };

    const openEditModal = (detail) => {
        console.log(detail);
        setEditingRowId(detail.ID);
        setEditingRowData({ ...detail });
        setIsEditModalOpen(true);
    };

    const closeEditModal = () => {
        setIsEditModalOpen(false);
        setEditingRowId(null);
    };

    const handleAddDetail = (newDetail) => {
        setWarehouses(prevWarehouses => [...prevWarehouses, newDetail]);
    };

    const handleUpdateDetail = (updatedDetail) => {
        setWarehouses((prevDetails) =>
            prevDetails.map((detail) =>
                detail.ID === updatedDetail.ID ? updatedDetail : detail
            )
        );
        closeEditModal();
    };

    const hasManufacturerRole = ['manufacturer', 'admin'].includes(localStorage.getItem('role'));

    return (
        <div className="table-container">
            <div className="header">
                <h2 className="title">Warehouse Table</h2>
                {hasManufacturerRole && (
                    <button className="button add-button" onClick={openAddModal}>
                        Add Record
                    </button>
                )}
            </div>
            {isAddModalOpen && (
                <AddWarehouseModal
                    isOpen={isAddModalOpen}
                    onClose={closeAddModal}
                    onAddWarehouse={handleAddDetail}
                />
            )}
            {isEditModalOpen && (
                <EditWarehouseModal
                    isOpen={isEditModalOpen}
                    onClose={closeEditModal}
                    detail={editingRowData}
                    onUpdateWarehouse={handleUpdateDetail}
                />
            )}
            {isShowModalOpen && selectedWarehouse && (
                <ShowWarehouseModal
                    isOpen={isShowModalOpen}
                    onClose={closeShowModal}
                    detail={selectedWarehouse}
                />
            )}
            <div className="scrollable-table">
                <table className="table">
                    <thead>
                        <tr>
                            <th className="th">Drug Name</th>
                            <th className="th">Manufacturer</th>
                            <th className="th">Batch Number</th>
                            <th className="th">Quantity</th>
                            <th className="th">Certificate</th>
                            <th className="th">Date</th>
                            <th className="th">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {warehouses.length > 0 ? (
                            warehouses.map((warehouse, index) => (
                                <tr key={index}>
                                    <td className="td">{warehouse.Drug ? warehouse.Drug.DrugName : ''}</td>
                                    <td className="td">{warehouse.Manufacturer ? warehouse.Manufacturer.ManufacturerName : ''}</td>
                                    <td className="td">{warehouse.BatchNumber}</td>
                                    <td className="td">{warehouse.Quantity}</td>
                                    <td className="td">{warehouse.CertificateCode}</td>
                                    <td className="td">{warehouse.Timestamp}</td>
                                    <td className="td">
                                        <button 
                                            className="button small-button" 
                                            onClick={() => {
                                                hasManufacturerRole 
                                                    ? openEditModal(warehouse) 
                                                    : openShowModal(warehouse);
                                            }}
                                        >
                                            {hasManufacturerRole ? 'Edit' : 'Show'}
                                        </button>
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

export default WarhouseTable;