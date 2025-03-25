import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getTelemetry } from "../api";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from "recharts";

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
        <div>
            <h1>Телеметрия устройства</h1>
            <div>
                <label>
                    Начальная дата:
                    <input
                        type="date"
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                    />
                </label>
                <label>
                    Конечная дата:
                    <input
                        type="date"
                        value={endDate}
                        onChange={(e) => setEndDate(e.target.value)}
                    />
                </label>
                <button onClick={fetchTelemetry}>Обновить</button>
            </div>
            <LineChart width={600} height={300} data={telemetry}>
                <XAxis dataKey="timestamp" />
                <YAxis />
                <CartesianGrid stroke="#eee" strokeDasharray="5 5" />
                <Line type="monotone" dataKey="temperature" stroke="#8884d8" />
                <Line type="monotone" dataKey="humidity" stroke="#82ca9d" />
                <Tooltip />
                <Legend />
            </LineChart>
        </div>
    );
};

export default TelemetryChart;