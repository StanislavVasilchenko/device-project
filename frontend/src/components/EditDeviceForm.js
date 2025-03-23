import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { getDeviceDetails, updateDevice } from "../api";

const EditDeviceForm = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [device, setDevice] = useState({ serial_number: "", model: "", address: "" });

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

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await updateDevice(id, device);
            navigate(`/devices/${id}`);
        } catch (error) {
            console.error("Ошибка при обновлении устройства:", error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Серийный номер"
                value={device.serial_number}
                onChange={(e) => setDevice({ ...device, serial_number: e.target.value })}
            />
            <input
                type="text"
                placeholder="Модель"
                value={device.model}
                onChange={(e) => setDevice({ ...device, model: e.target.value })}
            />
            <input
                type="text"
                placeholder="Адрес"
                value={device.address}
                onChange={(e) => setDevice({ ...device, address: e.target.value })}
            />
            <button type="submit">Сохранить</button>
        </form>
    );
};

export default EditDeviceForm;