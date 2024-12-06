import React, { useEffect, useState } from 'react';
import { fetchDrugs } from '../api';
import './Tables.css';

const DrugsTable = () => {
    const [drugs, setDrugs] = useState([]);
    const [filteredDrugs, setFilteredDrugs] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');

    // Завантаження даних про ліки
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetchDrugs();
                setDrugs(response.data);
                setFilteredDrugs(response.data);
            } catch (error) {
                console.error('Error fetching drugs:', error);
            }
        };

        fetchData();
    }, []);

    useEffect(() => {
        const filtered = drugs.filter((drug) =>
            drug.DrugName.toLowerCase().includes(searchTerm.toLowerCase()) ||
            drug.Type.TypeName.toLowerCase().includes(searchTerm.toLowerCase()) ||
            drug.Manufacturer.ManufacturerName.toLowerCase().includes(searchTerm.toLowerCase())
        );
        setFilteredDrugs(filtered);
    }, [searchTerm, drugs]);

    return (
        <div className="table-container">
            <h2 className="title">Search Drugs</h2>
            <div className="search-container">
                <div className="search-input-wrapper">
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Enter drug name, type, or manufacturer..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                    <span className="search-icon">&#128269;</span> {/* Unicode для значка лупи */}
                </div>
            </div>
            <div className="scrollable-table">
                <table className="table">
                    <thead>
                        <tr>
                            <th className="th">DrugID</th>
                            <th className="th">Drug Name</th>
                            <th className="th">Type</th>
                            <th className="th">Form</th>
                            <th className="th">Manufacturer</th>
                            <th className="th">Active Ingredient</th>
                            <th className="th">Concentration</th>
                            <th className="th">PackSize</th>
                            <th className="th">Dosage</th>
                            <th className="th">Prescription Required</th>
                            <th className="th">Storage Conditions</th>
                            <th className="th">ShelfLife</th>
                            <th className="th">Year Released</th>
                            <th className="th">Price</th>
                            <th className="th">Additional Notes</th>
                        </tr>
                    </thead>
                    <tbody>
                        {filteredDrugs.length > 0 ? (
                            filteredDrugs.map((drug) => (
                                <tr key={drug.DrugID}>
                                    <td className="td">{drug.DrugID}</td>
                                    <td className="td">{drug.DrugName}</td>
                                    <td className="td">{drug.Type.TypeName}</td>
                                    <td className="td">{drug.Form.FormName}</td>
                                    <td className="td">{drug.Manufacturer.ManufacturerName}</td>
                                    <td className="td">{drug.ActiveIngredient}</td>
                                    <td className="td">{drug.Concentration}</td>
                                    <td className="td">{drug.PackSize}</td>
                                    <td className="td">{drug.Dosage}</td>
                                    <td className="td">{drug.PrescriptionRequired ? 'Yes' : 'No'}</td>
                                    <td className="td">{drug.StorageConditions}</td>
                                    <td className="td">{drug.ShelfLife}</td>
                                    <td className="td">{drug.YearReleased}</td>
                                    <td className="td">{drug.Price}</td>
                                    <td className="td">{drug.AdditionalNotes}</td>
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan="15">No matching drugs found</td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default DrugsTable;
