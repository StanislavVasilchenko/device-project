import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { addDevice } from "../api";

const AddDeviceForm = () => {
    const [serialNumber, setSerialNumber] = useState("");
    const [model, setModel] = useState("");
    const [address, setAddress] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await addDevice({ serial_number: serialNumber, model, address });
            navigate("/");
        } catch (error) {
            console.error("Ошибка при добавлении устройства:", error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Серийный номер"
                value={serialNumber}
                onChange={(e) => setSerialNumber(e.target.value)}
            />
            <input
                type="text"
                placeholder="Модель"
                value={model}
                onChange={(e) => setModel(e.target.value)}
            />
            <input
                type="text"
                placeholder="Адрес"
                value={address}
                onChange={(e) => setAddress(e.target.value)}
            />
            <button type="submit">Добавить устройство</button>
        </form>
    );
};

export default AddDeviceForm;