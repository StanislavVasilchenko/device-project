import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navbar from "./components/Navbar";
import DeviceList from "./components/DeviceList";
import DeviceDetails from "./components/DeviceDetails";
import AddDeviceForm from "./components/AddDeviceForm";
import EditDeviceForm from "./components/EditDeviceForm";
import TelemetryChart from "./components/TelemetryChart";
import Login from "./components/Login";

const App = () => {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route path="/" element={<DeviceList />} />
                <Route path="/devices/:id" element={<DeviceDetails />} />
                <Route path="/devices/add" element={<AddDeviceForm />} />
                <Route path="/devices/:id/edit" element={<EditDeviceForm />} />
                <Route path="/devices/:id/telemetry" element={<TelemetryChart />} />
                <Route path="/login" element={<Login />} />
            </Routes>
        </Router>
    );
};

export default App;