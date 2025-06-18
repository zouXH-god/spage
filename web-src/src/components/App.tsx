"use client";

import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { useState, useEffect } from 'react';
import MainView from '@/views/MainView';
import OwnerView from '@/views/entities/OwnerView';
import ProjectView from '@/views/entities/ProjectView';
import LoginView from '@/views/auth/LoginView';
import MainLayout from '@/layouts/MainLayout';
import EmptyLayout from '@/layouts/EmptyLayout';

export default function App() {
    const [isClient, setIsClient] = useState(false);

    useEffect(() => {
        setIsClient(true);
    }, []);

    if (!isClient) {
        return (
            <div className="flex items-center justify-center min-h-screen">
                <div className="animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent"></div>
            </div>
        );
    }
    return (
        <BrowserRouter>
            <Routes>
                {/* 带导航栏的主布局 */}
                <Route element={<MainLayout />}>
                    <Route path="/" element={<MainView />} />
                    <Route path="/:owner" element={<OwnerView />} />
                    <Route path="/:owner/:project" element={<ProjectView />} />
                </Route>
                
                {/* 无导航栏的认证布局 */}
                <Route element={<EmptyLayout />}>
                    <Route path="/-/login" element={<LoginView />} />
                </Route>
            </Routes>
        </BrowserRouter>
    );
}