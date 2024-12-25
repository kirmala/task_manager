import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <h1 className="text-4xl font-bold mb-6">Welcome to Task Manager</h1>
        <p className="text-gray-700 mb-8">Manage your tasks efficiently and stay organized!</p>
        <div className="flex items-center justify-center min-h-screen">
            <div className="space-x-4">
                <Link
                    to="/login"
                    className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
                >
                    Login
                </Link>
                <Link
                    to="/register"
                    className="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition"
                >
                    Register
                </Link>
            </div>
        </div>
    </div>
    );
};

export default Home;