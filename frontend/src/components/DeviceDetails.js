import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { getDeviceDetails, deleteDevice } from "../api";

const DeviceDetails = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [device, setDevice] = useState(null);

    useEffect(() => {
        const fetchDeviceDetails = async () => {
            try {
                const data = await getDeviceDetails(id);
                setDevice(data);
            } catch (error) {
                console.error("Ошибка при загрузке устройства:", error);
            }
        };
        fetchDeviceDetails();
    }, [id]);

    const handleDelete = async () => {
        try {
            await deleteDevice(id);
            navigate("/");
        } catch (error) {
            console.error("Ошибка при удалении устройства:", error);
        }
    };

    if (!device) return <div>Загрузка...</div>;

    return (
        <div>
            <h1>Устройство: {device.serialNumber}</h1>
            <p>Модель: {device.model}</p>
            <p>Адрес: {device.address}</p>
            <button onClick={() => navigate(`/devices/${id}/edit`)}>Редактировать</button>
            <button onClick={handleDelete}>Удалить</button>
            <button onClick={() => navigate(`/devices/${id}/telemetry`)}>Показать телеметрию</button>
        </div>
    );
};

export default DeviceDetails;