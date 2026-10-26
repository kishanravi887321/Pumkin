import { api } from './apiClient.js'

export async function register({ name, email, password }) {
  const data = await api.post('/user/register', { name, email, password })
  return data?.user || null
}

export async function login({ email, password }) {
  const data = await api.post('/auth/login', { email, password })
  return data?.user || null
}

export async function loginWithGoogle(googleToken) {
  const data = await api.post('/user/google-login', { google_token: googleToken })
  return data?.user || data?.payload || null
}

export function logout() {
  return api.post('/auth/logout')
}

export async function getCurrentUser() {
  const data = await api.get('/user/me')
  return data?.user || null
}
