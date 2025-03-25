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

    if (!device) return <div style={{ padding: '20px' }}>Загрузка...</div>;

    return (
        <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
            <h1 style={{ marginBottom: '20px' }}>Устройство: {device.serialNumber}</h1>

            <div style={{
                backgroundColor: '#f8f9fa',
                padding: '20px',
                borderRadius: '8px',
                marginBottom: '20px'
            }}>
                <p style={{ margin: '10px 0' }}><strong>Модель:</strong> {device.model}</p>
                <p style={{ margin: '10px 0' }}><strong>Адрес:</strong> {device.address}</p>
                <p style={{ margin: '10px 0' }}><strong>Статус:</strong> {device.status || 'Не указан'}</p>
            </div>

            <div style={{ display: 'flex', gap: '10px' }}>
                <button
                    onClick={() => navigate(`/devices/${id}/edit`)}
                    style={{
                        padding: '10px 15px',
                        backgroundColor: '#2196F3',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Редактировать
                </button>
                <button
                    onClick={handleDelete}
                    style={{
                        padding: '10px 15px',
                        backgroundColor: '#f44336',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Удалить
                </button>
                <button
                    onClick={() => navigate(`/devices/${id}/telemetry`)}
                    style={{
                        padding: '10px 15px',
                        backgroundColor: '#4CAF50',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Показать телеметрию
                </button>
            </div>
        </div>
    );
};

export default DeviceDetails;