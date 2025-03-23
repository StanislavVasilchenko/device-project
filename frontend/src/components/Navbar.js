import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav>
            <Link to="/">Устройства</Link>
            <Link to="/login">Войти</Link>
        </nav>
    );
};

export default Navbar;