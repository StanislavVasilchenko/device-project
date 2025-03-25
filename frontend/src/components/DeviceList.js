import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getDevices, deleteDevice } from "../api";

const DeviceList = () => {
    const [devices, setDevices] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchDevices = async () => {
            try {
                const data = await getDevices();
                setDevices(data);
            } catch (error) {
                console.error("Ошибка при загрузке устройств:", error);
            }
        };
        fetchDevices();
    }, []);

    const handleDelete = async (id) => {
        try {
            await deleteDevice(id);
            setDevices(devices.filter((device) => device.id !== id));
        } catch (error) {
            console.error("Ошибка при удалении устройства:", error);
        }
    };

    return (
        <div>
            <h1>Устройства</h1>
            <button onClick={() => navigate("/devices/add")}>Добавить устройство</button>
            <ul>
                {devices.map((device) => (
                    <li key={device.id}>
                        {device.serialNumber} - {device.model}
                        <button onClick={() => navigate(`/devices/${device.id}`)}>Подробнее</button>
                        <button onClick={() => handleDelete(device.id)}>Удалить</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default DeviceList;