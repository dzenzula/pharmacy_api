import React, { useEffect, useState } from 'react';
import { fetchDrugs } from '../api';
import './Tables.css'; // Импортируем CSS-файл для таблиц

const DrugsTable = () => {
    const [drugs, setDrugs] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetchDrugs(); // Используем уже настроенный axios экземпляр
                setDrugs(response.data); // Данные из API
            } catch (error) {
                console.error('Error fetching drugs:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <div className="table-container">
            <h2 className="title">Drugs Table</h2>
            <div className="scrollable-table">
                <table className="table">
                    <thead>
                        <tr>
                            <th className="th">DrugID</th>
                            <th className="th">Drug Name</th>
                            <th className="th">TypeID</th>
                            <th className="th">FormID</th>
                            <th className="th">ManufacturerID</th>
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
                        {drugs.length > 0 ? (
                            drugs.map((drug) => (
                                <tr key={drug.DrugID}>
                                    <td className="td">{drug.DrugID}</td>
                                    <td className="td">{drug.DrugName}</td>
                                    <td className="td">{drug.TypeID}</td>
                                    <td className="td">{drug.FormID}</td>
                                    <td className="td">{drug.ManufacturerID}</td>
                                    <td className="td">{drug.ActiveIngredient}</td>
                                    <td className="td">{drug.Concentration}</td>
                                    <td className="td">{drug.PackSize}</td>
                                    <td className="td">{drug.Dosage}</td>
                                    <td className="td">{drug.PrescriptionRequired ? "Yes" : "No"}</td>
                                    <td className="td">{drug.StorageConditions}</td>
                                    <td className="td">{drug.ShelfLife}</td>
                                    <td className="td">{drug.YearReleased}</td>
                                    <td className="td">{drug.Price}</td>
                                    <td className="td">{drug.AdditionalNotes}</td>
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan="15">No data available</td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default DrugsTable;
