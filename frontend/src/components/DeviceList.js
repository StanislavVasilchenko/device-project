import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getDevices, deleteDevice } from "../api";

const DeviceList = () => {
    const [devices, setDevices] = useState([]);
    const [filteredDevices, setFilteredDevices] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [currentPage, setCurrentPage] = useState(1);
    const [devicesPerPage] = useState(5); // Количество устройств на странице
    const navigate = useNavigate();

    useEffect(() => {
        const fetchDevices = async () => {
            try {
                const data = await getDevices();
                setDevices(data);
                setFilteredDevices(data);
            } catch (error) {
                console.error("Ошибка при загрузке устройств:", error);
            }
        };
        fetchDevices();
    }, []);

    // Фильтрация устройств
    useEffect(() => {
        const filtered = devices.filter(device =>
            device.serialNumber.toLowerCase().includes(searchTerm.toLowerCase())
        );
        setFilteredDevices(filtered);
        setCurrentPage(1); // Сброс на первую страницу при изменении фильтра
    }, [searchTerm, devices]);

    // Пагинация
    const indexOfLastDevice = currentPage * devicesPerPage;
    const indexOfFirstDevice = indexOfLastDevice - devicesPerPage;
    const currentDevices = filteredDevices.slice(indexOfFirstDevice, indexOfLastDevice);
    const totalPages = Math.ceil(filteredDevices.length / devicesPerPage);

    const handleDelete = async (id) => {
        try {
            await deleteDevice(id);
            const updatedDevices = devices.filter(device => device.id !== id);
            setDevices(updatedDevices);
        } catch (error) {
            console.error("Ошибка при удалении устройства:", error);
        }
    };

    const paginate = (pageNumber) => setCurrentPage(pageNumber);

    return (
        <div style={{ padding: '20px' }}>
            <h1 style={{ marginBottom: '20px' }}>Список устройств</h1>

            {/* Фильтр по серийному номеру */}
            <div style={{ marginBottom: '20px' }}>
                <input
                    type="text"
                    placeholder="Поиск по серийному номеру"
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    style={{
                        padding: '8px',
                        width: '300px',
                        marginRight: '10px'
                    }}
                />
                <button
                    onClick={() => navigate("/devices/add")}
                    style={{
                        padding: '8px 16px',
                        backgroundColor: '#4CAF50',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Добавить устройство
                </button>
            </div>

            {/* Список устройств */}
            <ul style={{ listStyle: 'none', padding: 0 }}>
                {currentDevices.map((device) => (
                    <li
                        key={device.id}
                        style={{
                            padding: '10px',
                            margin: '10px 0',
                            border: '1px solid #ddd',
                            borderRadius: '4px',
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'center'
                        }}
                    >
                        <span>
                            <strong>Серийный номер:</strong> {device.serialNumber} |
                            <strong> Модель:</strong> {device.model}
                        </span>
                        <div>
                            <button
                                onClick={() => navigate(`/devices/${device.id}`)}
                                style={{
                                    marginRight: '10px',
                                    padding: '5px 10px',
                                    backgroundColor: '#2196F3',
                                    color: 'white',
                                    border: 'none',
                                    borderRadius: '4px',
                                    cursor: 'pointer'
                                }}
                            >
                                Подробнее
                            </button>
                            <button
                                onClick={() => handleDelete(device.id)}
                                style={{
                                    padding: '5px 10px',
                                    backgroundColor: '#f44336',
                                    color: 'white',
                                    border: 'none',
                                    borderRadius: '4px',
                                    cursor: 'pointer'
                                }}
                            >
                                Удалить
                            </button>
                        </div>
                    </li>
                ))}
            </ul>

            {/* Пагинация */}
            <div style={{ marginTop: '20px', display: 'flex', justifyContent: 'center' }}>
                {Array.from({ length: totalPages }, (_, i) => i + 1).map(number => (
                    <button
                        key={number}
                        onClick={() => paginate(number)}
                        style={{
                            margin: '0 5px',
                            padding: '5px 10px',
                            backgroundColor: currentPage === number ? '#4CAF50' : '#ddd',
                            color: currentPage === number ? 'white' : 'black',
                            border: 'none',
                            borderRadius: '4px',
                            cursor: 'pointer'
                        }}
                    >
                        {number}
                    </button>
                ))}
            </div>

            {/* Информация о пагинации */}
            <div style={{ marginTop: '10px', textAlign: 'center' }}>
                Показано {currentDevices.length} из {filteredDevices.length} устройств
            </div>
        </div>
    );
};

export default DeviceList;