import React from 'react'
import './styles/main.css'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import LoginPage from './pages/LoginPage'
import ListsPage from './pages/ListsPage'
import ListDetailsPage from './pages/ListDetailsPage'
import NotePage from './pages/NotePage'
import SignupPage from './pages/SignUpPage'

export default function App() {
  const redirectPath = localStorage.getItem('token') ? '/lists' : '/login'
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/lists" element={<ListsPage />} />
        <Route path="/lists/:listId" element={<ListDetailsPage />} />
        <Route path="/lists/:listId/items/:itemId" element={<NotePage />} />
        <Route path="*" element={<Navigate to={redirectPath} replace />} />
      </Routes>
    </BrowserRouter>
  )
}