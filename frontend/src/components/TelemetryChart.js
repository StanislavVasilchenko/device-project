import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getTelemetry } from "../api";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from "recharts";

const TelemetryChart = () => {
    const { id } = useParams();
    const [telemetry, setTelemetry] = useState([]);
    const [startDate, setStartDate] = useState("2023-01-01");
    const [endDate, setEndDate] = useState("2023-01-31");

    const fetchTelemetry = async () => {
        try {
            const data = await getTelemetry(id, startDate, endDate);
            setTelemetry(data);
        } catch (error) {
            console.error("Ошибка при загрузке телеметрии:", error);
        }
    };

    useEffect(() => {
        fetchTelemetry();
    }, [id, startDate, endDate]);

    return (
        <div style={{ padding: '20px', maxWidth: '1200px', margin: '0 auto' }}>
            <h1 style={{ marginBottom: '20px' }}>Телеметрия устройства</h1>

            <div style={{
                backgroundColor: '#f8f9fa',
                padding: '20px',
                borderRadius: '8px',
                marginBottom: '20px',
                display: 'flex',
                alignItems: 'center',
                gap: '15px'
            }}>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px', fontWeight: 'bold' }}>
                        Начальная дата:
                    </label>
                    <input
                        type="date"
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                        style={{
                            padding: '8px',
                            borderRadius: '4px',
                            border: '1px solid #ddd'
                        }}
                    />
                </div>

                <div>
                    <label style={{ display: 'block', marginBottom: '5px', fontWeight: 'bold' }}>
                        Конечная дата:
                    </label>
                    <input
                        type="date"
                        value={endDate}
                        onChange={(e) => setEndDate(e.target.value)}
                        style={{
                            padding: '8px',
                            borderRadius: '4px',
                            border: '1px solid #ddd'
                        }}
                    />
                </div>

                <button
                    onClick={fetchTelemetry}
                    style={{
                        padding: '8px 15px',
                        backgroundColor: '#2196F3',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer',
                        alignSelf: 'flex-end'
                    }}
                >
                    Обновить
                </button>
            </div>

            <div style={{ height: '400px', width: '100%' }}>
                <ResponsiveContainer width="100%" height="100%">
                    <LineChart data={telemetry} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="timestamp" />
                        <YAxis />
                        <Tooltip />
                        <Legend />
                        <Line type="monotone" dataKey="temperature" stroke="#8884d8" activeDot={{ r: 8 }} />
                        <Line type="monotone" dataKey="humidity" stroke="#82ca9d" />
                    </LineChart>
                </ResponsiveContainer>
            </div>
        </div>
    );
};

export default TelemetryChart;