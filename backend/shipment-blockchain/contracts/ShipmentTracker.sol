// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ShipmentTracker {
    struct Shipment {
        string batchId;
        string fromCity;
        string toCity;
        string status;
        string timestamp;
    }

    mapping(string => Shipment[]) public shipments;

    event ShipmentAdded(
        string batchId,
        string fromCity,
        string toCity,
        string status,
        string timestamp
    );

    event ShipmentUpdated(string batchId, string status, string timestamp);

    // Додавання нової поставки
    function addShipment(
        string memory batchId,
        string memory fromCity,
        string memory toCity,
        string memory status,
        string memory timestamp
    ) public {
        require(bytes(batchId).length > 0, "Batch ID cannot be empty");
        require(bytes(fromCity).length > 0, "From city cannot be empty");
        require(bytes(toCity).length > 0, "To city cannot be empty");
        require(bytes(status).length > 0, "Status cannot be empty");
        require(bytes(timestamp).length > 0, "Timestamp cannot be empty");

        Shipment memory newShipment = Shipment(
            batchId,
            fromCity,
            toCity,
            status,
            timestamp
        );
        shipments[batchId].push(newShipment);

        emit ShipmentAdded(batchId, fromCity, toCity, status, timestamp);
    }

    // Оновлення статусу вантажа
    function updateShipment(
        string memory batchId,
        string memory status,
        string memory timestamp
    ) public {
        require(bytes(batchId).length > 0, "BatchID cannot be empty");
        require(bytes(status).length > 0, "Status cannot be empty");
        require(shipments[batchId].length > 0, "Shipment does not exist");

        Shipment memory updatedShipment = Shipment(
            batchId,
            shipments[batchId][0].fromCity,
            shipments[batchId][0].toCity,
            status,
            timestamp
        );

        shipments[batchId].push(updatedShipment);

        emit ShipmentUpdated(batchId, status, timestamp);
    }

    // Отримати історію змін
    function getShipmentHistory(
        string memory batchId
    ) public view returns (Shipment[] memory) {
        require(bytes(batchId).length > 0, "Batch ID cannot be empty");
        return shipments[batchId];
    }
}
