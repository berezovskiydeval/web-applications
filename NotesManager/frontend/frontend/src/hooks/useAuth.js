import { useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import api from '../shared/api'
const TOKEN_KEY = 'token'

export const useAuth = () => {
  const [token, setToken] = useState(() => localStorage.getItem(TOKEN_KEY))

  const login = t => {
    localStorage.setItem(TOKEN_KEY, t)
    setToken(t)
  }

  const logout = () => {
    localStorage.removeItem(TOKEN_KEY)
    setToken(null)
  }

  return { token, login, logout, isAuth: !!token }
}
export function useSignIn() {
  return useMutation({
    mutationFn: ({ username, password }) =>
      api.post('/auth/sign-in', { username, password }),
    onSuccess: ({ data }) => {
      localStorage.setItem('token', data.accessToken ?? data.token)
    },
  })
}

