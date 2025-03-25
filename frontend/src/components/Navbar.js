import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav style={{
            backgroundColor: '#333',
            padding: '15px 20px',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center'
        }}>
            <div>
                <Link
                    to="/"
                    style={{
                        color: 'white',
                        textDecoration: 'none',
                        marginRight: '20px',
                        fontWeight: 'bold'
                    }}
                >
                    Устройства
                </Link>
            </div>
            <div>
                <Link
                    to="/login"
                    style={{
                        color: 'white',
                        textDecoration: 'none',
                        padding: '8px 15px',
                        backgroundColor: '#4CAF50',
                        borderRadius: '4px'
                    }}
                >
                    Войти
                </Link>
            </div>
        </nav>
    );
};

export default Navbar;