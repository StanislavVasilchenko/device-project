import axios from "axios";

const API_URL = "http://localhost:8080/api"; // Базовый URL бэкенда

// Утилита для авторизации
export const login = async (username, password) => {
    try {
        const response = await axios.post(`${API_URL}/auth/login`, { username, password });
        return response.data.token; // Возвращаем токен
    } catch (error) {
        console.error("Ошибка авторизации:", error);
        throw error;
    }
};

// Утилита для получения списка устройств
export const getDevices = async (page = 1, limit = 10) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.get(`${API_URL}/devices`, {
            params: { page, limit },
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при загрузке устройств:", error);
        throw error;
    }
};

// Утилита для получения деталей устройства
export const getDeviceDetails = async (id) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.get(`${API_URL}/devices/${id}`, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при загрузке устройства:", error);
        throw error;
    }
};

// Утилита для добавления устройства
export const addDevice = async (device) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.post(`${API_URL}/devices`, device, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при добавлении устройства:", error);
        throw error;
    }
};

// Утилита для обновления устройства
export const updateDevice = async (id, device) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.put(`${API_URL}/devices/${id}`, device, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при обновлении устройства:", error);
        throw error;
    }
};

// Утилита для удаления устройства
export const deleteDevice = async (id) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.delete(`${API_URL}/devices/${id}`, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при удалении устройства:", error);
        throw error;
    }
};

// Утилита для получения телеметрии
export const getTelemetry = async (deviceId, start, end) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.get(`${API_URL}/devices/${deviceId}/telemetry`, {
            params: { start, end },
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при загрузке телеметрии:", error);
        throw error;
    }
};

// Утилита для добавления телеметрии
export const addTelemetry = async (telemetry) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.post(`${API_URL}/telemetry`, telemetry, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при добавлении телеметрии:", error);
        throw error;
    }
};

// Утилита для удаления телеметрии
export const deleteTelemetry = async (id) => {
    const token = localStorage.getItem("token"); // Получаем токен из localStorage
    try {
        const response = await axios.delete(`${API_URL}/telemetry/${id}`, {
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
        });
        return response.data;
    } catch (error) {
        console.error("Ошибка при удалении телеметрии:", error);
        throw error;
    }
};